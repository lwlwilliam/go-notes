// 工具包
package tool
import (
	"io/ioutil"
	"log"
	"os"
	"encoding/json"
)

const (
	configPath		= "./config/config.json"
	adventurePath	= "./config/adventure.json"
)

// config 格式
type Config struct {
	HeroListURL			string	`json:heroListURL`
	AdventureURL		string	`json:adventureURL`
	EquipmentURL		string	`json:equipmentURL`
	SellEquipmentURL	string 	`json:sellEquipmentURL`
	Cookie				string	`json:cookie`
}

// adventure 配置格式
type AdventureConfig [][]string

// 获取通用配置
func (this Config) Res() Config {
	var cfg Config
	config, err := ioutil.ReadFile(configPath)
	errHandler(err)

	err = json.Unmarshal(config, &cfg)
	errHandler(err)

	return cfg
}

// 获取冒险配置
func (this AdventureConfig) Res(len int) AdventureConfig {
	cfg := make(AdventureConfig, len)
	config, err := ioutil.ReadFile(adventurePath)
	errHandler(err)

	err = json.Unmarshal(config, &cfg)
	errHandler(err)

	return cfg
}

// 错误处理
func errHandler(err error) {
	if err != nil {
		log.Println("Error:", err.Error())
		os.Exit(-1)
	}
}
