package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/tal-tech/go-zero/tools/goctl/api/spec"
	"io/ioutil"
)

type Parser struct {
	r  *bufio.Reader
	st string
}

func NewParser(filename string) (*Parser, error) {

	api, err := ioutil.ReadFile(filename) //读取文件内容
	if err != nil {
		return nil, err
	}

	info, body, service, err := MatchStruct(string(api)) //匹配结构体正则，获取到对应配置内容

	if err != nil {
		return nil, err
	}
	var buffer = new(bytes.Buffer)
	buffer.WriteString(info)
	buffer.WriteString(service)
	return &Parser{
		r:  bufio.NewReader(buffer),
		st: body,
	}, nil
}

func (p *Parser) Parse() (api *spec.ApiSpec, err error) {
	api = new(spec.ApiSpec)
	types, err := parseStructAst(p.st)	//解析结构体部分语法结构
	if err != nil {
		return nil, err
	}
	fmt.Println(types)

	/*
	api.Types = types
	var lineNumber = 1
	st := newRootState(p.r, &lineNumber)
	for {
		st, err = st.process(api)
		if err == io.EOF {
			return api, p.validate(api)
		}
		if err != nil {
			return nil, fmt.Errorf("near line: %d, %s", lineNumber, err.Error())
		}
		if st == nil {
			return api, p.validate(api)
		}
	}
	 */
	return nil,nil
}
