# sandy_log
引包   
`
go get github.com/zyjblockchain/sandy_log/log
`

使用   
`
// 需要初始化
// args[0]： 代码打印级别
// args[1]: 是否保存到文件
// args[2]: 是否显示日志打印位置
func init() {
    log.Setup(log.LevelInfo, true, true)
}
`

