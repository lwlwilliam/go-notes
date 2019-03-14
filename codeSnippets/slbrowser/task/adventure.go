// 冒险
package task

import (
	"regexp"
	"sync"
	"strconv"
	url2 "net/url"
	"log"
	"fmt"
	"github.com/lwlwilliam/Golang/codeSnippets/slbrowserr/tool"
)

type heroInfo struct {
	Name	string
	CityId	int
	HeroId	int
	//Grade	int
}

// 匹配英雄并分配任务
func HeroTask(heroList []byte, config tool.Config, done chan bool) {
	var wg sync.WaitGroup
	log.Println("task.HeroTask: come in")

	// 找出英雄名称、英雄 ID、所在城市 ID
	heroRegex, _ := regexp.Compile(`class="t_color03"><strong>([^<]+)(.|\n)*?view_mission_city\.alc\?[^&]+&cityId=(\d+)(.|\n)*?showSendHero\((\d+)`)
	res := heroRegex.FindAllSubmatch(heroList, -1)
	queue := make(chan bool, len(res))

	// 冒险配置
	var heroAdventure tool.AdventureConfig
	heroAdventure = heroAdventure.Res(len(res))

	for _, val := range res {
		var hero heroInfo
		for k, v := range val {
			if k == 1 || k == 3 || k == 5 {
				// 英雄名称
				if k == 1 {
					hero.Name = string(v)
					// 所在城市 ID
				} else if k == 3 {
					hero.CityId, _ = strconv.Atoi(string(v))
					// 英雄 ID
				} else {
					hero.HeroId, _ = strconv.Atoi(string(v))
					wg.Add(1)
					go adventure(hero, queue, &wg, config, heroAdventure)
				}
			}
		}
	}
	wg.Wait()
	close(queue)

	for {
		if _, ok := <- queue; !ok {
			done <- true
		}
	}
}

// 冒险任务
func adventure(hero heroInfo, queue chan bool, wg *sync.WaitGroup, config tool.Config, heroAdventure tool.AdventureConfig) {
	var roomId int
	var routeId int
	for _, val := range heroAdventure {
		if val[0] == hero.Name {
			temp, _ := strconv.Atoi(val[1])
			//roomId = temp
			//routeId = roomId * 10 + 1
			roomId, routeId = temp, temp * 10 + 1
			break
		}
	}

	var u url2.Values = make(url2.Values)
	u.Set("heroId", strconv.Itoa(hero.HeroId))
	u.Set("cityId", strconv.Itoa(hero.CityId))
	u.Set("roomName", "新任务")
	u.Set("routeId", strconv.Itoa(routeId))
	u.Set("missionId", strconv.Itoa(roomId))
	u.Set("heroVocation", "")
	params := u.Encode()

	res := tool.Request("POST", config.AdventureURL, config.Cookie, params)

	if string(res) == "1" {
		log.Println(hero.Name, "冒险成功")
	} else {
		log.Println(hero.Name, "冒险失败")
		fmt.Println(string(res))
	}

	queue <- true
	wg.Done()
}
