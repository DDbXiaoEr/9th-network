package store

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

const srcDir = "../../../../public/config"

func copyConfig(t *testing.T, dst string) {
	t.Helper()
	for _, name := range []string{"site.json", "services.json", "activities.json", "directions.json"} {
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
