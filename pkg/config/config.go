package config

import (
	"api/performance_tune/model"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

var setting *model.Setting
var testSetting *model.Setting

// Config 讀取config.yaml
func ConfigInit() {

	// 获取当前运行程序的路径
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("Error getting executable path: %s", err)
	}
	exeDir := filepath.Dir(exePath)

	// 配置文件相对路径，从 src/pkg/apilib 到 src/cmd/config
	configPath := filepath.Join(exeDir)
	fmt.Println("configPath", configPath)
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
	if setting == nil {
		ConfigInit()
	}
	return setting
}

// Config 讀取config.yaml
func ConfigTestInit() {

	// 确定当前文件的目录
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)

	// 配置文件相对路径，从 pkg/router 到 cmd/config
	configPath := filepath.Join(basePath, "..", "..", "cmd", "config")

	// 设置 Viper 读取的文件名称（不包含扩展名）
	viper.SetConfigName("config")
	// 设置 Viper 读取的文件路径
	viper.AddConfigPath(configPath)
	// 设置配置文件类型
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("无法读取配置文件: %s\n", err)
		return
	}

	if err := viper.Unmarshal(&testSetting); err != nil {
		fmt.Printf("无法映射配置到结构体: %s\n", err)
		return
	}

	fmt.Println("confing setting", testSetting)
}

func GetTestSetting() *model.Setting {
	if testSetting == nil {
		ConfigTestInit()
	}
	return testSetting
}
