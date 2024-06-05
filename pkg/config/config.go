package config

import (
	"api/performance_tune/model"
	"fmt"

	"github.com/spf13/viper"
)

var setting *model.Setting

// Config 讀取config.yaml
func ConfigInit() {

	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	viper.SetConfigName("config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("无法读取配置文件: %s\n", err)
		return
	}

	if err := viper.Unmarshal(&setting); err != nil {
		fmt.Printf("无法映射配置到结构体: %s\n", err)
		return
	}

	fmt.Println("confing setting", setting)
}

func GetSetting() *model.Setting {
	return setting
}
