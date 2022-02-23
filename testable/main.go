package main

import "github.com/op/go-logging"

// 编写可测试的代码

func main() {

	//app := &App{}
	//app.Start()

	var log = logging.MustGetLogger("example")
	app := NewApp(log)
	app.Start()

}
