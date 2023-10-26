package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

var TargetBits int
var BDFile string

func init() {
	viper.AddConfigPath("GoApp/configs/")
	viper.SetConfigName("config") // имя файла конфигурации без расширения
	viper.SetConfigType("json")   // формат файла конфигурации

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Ошибка при чтении конфигурации:", err)
		panic(err)
	}

	TargetBits = viper.GetInt("targetBits")
	BDFile = viper.GetString("dbFile")

}
