package config

import (
	"github.com/spf13/viper"
	"strings"
)

type ApiConfig struct {
	Port       int    `mapstructure:"PORT"`
	NodeEnv    string `mapstructure:"NODE_ENV"`
	ApiContext string `mapstructure:"API_CONTEXT"`

	AccessTokenPrivateKey string `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey  string `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	AccessTokenExpiresIn  int    `mapstructure:"ACCESS_TOKEN_EXPIRES_IN"`

	RefreshTokenPrivateKey string `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefreshTokenPublicKey  string `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
	RefreshTokenExpiresIn  int    `mapstructure:"REFRESH_TOKEN_EXPIRES_IN"`
}

func Load[C any](path string) (*C, error) {
	var cfg *C
	var err error

	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
