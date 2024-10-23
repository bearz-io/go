package vaults

type VaultSection struct {
	Uri     string                 `json:"uri" yaml:"uri"`
	Kind    string                 `json:"kind" yaml:"kind"`
	With    map[string]interface{} `json:"with" yaml:"with"`
	Secrets []SecretItem           `json:"secrets" yaml:"secrets"`
}

type SecretItem struct {
	Name     string `json:"name" yaml:"name"`
	Path     string `json:"path" yaml:"path"`
	Required bool   `json:"required" yaml:"required"`
	Size     int    `json:"size" yaml:"size"`
	Upper    bool   `json:"upper" yaml:"upper"`
	Lower    bool   `json:"lower" yaml:"lower"`
	Digits   bool   `json:"digits" yaml:"digits"`
	Special  string `json:"special" yaml:"special"`
}
