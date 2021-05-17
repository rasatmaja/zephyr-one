package config

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

// Config ...
type Config struct {
	ServerHost    string `mapstructure:"SERVER_HOST"`
	ServerPort    int    `mapstructure:"SERVER_PORT"`
	ServerReadTO  int    `mapstructure:"SERVER_READ_TIMEOUT"`
	ServerWriteTO int    `mapstructure:"SERVER_WRITE_TIMEOUT"`
	ServerIdleTO  int    `mapstructure:"SERVER_IDLE_TIMEOUT"`
}

// Viper ...
type Viper struct {
	*viper.Viper
}

// BuildConfig ...
func BuildConfig() *Config {
	cfg := &Config{}

	vpr := Viper{viper.New()}
	vpr.AddConfigPath(".")
	vpr.SetConfigName("app")
	vpr.SetConfigType("env")

	vpr.AutomaticEnv()
	vpr.bindenvs(cfg)

	if err := vpr.ReadInConfig(); err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	if err := vpr.Unmarshal(&cfg); err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	vpr.Debug()

	fmt.Println(cfg)
	return cfg
}

// bindenvs: workaround to make the unmarshal work with environment variables
// Inspired from solution found here : https://github.com/spf13/viper/issues/188#issuecomment-399884438
// reference: https://github.com/spf13/viper/issues/761#issuecomment-626122696
func (b *Viper) bindenvs(iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	if ifv.Kind() == reflect.Ptr {
		ifv = ifv.Elem()
	}
	for i := 0; i < ifv.NumField(); i++ {
		v := ifv.Field(i)
		t := ifv.Type().Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}
		if tv == ",squash" {
			b.bindenvs(v.Interface(), parts...)
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			b.bindenvs(v.Interface(), append(parts, tv)...)
		default:
			b.BindEnv(strings.Join(append(parts, tv), "."))
		}
	}
}
