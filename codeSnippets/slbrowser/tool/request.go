// 工具包
package tool
import (
	"net/http"
	"strings"
	"io/ioutil"
)

// 请求
func Request(method string, url string, cookie string, params string) []byte {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, strings.NewReader(params))
	errHandler(err)
	req.Header.Set("Cookie", cookie)
	if method == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	errHandler(err)

	return body
}
