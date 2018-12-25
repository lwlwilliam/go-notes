### goconfig 学习

项目地址：[https://github.com/Unknwon/goconfig](https://github.com/Unknwon/goconfig)
API 文档地址：[https://gowalker.org/github.com/Unknwon/goconfig](https://gowalker.org/github.com/Unknwon/goconfig)


#### read.go: 读配置

```
// 创建 ConfigFile 结构，打开一个或多个 ini 文件并把文件信息及内容保存到 ConfigFile 中
func LoadConfigFile(fileName string, moreFiles ...string) (c *ConfigFile, err error)


// 加载文件
func (c *ConfigFile) loadFile(fileName string) (err error)


// 读取文件内容。这里对 BOM 头进行了处理，逐行从缓冲中读取内容。首先每行首尾空白去除，空行忽略
//      1. 当首字符为"#"或";"时，按注释处理；
//      2. 当首字符为"["且尾字符为"]"时，去首尾字符空白后按 section 处理，并保存到 ConfigFile 中；
//      3. section 为空时，出错；
//      4. 默认分支：首字符为"\""或"`"时进行特殊处理；
func (c *ConfigFile) read(reader io.Reader) (err error)


// 从输入数据中构建 ConfigFile 变量并保存到临时文件中，由于数据会保存到系统临时文件，所以不要用这个
// 方法处理敏感数据
func LoadFromData(data []byte) (c *ConfigFile, err error)


// 从 io.Reader 中直接接收数据并构建 ConfigFile 变量。必须要使用 ReloadData 刷新(这里暂时不清楚为什么这么说，有待测试)，
// 而且不能追加文件
func LoadFromReader(in io.Reader) (c *ConfigFile, err error)


// 当配置文件改变时使用该方法进行刷新
func (c *ConfigFile) Reload() (err error)
```
