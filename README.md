# lredis

花了一下午时间读了下redisgo的源码，除了连接池那里还没有彻底搞明白，其它大概明白了它的实现方式。
为了练手，写了这个小例子，只是简单的实现了功能，漏洞很多，容错也很差，后面会继续更新。

安装:

	go get github.com/jfeige/lredis
	
使用:


	conn, err := NewConn("127.0.0.1:6379","lifei")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
  
  
  
  
                
方法列表:

	Send(command string, args ...interface{})error
	
	Do(command string,args ...interface{})(replay interface{},err error){
	
	String(replay interface{},err error)(string,error){
	
	Strings(replay interface{},err error)([]string,error){
	
	StringMap(replay interface{},err error)(map[string]string,error){
	
	Bool(replay interface{},err error)(bool,error){
	
	Int(replay interface{},err error)(int,error){
  
  ......
