package configuration

import (
	"log"

	"github.com/spf13/viper"
)

var config = new(configuration)

func init() {
	v := viper.New()

	if err := setEnvsBindings(v); err != nil {
		log.Fatalf("configuration: %v", err)
	}

	if err := v.Unmarshal(config); err != nil {
		log.Fatalf("configuration: %v", err)
	}
}

type configuration struct {
	Environment    string `mapstructure:"env"`             // Current environment mode
	Bind           string `mapstructure:"bind"`            // Server bind address
	ImageProcessor string `mapstructure:"image_processor"` // Image transform processor
}

func setEnvsBindings(v *viper.Viper) error {
	v.SetEnvPrefix("server")
	for env, def := range map[string]interface{}{
		"env":             "development",
		"bind":            "0.0.0.0:8080",
		"image_processor": "govips",
	} {
		if err := v.BindEnv(env); err != nil {
			return err
		}
		v.SetDefault(env, def)
	}

	return nil
}

// Bind server binding
func Bind() string {
	return config.Bind
}

// Environment current environment mode
func Environment() string {
	return config.Environment
}

// IsDevelopment run in development mode?
func IsDevelopment() bool {
	return config.Environment == "development"
}

// ImageProcessor return string representation of image processor
func ImageProcessor() string {
	return config.ImageProcessor
}
