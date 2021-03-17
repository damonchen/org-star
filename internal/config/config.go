package config

// DB db
type DB struct {
	Name     string `yaml:"name,omitempty"`
	Host     string `yaml:"host,omitempty"`
	Port     string `yaml:"port,omitempty"`
	User     string `yaml:"user,omitempty"`
	Password string `yaml:"password,omitempty"`
	Charset  string `yaml:"charset,omitempty"`
}

// Github github
type Github struct {
	AppID  string `yaml:"appId,omitempty"`
	AppKey string `yaml:"appKey,omitempty"`
}

// Configuration configuration
type Configuration struct {
	Host   string `yaml:"host,omitempty"`
	Port   string `yaml:"port,omitempty"`
	Server string `yaml:"server,omitempty"`
	DB     DB     `yaml:"db,omitempty"`
	Github Github `yaml:"github,omitempty"`
}
