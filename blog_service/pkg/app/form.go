package app

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	val "github.com/go-playground/validator/v10"
	"strings"
)

type ValidError struct { //成为errors的对象
	Key     string
	Message string
}

func (v *ValidError) Error() string {
	return v.Message
}

type ValidErrors []*ValidError //成为errors的对象

func (vs ValidErrors) Errors() []string {
	var errs []string
	for _, err := range vs {
		errs = append(errs, err.Error())
	}
	return errs
}

func (vs ValidErrors) Error() string {
	return strings.Join(vs.Errors(), ",")
}

func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
	var errs ValidErrors
	err := c.ShouldBind(v)	//绑定、校验入参
	if err != nil {
		v := c.Value("trans")
		trans, _ := v.(ut.Translator)
		verrs, ok := err.(val.ValidationErrors)
		if !ok {
			return true, nil
		}
		for key, value := range verrs.Translate(trans) {
			errs = append(errs, &ValidError{
				Key:     key,
				Message: value,
			})
		}
		return true, errs
	}
	return false, nil

}
