## time-format-alfred

一款使用了`github.com/noaway/dateparse`所提供的时间解析函数实现的alfred时间解析workflow，可以针对多种时间输入格式进行相应的format。


## download

1. `repository -> release -> 选择最新版本`


## quick start

1. 下载并安装workflow

2. 配置常用时区，格式为 `./time-format-alfred -time={query} UTC America/Los_Angeles`,直接在后面追加即可，使用逗号隔开

![](http://imgblog.mrdear.cn/1548854492.png?imageMogr2/thumbnail/!100p)

3. 输入时间`time now`，可以使用now代指当前时间

![](http://imgblog.mrdear.cn/1548854370.png?imageMogr2/thumbnail/!100p)

4. 输入时间`time 1548854618000`

![](http://imgblog.mrdear.cn/1548854650.png?imageMogr2/thumbnail/!100p)

5. 输入时间以及指定该时间所属时区`time 2019-01-30 21:24:44,gmt-7`,表示当前时间是GMT-7时区的时间

![](http://imgblog.mrdear.cn/1548854736.png?imageMogr2/thumbnail/!100p)
