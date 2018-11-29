// 装备包
package task
import (
	"sync"
	"log"
	"regexp"
	"strconv"
	"bytes"
	"net/url"
	"fmt"
	"github.com/lwlwilliam/Golang/miniProgram/slbrowser/tool"
)

// 装备类型
type equipmentInfo struct {
	// 装备 id
	Id			int
	// 装备拼音
	Name		string
	// 级别：白色为 1，蓝色为 2，绿色为 3，紫色为 4，金色为 5
	Level		int
	// 耐久度，格式为"100/100"，前面的整数为当前耐久度，后面的为当前最大耐久度
	Durability	int
}

func EquipmentTask(equipmentList []byte, config tool.Config, done chan bool) {
	var wg sync.WaitGroup
	log.Println("task.EquipmentTask: come in")

	// 找出所有装备
	equipmentRegex, _ := regexp.Compile(`<tr id="equip_(\d+)">(.|\n)*?<img.*?src=.*?heroitem/(\w+)\.gif(.|\n)*?class="item_level(\d+).*?".*?(\d+/\d+)(.|\n)*?</tr>`)
	res := equipmentRegex.FindAllSubmatch(equipmentList, -1)
	queue := make(chan bool, len(res))

	for _, val := range res {
		var equipment equipmentInfo
		for k, v := range val {
			if k == 1 || k == 3 || k == 5 || k == 6 {
				// 装备 id
				if k == 1 {
					equipment.Id, _ = strconv.Atoi(string(v))
				// 装备拼音名称
				} else if k == 3 {
					equipment.Name = string(v)
				// 装备级别
				} else if k == 5 {
					equipment.Level, _ = strconv.Atoi(string(v))
				// 装备耐久度
				} else {
					durability := bytes.Split(v, []byte("/"))
					equipment.Durability, _ = strconv.Atoi(string(durability[1]))
					wg.Add(1)
					go sell(equipment, queue, &wg, config)
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

// 出售
func sell(equipment equipmentInfo, queue chan bool, wg *sync.WaitGroup, config tool.Config) {
	// 写死的
	cid := "5330"
	itemSellType := "2"
	heroId := "1354705318819751"
	pwd := "123456"

	if (equipment.Level < 5) {
		var u url.Values = make(url.Values)
		u.Set("itemId", strconv.Itoa(equipment.Id))
		u.Set("heroId", heroId)
		u.Set("itemSellType", itemSellType)
		u.Set("cid", cid)
		u.Set("pwd", pwd)
		params := u.Encode()

		res := tool.Request("POST", config.SellEquipmentURL, config.Cookie, params)

		if string(res) == "1" {
			log.Println(equipment.Name, "出售成功")
		} else {
			log.Println(equipment.Name, "出售失败")
			fmt.Println(string(res))
		}
	}

	queue <- true
	wg.Done()
}
