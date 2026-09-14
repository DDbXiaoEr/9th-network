package tui

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func runWithKeys(t *testing.T, m tea.Model, keys string) tea.Model {
	t.Helper()
	var out bytes.Buffer
	p := tea.NewProgram(m, tea.WithInput(strings.NewReader(keys)), tea.WithOutput(&out))
	final, err := p.Run()
	if err != nil {
		t.Fatalf("程序运行失败: %v", err)
	}
	return final
}

func TestMenuPushAndQuit(t *testing.T) {
	ctx := &Ctx{}
	name := "第九网络组"
	tab := "社团简介"
	form := newFormScreen(ctx, "品牌信息", []formRow{
		rowText("名称", &name),
		rowText("标签", &tab),
	})
	menu := newMenuScreen(ctx, "选择", []menuItem{
		{label: "品牌", desc: "brand", open: func() screen { return form }},
	})

	m := New(ctx, menu)

	m2, cmd := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if cmd == nil {
		t.Fatal("enter 未返回压栈命令")
	}
	m3, _ := m2.Update(cmd())
	if got := len(m3.(*root).stack); got != 2 {
		t.Fatalf("压栈后栈深应为 2，实际 %d", got)
	}

	m4, cmd := m3.Update(tea.KeyMsg{Type: tea.KeyEsc})
	if cmd == nil {
		t.Fatal("esc 未返回出栈命令")
	}
	m5, _ := m4.Update(cmd())
	if got := len(m5.(*root).stack); got != 1 {
		t.Fatalf("出栈后栈深应为 1，实际 %d", got)
	}
}

func TestTextListAdd(t *testing.T) {
	ctx := &Ctx{}
	list := []string{"a", "b"}
	final := runWithKeys(t, New(ctx, newStringList(ctx, "列表", &list)), "a\x03")
	_ = final
	if len(list) != 3 {
		t.Fatalf("期望新增一行后为 3 行，实际 %d", len(list))
	}
}

func TestChoiceRowCycle(t *testing.T) {
	value := "modal"
	row := choiceRow("方式", &value, opt("modal", "弹窗提示"), opt("link", "链接跳转"))
	if got := row.choiceLabel(); got != "弹窗提示" {
		t.Fatalf("期望展示「弹窗提示」，实际 %q", got)
	}
	row.nextChoice()
	if value != "link" {
		t.Fatalf("切换后期望 link，实际 %q", value)
	}
	row.nextChoice()
	if value != "modal" {
		t.Fatalf("循环切换后期望 modal，实际 %q", value)
	}
	value = "unknown"
	row.nextChoice()
	if value != "modal" {
		t.Fatalf("未知值切换后期望回到首项 modal，实际 %q", value)
	}
}

func TestRunWithRealConfig(t *testing.T) {
	srcDir := "../../../../public/config"
	dst := t.TempDir()
	for _, name := range []string{"site.json", "services.json", "activities.json", "directions.json"} {
		b, err := os.ReadFile(filepath.Join(srcDir, name))
		if err != nil {
			t.Fatalf("读取 %s 失败: %v", name, err)
		}
		if err := os.WriteFile(filepath.Join(dst, name), b, 0o644); err != nil {
			t.Fatalf("写入 %s 失败: %v", name, err)
		}
	}

	var out bytes.Buffer
	err := run(Options{ConfigDir: dst},
		tea.WithInput(strings.NewReader("\r\x03")),
		tea.WithOutput(&out),
	)
	if err != nil {
		t.Fatalf("run 失败: %v", err)
	}
	if !strings.Contains(out.String(), "品牌信息") {
		t.Fatalf("输出中未渲染 site.json 界面")
	}
}
