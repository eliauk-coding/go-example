package model

type ServerConfig struct {
	Ip   string `json:"ip"`
	Port int    `json:"port"`
}

type MysqlConfig struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Database string  `json:"database"`
	Host     string  `json:"host"`
	Port     int     `json:"port"`
	Timeout  float64 `json:"timeout"`
}

type Config struct {
	ServerConf ServerConfig `json:"server"`
	MysqlConf  MysqlConfig  `json:"mysql"`
}
