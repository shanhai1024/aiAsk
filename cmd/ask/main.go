package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"ask/internal/ai"
	"ask/internal/config"
	"ask/internal/i18n"
	"ask/internal/shell"
	"ask/internal/ui"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	// Version 由构建时的 ldflags 注入，请勿手动修改
	Version = "dev"

	cfg        config.Config
	configPath string

	// UI 风格
	cyan    = lipgloss.NewStyle().Foreground(lipgloss.Color("#00D7FF"))
	green   = lipgloss.NewStyle().Foreground(lipgloss.Color("#04B575")).Bold(true)
	red     = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF4444")).Bold(true)
	warnBox = lipgloss.NewStyle().Padding(1, 2).Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#FF0000")).Bold(true)
)

func main() {
	i18n.Detect()
	config.LoadOrCreate(&configPath, &cfg)

	// 自动引导
	if cfg.APIKey == "" && !isConfigCmd(os.Args) {
		config.SetupGuide(configPath, &cfg)
	}

	rootCmd := &cobra.Command{
		Use:     "ask [" + i18n.T("requirement") + "]",
		Version: Version,
		Short:   i18n.T("root_short"),
		Example: i18n.T("root_example"),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				executeAsk(strings.Join(args, " "))
			} else {
				_ = cmd.Help()
			}
		},
	}

	// 注册分类帮助模板
	cobra.AddTemplateFunc("i18n", i18n.T)
	rootCmd.SetUsageTemplate(`{{i18n "usage"}}: {{.CommandPath}} [{{i18n "requirement"}}]

{{i18n "group_start"}}
  {{rpad "[requirement]" 14}} {{i18n "root_short"}}

{{i18n "group_config"}}
{{range .Commands}}{{if eq .Name "set"}}  {{rpad .Name 14}} {{.Short}}{{end}}{{end}}

{{i18n "group_help"}}
{{range .Commands}}{{if eq .Name "help"}}  {{rpad .Name 14}} {{.Short}}{{end}}{{end}}
`)

	setCmd := &cobra.Command{
		Use:   "set",
		Short: "⚙️  " + i18n.T("set_short"),
	}

	setCmd.AddCommand(&cobra.Command{
		Use:   "init",
		Short: "🔄 " + i18n.T("set_init_short"),
		Run: func(c *cobra.Command, a []string) {
			_ = os.Remove(configPath)
			newCfg := &config.Config{}
			config.SetupGuide(configPath, newCfg)
			cfg = *newCfg
		},
	})

	rootCmd.AddCommand(setCmd)
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	// 区分直接提问与子命令
	if len(os.Args) > 1 && !isConfigCmd(os.Args) {
		executeAsk(strings.Join(os.Args[1:], " "))
		return
	}

	_ = rootCmd.Execute()
}

func isConfigCmd(args []string) bool {
	if len(args) < 2 {
		return false
	}
	base := args[1]
	return base == "set" || base == "help" || base == "version" || strings.HasPrefix(base, "-")
}

func executeAsk(query string) {
	// 思考动画
	stopAnim := make(chan bool)
	go func() {
		spin := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
		for i := 0; ; i++ {
			select {
			case <-stopAnim:
				fmt.Print("\r\033[K")
				return
			default:
				fmt.Printf("\r%s %s", cyan.Render(spin[i%len(spin)]), i18n.T("ai_thinking"))
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	// AI调用
	raw, err := ai.FetchCommand(query, cfg.APIURL, cfg.APIKey, cfg.AIModel, string(i18n.CurrentLang))
	stopAnim <- true

	if err != nil {
		handleError(err)
		return
	}

	// TUI 选择
	items := ui.ParseLines(raw)
	if len(items) == 0 {
		return
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 15)
	l.Title = i18n.T("ui_title")

	m, err := tea.NewProgram(ui.NewModel(l), tea.WithAltScreen()).Run()
	if err != nil {
		return
	}

	res, ok := m.(ui.Model)
	if !ok || res.Choice == "" || isExit(res.Choice) {
		return
	}

	cmd := res.Choice

	// 风险拦截
	if shell.CheckRisk(cmd) == 1 {
		fmt.Println("\n" + warnBox.Render("⚠️  "+i18n.T("risk_warning")))
		fmt.Printf("Action: %s\n", red.Render(cmd))
		fmt.Print("Confirm execution? (y/N): ")

		var confirm string
		_, _ = fmt.Scanln(&confirm)
		if strings.ToLower(strings.TrimSpace(confirm)) != "y" {
			fmt.Println("Cancelled.")
			return
		}
	}

	//最终交付
	if !shell.CheckBin(cmd) {
		if install := shell.GetInstallCmd(cmd); install != "" {
			fmt.Printf("\n💡 Command not found. Suggestion: %s\n", install)
			fmt.Print("Install and run? (y/n): ")
			var cf string
			_, _ = fmt.Scanln(&cf)
			if strings.ToLower(strings.TrimSpace(cf)) == "y" {
				shell.Execute(install)
			}
		}
	}

	fmt.Printf("\n🚀 %s\n", green.Render("Running: "+cmd))
	shell.Execute(cmd)

	fmt.Print(cyan.Render(i18n.T("done_msg")))
	_, _ = fmt.Scanln()
}

func handleError(err error) {
	if err.Error() == "AUTH_FAILURE" {
		fmt.Println(red.Render("\n" + i18n.T("auth_err")))
		return
	}
	fmt.Printf("\n❌ Error: %v\n", err)
}

func isExit(choice string) bool {
	return choice == "退出" || choice == "Quit"
}
