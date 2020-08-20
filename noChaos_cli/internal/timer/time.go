package timer

import (
	"time"
)

func GetNowTime() time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai") //设置时区
	return time.Now().In(location)
}

//获取推算时间
func GetCalculateTime(currentTime time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d) //获取持续时间 格式"3h23m3s12ms"
	if err != nil {
		return time.Time{}, nil
	}
	return currentTime.Add(duration), nil
}
