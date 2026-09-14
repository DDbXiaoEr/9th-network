package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func pad(s string, w int) string {
	return lipgloss.NewStyle().Width(w).Render(s)
}

func newInput(ctx *Ctx) textinput.Model {
	ti := textinput.New()
	ti.Prompt = "› "
	ti.CharLimit = 0
	ti.Width = inputWidth(ctx)
	return ti
}

func inputWidth(ctx *Ctx) int {
	if ctx.W > 24 {
		return ctx.W - 16
	}
	return 40
}

// ---------------------------------------------------------------- menuScreen

type menuItem struct {
	label string
	desc  string
	open  func() screen
}

type menuScreen struct {
	ctx     *Ctx
	heading string
	items   []menuItem
	cursor  int
}

func newMenuScreen(ctx *Ctx, heading string, items []menuItem) *menuScreen {
	return &menuScreen{ctx: ctx, heading: heading, items: items}
}

func (m *menuScreen) Init() tea.Cmd { return nil }
func (m *menuScreen) Title() string { return m.heading }
func (m *menuScreen) commit()       {}

func (m *menuScreen) Update(msg tea.Msg) (screen, tea.Cmd) {
	key, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}
	switch key.String() {
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < len(m.items)-1 {
			m.cursor++
		}
	case "enter":
		if m.cursor < len(m.items) && m.items[m.cursor].open != nil {
			return m, push(m.items[m.cursor].open())
		}
	case "esc":
		return m, pop()
	}
	return m, nil
}

func (m *menuScreen) View() string {
	width := 0
	for _, it := range m.items {
		if w := lipgloss.Width(it.label); w > width {
			width = w
		}
	}
	var b strings.Builder
	for i, it := range m.items {
		if i == m.cursor {
			b.WriteString(cursorStyle.Render("❯ ") + cursorStyle.Render(pad(it.label, width)))
		} else {
			b.WriteString("  " + labelStyle.Render(pad(it.label, width)))
		}
		if it.desc != "" {
			b.WriteString("  " + dimStyle.Render(it.desc))
		}
		b.WriteString("\n")
	}
	b.WriteString("\n" + helpStyle.Render("  ↑/↓ 选择    enter 进入"))
	return b.String()
}

// ---------------------------------------------------------------- arrayListScreen

type arrayListScreen struct {
	ctx     *Ctx
	heading string
	empty   string
	count   func() int
	labelAt func(i int) (string, string)
	editAt  func(i int) screen
	add     func()
	del     func(i int)
	move    func(i int, delta int) bool
	cursor  int
}

func newArrayList(ctx *Ctx, heading, empty string,
	count func() int,
	labelAt func(i int) (string, string),
	editAt func(i int) screen,
	add func(),
	del func(i int),
	move func(i int, delta int) bool,
) *arrayListScreen {
	return &arrayListScreen{ctx: ctx, heading: heading, empty: empty, count: count,
		labelAt: labelAt, editAt: editAt, add: add, del: del, move: move}
}

func (l *arrayListScreen) Init() tea.Cmd { return nil }
func (l *arrayListScreen) Title() string { return l.heading }
func (l *arrayListScreen) commit()       {}

func (l *arrayListScreen) clamp() {
	n := l.count()
	if l.cursor >= n {
		l.cursor = n - 1
	}
	if l.cursor < 0 {
		l.cursor = 0
	}
}

func (l *arrayListScreen) Update(msg tea.Msg) (screen, tea.Cmd) {
	key, ok := msg.(tea.KeyMsg)
	if !ok {
		return l, nil
	}
	switch key.String() {
	case "up", "k":
		if l.cursor > 0 {
			l.cursor--
		}
	case "down", "j":
		if l.cursor < l.count()-1 {
			l.cursor++
		}
	case "enter":
		if l.editAt != nil && l.count() > 0 {
			return l, push(l.editAt(l.cursor))
		}
	case "a":
		if l.add != nil {
			l.add()
			l.cursor = l.count() - 1
		}
	case "d":
		if l.del != nil && l.count() > 0 {
			l.del(l.cursor)
			l.clamp()
		}
	case "J":
		if l.move != nil && l.move(l.cursor, 1) {
			l.cursor++
		}
	case "K":
		if l.move != nil && l.move(l.cursor, -1) {
			l.cursor--
		}
	case "esc":
		return l, pop()
	}
	return l, nil
}

func (l *arrayListScreen) View() string {
	n := l.count()
	var b strings.Builder
	if n == 0 {
		msg := l.empty
		if msg == "" {
			msg = "暂无条目，按 a 新增。"
		}
		b.WriteString(dimStyle.Render("  " + msg))
		b.WriteString("\n")
	} else {
		for i := 0; i < n; i++ {
			title, value := l.labelAt(i)
			marker := "  "
			t := labelStyle
			if i == l.cursor {
				marker = cursorStyle.Render("❯ ")
				t = cursorStyle
			}
			b.WriteString(fmt.Sprintf("%s%s %s", marker, dimStyle.Render(fmt.Sprintf("%2d.", i+1)), t.Render(title)))
			if value != "" {
				b.WriteString("  " + dimStyle.Render(value))
			}
			b.WriteString("\n")
		}
	}

	var keys []string
	if l.editAt != nil {
		keys = append(keys, "enter 编辑")
	}
	if l.add != nil {
		keys = append(keys, "a 新增")
	}
	if l.del != nil {
		keys = append(keys, "d 删除")
	}
	if l.move != nil {
		keys = append(keys, "J/K 上移/下移")
	}
	b.WriteString("\n" + helpStyle.Render("  "+strings.Join(keys, "    ")))
	return b.String()
}

// ---------------------------------------------------------------- formScreen

type rowKind int

const (
	rowString rowKind = iota
	rowBool
	rowOpen
	rowChoice
)

// choiceOption 是 rowChoice 的可选项：value 写入配置，label 用于界面展示。
type choiceOption struct {
	value string
	label string
}

func opt(value, label string) choiceOption {
	return choiceOption{value: value, label: label}
}

type formRow struct {
	label   string
	kind    rowKind
	ptr     *string
	bptr    *bool
	summary func() string
	open    func() screen
	options []choiceOption
	long    bool
}

func rowText(label string, p *string) formRow {
	return formRow{label: label, kind: rowString, ptr: p, long: len(*p) > 48}
}

func rowLong(label string, p *string) formRow {
	return formRow{label: label, kind: rowString, ptr: p, long: true}
}

func boolRow(label string, p *bool) formRow {
	return formRow{label: label, kind: rowBool, bptr: p}
}

func choiceRow(label string, p *string, options ...choiceOption) formRow {
	return formRow{label: label, kind: rowChoice, ptr: p, options: options}
}

func openRow(label string, summary func() string, open func() screen) formRow {
	return formRow{label: label, kind: rowOpen, summary: summary, open: open}
}

// choiceLabel 返回当前值对应的展示文案；值不在选项中时原样返回。
func (r formRow) choiceLabel() string {
	if *r.ptr == "" {
		return ""
	}
	for _, o := range r.options {
		if o.value == *r.ptr {
			return o.label
		}
	}
	return *r.ptr
}

// nextChoice 在当前选项之间循环切换；当前值不在选项中时从第一项开始。
func (r formRow) nextChoice() {
	if len(r.options) == 0 {
		return
	}
	idx := -1
	for i, o := range r.options {
		if o.value == *r.ptr {
			idx = i
			break
		}
	}
	*r.ptr = r.options[(idx+1)%len(r.options)].value
}

type formScreen struct {
	ctx     *Ctx
	heading string
	rows    []formRow
	cursor  int
	editing bool
	input   textinput.Model
}

func newFormScreen(ctx *Ctx, heading string, rows []formRow) *formScreen {
	return &formScreen{ctx: ctx, heading: heading, rows: rows, input: newInput(ctx)}
}

func (f *formScreen) Init() tea.Cmd { return nil }
func (f *formScreen) Title() string { return f.heading }

func (f *formScreen) commit() {
	if f.editing && f.cursor < len(f.rows) && f.rows[f.cursor].ptr != nil {
		*f.rows[f.cursor].ptr = f.input.Value()
		f.editing = false
	}
}

func (f *formScreen) value(i int) string {
	r := f.rows[i]
	switch r.kind {
	case rowString:
		if *r.ptr == "" {
			return dimStyle.Render("（空）")
		}
		return labelStyle.Render(*r.ptr)
	case rowBool:
		if *r.bptr {
			return badgeStyle.Render("[x] 是")
		}
		return dimStyle.Render("[ ] 否")
	case rowChoice:
		if label := r.choiceLabel(); label != "" {
			return badgeStyle.Render(label)
		}
		return dimStyle.Render("（未设置）")
	case rowOpen:
		if r.summary != nil {
			return dimStyle.Render(r.summary())
		}
	}
	return ""
}

func (f *formScreen) Update(msg tea.Msg) (screen, tea.Cmd) {
	if f.editing {
		if key, ok := msg.(tea.KeyMsg); ok {
			switch key.String() {
			case "esc":
				f.editing = false
				return f, nil
			case "enter":
				*f.rows[f.cursor].ptr = f.input.Value()
				f.editing = false
				return f, nil
			}
		}
		var cmd tea.Cmd
		f.input, cmd = f.input.Update(msg)
		return f, cmd
	}

	key, ok := msg.(tea.KeyMsg)
	if !ok {
		return f, nil
	}
	switch key.String() {
	case "up", "k":
		if f.cursor > 0 {
			f.cursor--
		}
	case "down", "j":
		if f.cursor < len(f.rows)-1 {
			f.cursor++
		}
	case "enter":
		r := f.rows[f.cursor]
		switch r.kind {
		case rowBool:
			*r.bptr = !*r.bptr
		case rowChoice:
			r.nextChoice()
		case rowString:
			if r.long {
				return f, push(newTextScreen(f.ctx, f.heading+" · "+r.label, r.ptr))
			}
			f.editing = true
			f.input.SetValue(*r.ptr)
			f.input.CursorEnd()
			f.input.Focus()
		case rowOpen:
			if r.open != nil {
				return f, push(r.open())
			}
		}
	case "ctrl+e":
		r := f.rows[f.cursor]
		if r.kind == rowString {
			return f, push(newTextScreen(f.ctx, f.heading+" · "+r.label, r.ptr))
		}
	case "esc":
		return f, pop()
	}
	return f, nil
}

func (f *formScreen) View() string {
	width := 0
	for _, r := range f.rows {
		if w := lipgloss.Width(r.label); w > width {
			width = w
		}
	}
	var b strings.Builder
	for i, r := range f.rows {
		marker := "  "
		lab := labelStyle
		if i == f.cursor {
			marker = cursorStyle.Render("❯ ")
			lab = cursorStyle
		}
		if f.editing && i == f.cursor {
			b.WriteString(marker + lab.Render(pad(r.label, width)) + "  " + f.input.View() + "\n")
			continue
		}
		b.WriteString(marker + lab.Render(pad(r.label, width)) + "  " + f.value(i) + "\n")
	}
	b.WriteString("\n" + helpStyle.Render("  enter 编辑/切换    ctrl+e 多行编辑    esc 返回"))
	return b.String()
}

// ---------------------------------------------------------------- stringListScreen

type stringListScreen struct {
	ctx     *Ctx
	heading string
	list    *[]string
	cursor  int
	editing bool
	input   textinput.Model
}

func newStringList(ctx *Ctx, heading string, list *[]string) *stringListScreen {
	return &stringListScreen{ctx: ctx, heading: heading, list: list, input: newInput(ctx)}
}

func (s *stringListScreen) Init() tea.Cmd { return nil }
func (s *stringListScreen) Title() string { return s.heading }

func (s *stringListScreen) commit() {
	if s.editing && s.cursor < len(*s.list) {
		(*s.list)[s.cursor] = s.input.Value()
		s.editing = false
	}
}

func (s *stringListScreen) clamp() {
	if s.cursor >= len(*s.list) {
		s.cursor = len(*s.list) - 1
	}
	if s.cursor < 0 {
		s.cursor = 0
	}
}

func (s *stringListScreen) Update(msg tea.Msg) (screen, tea.Cmd) {
	if s.editing {
		if key, ok := msg.(tea.KeyMsg); ok {
			switch key.String() {
			case "esc":
				s.editing = false
				return s, nil
			case "enter":
				(*s.list)[s.cursor] = s.input.Value()
				s.editing = false
				return s, nil
			}
		}
		var cmd tea.Cmd
		s.input, cmd = s.input.Update(msg)
		return s, cmd
	}

	key, ok := msg.(tea.KeyMsg)
	if !ok {
		return s, nil
	}
	switch key.String() {
	case "up", "k":
		if s.cursor > 0 {
			s.cursor--
		}
	case "down", "j":
		if s.cursor < len(*s.list)-1 {
			s.cursor++
		}
	case "enter":
		if len(*s.list) > 0 {
			s.editing = true
			s.input.SetValue((*s.list)[s.cursor])
			s.input.CursorEnd()
			s.input.Focus()
		}
	case "ctrl+e":
		if len(*s.list) > 0 {
			return s, push(newTextScreen(s.ctx, s.heading, &(*s.list)[s.cursor]))
		}
	case "a":
		*s.list = append(*s.list, "")
		s.cursor = len(*s.list) - 1
	case "d":
		if len(*s.list) > 0 {
			*s.list = append((*s.list)[:s.cursor], (*s.list)[s.cursor+1:]...)
			s.clamp()
		}
	case "J":
		if s.cursor < len(*s.list)-1 {
			(*s.list)[s.cursor], (*s.list)[s.cursor+1] = (*s.list)[s.cursor+1], (*s.list)[s.cursor]
			s.cursor++
		}
	case "K":
		if s.cursor > 0 {
			(*s.list)[s.cursor], (*s.list)[s.cursor-1] = (*s.list)[s.cursor-1], (*s.list)[s.cursor]
			s.cursor--
		}
	case "esc":
		return s, pop()
	}
	return s, nil
}

func (s *stringListScreen) View() string {
	var b strings.Builder
	if len(*s.list) == 0 {
		b.WriteString(dimStyle.Render("  暂无内容，按 a 新增。") + "\n")
	}
	for i, v := range *s.list {
		marker := "  "
		text := labelStyle
		if i == s.cursor {
			marker = cursorStyle.Render("❯ ")
			text = cursorStyle
		}
		if s.editing && i == s.cursor {
			b.WriteString(marker + s.input.View() + "\n")
			continue
		}
		display := v
		if display == "" {
			display = dimStyle.Render("（空）")
		} else {
			display = text.Render(display)
		}
		b.WriteString(fmt.Sprintf("%s%s %s\n", marker, dimStyle.Render(fmt.Sprintf("%2d.", i+1)), display))
	}
	b.WriteString("\n" + helpStyle.Render("  enter 编辑    ctrl+e 多行    a 新增    d 删除    J/K 上移/下移    esc 返回"))
	return b.String()
}

// ---------------------------------------------------------------- textScreen

type textScreen struct {
	ctx     *Ctx
	heading string
	ptr     *string
	ta      textarea.Model
}

func newTextScreen(ctx *Ctx, heading string, ptr *string) *textScreen {
	ta := textarea.New()
	ta.ShowLineNumbers = false
	ta.SetValue(*ptr)
	ta.Focus()
	w, h := 60, 6
	if ctx.W > 24 {
		w = ctx.W - 8
	}
	if ctx.H > 14 {
		h = ctx.H - 12
	}
	ta.SetWidth(w)
	ta.SetHeight(h)
	return &textScreen{ctx: ctx, heading: heading, ptr: ptr, ta: ta}
}

func (t *textScreen) Init() tea.Cmd { return textarea.Blink }
func (t *textScreen) Title() string { return t.heading }

func (t *textScreen) commit() { *t.ptr = t.ta.Value() }

func (t *textScreen) Update(msg tea.Msg) (screen, tea.Cmd) {
	if key, ok := msg.(tea.KeyMsg); ok && key.String() == "esc" {
		*t.ptr = t.ta.Value()
		return t, pop()
	}
	var cmd tea.Cmd
	t.ta, cmd = t.ta.Update(msg)
	return t, cmd
}

func (t *textScreen) View() string {
	var b strings.Builder
	b.WriteString(boxStyle.Render(t.ta.View()))
	b.WriteString("\n" + helpStyle.Render("  esc 完成    ctrl+s 保存文件"))
	return b.String()
}

// ---------------------------------------------------------------- pathScreen

type pathScreen struct {
	ctx      *Ctx
	heading  string
	hint     string
	input    textinput.Model
	onSubmit func(string) error
	next     func() screen
}

func newPathScreen(ctx *Ctx, heading, hint, initial string,
	onSubmit func(string) error,
	next func() screen,
) *pathScreen {
	ti := newInput(ctx)
	ti.SetValue(initial)
	ti.CursorEnd()
	ti.Focus()
	return &pathScreen{ctx: ctx, heading: heading, hint: hint, input: ti, onSubmit: onSubmit, next: next}
}

func (p *pathScreen) Init() tea.Cmd { return textinput.Blink }
func (p *pathScreen) Title() string { return p.heading }
func (p *pathScreen) commit()       {}

func (p *pathScreen) Update(msg tea.Msg) (screen, tea.Cmd) {
	if key, ok := msg.(tea.KeyMsg); ok {
		switch key.String() {
		case "esc":
			return p, pop()
		case "enter":
			value := strings.TrimSpace(p.input.Value())
			if p.onSubmit != nil {
				if err := p.onSubmit(value); err != nil {
					p.ctx.SetStatus("✗ " + err.Error())
					return p, nil
				}
			}
			p.ctx.SetStatus("✓ 已切换配置目录")
			next := p.next
			return p, func() tea.Msg { return resetMsg{initial: next()} }
		}
	}
	var cmd tea.Cmd
	p.input, cmd = p.input.Update(msg)
	return p, cmd
}

func (p *pathScreen) View() string {
	var b strings.Builder
	b.WriteString("  " + dimStyle.Render(p.hint) + "\n\n")
	b.WriteString("  " + labelStyle.Render("路径") + "  " + p.input.View() + "\n\n")
	b.WriteString(helpStyle.Render("  enter 应用    esc 取消"))
	return b.String()
}
