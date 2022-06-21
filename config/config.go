package config

import (
	"github.com/spf13/viper"
	"path"
	"runtime"
	"strings"
)

const FileName = "config"

func initialConfig() *viper.Viper {

	v := viper.New()

	v.SetConfigType("yml")

	_, b, _, _ := runtime.Caller(0)

	configPath := path.Join(path.Dir(b))

	v.AddConfigPath(configPath)

	v.SetConfigName(FileName)

	err := v.ReadInConfig()

	if err != nil {
		panic("can not read config file" + err.Error())
	}

	v.AutomaticEnv()

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	_ = v.AllSettings()

	return v
}

type Config interface {
	GetString(name string) string
	GetBool(name string) bool
	GetArray(name string) []string
	GetMap(name string) map[string]interface{}
	GetMapString(name string) map[string]string
	GetInt(name string) int
}

type config struct {
	v *viper.Viper
}

func (c *config) GetString(name string) string {
	return c.v.GetString(name)
}

func (c *config) GetBool(name string) bool {
	return c.v.GetBool(name)
}

func (c *config) GetArray(name string) []string {
	return c.v.GetStringSlice(name)
}

func (c *config) GetInt(name string) int {
	return c.v.GetInt(name)
}

func (c *config) GetMap(name string) map[string]interface{} {
	return c.v.GetStringMap(name)
}

func (c *config) GetMapString(name string) map[string]string {
	return c.v.GetStringMapString(name)
}

func NewConfig() Config {
	return &config{initialConfig()}
}
