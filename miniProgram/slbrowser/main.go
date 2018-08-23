package main

import (
	"./tool"
	"./task"
)

func main() {
	var config tool.Config
	config = config.Res()

	// 冒险
	heroList := tool.Request("GET", config.HeroListURL, config.Cookie, "")
	adventureDone := make(chan bool)
	go task.HeroTask(heroList, config, adventureDone)
	<- adventureDone

	// 卖装备
	equipmentList := tool.Request("POST", config.EquipmentURL, config.Cookie, "itemType=1&itemSellType=2")
	equipmentDone := make(chan bool)
	go task.EquipmentTask(equipmentList, config, equipmentDone)
	<- equipmentDone
}
