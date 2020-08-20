/**
放置单词格式转换的子命令word

测试方法：
	打包后，使用指令
	输入：./noChaos_cli word -s=wangsuchao -m=1
	输出：2020/08/19 21:30:26 输出结果:WANGSUCHAO
*/
package word

import (
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
