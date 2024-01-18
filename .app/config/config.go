package config

import (
	"strings"

	"github.com/spf13/viper"
)

func DatabaseName() string {
	return viper.GetString("DATABASE")
}
func MongoConnectionString() string {
	databaseUrl := strings.ReplaceAll(viper.GetString("DATABASEURL"), "mongodb://", "")
	return "mongodb://" + viper.GetString("DATABASEADMIN") + ":" + viper.GetString("DATABASEPASSWORD") + "@" + databaseUrl
}
func Setup(envPath string) {
	viper.SetConfigFile(envPath)
	viper.AutomaticEnv()
	viper.ReadInConfig()

}
