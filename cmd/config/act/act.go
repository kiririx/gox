package act

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"kiririx/gox/cmd/prop"
	"os"
	"strings"
)

var CmdLs = &cobra.Command{
	Use: "ls",
	RunE: func(cmd *cobra.Command, args []string) error {
		for k, v := range prop.Props {
			fmt.Println(k + "=" + v)
		}
		return nil
	},
}

var CmdRm = &cobra.Command{
	Use: "rm",
	RunE: func(cmd *cobra.Command, args []string) error {
		k := strings.Split(args[0], "=")[0]
		if k == "" {
			return errors.New("要删除的库别名不存在")
		}
		delete(prop.Props, k)
		err := writeToFile()
		if err != nil {
			return err
		}
		return nil
	},
}

var CmdSet = &cobra.Command{
	Use: "set",
	RunE: func(cmd *cobra.Command, args []string) error {
		k := strings.Split(args[0], "=")[0]
		v := strings.Split(args[0], "=")[1]
		if k == "" || v == "" {
			return errors.New("格式错误，正确用法:  [gox config set 库别名=库地址], 如果库地址有多个，可以用英文逗号分隔")
		}
		prop.Props[k] = v
		err := writeToFile()
		if err != nil {
			return err
		}
		return nil
	},
}

func writeToFile() error {
	file, err := os.Create(prop.PropPath())
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for key, value := range prop.Props {
		line := key + "=" + value + "\n"
		_, err := writer.WriteString(line)
		if err != nil {
			return err
		}
	}
	return nil
}
