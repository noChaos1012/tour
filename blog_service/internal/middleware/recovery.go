package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/noChaos1012/tour/blog_service/global"
	"github.com/noChaos1012/tour/blog_service/pkg/Email"
	"github.com/noChaos1012/tour/blog_service/pkg/app"
	"github.com/noChaos1012/tour/blog_service/pkg/errcode"
	"time"
)

//异常捕获
func Recovery() gin.HandlerFunc {
	defaultMailer := Email.NewEmail(&Email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().Errorf("panic recover err:%v", err)

				emailErr := defaultMailer.SendMail(global.EmailSetting.To,
					fmt.Sprintf("异常抛出，发生时间:%d", time.Now().Unix()),
					fmt.Sprintf("异常信息:%v", err))
				if emailErr !=  nil{
					global.Logger.Panicf("mail.SendMail err:%v",err)
				}
				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
