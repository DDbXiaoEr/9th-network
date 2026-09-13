package store

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

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
