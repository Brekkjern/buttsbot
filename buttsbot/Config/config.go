package Config

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

	TwitterAPIToken string
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	viper.SetDefault("Loglevel", "INFO")
	viper.SetDefault("Nick", "buttsbot")
	viper.SetDefault("IrcUseSSL", false)
	viper.SetDefault("NickservPass", "")
	viper.SetDefault("TwitterAPIToken", "")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
