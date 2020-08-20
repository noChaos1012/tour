/**
放置单词格式转换的子命令word

测试方法：
	打包后，使用指令
	输入：./noChaos_cli word -s=wangsuchao -m=1
	输出：2020/08/19 21:30:26 输出结果:WANGSUCHAO
*/
package cmd

import (
	"github.com/noChaos1012/tour/noChaos_cli/internal/word"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

const (
	MODE_UPPER                        = iota + 1 //单词全部转换为大写
	MODE_LOWER                                   //单词全部转换为小写
	MODE_UNDERSCORE_TO_UPPERCAMELCASE            //下划线单词转换为大写驼峰单词
	MODE_UNDERSCORE_TO_LOWERCAMELCASE            //下划线单词转换为小写(首字母)驼峰单词
	MODE_CAMELCASE_TO_UNDERSCORE                 //驼峰单词转下划线单词 使用govalidator库的转换方法
)

var desc = strings.Join([]string{
	"该子命令支持各种单次格式转换，模式如下:",
	"1:全部单词转换为大写",
	"2:全部单词转换为小写",
	"3:下划线单词转换为大写驼峰单词",
	"4:下划线单词转换为小写驼峰单词",
	"5:驼峰单词转换成下划线单词",
}, "\n")

var mode int8  //输入模式
var str string //输入文字

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string

		switch mode {
		case MODE_UPPER:
			content = word.ToUpper(str)
		case MODE_LOWER:
			content = word.ToLower(str)
		case MODE_UNDERSCORE_TO_UPPERCAMELCASE:
			content = word.UnderscoreToUpperCamelCase(str)
		case MODE_UNDERSCORE_TO_LOWERCAMELCASE:
			content = word.UnderscoreToLowerCamelCase(str)
		case MODE_CAMELCASE_TO_UNDERSCORE:
			content = word.CamelCaseToUnderscore(str)

		default:
			log.Fatalf("站不支持该转换模式，请执行help word 查看帮助文档")
		}
		log.Printf("输出结果:%s", content)

	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的模式")
}
