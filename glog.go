package glog

type logger interface {
	Log(...interface{})
	Logf(string, ...interface{})
	Mute()
	Unmute()
	To(Output, string)
}
