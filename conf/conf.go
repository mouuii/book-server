package conf

import (
	"github.com/spf13/viper"
)

type Config struct {
	File string
	Path string
	Name string
	Type string
}

func (this *Config) InitConfig() error {
	if this.Type == "" {
		this.Type = "yaml"
	}
	viper.SetConfigType(this.Type)
	this.setConfigFile()
	viper.AutomaticEnv()
	return viper.ReadInConfig()
}

func (this *Config) setConfigFile() {
	if this.File != "" {
		viper.SetConfigFile(this.File)
		return
	}
	viper.AddConfigPath(this.Path)
	viper.SetConfigName(this.Name)
}

type DbConfig struct {
	Dialect  string
	Name     string
	UserName string
	Password string
	Host     string
	Port     string
}

func GetDbConfig() DbConfig {
	db := viper.GetStringMapString("database")
	return DbConfig{
		db["dialect"],
		db["database_name"],
		db["username"],
		db["password"],
		db["host"],
		db["port"],
	}
}
