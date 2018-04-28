package lredis

import (
	"fmt"
	"testing"
)

func Test_Send(t *testing.T) {
	conn, err := NewConn("182.92.158.94:6379","lifei")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	conn.Send("ZADD","kafka",123,"abc")
	conn.Send("ZADD","kafka",234,"bcd")
	conn.Send("ZADD","kafka",890,"def")
	ret_strings,err := conn.Strings(conn.Do("zrange", "kafka",0,-1))
	fmt.Println("strings----",ret_strings,err)

	ret_string_map,err := conn.StringMap(conn.Do("zrange","kafka",0,-1,"withscores"))
	fmt.Println("string_map----",ret_string_map,err)


	conn.Send("set","name","feige")

	ret_string,err := conn.String(conn.Do("get","name"))
	fmt.Println("string----",ret_string,err)

	conn.Send("set","age",30)

	ret_int,err := conn.Int(conn.Do("get","age"))
	fmt.Println("int----",ret_int,err)


	aa,e := conn.Do("exists","name")
	fmt.Println(string(aa.([]byte)),e)

	ret_bool,err := conn.Bool(conn.Do("exists","name"))
	fmt.Println("bool----",ret_bool)




}
