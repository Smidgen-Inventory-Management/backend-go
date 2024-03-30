package smidgen

type ServerConfig struct {
	Environments map[string]struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Debug    bool   `yaml:"debug"`
		RootPath string `yaml:"root_path"`
	} `yaml:",inline"`
}