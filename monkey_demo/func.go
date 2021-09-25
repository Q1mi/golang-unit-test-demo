package monkey_demo

import (
	"fmt"
	"varys"
)

// 假设你们公司中台提供了一个用户中心的库：varys
// 使用这个库可以很方便的根据uid获取用户相关信息


func MyFunc(uid int64)string{
	u, err := varys.GetInfoByUID(uid)
	if err != nil {
		return "welcome"
	}

	// 这里是一些逻辑代码...
fmt.Printf("-->%#v\n", u)
	return fmt.Sprintf("hello %s\n", u.Name)
}
