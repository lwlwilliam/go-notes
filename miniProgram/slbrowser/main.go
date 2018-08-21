package main

import (
	"./tool"
	"./task"
)

func main() {
	done := make(chan bool)

	var config tool.Config
	config = config.Res()

	heroList := tool.Request("GET", config.HeroListURL, config.Cookie, "")

	go task.HeroTask(heroList, config, done)
	<- done
}

