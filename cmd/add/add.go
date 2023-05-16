package add

import (
	"errors"
	"fmt"
	"github.com/kiririx/krutils/logx"
	"github.com/spf13/cobra"
	"kiririx/gox/cmd/prop"
	"os/exec"
	"strings"
)

var libs string

func init() {
	CmdAdd.PersistentFlags().StringVar(&libs, "lib", "", "这是你要添加的库的别名，别名通过 [gox config set 库别名=库地址] 设置")
}

var CmdAdd = &cobra.Command{
	Use: "add",
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, lib := range strings.Split(libs, ",") {
			libVal := prop.Props[lib]
			if libVal == "" {
				return errors.New(lib + "不存在，你可以用 [gox config set " + lib + "=库地址] 这种方式去配置")
			}

			for _, v := range strings.Split(libVal, ",") {
				logx.INFO("sh: " + "go get -u " + v)
				cmd := exec.Command("go", "get", "-u", v)
				output, err := cmd.CombinedOutput()
				if err != nil {
					return err
				}
				fmt.Println(string(output))
			}
		}
		return nil
	},
}
