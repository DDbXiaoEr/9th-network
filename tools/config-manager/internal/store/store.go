package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"the9thnet/config-manager/internal/model"
)

// Store 负责从磁盘读写各配置文件，写回时保持两空格缩进与 JSON 字段顺序。
type Store struct {
	Dir string
}

func New(dir string) *Store {
	return &Store{Dir: dir}
}

func (s *Store) path(name string) string {
	return filepath.Join(s.Dir, name)
}

func (s *Store) load(name string, v any) error {
	b, err := os.ReadFile(s.path(name))
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, v); err != nil {
		return fmt.Errorf("解析 %s 失败: %w", name, err)
	}
	return nil
}

func (s *Store) save(name string, v any) error {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	b = append(b, '\n')
	return writeFileAtomic(s.path(name), b)
}

func writeFileAtomic(path string, data []byte) error {
	dir := filepath.Dir(path)
	tmp, err := os.CreateTemp(dir, ".config-*.tmp")
	if err != nil {
		return err
	}
	tmpName := tmp.Name()
	defer os.Remove(tmpName)

	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	if err := os.Chmod(tmpName, 0o644); err != nil {
		return err
	}
	return os.Rename(tmpName, path)
}

func (s *Store) LoadSite() (*model.Site, error) {
	v := &model.Site{}
	if err := s.load("site.json", v); err != nil {
		return nil, err
	}
	return v, nil
}

func (s *Store) SaveSite(v *model.Site) error { return s.save("site.json", v) }

func (s *Store) LoadServices() (*model.Services, error) {
	v := &model.Services{}
	if err := s.load("services.json", v); err != nil {
		return nil, err
	}
	return v, nil
}

func (s *Store) SaveServices(v *model.Services) error { return s.save("services.json", v) }

func (s *Store) LoadActivities() (*model.Activities, error) {
	v := &model.Activities{}
	if err := s.load("activities.json", v); err != nil {
		return nil, err
	}
	return v, nil
}

func (s *Store) SaveActivities(v *model.Activities) error { return s.save("activities.json", v) }

func (s *Store) LoadDirections() (*model.Directions, error) {
	v := &model.Directions{}
	if err := s.load("directions.json", v); err != nil {
		return nil, err
	}
	return v, nil
}

func (s *Store) SaveDirections(v *model.Directions) error { return s.save("directions.json", v) }

func (s *Store) LoadArticlesConfig() (*model.Articles, error) {
	v := &model.Articles{}
	err := s.load("articles.json", v)
	if err == nil {
		return v, nil
	}
	if !os.IsNotExist(err) {
		return nil, err
	}
	legacy, legacyErr := s.loadLegacyArticlesFromSite()
	if legacyErr != nil || legacy == nil {
		return nil, err
	}
	return legacy, nil
}

type legacySiteArticles struct {
	ResourceGroups   []model.ResourceGroup   `json:"resourceGroups"`
	ResourceArticles []model.ResourceArticle `json:"resourceArticles"`
}

func (s *Store) loadLegacyArticlesFromSite() (*model.Articles, error) {
	b, err := os.ReadFile(s.path("site.json"))
	if err != nil {
		return nil, err
	}
	legacy := &legacySiteArticles{}
	if err := json.Unmarshal(b, legacy); err != nil {
		return nil, err
	}
	if len(legacy.ResourceGroups) == 0 && len(legacy.ResourceArticles) == 0 {
		return nil, nil
	}
	return &model.Articles{Groups: legacy.ResourceGroups, Articles: legacy.ResourceArticles}, nil
}

func (s *Store) SaveArticlesConfig(v *model.Articles) error { return s.save("articles.json", v) }

func ArticleFile(a model.ResourceArticle) string {
	file := strings.TrimSpace(a.File)
	if file != "" {
		if !strings.HasPrefix(file, "/") {
			file = "/" + file
		}
		return file
	}
	if id := strings.TrimSpace(a.ID); id != "" {
		return "/articles/" + id + ".md"
	}
	return ""
}

func NormalizeArticleFiles(articles []model.ResourceArticle) {
	for i := range articles {
		if path := ArticleFile(articles[i]); path != "" {
			articles[i].File = path
		}
	}
}

func ResolveArticlePath(webRoot, file string) (string, error) {
	if strings.TrimSpace(webRoot) == "" {
		return "", errors.New("未设置网站根目录，无法读写文章")
	}
	rel := strings.ReplaceAll(strings.TrimSpace(file), "\\", "/")
	rel = strings.TrimPrefix(rel, "/")
	if rel == "" {
		return "", errors.New("文章路径为空")
	}
	if strings.Contains(rel, "..") {
		return "", fmt.Errorf("非法文章路径: %s", file)
	}
	cleaned := path.Clean(rel)
	if !strings.HasPrefix(cleaned, "articles/") {
		return "", fmt.Errorf("文章必须放在 /articles/ 目录: %s", file)
	}
	if path.Ext(cleaned) != ".md" {
		return "", fmt.Errorf("文章必须是 .md 文件: %s", file)
	}

	absRoot, err := filepath.Abs(webRoot)
	if err != nil {
		return "", err
	}
	full, err := filepath.Abs(filepath.Join(absRoot, filepath.FromSlash(cleaned)))
	if err != nil {
		return "", err
	}
	relToRoot, err := filepath.Rel(absRoot, full)
	if err != nil || strings.HasPrefix(relToRoot, "..") {
		return "", fmt.Errorf("文章路径超出网站根目录: %s", file)
	}
	return full, nil
}

func LoadArticles(webRoot string, articles []model.ResourceArticle) error {
	for i := range articles {
		file := ArticleFile(articles[i])
		if file == "" {
			continue
		}
		full, err := ResolveArticlePath(webRoot, file)
		if err != nil {
			return err
		}
		b, err := os.ReadFile(full)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return fmt.Errorf("读取 %s 失败: %w", file, err)
		}
		articles[i].Markdown = string(b)
		articles[i].File = file
	}
	return nil
}

func SyncResourceGroupLinks(groups []model.ResourceGroup, articles []model.ResourceArticle) {
	seenByGroup := make([]map[string]struct{}, len(groups))
	indexByTitle := make(map[string]int, len(groups))
	for i := range groups {
		seenByGroup[i] = make(map[string]struct{}, len(groups[i].Links))
		indexByTitle[groups[i].Title] = i
		for _, link := range groups[i].Links {
			seenByGroup[i][link] = struct{}{}
		}
	}
	for _, article := range articles {
		title := strings.TrimSpace(article.Title)
		if title == "" {
			continue
		}
		idx, ok := indexByTitle[strings.TrimSpace(article.Category)]
		if !ok {
			continue
		}
		if _, exists := seenByGroup[idx][title]; exists {
			continue
		}
		groups[idx].Links = append(groups[idx].Links, title)
		seenByGroup[idx][title] = struct{}{}
	}
}

func SaveArticles(webRoot string, articles []model.ResourceArticle) error {
	NormalizeArticleFiles(articles)
	for i := range articles {
		file := articles[i].File
		if file == "" {
			continue
		}
		full, err := ResolveArticlePath(webRoot, file)
		if err != nil {
			return err
		}
		if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
			return err
		}
		data := []byte(articles[i].Markdown)
		if len(data) > 0 && data[len(data)-1] != '\n' {
			data = append(data, '\n')
		}
		if err := writeFileAtomic(full, data); err != nil {
			return fmt.Errorf("写入 %s 失败: %w", file, err)
		}
	}
	return nil
}
