package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"
)

// getDirectoryContext 获取当前目录的文件列表，帮助 AI 生成更精准的命令
func getDirectoryContext() string {
	files, err := os.ReadDir(".")
	if err != nil {
		return "Unknown"
	}
	var fileNames []string
	for i, f := range files {
		// 限制数量，防止 Prompt 溢出
		if i >= 15 {
			break
		}
		// 加上简单的类型标识
		if f.IsDir() {
			fileNames = append(fileNames, f.Name()+"/")
		} else {
			fileNames = append(fileNames, f.Name())
		}
	}
	if len(fileNames) == 0 {
		return "Empty directory"
	}
	return strings.Join(fileNames, ", ")
}

// FetchCommand 请求 AI 生成命令
func FetchCommand(query, apiURL, apiKey, model, lang string) (string, error) {
	// 1. 收集环境变量与目录上下文
	dirFiles := getDirectoryContext()
	osInfo := runtime.GOOS
	userShell := os.Getenv("SHELL")
	if userShell == "" && osInfo == "windows" {
		userShell = "cmd/powershell"
	}

	// 2. 构建 System Prompt
	systemPrompt := fmt.Sprintf(`You are a CLI expert. 
OS: %s, Shell: %s.
Files in current directory: %s.
Output ONLY: [command]||[explanation].
Language: %s. No markdown. No quotes.`, osInfo, userShell, dirFiles, lang)

	// 3. 准备请求 Payload
	payload := map[string]interface{}{
		"model": model,
		"messages": []map[string]string{
			{"role": "system", "content": systemPrompt},
			{"role": "user", "content": query},
		},
		"temperature": 0.3, // 降低随机性，保证命令准确
	}

	body, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// 4. 执行请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return "", fmt.Errorf("AUTH_FAILURE")
	}

	respBody, _ := io.ReadAll(resp.Body)
	var res struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.Unmarshal(respBody, &res); err != nil {
		return "", err
	}

	if len(res.Choices) == 0 {
		return "", fmt.Errorf("AI returned no results")
	}

	return res.Choices[0].Message.Content, nil
}
