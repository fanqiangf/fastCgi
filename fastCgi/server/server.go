package server

import (
	"bufio"
	"fastCgi/common"
	"fastCgi/phpCgi"
	"fmt"
	"net"
	"strings"
)

func HandleClient(conn net.Conn) {
	defer conn.Close()
	data, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		common.CheckError(err, 45)
	}
	if len(data) == 0 {
		return
	}
	fmt.Println("client的消息是", data)
	//获取method filename 协议名称
	res := strings.Split(data, " ")
	fmt.Println(res)
	method := res[0]
	url := res[1]
	agreement := res[2]
	strLen := strings.Index(url, "/")
	//截取问号之前的
	//判断是否有?
	endLen := strings.Index(url, "?")
	if endLen == -1 {
		endLen = 0
	}

	isHTTP := strings.Contains(agreement, "HTTP")
	if method == "GET" && isHTTP {
		fmt.Println("您的http请求已经正式进来了!")
		if url == "/" {
			common.WriteBuff("<h1>你没有说想去那个页面！</h1>", conn)
		}
		filename := common.Substring(url, strLen+1, endLen-1)
		isPhp := strings.Contains(filename, ".php")
		path := common.GetRead("go_src")
		fmt.Println(filename)
		filename = path + filename
		fmt.Println(filename, "\n")
		fileExists := common.PathExists(filename)
		fmt.Println(fileExists)
		if !fileExists {
			common.WriteBuff("<h1>我还是not 404 found 吧！</h1>", conn)
		}
		if fileExists {
			if isPhp {
				//获取问号后面的params
				parmasStart := strings.Index(url, "?")
				paramsScript := common.Substring(url, parmasStart+1, -1)
				//调用php-cgi程序去解析php代码
				defaultHTML := phpCgi.PhpCgi(filename, paramsScript)
				common.WriteBuff(defaultHTML, conn)
			} else {

				defaultHTML := common.Rear(filename)
				common.WriteBuff(defaultHTML, conn)

			}

		}
	}
}
