package configs

import "github.com/spf13/viper"

type Conf struct {
	DBDriver            string `mapstructure:"DB_DRIVER"`
	DBHost              string `mapstructure:"DB_HOST"`
	DBPort              string `mapstructure:"DB_PORT"`
	DBUser              string `mapstructure:"DB_USER"`
	DBPassword          string `mapstructure:"DB_PASSWORD"`
	DBName              string `mapstructure:"DB_NAME"`
	RABBITMQ_USER       string `mapstructure:"RABBITMQ_USER"`
	RABBITMQ_PASSWORD   string `mapstructure:"RABBITMQ_PASSWORD"`
	RABBITMQ_HOST       string `mapstructure:"RABBITMQ_HOST"`
	RABBITMQ_PORT       string `mapstructure:"RABBITMQ_PORT"`
	RABBITMQ_QUEUE_NAME string `mapstructure:"RABBITMQ_QUEUE_NAME"`
	WebServerPort       string `mapstructure:"WEBSERVER_PORT"`
}

func LoadConfig(path, envFile string) (*Conf, error) {
	var cfg *Conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(envFile)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, err
}
