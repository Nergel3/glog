package glog

/*
 * Logger interface
 */
type Logger interface {
	Log(...interface{})
	Logf(string, ...interface{})
	Mute()
	Unmute()
	To(Output, string)
}
