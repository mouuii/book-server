package core

import "github.com/spf13/viper"

type Config struct {
	File string
	Path string
	Name string
	Type string
}

func (this *Config) Init() error {

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
	dialect := viper.GetString("database.dialect")
	name := viper.GetString("database.database_name")
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")

	return DbConfig{
		dialect,
		name,
		username,
		password,
		host,
		port,
	}
}
