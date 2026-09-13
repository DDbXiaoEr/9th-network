package tui

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// Ctx 在各级界面之间共享：窗口尺寸、当前文件、保存/重载回调与状态提示。
type Ctx struct {
	ConfigDir string
	WebRoot   string
	File      string
	W, H      int

	SaveFn   func() error
	ReloadFn func() error

	status   string
	statusAt time.Time
}

func (c *Ctx) SetStatus(s string) {
	c.status = s
	c.statusAt = time.Now()
}

func (c *Ctx) Status() (string, bool) {
	if c.status == "" || time.Since(c.statusAt) > 4*time.Second {
		return "", false
	}
	return c.status, true
}

// screen 是压栈式导航中的一个界面。
type screen interface {
	Init() tea.Cmd
	Update(tea.Msg) (screen, tea.Cmd)
	View() string
	Title() string
}

// committer 在保存前把界面内未提交的编辑写回模型。
type committer interface{ commit() }

type pushMsg struct{ s screen }
type popMsg struct{}
type resetMsg struct{ initial screen }
type savedMsg struct{ err error }
type reloadedMsg struct{ err error }

func push(s screen) tea.Cmd { return func() tea.Msg { return pushMsg{s} } }
func pop() tea.Cmd          { return func() tea.Msg { return popMsg{} } }

type root struct {
	ctx   *Ctx
	stack []screen
}

// New 创建根模型，initial 为初始界面（通常是配置文件菜单）。
func New(ctx *Ctx, initial screen) tea.Model {
	return &root{ctx: ctx, stack: []screen{initial}}
}

func (r *root) Init() tea.Cmd { return r.top().Init() }

func (r *root) top() screen { return r.stack[len(r.stack)-1] }

func (r *root) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m := msg.(type) {
	case tea.WindowSizeMsg:
		r.ctx.W, r.ctx.H = m.Width, m.Height
		return r, nil
	case pushMsg:
		r.stack = append(r.stack, m.s)
		return r, m.s.Init()
	case popMsg:
		if len(r.stack) > 1 {
			r.stack = r.stack[:len(r.stack)-1]
		}
		return r, nil
	case resetMsg:
		if m.initial != nil {
			r.stack = []screen{m.initial}
			return r, m.initial.Init()
		}
		return r, nil
	case savedMsg:
		if m.err != nil {
			r.ctx.SetStatus("✗ 保存失败: " + m.err.Error())
		} else {
			r.ctx.SetStatus("✓ 已保存 " + r.ctx.File)
		}
		return r, nil
	case reloadedMsg:
		if m.err != nil {
			r.ctx.SetStatus("✗ 重新加载失败: " + m.err.Error())
		} else {
			r.ctx.SetStatus("↻ 已从磁盘重新加载 " + r.ctx.File)
		}
		return r, nil
	case tea.KeyMsg:
		switch m.String() {
		case "ctrl+c":
			return r, tea.Quit
		case "ctrl+s":
			if c, ok := r.top().(committer); ok {
				c.commit()
			}
			if r.ctx.SaveFn != nil {
				return r, func() tea.Msg { return savedMsg{err: r.ctx.SaveFn()} }
			}
		case "ctrl+r":
			if r.ctx.ReloadFn != nil {
				return r, func() tea.Msg { return reloadedMsg{err: r.ctx.ReloadFn()} }
			}
		}
	}

	next, cmd := r.top().Update(msg)
	r.stack[len(r.stack)-1] = next
	return r, cmd
}

func (r *root) View() string {
	var b strings.Builder
	b.WriteString("\n")
	b.WriteString(titleStyle.Render(" 第九网络组 · 配置管理器 "))
	b.WriteString("\n")
	b.WriteString(dimStyle.Render("  " + r.top().Title()))
	b.WriteString("\n\n")
	b.WriteString(r.top().View())
	b.WriteString("\n")
	b.WriteString(helpStyle.Render(" ctrl+s 保存   ctrl+r 重新加载   esc 返回   ctrl+c 退出   "))
	b.WriteString("\n")
	if s, ok := r.ctx.Status(); ok {
		if strings.HasPrefix(s, "✗") {
			b.WriteString(errStyle.Render(" " + s))
		} else {
			b.WriteString(okStyle.Render(" " + s))
		}
		b.WriteString("\n")
	}
	return b.String()
}
