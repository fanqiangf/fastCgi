package phpCgi

import (
	"fastCgi/common"
	"fmt"
	"os"
	"os/exec"
)

//实现php-cgi（暂时用cgi程序去解析）
func PhpCgi(filename string, params string) string {
	phpCgi := common.GetRead("php-cgi")
	cmd := exec.Command(phpCgi, filename)
	out, err := cmd.CombinedOutput()
	os.Setenv("QUERY_STRING", params)
	if err != nil {
		fmt.Println(err)
	}
	return string(out)
}
