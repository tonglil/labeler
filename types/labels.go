package types

type LabelFile struct {
	Repo   string  `yaml:"repo"`
	Labels []Label `yaml:"labels"`
}

type Label struct {
	Name  string `yaml:"name"`
	Color string `yaml:"color"`
	From  string `yaml:"from,omitempty"`
}
