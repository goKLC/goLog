# goLog
Logger for go projects

#### **Install**

`go get github.com/mkilic91/goLog`

#### **Import**
   
    import "github.com/mkilic91/goLog"

#### **Usage**

    var log goLog.LogInterface
    var conf *goLog.Config
    log, conf = goLog.New()
    

##### Config
    conf.Path = "logs"
    conf.FileName = "app.log"
    conf.TimeFormat = "2006-01-02 15:04:05"
    conf.PrintTerminal = true
    conf.Partition = true
    conf.PartitionRange = goLog.DAILY
    
Path = log folder string e.g. myProject/logs

FileName = log file name string 

TimeFormat = log time format golang time format

PritTerminal = log print for runtime terminal true or false

Partition = log file partition mode true or false

PartitionRange = log file partition range HOURLY, DAILY, MONTHLY, and YEARLY


##### Create Log

This example:

    log.Info("foo message", map[string]interface{}{"foo":"bar", "baz": "foo", "subContext": map[string]interface{}{"subFoo": "subBar"}})
    log.Error("foo", nil)
    
Output:

    file path : logs/2019-11-5-app.log

    [2019-11-03 00:19:00] Info: foo message
	        {
	          "baz": "foo",
	          "foo": "bar",
	          "subContext": {
	            "subFoo": "subBar"
	            }
	         }
    [2019-11-03 00:19:00] Error: foo

