package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"the9thnet/config-manager/internal/settings"
	"the9thnet/config-manager/internal/tui"
)

var configFiles = []string{"site.json", "services.json", "activities.json", "directions.json"}

func main() {
	dir := flag.String("dir", "", "配置文件目录（直接读写该目录，优先级最高）")
	web := flag.String("web", "", "网站根目录（读写 <网站根目录>/config；会记住该设置）")
	flag.Parse()

	opts, err := resolve(*dir, *web)
	if err != nil {
		fmt.Fprintln(os.Stderr, "错误: "+err.Error())
		os.Exit(1)
	}

	if err := tui.Run(opts); err != nil {
		fmt.Fprintln(os.Stderr, "错误: "+err.Error())
		os.Exit(1)
	}
}

func resolve(dir, web string) (tui.Options, error) {
	saved := settings.Load()
	if web == "" {
		web = saved.WebRoot
	}

	// 1. 显式 -dir：读写同一目录。
	if dir != "" {
		abs, err := mustConfigDir(dir)
		if err != nil {
			return tui.Options{}, err
		}
		return tui.Options{ConfigDir: abs, WebRoot: absOrEmpty(web)}, nil
	}

	sourceDir, sourceOK := discoverSource()

	// 2. 指定了网站根目录：读写 <web>/config；若其中已有配置则直接使用，
	//    否则从源目录读取、写入该网站目录（实现「生成到网站 config 目录」）。
	if web != "" {
		absWeb, err := filepath.Abs(web)
		if err != nil {
			return tui.Options{}, err
		}
		target := filepath.Join(absWeb, "config")
		if isConfigDir(target) {
			return tui.Options{ConfigDir: target, TargetDir: target, WebRoot: absWeb}, nil
		}
		if sourceOK {
			if err := os.MkdirAll(target, 0o755); err != nil {
				return tui.Options{}, err
			}
			return tui.Options{ConfigDir: sourceDir, TargetDir: target, WebRoot: absWeb}, nil
		}
		return tui.Options{}, fmt.Errorf("网站根目录 %s 下没有可用的 config 配置，也未找到源配置目录", absWeb)
	}

	// 3. 自动查找。
	if sourceOK {
		return tui.Options{ConfigDir: sourceDir, TargetDir: sourceDir}, nil
	}
	return tui.Options{}, fmt.Errorf("未找到配置目录，请在项目根目录运行，或用 -web / -dir 指定")
}

func discoverSource() (string, bool) {
	candidates := []string{
		"public/config",
		"../public/config",
		"../../public/config",
		"config",
		"dist/config",
	}
	for _, c := range candidates {
		if isConfigDir(c) {
			abs, err := filepath.Abs(c)
			if err != nil {
				return "", false
			}
			return abs, true
		}
	}
	return "", false
}

func mustConfigDir(dir string) (string, error) {
	if !isConfigDir(dir) {
		return "", fmt.Errorf("目录 %s 中没有完整的配置文件（需要 %v）", dir, configFiles)
	}
	return filepath.Abs(dir)
}

func isConfigDir(dir string) bool {
	info, err := os.Stat(dir)
	if err != nil || !info.IsDir() {
		return false
	}
	for _, f := range configFiles {
		if _, err := os.Stat(filepath.Join(dir, f)); err != nil {
			return false
		}
	}
	return true
}

func absOrEmpty(p string) string {
	if p == "" {
		return ""
	}
	abs, err := filepath.Abs(p)
	if err != nil {
		return ""
	}
	return abs
}
