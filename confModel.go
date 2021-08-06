package main

type Config struct {
	ServerConf 	ServerConfig `ini:"SERVER"`
	ClientConf 	ClientConfig `ini:"CLIENT"`
}

type ServerConfig struct {
	IP 		string 	`ini:"ip"`
	Port 	int 	`ini:"port"`
}

type ClientConfig struct {
	Username 	string 		`ini:"username"`
	Password 	string		`ini:"password"`
}
