package vaults

type Features struct {
	exists map[string]bool
}

func (f *Features) Has(name string) bool {
	enabled, ok := f.exists[name]
	return ok && enabled
}

func NewFeatures(features map[string]bool) *Features {
	return &Features{exists: features}
}
