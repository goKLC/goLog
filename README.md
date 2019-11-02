# goLog
Logger for go projects

#### **Install**

`go get github.com/mkilic91/goLog`

#### **Usage**

    log, conf := goLog.New()
    

##### Config
    conf.Path = "logs"
    conf.FileName = "app.log"
    conf.TimeFormat = "2006-01-02 15:04:05"
    conf.PrintTerminal = true
    
Path = log folder e.g. myProject/logs

FileName = log file name

TimeFormat = log time format 

PritTerminal = log print for runtime terminal


##### Create Log

This example:

    log.Info("foo message", goLog.Data{"foo":"bar", "baz": "foo", "subData": goLog.Data{"subFoo": "subBar"}})
    log.Error("foo", goLog.Data{})
    
Output:

    [2019-11-03 00:19:00] Info: foo message
	        {
	          "baz": "foo",
	          "foo": "bar",
	          "subData": {
	            "subFoo": "subBar"
	            }
	         }
    [2019-11-03 00:19:00] Error: foo

