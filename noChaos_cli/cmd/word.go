/**
放置单词格式转换的子命令word
*/
package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"strings"
	"unicode"
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
			content = ToUpper(str)
		case MODE_LOWER:
			content = ToLower(str)
		case MODE_UNDERSCORE_TO_UPPERCAMELCASE:
			content = UnderscoreToUpperCamelCase(str)
		case MODE_UNDERSCORE_TO_LOWERCAMELCASE:
			content = UnderscoreToLowerCamelCase(str)
		case MODE_CAMELCASE_TO_UNDERSCORE:
			content = CamelCaseToUnderscore(str)

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

//单词全部转换为大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

//单词全部转换为小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

//下划线单词转换为大写驼峰单词
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

//下划线单词转换为小写(首字母)驼峰单词
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

//驼峰单词转下划线单词 使用govalidator库的转换方法
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}
