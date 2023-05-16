package prop

import (
	"fmt"
	"github.com/kiririx/krutils/confx"
	"github.com/kiririx/krutils/logx"
	"os"
	"os/user"
)

var homeDir = ""

var Props map[string]string

func init() {
	// 获取当前用户的账户信息
	currentUser, err := user.Current()
	if err != nil {
		logx.ERR(err)
		return
	}

	// 获取用户目录路径
	homeDir = currentUser.HomeDir

	// 判断文件是否存在
	if _, err := os.Stat(PropPath()); os.IsNotExist(err) {
		// 创建文件
		err := os.MkdirAll(homeDir+"/.gox", os.ModePerm)
		if err != nil {
			return
		}
		file, err := os.Create(PropPath())
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
	}

	Props, err = confx.ResolveProperties(PropPath())
	if err != nil {
		logx.ERR(err)
		return
	}
}

func PropPath() string {
	return homeDir + "/.gox/gox.properties"
}
