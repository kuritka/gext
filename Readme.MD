# Go Extensions

**gext** extends standard go libraries for commonly used programming patterns and utilities

### log


```go
package main

import "github.com/kuritka/gext/log"

var logger = log.Log

func main(){
	logger.Panic().Msg("panic!")
}
```

```bash
{"level":"panic","time":"2020-03-02T16:29:20+01:00","caller":"/home/michal/go/src/gext/main.go:8","message":"panic!"}
```