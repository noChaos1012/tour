package main

import (
	"github.com/noChaos1012/tour/noChaos_cli/cmd"
	"log"
)

func main() {
	err := cmd.Execute() //启用cmd根指令
	if err != nil {
		log.Fatalf("cmd.Execute error:%v", err)
	}
}

/**
type Name string

//实现Value的接口就可以算作Value类型
func (i *Name) String() string {
	return fmt.Sprint(*i)
}

func (i *Name) Set(value string) error {
	if len(*i) > 0 {
		return errors.New("name flage already set")
	}
	*i = Name("noChaos_custom_struct-" + value)
	return nil
}

func main() {
	var name Name
	goCmd := flag.NewFlagSet("go", flag.ExitOnError) 		//解析失败退出程序
	goCmd.Var(&name, "name", "帮助信息")           	//使用自定义结构体接收数据
	flag.Parse()

	args := flag.Args()
	switch args[0] {

	case "go":
		_ = goCmd.Parse(args[1:]) //./noChaos_cli go -name=nochaos

	default:
		name = "type undefined" //无法识别是传入参数
	}

	log.Printf("name:%s", name)
}
*/
