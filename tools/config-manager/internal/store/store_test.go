package store

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"the9thnet/config-manager/internal/model"
)

const srcDir = "../../../../public/config"

func copyConfig(t *testing.T, dst string) {
	t.Helper()
	for _, name := range []string{"site.json", "services.json", "activities.json", "directions.json", "articles.json"} {
		b, err := os.ReadFile(filepath.Join(srcDir, name))
		if err != nil {
			t.Fatalf("读取源配置 %s 失败: %v", name, err)
		}
		if err := os.WriteFile(filepath.Join(dst, name), b, 0o644); err != nil {
			t.Fatalf("写入临时配置 %s 失败: %v", name, err)
		}
	}
}

func decodeAny(t *testing.T, path string) any {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var v any
	if err := json.Unmarshal(b, &v); err != nil {
		t.Fatal(err)
	}
	return v
}

// TestRoundTrip 校验「结构体 -> JSON」不会丢失或改变原有配置内容。
func TestRoundTrip(t *testing.T) {
	src := copyToTemp(t)
	st := New(src)
	out := filepath.Join(src, "out")
	if err := os.MkdirAll(out, 0o755); err != nil {
		t.Fatal(err)
	}

	orig := func(name string) any { return decodeAny(t, filepath.Join(src, name)) }
	after := func(name string) any { return decodeAny(t, filepath.Join(out, name)) }

	site, err := st.LoadSite()
	if err != nil {
		t.Fatal(err)
	}
	writeJSON(t, filepath.Join(out, "site.json"), site)
	if !reflect.DeepEqual(orig("site.json"), after("site.json")) {
		t.Errorf("site.json 往返后内容不一致")
	}

	services, err := st.LoadServices()
	if err != nil {
		t.Fatal(err)
	}
	writeJSON(t, filepath.Join(out, "services.json"), services)
	if !reflect.DeepEqual(orig("services.json"), after("services.json")) {
		t.Errorf("services.json 往返后内容不一致")
	}

	activities, err := st.LoadActivities()
	if err != nil {
		t.Fatal(err)
	}
	writeJSON(t, filepath.Join(out, "activities.json"), activities)
	if !reflect.DeepEqual(orig("activities.json"), after("activities.json")) {
		t.Errorf("activities.json 往返后内容不一致")
	}

	directions, err := st.LoadDirections()
	if err != nil {
		t.Fatal(err)
	}
	writeJSON(t, filepath.Join(out, "directions.json"), directions)
	if !reflect.DeepEqual(orig("directions.json"), after("directions.json")) {
		t.Errorf("directions.json 往返后内容不一致")
	}

	articles, err := st.LoadArticlesConfig()
	if err != nil {
		t.Fatal(err)
	}
	writeJSON(t, filepath.Join(out, "articles.json"), articles)
	if !reflect.DeepEqual(orig("articles.json"), after("articles.json")) {
		t.Errorf("articles.json 往返后内容不一致")
	}
}

func copyToTemp(t *testing.T) string {
	t.Helper()
	dst := t.TempDir()
	copyConfig(t, dst)
	return dst
}

func writeJSON(t *testing.T, path string, v any) {
	t.Helper()
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, append(b, '\n'), 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestArticleRoundTrip(t *testing.T) {
	root := t.TempDir()
	articles := []model.ResourceArticle{{
		ID:       "windows-install",
		Title:    "Windows 系统安装教程",
		File:     "/articles/windows-install.md",
		Markdown: "# 标题\n\n正文\n",
	}}
	if err := SaveArticles(root, articles); err != nil {
		t.Fatal(err)
	}
	got := []model.ResourceArticle{{ID: "windows-install", File: "/articles/windows-install.md"}}
	if err := LoadArticles(root, got); err != nil {
		t.Fatal(err)
	}
	if got[0].Markdown != "# 标题\n\n正文\n" {
		t.Fatalf("正文往返不一致: %q", got[0].Markdown)
	}
}

func TestSyncResourceGroupLinks(t *testing.T) {
	groups := []model.ResourceGroup{{Title: "系统教学", Links: []string{"Windows 系统安装教程"}}}
	articles := []model.ResourceArticle{
		{Title: "Windows 系统安装教程", Category: "系统教学"},
		{Title: "ssl证书自签名快捷制作", Category: "系统教学"},
		{Title: "无关文章", Category: "不存在的分类"},
	}
	SyncResourceGroupLinks(groups, articles)
	if len(groups[0].Links) != 2 || groups[0].Links[1] != "ssl证书自签名快捷制作" {
		t.Fatalf("未把新文章同步进分组: %#v", groups[0].Links)
	}
}

func TestLoadArticlesConfigFromLegacySite(t *testing.T) {
	dir := t.TempDir()
	legacy := map[string]any{
		"resourceGroups": []map[string]any{
			{"title": "系统教学", "icon": "/images/1.png", "accent": "cyan", "links": []string{"旧标题"}},
		},
		"resourceArticles": []map[string]any{
			{"id": "legacy", "category": "系统教学", "title": "旧标题", "file": "/articles/legacy.md"},
		},
	}
	writeJSON(t, filepath.Join(dir, "site.json"), legacy)
	got, err := New(dir).LoadArticlesConfig()
	if err != nil {
		t.Fatal(err)
	}
	if len(got.Groups) != 1 || got.Groups[0].Title != "系统教学" {
		t.Fatalf("未从旧 site.json 读取分组: %#v", got.Groups)
	}
	if len(got.Articles) != 1 || got.Articles[0].ID != "legacy" {
		t.Fatalf("未从旧 site.json 读取文章: %#v", got.Articles)
	}
}

func TestResolveArticlePathRejectsTraversal(t *testing.T) {
	root := t.TempDir()
	if _, err := ResolveArticlePath(root, "/articles/../site.json"); err == nil {
		t.Fatal("期望拒绝路径穿越")
	}
	if _, err := ResolveArticlePath(root, "/config/site.json"); err == nil {
		t.Fatal("期望拒绝非 articles 目录")
	}
}
