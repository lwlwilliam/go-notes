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


// 刷新从内存中获取的配置
func (c *ConfigFile) ReloadData(in io.Reader) (err errpr)


// 往 ConfigFile 中追加文件，追加后会自动调用 Reload 方法，无需额外调用
func (c *ConfigFile) AppendFiles(files ...string) error
```

#### write.go: 写配置

```
// 保存配置。把 ConfigFile 内容写到 io.Writer 中
func SaveConfigData(c *ConfigFile, out io.Writer) (err error)


// 把配置写到文件系统中
func SaveConfigFile(c *ConfigFile, filename string) (err error)
```

#### conf.go: 配置操作

```
// 首先定义所需的常量，以及根据操作系统确定换行符等


// ConfigFile 表示 ini 格式的配置文件
type ConfigFile struct {
    lock            sync.RWMutex                    // Go map is not safe.
    fileNames       []string                        // Support mutil-files.
    data            map[string]map[string]string    // Section -> key : value
    
    // List can keep sections and keys in order.
    sectionList     []string                        // Section name list.
    keyList         map[string]string               // Section -> Key name list
    
    sectionComments map[string]string               // Sections comments.
    keyComments     map[string]map[string]string    // Keys comments.
    BlockMode       bool                            // Indicates whether use lock or not.
}


// 初始化一个 ConfigFile 变量
func newConfigFile(fileNames []string) *ConfigFile


// 添加一个 section-key-value 到配置中，需要先判断 section 和 key 是否已存在，注意要锁
func (c *ConfigFile) SetValue(section, key, value string) bool


// 删除指定 section 中的 key 值，先判断 section 和 key 是否已存在。注意要锁
func (c *ConfigFile) DeleteKey(section, key string) bool


// 获取 section 中的 key 值。先判断 section 和 key 是否存在。注意要锁。如果不存在，再判断是否为子 section
// TODO: 有待深入了解 ini 格式
func (c *ConfigFile) GetValue(section, key string) (string, error)


// 返回 value 对应的布尔类型值
func (c *ConfigFile) Bool(section, key string) (bool, error)


// 返回 value 对应的 Float64 类型值
func (c *ConfigFile) Float64(section, key string) (float64, error)


// 返回 value 对应的 Int 类型值
func (c *ConfigFile) Int(section, key string) (int, error)


// 返回 value 对应的 Int64 类型值
func (c *ConfigFile) Int64(section, key string) (int64, error)
```
