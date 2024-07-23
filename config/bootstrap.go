package config

import (
	"fmt"
	"github.com/joho/godotenv"
	common "github.com/saas0503/factory-common"
	"log"
	"os"
	"reflect"
	"strconv"
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

func Load[C any](c C) (*C, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	st := common.MergeStructs(c, ApiConfig{})
	cfg := make(map[string]interface{})

	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		tag := field.Tag.Get("mapstructure")
		val := os.Getenv(tag)
		if val == "" {
			continue
		}

		if field.Type == reflect.TypeOf(1) {
			i, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			cfg[field.Name] = i
		} else {
			cfg[field.Name] = os.Getenv(tag)
		}
	}
	fmt.Println(cfg)
	var result = &c
	for k, v := range cfg {
		err := common.SetField(result, k, v)
		if err != nil {
			return nil, err
		}
	}

	return &c, nil
}
