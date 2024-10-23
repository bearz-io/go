package configblocks

type Import struct {
	Uri  string `json:"uri" yaml:"uri"`
	Hash string `json:"hash" yaml:"hash"`
}
