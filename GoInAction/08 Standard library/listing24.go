// 这个示例程序展示如休使用 json 包和 NewDecoder 函数来解码 JSON 响应
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
)

type (
	// gResult 映射到从搜索拿到的结果文档
	gResult struct {
		GsearchResultClass	string	`json:"GsearchResultClass"`
		UnescapedURL		string	`json:"unescapedUrl"`
		URL					string	`json:"url"`
		VisibleURL			string	`json:"visibleUrl"`
		CacheURL			string	`json:"cacheUrl"`
		Title				string	`json:"title"`
		TitleNoFormatting	string	`json:"titleNoFormatting"`
		Content				string	`json:"content"`
	}

	 //gResponse 包含顶级的文档
	gResponse struct {
		ResponseData struct {
			Results []gResult `json:"results"`
		} `json:"responseData"`
	}
)

func main() {
	uri := "http://ajax.googleapis.com/ajax/services/search/web?v=1.0&rsz=8&q=golang"

	// 向 Google 发起搜索
	resp, err := http.Get(uri)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	defer resp.Body.Close()

	// 将 JSON 响应解码到结构类型
	var gr gResponse
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(gr)
}
