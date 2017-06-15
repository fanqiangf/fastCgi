package common

//在本文件中配置基本的环境变量目录
func getConfig() map[string]string {
	var config = make(map[string]string)
	//tcp监听的端口
	config["port"] = ":4445"
	//本机项目目录的src地址
	config["go_src"] = "/Users/qiangf/go/src/"
	//本机php bin命令下的php-cgi
	config["php-cgi"] = "/usr/local/Cellar/php56/5.6.30_6.tmp/bin/php-cgi"
	return config
}

//只开放读的接口

func GetRead(key string) string {
	confMap := getConfig()
	return confMap[key]
}
