package time_ext

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeFormat(t *testing.T) {
	now := time.Now()
	format := "yyyy-MM-dd HH:mm:ss"
	value := FormatTime(now, format)
	fmt.Printf("now of format %v is %v\n", format, value)

	format = "yy-M-d H:m:s"
	value = FormatTime(now, format)
	fmt.Printf("now of format %v is %v\n", format, value)

	format = "今天天气好晴朗，当前时间是：yyyy年MM月dd日 HH时mm分ss秒"
	value = FormatTime(now, format)
	fmt.Printf("now of format %v is %v\n", format, value)
}

func TestTimeParse(t *testing.T) {
	value := "2017-05-30 15:30:59"
	format := "yyyy-MM-dd HH:mm:ss"
	time, _ := Parse(value, format)
	fmt.Println(time)
}
