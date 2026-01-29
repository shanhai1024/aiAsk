package ui

import (
	"ask/internal/i18n"
	"ask/internal/shell"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	docStyle    = lipgloss.NewStyle().Margin(1, 2)
	dangerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF4444")).Bold(true)
)

type Item struct{ CmdStr, Desc string }

func (i Item) FilterValue() string { return i.CmdStr }
func (i Item) Title() string {
	// 如果是危险命令，在列表中标红显示
	if shell.CheckRisk(i.CmdStr) == 1 {
		return dangerStyle.Render("⚠️  " + i.CmdStr)
	}
	return i.CmdStr
}
func (i Item) Description() string { return i.Desc }

type Model struct {
	List   list.Model
	Choice string
}

func NewModel(l list.Model) Model { return Model{List: l} }

func ParseLines(raw string) []list.Item {
	raw = strings.ReplaceAll(raw, "```bash", "")
	raw = strings.ReplaceAll(raw, "```", "")
	raw = strings.ReplaceAll(raw, "`", "")
	lines := strings.Split(strings.TrimSpace(raw), "\n")
	var items []list.Item
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if parts := strings.Split(l, "||"); len(parts) >= 2 {
			items = append(items, Item{CmdStr: strings.TrimSpace(parts[0]), Desc: strings.TrimSpace(parts[1])})
		}
	}
	if len(items) > 0 {
		exit := "Quit"
		if i18n.CurrentLang == i18n.ZH {
			exit = "退出"
		}
		items = append(items, Item{CmdStr: exit, Desc: "Exit"})
	}
	return items
}

func (m Model) Init() tea.Cmd { return nil }
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			if i, ok := m.List.SelectedItem().(Item); ok {
				m.Choice = i.CmdStr
			}
			return m, tea.Quit
		case "c":
			if i, ok := m.List.SelectedItem().(Item); ok {
				clipboard.WriteAll(i.CmdStr)
			}
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.List.SetSize(msg.Width-h, msg.Height-v)
	}
	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return docStyle.Render(m.List.View())
}
