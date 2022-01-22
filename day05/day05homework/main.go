package main
/*
1.把课上读文件的三种方式写一遍
2.把课上写文件的三种方式和copyfile函数写一遍
3.写一个日志库
logger.Trace()
logger.debug()
logger.warning()
logger.Info()
logger.Error("日志的内容")

写日志

fmt.fprintf("")

需求：
可以往不同的输出位置记录日志
日志分为五种级别

接口：日志可以输出到终端，也可以输出到文件中
文件操作
*/
type Db struct{
	name string
}
func (d Db)Trace(){
	//查询所有日志
}
func (d Db)Debug(){
	//级别是debug（错误存储位置）
}
func (d Db)Warning(){
	
}


func main() {
	
}