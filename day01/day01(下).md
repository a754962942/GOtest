# 整型
整型分为两个大类	
按长度分为int8,int16,int32,int64分别对应uint8,uint16,uint32,uint64
#浮点型
float32和float64
#bool布尔型
	b1:=true
	var b2 bool 
	bool默认值为false
#字符串
Go语言中的字符串以猿声数据类型出现，使用字符串就像使用其他原生数据类型（int,bool,float32,float64等）一样。Go语言里的字符串的内部**实现使用UTF-8编码**，字符串的值为**双引号("")**中的内容，可以在Go语言的源码中直接添加ASCII码字符，例如
	s1:="hello"
	s2:="你好"
**GO语言中字符串使用双引号("")包裹的**

**GO语言中的单引号包裹的是字符**

	//字符串
	s:="hello "
	//单独的字母、汉字、符号表示一个字符
	c1:='h'
	c2:='1'
	c3:='啊'
	//字节：1字节=8Bit(8个二进制位)
	//1个字符'A'=1个字节
	//1个utf-8编码的汉字'啊'=一般占3个字节

>字符串转义符

>\r 回车符（返回行首）

>\n 换行符（直接跳到下一行的同列位置）

>\t 制表符

> \' 单引号

> \" 双引号

> \\ 反斜杠

多行字符串使用
例

	s:=`
		123
		234
		567
	`
##字符串操作
strings.Split(s,"")分割字符串

strings.Join(s,"")拼接字符串

strings.Contains(s,"")判断字符串是否包含

string.HasPrefix(s,"")判断字符串是否以·开始

strings.HasSuffix(s,"")判断字符串是否以·结束

strings.Index(s,"")取出字符在字符串的位置

strings.LastIndex(s,"")取出字符在字符串最后出现的位置

----------
#byte 和rune

byte（uint8）

rune（int32）

Go语言中为了除了非ASCII码的字符，定义了rune类型
可以通过rune类型来进行字符串的修改