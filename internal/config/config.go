package config

import "github.com/spf13/viper"

const (
	KubeConfigPath 	= "KUBE_CONFIG_PATH"
	K8sURL			= "K8s_URL"
	LogLevel		= "LOG_LEVEL"
)

func init() {
	viper.AutomaticEnv()
	viper.SetDefault(KubeConfigPath, "")
	viper.SetDefault(K8sURL, "")
	viper.SetDefault(LogLevel, "info")
}

func ReadConfig(env string) error {
	viper.SetConfigFile("app-" + env + ".yml")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}