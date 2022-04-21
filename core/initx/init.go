package initx

import (
	"os"
	"log"

	"github.com/spf13/viper"
)

func InitX() {
	viper.AddConfigPath(".")
	viper.SetConfigName("gomo")
	viper.SetConfigType("json")

	viper.SetDefault("modules", []string{})

	if err := viper.SafeWriteConfig(); err != nil {
		if os.IsNotExist(err) {
			err = viper.WriteConfig()

			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal(err)
		}
	}
}
