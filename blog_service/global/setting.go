package global

import (
	"github.com/noChaos1012/tour/blog_service/pkg/logger"
	"github.com/noChaos1012/tour/blog_service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSettingS
	Logger          *logger.Logger
)
