package rds

//Config DB configs
type Config struct {
	Host     string `yaml:"host" default:"localhost"`
	Port     int    `yaml:"port" default:"25432"`
	DBName   string `yaml:"db_name" default:"common"`
	User     string `yaml:"user" default:"common"`
	Password string `yaml:"password" default:"123456"`
	SslMode  string `yaml:"ssl_mode" default:"disable"`
}
