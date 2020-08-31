## logs
logs is a Go logs manager. It can use many logs adapters.


## How to install?

	go get github.com/ha666/logs


## What adapters are supported?

As of now this logs support console, file .


## How to use it?

First you must import it

```golang
import (
	"github.com/ha666/logs"
)
```

Then init a Log (example with console adapter)

```golang
logs.SetLogger(logs.AdapterConsole, `{"level":7}`)
```

> the first params stand for how many channel

Use it like this:

```golang
logs.Alert("alert")
logs.Error("error")
logs.Warning("warning")
logs.Warn("warn")
logs.Info("info")
logs.Debug("debug")
logs.Trace("trace")
```

## File adapter

Configure file adapter like this:

```golang
logs.SetLogger(logs.AdapterFile, `{"filename":"./log/log.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":100}`)
```

## Logging caller information (file name & line number)

The module can be configured to include the file & line number of the log calls in the logging output. This functionality is disabled by default, but can be enabled using the following code:
```golang
logs.EnableFuncCallDepth(true)
```
Use true to turn file & line number logging on, and false to turn it off. Default is false.

If your application encapsulates the call to the log methods, you may need use SetLogFuncCallDepth to set the number of stack frames to be skipped before the caller information is retrieved. The default is 2.
```golang
logs.SetLogFuncCallDepth(3)
```


