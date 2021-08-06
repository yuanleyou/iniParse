package main

func main() {

	// 制作测试数据
	var conf Config
	conf.ServerConf.IP = "192.168.0.1"
	conf.ServerConf.Port = 8080
	conf.ClientConf.Username = "Yuan"
	conf.ClientConf.Password = "Abcd123456"

	//将配置信息写入配置文件
	StructToFile("./config.ini", conf)

	var config Config
	// 从配置文件解析到结构体中
	FileToStruct("./config.ini", &config)
	logger.Printf("%#v", config)
}