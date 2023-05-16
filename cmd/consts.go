package cmd

const help = `
gox是管理go依赖的cli工具，例如你可以通过 [gox add --lib=gorm] 这条命令来快速添加gorm依赖。

gox <command> <option>

Usage:

gox config set <库别名>=<库地址>    设置库别名和地址关系
gox config rm <库别名>              移除库别名
gox config ls                       关系列表
gox add --lib=<库别名>              快速添加依赖，多个依赖用英文逗号分割`
