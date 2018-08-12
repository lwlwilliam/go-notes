// 找回了第一次玩的网游，写了个小程序自动练级。
// 当然还不完善，毕竟只花了一天构思和实现，而且对 Go 还不熟悉。
// 在这个过程中发现了自己的一些不足。这些算是练手吧，以后空闲继续完善。
// 暂时就把程序都放在一个文件里，需要配置文件的配合，配置就不上传了。
// 打算用这个练习一下 Go。
package main

import (
	"regexp"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"strconv"
	url2 "net/url"
	//"golang.org/x/text/encoding/simplifiedchinese"
	//"golang.org/x/text/transform"
	"sync"
	"fmt"
)

func throwErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Config struct{
	HeroListURL		string	`json:"heroListURL"`
	AdventureURL	string	`json:"adventureURL"`
	Cookie			string	`json:"cookie"`
}

type AdventureCfg [][]string

// 从英雄列表中保存有用英雄信息
type heroInfo struct {
	Name	string
	CityId	int
	HeroId	int
	//Grade	int
}

var globalConfig Config

// 获取配置
func getConfig() Config {
	cfg := Config{}

	file, err := ioutil.ReadFile("./config.json")
	throwErr(err)

	err = json.Unmarshal(file, &cfg)
	throwErr(err)

	return cfg
}

// 获取英雄冒险房间
func getAdventureCfg(len int) AdventureCfg {
	cfg := make(AdventureCfg, len)

	file, err := ioutil.ReadFile("./heroAdventure.json")
	throwErr(err)

	err = json.Unmarshal(file, &cfg)
	throwErr(err)

	return cfg
}

// 请求
func request(method string, url string, cookie string, params string) []byte {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, strings.NewReader(params))
	throwErr(err)
	req.Header.Set("Cookie", cookie)
	if method == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	throwErr(err)

	return body
}

// 为英雄创建任务
func createTask(heroList []byte, done chan bool) {
	// 找出英雄名称、英雄 ID、所在城市 ID
	heroRegex, _ := regexp.Compile(`class="t_color03"><strong>([^<]+)(.|\n)*?view_mission_city\.alc\?[^&]+&cityId=(\d+)(.|\n)*?showSendHero\((\d+)`)
	res := heroRegex.FindAllSubmatch(heroList, -1)
	queue := make(chan bool, len(res))
	var wg sync.WaitGroup

	heroAdventure := getAdventureCfg(len(res))

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
					go adventure(hero, queue, &wg, heroAdventure)
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

// 为英雄分发冒险任务
func adventure(hero heroInfo, queue chan bool, wg *sync.WaitGroup, heroAdventure AdventureCfg) {
	var roomId int
	var routeId int
	for _, val := range heroAdventure {
		if val[0] == hero.Name {
			temp, _ := strconv.Atoi(val[1])
			roomId = temp
			//roomId, _ := strconv.Atoi(val[1])
			routeId = roomId * 10 + 1
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

	res := request("POST", globalConfig.AdventureURL, globalConfig.Cookie, params)
	fmt.Println(string(res))


	queue <- true
	wg.Done()
}

func main() {
	var done = make(chan bool)
	cfg := getConfig()
	globalConfig = cfg
	heroList := request("GET", cfg.HeroListURL, cfg.Cookie, "")

	go createTask(heroList, done)
	<- done
}
