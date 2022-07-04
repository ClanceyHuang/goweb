package config

import (
	"encoding/json"
	"os"
)

// 定义配置结构体
type config struct {
	Static              string    `json:"static"`
	Template            string    `json:"template"`
	Address             string    `json:"address"`
	Port                int       `json:"port"`
	PprofPort           int       `json:"pprofPort"`
	HandleTimeoutSecond int       `json:"handleTimeoutSecond"`
	Log                 logConfig `json:"log"`
}

// 定义日志配置结构体
type logConfig struct {
	Trace   string `json:"trace"`
	Info    string `json:"info"`
	Warning string `json:"warning"`
	Error   string `json:"error"`
}

var Config config

/**
 * 加载配置的函数
 * @param string configFile 配置文件路径
 * @return mixed 返回抛错或者对应的配置文件数据
 */
func LoadConfig(configFile string) error {
	configData, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}
	return json.Unmarshal(configData, &Config)
}

// 测试用方法
func exportDefaultConfig(configFile string) error {
	defaultConfig := config{}
	configData, err := json.MarshalIndent(&defaultConfig, "", "\t")
	if err != nil {
		return err
	}
	err = os.WriteFile(configFile, configData, 0666)
	if err != nil {
		return err
	}
	return nil
}
