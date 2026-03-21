package config

import (
	viper "github.com/spf13/viper"
)

type Config struct {
	Loglevel string

	Nick         string
	IrcServer    string
	IrcUseSSL    bool
	Channels     string
	NickservPass string

}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("buttsbot")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	viper.SetDefault("Loglevel", "INFO")
	viper.SetDefault("Nick", "buttsbot")
	viper.SetDefault("IrcUseSSL", false)
	viper.SetDefault("NickservPass", "")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
