package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/noaway/dateparse"
	"time"
	"time-format-alfred/model"
)

var (
	paramTime string
	help      bool
	icon      = model.Icon{
		Path: "./icon.png",
	}
	resultItems = model.Items{
		Items: make([]model.Item, 0, 3),
	}
)

func init() {
	flag.StringVar(&paramTime, "time", "", "时间信息，支持多种格式")
	flag.BoolVar(&help, "h", false, "this help")
}

func main() {
	flag.Parse()
	if help {
		flag.Usage()
		return
	}

	result, e := dateparse.ParseAny(paramTime)
	if e != nil {
		formatError(e)
		return
	}

	zoneArgs := []string{time.Local.String()}
	zoneArgs = append(zoneArgs, flag.Args()...)
	formatTimestamp(result.UnixNano(), zoneArgs)
}

// 错误信息输出
func formatError(e error) {
	item := model.Item{
		Uid:      "1",
		Title:    "无法解析该格式",
		Subtitle: e.Error(),
		Icon:     icon,
	}
	resultItems.Items = append(resultItems.Items, item)
	bytes, _ := json.Marshal(resultItems)
	fmt.Println(string(bytes))
}

// 按照指定时区输出
func formatTimestamp(timeNano int64, timeZones []string) {
	unix := time.Unix(convertSecond(timeNano), timeNano%1000000)
	for _, zone := range timeZones {
		loc, _ := time.LoadLocation(zone)
		result := unix.In(loc).Format("2006-01-02T15:04:05 -07:00 MST")
		item := model.Item{
			Uid:      "1",
			Title:    loc.String(),
			Subtitle: result,
			Icon:     icon,
		}
		resultItems.Items = append(resultItems.Items, item)
	}
	bytes, _ := json.Marshal(resultItems)
	fmt.Println(string(bytes))
}

func convertSecond(timeNano int64) int64 {
	return timeNano / int64(time.Second)
}
