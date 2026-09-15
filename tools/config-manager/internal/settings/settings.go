package settings

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Settings 是跨会话持久化的用户设置。
type Settings struct {
	WebRoot string `json:"webRoot"`
}

func filePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".config", "9thnet-config-manager", "settings.json"), nil
}

// Load 读取设置；文件不存在或损坏时返回零值，不报错。
func Load() Settings {
	var s Settings
	p, err := filePath()
	if err != nil {
		return s
	}
	b, err := os.ReadFile(p)
	if err != nil {
		return s
	}
	_ = json.Unmarshal(b, &s)
	return s
}

func Save(s Settings) error {
	p, err := filePath()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		return err
	}
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(p, append(b, '\n'), 0o644)
}
