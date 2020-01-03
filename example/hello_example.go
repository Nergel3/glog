package example

import "github.com/nergel3/glog"

func ExampleHello() {
	i := glog.INFO
	i.Log("Hello World")
	// Output :
	// 2020-01-01 00:00:00 +0000 UTC [INFO] - Hello World
}
