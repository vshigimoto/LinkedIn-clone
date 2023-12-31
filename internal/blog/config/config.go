package config

type Config struct {
	Database   Database   `yaml:"Database"`
	HttpServer HttpServer `yaml:"HttpServer"`
}

type Database struct {
	Main    DbNode `yaml:"Main"`
	Replica DbNode `yaml:"Replica"`
}

type DbNode struct {
	Host     string `yaml:"Host"`
	User     string `yaml:"User"`
	Port     int    `yaml:"Port"`
	Password string `yaml:"Password"`
	DbName   string `yaml:"DbName"`
	SslMode  string `yaml:"SslMode"`
}

type HttpServer struct {
	Port      int `yaml:"Port"`
	AdminPort int `yaml:"AdminPort"`
}
