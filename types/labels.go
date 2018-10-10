package types

// LabelFile represents a labels.yaml file.
type LabelFile struct {
	Repo   string   `yaml:"repo"`
	Labels []*Label `yaml:"labels"`
}

// Label represents a label in the labels.yaml file.
type Label struct {
	Name        string `yaml:"name"`
	Color       string `yaml:"color"`
	From        string `yaml:"from,omitempty"`
	Description string `yaml:"description,omitempty"`
}
