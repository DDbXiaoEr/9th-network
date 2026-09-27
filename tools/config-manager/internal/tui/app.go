package tui

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"the9thnet/config-manager/internal/model"
	"the9thnet/config-manager/internal/settings"
	"the9thnet/config-manager/internal/store"
)

// Options 是启动参数。
type Options struct {
	// ConfigDir 是读取配置的来源目录。
	ConfigDir string
	// TargetDir 是写入配置的目标目录；为空时与 ConfigDir 相同。
	TargetDir string
	// WebRoot 是网站根目录，用于界面展示与持久化。
	WebRoot string
}

type app struct {
	ctx        *Ctx
	store      *store.Store
	configDir  string
	targetDir  string
	sourceWeb  string
	webRoot    string
	site       *model.Site
	services   *model.Services
	activities *model.Activities
	directions *model.Directions
	articles   *model.Articles
}

func websiteRoot(configDir, fallback string) string {
	if strings.TrimSpace(fallback) != "" {
		return fallback
	}
	if filepath.Base(configDir) == "config" {
		return filepath.Dir(configDir)
	}
	return configDir
}

// Run 加载全部配置并启动 TUI。
func Run(opts Options) error {
	return run(opts)
}

func run(opts Options, programOpts ...tea.ProgramOption) error {
	target := opts.TargetDir
	if target == "" {
		target = opts.ConfigDir
	}
	webRoot := websiteRoot(target, opts.WebRoot)
	a := &app{
		ctx:       &Ctx{ConfigDir: opts.ConfigDir, WebRoot: webRoot},
		configDir: opts.ConfigDir,
		targetDir: target,
		sourceWeb: websiteRoot(opts.ConfigDir, ""),
		webRoot:   webRoot,
		store:     store.New(target),
	}
	if err := a.loadAll(); err != nil {
		return err
	}

	options := append([]tea.ProgramOption{tea.WithAltScreen()}, programOpts...)
	p := tea.NewProgram(New(a.ctx, a.fileMenu()), options...)
	_, err := p.Run()
	return err
}

func (a *app) loadAll() error {
	st := store.New(a.configDir)
	site, err := st.LoadSite()
	if err != nil {
		return fmt.Errorf("加载 site.json 失败: %w", err)
	}
	services, err := st.LoadServices()
	if err != nil {
		return fmt.Errorf("加载 services.json 失败: %w", err)
	}
	activities, err := st.LoadActivities()
	if err != nil {
		return fmt.Errorf("加载 activities.json 失败: %w", err)
	}
	directions, err := st.LoadDirections()
	if err != nil {
		return fmt.Errorf("加载 directions.json 失败: %w", err)
	}
	articles, err := st.LoadArticlesConfig()
	if err != nil {
		return fmt.Errorf("加载 articles.json 失败: %w", err)
	}
	if err := store.LoadArticles(a.sourceWeb, articles.Articles); err != nil {
		return fmt.Errorf("加载文章失败: %w", err)
	}
	a.site = site
	a.services = services
	a.activities = activities
	a.directions = directions
	a.articles = articles
	return nil
}

// switchDir 把工作目录切换到网站根目录下的 config，并尽量加载其中的配置；
// 缺失的文件保留当前内存中的数据，保存时即可生成到该目录。
func (a *app) switchDir(webRoot string) error {
	root := strings.TrimSpace(webRoot)
	if root == "" {
		return errors.New("路径不能为空")
	}
	abs, err := filepath.Abs(root)
	if err != nil {
		return err
	}
	info, err := os.Stat(abs)
	if err != nil {
		return fmt.Errorf("目录不存在: %s", abs)
	}
	if !info.IsDir() {
		return fmt.Errorf("不是目录: %s", abs)
	}

	cfgDir := filepath.Join(abs, "config")
	if err := os.MkdirAll(cfgDir, 0o755); err != nil {
		return err
	}
	st := store.New(cfgDir)

	site := *a.site
	if v, err := st.LoadSite(); err == nil {
		site = *v
	}
	services := *a.services
	if v, err := st.LoadServices(); err == nil {
		services = *v
	}
	activities := *a.activities
	if v, err := st.LoadActivities(); err == nil {
		activities = *v
	}
	directions := *a.directions
	if v, err := st.LoadDirections(); err == nil {
		directions = *v
	}
	articles := *a.articles
	if v, err := st.LoadArticlesConfig(); err == nil {
		articles = *v
	}

	if err := store.LoadArticles(abs, articles.Articles); err != nil {
		return fmt.Errorf("加载文章失败: %w", err)
	}

	a.store = st
	a.configDir = cfgDir
	a.targetDir = cfgDir
	a.sourceWeb = abs
	a.webRoot = abs
	a.ctx.ConfigDir = cfgDir
	a.ctx.WebRoot = abs
	*a.site = site
	*a.services = services
	*a.activities = activities
	*a.directions = directions
	*a.articles = articles

	_ = settings.Save(settings.Settings{WebRoot: abs})
	return nil
}

func (a *app) fileMenu() screen {
	a.ctx.File = ""
	a.ctx.SaveFn = nil
	a.ctx.ReloadFn = nil

	return newMenuScreen(a.ctx, "选择配置文件", []menuItem{
		{label: "site.json", desc: "品牌 / 导航 / 页脚 / 首屏 / 社团 / 加入", open: a.openSite},
		{label: "services.json", desc: "社团公共服务", open: a.openServices},
		{label: "activities.json", desc: "活动通知与回顾", open: a.openActivities},
		{label: "directions.json", desc: "兴趣方向 / CTF 战队", open: a.openDirections},
		{label: "articles.json", desc: "学习资源分组与文章", open: a.openArticles},
		{label: "⚙ 设置网站根目录", desc: "当前写入 " + a.webRoot, open: a.openSettings},
	})
}

func (a *app) openSettings() screen {
	return newPathScreen(a.ctx,
		"设置网站根目录",
		"配置写入 <网站根目录>/config/*.json，文章写入 <网站根目录>/articles/*.md；目录不存在会报错，子目录会自动创建。",
		a.webRoot,
		a.switchDir,
		a.fileMenu,
	)
}

func (a *app) openSite() screen {
	a.ctx.File = "site.json"
	a.ctx.SaveFn = func() error { return a.store.SaveSite(a.site) }
	a.ctx.ReloadFn = func() error {
		v, err := a.store.LoadSite()
		if err != nil {
			return err
		}
		*a.site = *v
		return nil
	}
	return siteRoot(a.ctx, a.site)
}

func (a *app) openServices() screen {
	a.ctx.File = "services.json"
	a.ctx.SaveFn = func() error { return a.store.SaveServices(a.services) }
	a.ctx.ReloadFn = func() error {
		v, err := a.store.LoadServices()
		if err != nil {
			return err
		}
		*a.services = *v
		return nil
	}
	return servicesRoot(a.ctx, a.services)
}

func (a *app) openActivities() screen {
	a.ctx.File = "activities.json"
	a.ctx.SaveFn = func() error { return a.store.SaveActivities(a.activities) }
	a.ctx.ReloadFn = func() error {
		v, err := a.store.LoadActivities()
		if err != nil {
			return err
		}
		*a.activities = *v
		return nil
	}
	return activitiesRoot(a.ctx, a.activities)
}

func (a *app) openDirections() screen {
	a.ctx.File = "directions.json"
	a.ctx.SaveFn = func() error { return a.store.SaveDirections(a.directions) }
	a.ctx.ReloadFn = func() error {
		v, err := a.store.LoadDirections()
		if err != nil {
			return err
		}
		*a.directions = *v
		return nil
	}
	return directionsRoot(a.ctx, a.directions)
}

func (a *app) openArticles() screen {
	a.ctx.File = "articles.json"
	a.ctx.SaveFn = func() error {
		store.NormalizeArticleFiles(a.articles.Articles)
		store.SyncResourceGroupLinks(a.articles.Groups, a.articles.Articles)
		if err := store.SaveArticles(a.webRoot, a.articles.Articles); err != nil {
			return err
		}
		return a.store.SaveArticlesConfig(a.articles)
	}
	a.ctx.ReloadFn = func() error {
		v, err := a.store.LoadArticlesConfig()
		if err != nil {
			return err
		}
		if err := store.LoadArticles(a.webRoot, v.Articles); err != nil {
			return err
		}
		*a.articles = *v
		return nil
	}
	return articlesRoot(a.ctx, a.articles)
}

// objectList 是一个通用的「结构体切片」列表界面：支持进入编辑、新增、删除与上下移。
func objectList[T any](ctx *Ctx, heading, empty string, list *[]T,
	label func(i int, v *T) (string, string),
	form func(ctx *Ctx, v *T) screen,
	blank func() T,
) *arrayListScreen {
	return newArrayList(ctx, heading, empty,
		func() int { return len(*list) },
		func(i int) (string, string) { return label(i, &(*list)[i]) },
		func(i int) screen { return form(ctx, &(*list)[i]) },
		func() { *list = append(*list, blank()) },
		func(i int) { *list = append((*list)[:i], (*list)[i+1:]...) },
		func(i, delta int) bool {
			j := i + delta
			if j < 0 || j >= len(*list) {
				return false
			}
			(*list)[i], (*list)[j] = (*list)[j], (*list)[i]
			return true
		},
	)
}

func itemCount(n int) string { return fmt.Sprintf("%d 项", n) }
