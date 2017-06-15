package common

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"
)

func CheckError(err error, line int) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s %d \n", err.Error(), line)
	}

} //读取文件内容
func Rear(filename string) string {
	fp, err := os.Open(filename)
	CheckError(err, 89)
	defer fp.Close()
	buf := bufio.NewReader(fp)
	var res string
	for {
		line, err := buf.ReadString('\n')
		if io.EOF == err {
			break
		}
		res += line
	}
	return res
}

func WriteBuff(DEFAULT_HTML string, conn net.Conn) {
	serve_time := time.Now()
	buffers := bytes.Buffer{}
	buffers.WriteString("HTTP/1.1 200 OK\r\n")
	buffers.WriteString("Server: Cyeam\r\n")
	buffers.WriteString("Date: " + serve_time.Format(time.RFC1123) + "\r\n")
	buffers.WriteString("Content-Type: text/html; charset=utf-8\r\n")
	buffers.WriteString("Content-length:" + fmt.Sprintf("%d", len(DEFAULT_HTML)) + "\r\n")
	buffers.WriteString("\r\n")
	buffers.WriteString(DEFAULT_HTML)
	//buffers.WriteString(buffers.String())
	conn.Write(buffers.Bytes())
}

//从父串中截取子串 支持中文字符串截取
func Substring(str string, begin, length int) string {
	// /index.php
	rs := []rune(str)

	lth := len(rs)
	if begin < 0 {
		begin = 0
	}
	if begin > lth {
		begin = lth
	}

	end := begin + length

	if end > lth || (length == -1) {
		end = lth
	}
	return string(rs[begin:end])
}

//获取文件扩展名
func GetFileTypeName(filename string, start string) string {
	startIndex := strings.Index(filename, start)
	startIndex++
	res := string([]byte(filename)[startIndex])
	return res
}

//查看文件是否存在

func PathExists(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}
	return false
}
