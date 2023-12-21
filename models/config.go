package models

// StoreCfg represents the configuration for the application's database layer.
type StoreCfg struct {
	Server   string `json:"server" yaml:"server" mapstructure:"server"`
	Database string `json:"database" yaml:"database" mapstructure:"database"`
	User     string `json:"user" yaml:"user" mapstructure:"user"`
	Password string `json:"password" yaml:"password" mapstructure:"password"`
	Port     int    `json:"port" yaml:"port" mapstructure:"port"`
	Timeout  int    `json:"timeout" yaml:"timeout" mapstructure:"timeout"`
}