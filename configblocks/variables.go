package configblocks

type VariableItem struct {
	Name       string `json:"name" yaml:"name"`
	Value      string `json:"value" yaml:"value"`
	Type       string `json:"type" yaml:"type"`
	Expression bool   `json:"expression"`
}
