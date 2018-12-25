// goconfig API demo
package main

import (
	"fmt"
	"log"
	"os"
	"github.com/lwlwilliam/goconfig"
	"io"
)

func main() {
	//tmpName := path.Join(os.TempDir(), "goconfig", fmt.Sprintf("%d", time.Now().Nanosecond()))
	//fmt.Println(tmpName)

	//configs, err := loadConfigFile("testdata/conf.ini")
	//configs, err := loadFromData()

	dir, err := os.Getwd()
	fatalErr(err)
	reader, err := os.Open(dir+ "/testdata/conf.ini")
	fatalErr(err)
	configs, err := loadFromReader(reader)
	// 这里 LoadFromReader 时不能追加
	//configs.AppendFiles("testdata/conf2.ini")

	fatalErr(err)
	for _, v := range configs.GetSectionList() {
		fmt.Println(v, ": ")
		section, _ := configs.GetSection(v)

		for k, v := range section {
			fmt.Println(k, v)
		}
	}

	//getValue(configs)
	//delete(configs)
	//reload(configs)
	//transform(configs)


}

// 从 io.Reader 获取配置
func loadFromReader(in io.Reader) (*goconfig.ConfigFile, error) {
	fmt.Println("==================== loadFromReader")
	configs, err := goconfig.LoadFromReader(in)

	return configs, err
}

// 从输入数据中获取配置
func loadFromData() (*goconfig.ConfigFile, error) {
	fmt.Println("==================== loadFromData")
	data := []byte("[test]\nhaha=abc")
	configs, err := goconfig.LoadFromData(data)

	return configs, err
}

// 可以加载一个或多个文件，返回一个类型为 ConfigFile 的指针
func loadConfigFile(files ...string) (*goconfig.ConfigFile, error) {
	fmt.Println("==================== loadConfigFile")
	configs, err := goconfig.LoadConfigFile(files[0], files[1:]...)

	return configs, err
}

// GetValue 为 ConfigFile 类型的方法，通过该方法可以获取指定的值
func getValue(configs *goconfig.ConfigFile) {
	fmt.Println("==================== getValue")
	value, _ := configs.GetValue("Demo", "key1")
	fmt.Println(";", value)
}

// 删除 ConfigFile 中指定 section 或 key
func delete(configs *goconfig.ConfigFile) {
	fmt.Println("==================== delete")
	configs.DeleteKey("Demo", "key1")
	value, _ := configs.GetValue("Demo", "key1")
	fmt.Println(";", value)

	configs.DeleteSection("Demo")
	value, _ = configs.GetValue("Demo", "key2")
	fmt.Println(";", value)
}

// 重新加载 ConfigFile，与 DeleteKey 或 DeleteSection 方法一起使用对比效果
func reload(configs *goconfig.ConfigFile) {
	fmt.Println("==================== reload")
	delete(configs)
	value, _ := configs.GetValue("Demo", "key2")
	fmt.Println(";", value)
	configs.Reload()
	value, _ = configs.GetValue("Demo", "key2")
	fmt.Println(";", value)

	delete(configs)
	value, _ = configs.GetValue("Demo", "key2")
	fmt.Println(";", value)
	file, err := os.Open("testdata/conf.ini")
	fatalErr(err)
	configs.ReloadData(file)
	value, _ = configs.GetValue("Demo", "key2")
	fmt.Println(";", value)
}

// 保存 ConfigFile 到文件中
func save(configs *goconfig.ConfigFile) {
	fmt.Println("==================== save")
	delete(configs)
	file, _ := os.OpenFile("testdata/saveConfigData.ini", os.O_CREATE|os.O_RDWR, 0755)
	goconfig.SaveConfigData(configs, file)

	goconfig.SaveConfigFile(configs, "testdata/saveConfigFile.ini")
}

// key 值类型转换
func transform(configs *goconfig.ConfigFile) {
	fmt.Println("==================== judge")
	b, _ := configs.Bool("Demo", "key1")
	fmt.Println(";", b)

	f, _ := configs.Float64("Demo", "key1")
	fmt.Println(";", f)
	f, _ = configs.Float64("parent", "money")
	fmt.Println(";", f)

	i, _ := configs.Int("parent", "age")
	fmt.Println(";", i)

	i64, _ := configs.Int64("parent", "age")
	fmt.Println(";", i64)
}

// 错误处理
func fatalErr(err error) {
	if err != nil {
		log.Println(err.Error())
		os.Exit(-1)
	}
}
