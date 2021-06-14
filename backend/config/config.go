package config

import "github.com/spf13/viper"

type Config struct {
	Host       string `mapstructure:"DB_HOST"`
	Port       string `mapstructure:"DB_PORT"`
	User       string `mapstructure:"DB_USER"`
	Password   string `mapstructure:"DB_PASSWORD"`
	Dbname     string `mapstructure:"DB_NAME"`
	DbDriver   string `mapstructure:"DB_DRIVER"`
	StaticRoot string `mapstructure:"STATIC_ROOT"`
	ApiSecret  string `mapstructure:"API_SECRET"`
	LogFile    string `mapstructure:"LOGFILE"`
}

func LoadConfig(path string, file string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(file)
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
