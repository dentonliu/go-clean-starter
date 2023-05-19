package config

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/spf13/viper"
)

type Config struct {
	ServerPort string `mapstructure:"server_port"`

	DSN     string `mapstructure:"dsn"`
	DSNTest string `mapstructure:"dsn_test"`

	JWTSigningMethod   string `mapstructure:"jwt_signing_method"`
	JWTSigningKey      string `mapstructure:"jwt_signing_key"`
	JWTVerificationKey string `mapstructure:"jwt_verification_key"`
}

func (c Config) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.DSN, validation.Required),
		validation.Field(&c.JWTSigningKey, validation.Required),
		validation.Field(&c.JWTVerificationKey, validation.Required))
}

func Load(configPaths ...string) (*Config, error) {
	c := Config{}

	v := viper.New()
	v.SetConfigName("app")
	v.SetConfigType("yaml")
	v.SetEnvPrefix("app")
	v.AutomaticEnv()
	v.SetDefault("jwt_signing_method", "HS256")

	for _, path := range configPaths {
		v.AddConfigPath(path)
	}

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("Failed to read the configuration file: %s", err)
	}

	if err := v.Unmarshal(&c); err != nil {
		return nil, err
	}

	return &c, c.Validate()
}
