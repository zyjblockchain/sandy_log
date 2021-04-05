# sandy_log

引包   
```
go get github.com/zyjblockchain/sandy_log/log
```

使用   
```golang
// 需要初始化
// args[0]: module 名称
// args[1]: 代码打印级别
// args[2]: 是否保存到文件
// args[3]: 是否显示日志打印位置
func init() {
    log.Setup("default", log.LevelInfo, true, true)
}
```

