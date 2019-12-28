# goLog
Logger for go projects

#### **Install**

    go get github.com/goKLC/goLog

#### **Import**
   
    import "github.com/goKLC/goLog"

#### **Usage**

    var log goLog.LogInterface
    var conf *goLog.Config
    log, conf = goLog.New()
    
    fileHandler := goLog.NewFileHandler()
    fileHandler.Partition = true
    fileHandler.PartitionRange = goLog.DAILY
    fileHandler.Path = "log"
    
    terminalHandler := goLog.NewTerminalHandler()
    
    redisClient := redis.NewClient(&redis.Options{
    		Addr:               "127.0.0.1:6379",
    		Password:           "",
    		DB:                 0,
    	})
    
    redisHandler := goLog.NewRedisHandler()
    redisHandler.Client = redisClient
    redisHandler.Key = "goLog"
    
	db, _ := gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/go_log?charset=utf8&parseTime=True")
    dbHandler := goLog.NewDBHandler(db)
    
    emailHandler := goLog.NewEmailHandler()
    emailHandler.Levels = []goLog.Level{goLog.CRITICAL, goLog.EMERGENCY, goLog.ERROR}
    emailHandler.Host = "smtp.mailtrap.io:25"
    emailHandler.Auth = smtp.PlainAuth("", "username", "password", "smtp.mailtrap.io")
    emailHandler.From = "info@goklc.org"
    emailHandler.To = []string{"admin@goklc.org", "mkilic@goklc.org"}
    	
    config.AddHandler(fileHandler)
    config.AddHandler(terminalHandler)
    config.AddHandler(redisHandler)
    config.AddHandler(dbHandler)
    config.AddHandler(emailHandler)
    
_*RedisHandler use `github.com/go-redis/redis`_

_*DBHandler use `github.com/jinzhu/gorm`_

##### Config
    conf.TimeFormat = "2006-01-02 15:04:05"
    
TimeFormat = log time format golang time format


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

