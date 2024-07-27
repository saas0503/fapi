package config

import "github.com/spf13/viper"

type Config struct {
	Port      int    `mapstructure:"PORT"`
	NodeEnv   string `mapstructure:"NODE_ENV"`
	ApiPrefix string `mapstructure:"API_PREFIX"`

	RedisAddr string `mapstructure:"REDIS_ADDR"`
	RedisPass string `mapstructure:"REDIS_PASS"`

	DbHost string `mapstructure:"DB_HOST"`
	DbPort int    `maptstructure:"DB_PORT"`
	DbUser string `mapstrucutre:"DB_USER"`
	DbPass string `mapstructure:"DB_PASS"`
	DbName string `mapstrucutre:"DB_NAME"`

	AccessTokenPrivateKey string `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey  string `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	AccessTokenExpiresIn  int    `mapstructure:"ACCESS_TOKEN_EXPIRES_IN"`

	RefreshTokenPrivateKey string `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefreshTokenPublicKey  string `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
	RefreshTokenExpiresIn  int    `mapstructure:"REFRESH_TOKEN_EXPIRES_IN"`
}

func Load(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
