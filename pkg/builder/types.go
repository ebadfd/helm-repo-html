package builder

import "time"

type MaintainerEntry struct {
	Email string `yaml:"email"`
	Name  string `yaml:"name"`
	Url   string `yaml:"url"`
}

type Dependencies struct {
	Condition  *string `yaml:"condition"`
	Name       *string `yaml:"name"`
	Repository *string `yaml:"repository"`
	Version    *string `yarml:"version"`
}

type ChartEntry struct {
	APIVersion    string            `yaml:"apiVersion"`
	AppVersion    string            `yaml:"appVersion"`
	Created       time.Time         `yaml:"created"`
	Dependencies  []Dependencies    `yaml:"dependencies"`
	Description   string            `yaml:"description"`
	Digest        string            `yaml:"digest"`
	Icon          string            `yaml:"icon"`
	Keywords      []string          `yaml:"keywords"`
	Maintainers   []MaintainerEntry `yaml:"maintainers"`
	Name          string            `yaml:"name"`
	Version       string            `yaml:"version"`
	KubeVersion   string            `yaml:"kubeVersion"`
	Home          string            `yaml:"home"`
	Type          string            `yarml:"type"`
	Sources       []string          `yaml:"sources"`
	Engine        string            `yaml:"engine"`
	Deprecated    bool              `yaml:"deprecated"`
	TillerVersion string            `yaml:"tillerVersion"`
	Urls          []string          `yaml:"urls"`
}

type Charts struct {
	APIVersion string                  `yaml:"apiVersion"`
	Entries    map[string][]ChartEntry `yaml:"entries"`
	Generated  time.Time               `yaml:"generated"`
}
