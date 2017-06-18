	###########目录结构描述
	├── Readme.md                   // help
	├── app                         // 应用
	├── common                      // 配置
	│   ├── common.go               //公共函数
	│   └── config.go              // 系统依赖的全局变量 （--配置本文件才能运行demo--）
	├── main                      
	│   └── main.go                 // 入口文件
	├── phpCgi                      
	│   └── phpCgi.go               // cgi解析器
	├── server                      
	│   └── server.go               // web服务解析
#####说明 v0.1
仅做学习使用，不依赖go的net/http包,用go的协程做了socket握手以及http协议的
封装，目前仅支持静态web（get方式）请求和php代码的解析，后期准备加入fast-cgi协议的支持。



运行本demo需要对config.go文件进行配置。
