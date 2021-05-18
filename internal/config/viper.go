package config

import (
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

// Viper is a stuct to hold viper struct
type Viper struct{ *viper.Viper }

// GetViper ia a function to initialize  viper
func GetViper() *Viper {
	return &Viper{viper.New()}
}

// BindEnvs is a workaround to make the unmarshal work with environment variables
// reference: https://github.com/spf13/viper/issues/761#issuecomment-626122696
func (vpr *Viper) BindEnvs(iface interface{}, parts ...string) {
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
			vpr.BindEnvs(v.Interface(), parts...)
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			vpr.BindEnvs(v.Interface(), append(parts, tv)...)
		default:
			vpr.BindEnv(strings.Join(append(parts, tv), "."))
		}
	}
}

// IsFileNotFoundError is a funtion warper to check error file not found
func (vpr *Viper) IsFileNotFoundError(err error) bool {
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		// Config file not found; ignore error if desired
		return true
	}
	return false
}
