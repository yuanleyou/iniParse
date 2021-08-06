# iniParse
A simple ini file serialization and deserialization tool is implemented using the reflect mechanism of the Go language 

# example

```go
package main

func main() {

	// 制作测试数据
	var conf Config
	conf.ServerConf.IP = "192.168.0.1"
	conf.ServerConf.Port = 8080
	conf.ClientConf.Username = "Yuan"
	conf.ClientConf.Password = "Abcd123456"

	//将结构体信息写入配置文件
	StructToFile("./config.ini", conf)

	var config Config
	// 从配置文件解析到结构体中
	FileToStruct("./config.ini", &config)
	logger.Printf("%#v", config)
}
```

# output

```shell
main.Config{ServerConf:main.ServerConfig{IP:"192.168.0.1", Port:8080}, ClientConf:main.ClientConfig{Username:"Yuan", Password:"Abcd123456"}}
```
