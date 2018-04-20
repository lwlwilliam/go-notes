// 这个示例程序展示如何写基础单元测试
// 测试文件名是 listing01_test.go，Go 语言的测试工具只会认为以 _test.go 结尾的文件是测试文件。
// 如果没有遵从这个约定，在包里运行 go test 的时候就可能会报告没有测试文件。一旦测试工具找到了测试
// 文件，就会查找里面的测试函数并执行。
package listing01_test

import (
	"net/http"
	"testing"	// 引入 testing 包，这个包提供了从测试框架到报告测试的输出和状态的各种测试功能的支持。
)

// 包含写测试输出时用到的对号和叉号
const checkMark = "\u2713"
const ballotX = "\u2717"

// TestDownload 确认 http 包的 Get 函数可以下载内容
// 一个测试函数必须是公开的函数，并且以 Test 单词开头。不但函数名字要以 Test 开头，而且函数的签名必须接收
// 一个指向 testing.T 类型的指针，并且不返回任何值。如果没有遵守这些约定，测试框架就不会认为这个函数是一个
// 测试函数，也不会让测试工具去执行它。
func TestDownload(t *testing.T) {	// 指向 testing.T 类型的指针很重要。这个指针提供的机制可以报告每个测试
									// 的输出和状态。测试的输出格式没有标准要求。

	// 要测试的 URL，以及期望从响应中返回的状态
	url := "http://www.goinggo.net/feeds/posts/default?alt=rss"
	statusCode := 200

	// t.Log 输出测试消息
	t.Log("Given the need to test downloading content.")
	{
		// 格式化消息
		// 每个测试函数都应该通过解释这个测试的给定要求(given need)，来说明为什么应该存在这个测试
		// 对这个例子来说，给定要求是测试能否成功下载数据。在声明了测试的给定要求后，测试应该说明被
		// 测试的代码应该在什么情况下被执行，以及如何执行。

		// 这里看到了测试执行条件的说明。它特别说明了要测试的值。
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"",
			url, statusCode)
		{

			// 在每种情况下，都会说明测试应有的结果。如果调用失败，除了结果，还会输出叉号及得到的错误值。
			// 如果测试成功，会输出对号
			resp, err := http.Get(url)
			if err != nil {
				t.Fatal("\t\tShould be able to make the Get call.",
					ballotX, err)
			}
			t.Log("\t\tShould be able to make the Get call.",
				checkMark)

			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t\tShould receive a \"%d\" status. %v",
					statusCode, checkMark)
			} else {
				t.Errorf("\t\tShould receive a \"%d\" status. %v %v",
					statusCode, ballotX, resp.StatusCode)
			}
		}
	}
}
