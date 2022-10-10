package models

type DBConfig struct {
	DriverName string
	Host       string
	Port       string
	Username   string
	Password   string
	Database   string
	Charset    string
}

type LogConfig struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"max_size"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
}

type SnowFlake struct {
	Startime  string 
	MachineID int64  `json:"machineID"`
}
