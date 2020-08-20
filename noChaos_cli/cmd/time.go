/**
时间的获取
1、获取当前时间
./noChaos_cli time now
输出结果:2020-08-19 22:51:53,1597848713

2、计算所需时间
./noChaos_cli time calc -c="2020-08-18 12:00:00" -d=23h
输出结果:2020-08-19 11:00:00,1597834800

*/
package cmd

import (
	"github.com/noChaos1012/tour/noChaos_cli/internal/timer"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
	"time"
)

var calculateTime string
var duration string

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calcutelateTimeCmd)

	calcutelateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", "需要计算的时间，有效单位为时间戳或已格式化的时间")
	calcutelateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", `持续时间，有效时间单位为"ns","us"(or "μs"),"ms","s","m","h"`)
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间（上海时间）",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("输出结果:%s,%d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix()) //输出格式化时间和时间戳
	},
}

var calcutelateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算所需时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2006-01-02 15:04:05"
		if calculateTime == "" {
			currentTimer = timer.GetNowTime()
		} else {
			var err error
			if !strings.Contains(calculateTime, " ") {
				layout = "2006-01-02"
			}
			currentTimer, err = time.Parse(layout, calculateTime)
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0)
			}
		}

		calculateTime, err := timer.GetCalculateTime(currentTimer, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculateTime err:%v", err)
		}
		log.Printf("输出结果:%s,%d", calculateTime.Format(layout), calculateTime.Unix())

	},
}
