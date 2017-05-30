package time_ext

import (
	"fmt"
	"time"
)

// time.Time wrapper
type TimeWrapper struct {
	time.Time
}

// 一个日期字符串 value 生成 time.Time
// format 仅支持正常的 年: yy/yyyy  月: M/MM 日: d/dd 小时: H/HH h/hh 分钟: m/mm 秒钟: s/ss 和其他字符的组合模式
// 不支持类似于 yyyyy  MMMM 非常规格式， 这些可能会有非预期结果哟~
func Parse(value, format string) (time.Time, error) {
	layout := get_std_format(format)
	return time.Parse(layout, value)
}

// TimeWrapper 生成格式串
func (t TimeWrapper) Format(format string) string {
	layout := get_std_format(format)
	return t.Time.Format(layout)
}

// time.Time 生成格式字符串
func FormatTime(t time.Time, format string) string {
	return TimeWrapper{t}.Format(format)
}

func inner_parse(format string) string {
	format = format + "\u0000"
	var inner_format []byte = make([]byte, 0, 14)
	var i, j int

	for ; j < len(format); j++ {
		if format[j] == format[i] {
			continue
		}
		if format[i] == 'y' || format[i] == 'M' || format[i] == 'd' || format[i] == 'H' || format[i] == 'h' || format[i] == 'm' || format[i] == 's' {
			inner_format = append(inner_format, format[i])
			inner_format = append(inner_format, fmt.Sprintf("%d", j-i)...)
		} else {
			inner_format = append(inner_format, format[i:j]...)
		}
		i = j
	}
	return string(inner_format)
}

// 将通用的格式串譬如 yyyy-MM-dd HH:mm:ss 转换为 go 标准 layout： 2016-01-02 15:04:05
// 转换过程分两步：
// (1) 将 yyyyMMddHHmmss 转换为 y4M2d2H2m2s2
// (2) 将 y4M2d2H2m2s2 转换为 20160102150405
func get_std_format(format string) string {
	var f = inner_parse(format) // 获取 yyyyMMddHHmmss 格式 --> y4M2d2H2m2s2
	var inner_format []byte = make([]byte, 0, 14)

	base := map[byte]int{
		'y': 2006, 'M': 01, 'd': 02, 'H': 15, 'h': 03, 'm': 04, 's': 05,
	}
	pow := map[byte]int{
		'1': 10, '2': 100, '3': 1000, '4': 10000,
	}

	// 从2016年01月02号15点04分05秒时间中按照 y4M2d2H2m2s2 生成诸如 20160102150405 的字符串
	for i := 0; i < len(f); i++ {
		switch f[i] {
		case 'y', 'M', 'd', 'h', 'm', 's':
			inner_format = append(inner_format, fmt.Sprintf("%0"+string(f[i+1])+"d", base[f[i]]%pow[f[i+1]])...)
			i++
		case 'H':
			inner_format = append(inner_format, fmt.Sprintf("%d", base[f[i]])...)
			i++
		default:
			inner_format = append(inner_format, f[i])
		}
	}
	return string(inner_format)
}
