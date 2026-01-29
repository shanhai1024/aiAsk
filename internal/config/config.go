package config

import (
	"ask/internal/i18n" // 引入 i18n 包
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	APIKey  string `json:"api_key"`
	APIURL  string `json:"api_url"`
	AIModel string `json:"ai_model"`
}

func SetupGuide(path string, cfg *Config) {
	fmt.Println(i18n.T("setup_welcome"))
	fmt.Println("------------------------------------------------")
	fmt.Println(i18n.T("setup_select"))
	fmt.Println("1. DeepSeek")
	fmt.Println("2. ChatGPT (OpenAI)")
	fmt.Println("3. Gemini (Google)")
	fmt.Println("4. Custom (Ollama / Others)")
	fmt.Print(i18n.T("setup_input"))

	var choice string
	fmt.Scanln(&choice)

	switch strings.TrimSpace(choice) {
	case "1":
		cfg.APIURL = "https://api.deepseek.com/v1/chat/completions"
		cfg.AIModel = "deepseek-chat"
	case "2":
		cfg.APIURL = "https://api.openai.com/v1/chat/completions"
		cfg.AIModel = "gpt-4o-mini"
	case "3":
		cfg.APIURL = "https://generativelanguage.googleapis.com/v1beta/openai/chat/completions"
		cfg.AIModel = "gemini-1.5-flash"
	case "4":
		fmt.Print(i18n.T("setup_url"))
		fmt.Scanln(&cfg.APIURL)
		fmt.Print(i18n.T("setup_model"))
		fmt.Scanln(&cfg.AIModel)
	default:
		fmt.Println(i18n.T("setup_invalid"))
		cfg.APIURL = "https://api.deepseek.com/v1/chat/completions"
		cfg.AIModel = "deepseek-chat"
	}

	fmt.Print(i18n.T("setup_key"))
	fmt.Scanln(&cfg.APIKey)
	cfg.APIKey = strings.TrimSpace(cfg.APIKey)

	Save(path, *cfg)
	fmt.Println(i18n.T("setup_done"))
}

// LoadOrCreate & Save 保持不变
func LoadOrCreate(path *string, cfg *Config) {
	home, _ := os.UserHomeDir()
	*path = filepath.Join(home, ".ask_config.json")
	data, err := os.ReadFile(*path)
	if err == nil {
		json.Unmarshal(data, cfg)
	}
}

func Save(path string, cfg Config) error {
	data, _ := json.MarshalIndent(cfg, "", "  ")
	return os.WriteFile(path, data, 0644)
}
