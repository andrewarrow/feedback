package util

import "github.com/spf13/viper"
import "fmt"

var AllConfig Config

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type PathConfig struct {
	Prefix string
	Videos string
}
type HttpConfig struct {
	Port string
	Host string
}
type DirectoryConfig struct {
	Active string
	Guid   string
}

type Config struct {
	Db          DatabaseConfig  `mapstructure:"database"`
	Path        PathConfig      `mapstructure:"paths"`
	Http        HttpConfig      `mapstructure:"http"`
	Directories DirectoryConfig `mapstructure:"directories"`
}

func InitConfig() bool {

	v := viper.New()
	v.AddConfigPath(".")
	v.AddConfigPath("/")
	v.AddConfigPath("..")
	v.SetConfigName("conf")
	v.SetConfigType("toml")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		fmt.Println(err)
		return false
	}
	if err := v.Unmarshal(&AllConfig); err != nil {
		fmt.Printf("couldn't read config: %s", err)
	}

	return true
}
