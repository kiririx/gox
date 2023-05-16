package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/kiririx/krutils/confx"
	"github.com/kiririx/krutils/logx"
	"os"
	"os/user"
	"strings"
)

var homeDir = ""

var props map[string]string

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
	if _, err := os.Stat(propPath()); os.IsNotExist(err) {
		// 创建文件
		err := os.MkdirAll(homeDir+"/.gox", os.ModePerm)
		if err != nil {
			return
		}
		file, err := os.Create(propPath())
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
	}

	props, err = confx.ResolveProperties(propPath())
	if err != nil {
		logx.ERR(err)
		return
	}

}

func propPath() string {
	return homeDir + "/.gox/gox.properties"
}

func create() {

}

func add(libs string) {

	logx.INFO(libs)
}

func goxSet(prop string) error {
	k := strings.Split(prop, "=")[0]
	v := strings.Split(prop, "=")[1]
	if k == "" || v == "" {
		return errors.New("the value to be set is incomplete")
	}
	props[k] = v
	err := writeToFile()
	if err != nil {
		return err
	}
	return nil
}

func writeToFile() error {
	file, err := os.Create(propPath())
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for key, value := range props {
		line := key + "=" + value + "\n"
		_, err := writer.WriteString(line)
		if err != nil {
			return err
		}
	}
	return nil
}

func goxRm(prop string) error {
	k := strings.Split(prop, "=")[0]
	if k == "" {
		return errors.New("the value to be set is incomplete")
	}
	delete(props, k)
	err := writeToFile()
	if err != nil {
		return err
	}
	return nil
}

func goxLs() error {
	for k, v := range props {
		fmt.Println(k + "=" + v)
	}
	return nil
}
