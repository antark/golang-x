##### golang扩展

###### 时间处理
<sup><sub>1. 支持 yyyy-MM-dd HH:mm:ss 格式的日期处理</sub></sup><br/>

<code><pre>
    now := time.Now()

    format := "yyyy-MM-dd HH:mm:ss"

    value := FormatTime(now, format)

    fmt.Printf("now of format %v is %v\n", format, value)
</pre></code>

<code><pre>
    value := "2017-05-30 15:30:59"

    format := "yyyy-MM-dd HH:mm:ss"

    time, _ := Parse(value, format)
</pre></code>


###### 排序
<sup><sub>1. </sub></sup><br/>

###### 字符编码转换
<sup><sub>1. </sub></sup><br/>

###### 加密处理
<sup><sub>1. </sub></sup><br/>