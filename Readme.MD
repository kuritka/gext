# Go Extensions

**gext** extends standard go libraries for commonly used programming patterns and utilities

### concurency
premise implementation and event aggregator

### data

#### Union 
Union merge two maps.If existing map and altered map contains
same key than value of existing is replaced by value of altered
```go
union := extensions.Union(existingMap, alteredMap)
```

### date

### env
reading string from environment variable
```go
env.MustGetStringFlagFromEnv(envLabels)
```


### guard
guardians, panics or throw errors. It is bad practice to panic within library packages, return error instead 
```go
err := s.ListenAndServeTLS(certPath, keyPath)
guards.FailOnError(err, "server stopped")
```

```go
if contentType := r.Header.Get("Content-Type"); contentType != httphead.ContentType.ApplicationJSON {
		guards.HttpThrowError(w, http.StatusInternalServerError, "Content-Type=%s, expect application/json", contentType)
		return
}
```

### httphead


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

### parser
Used for parsing basic types values 
```go
.ToString()
.ToFloat64()
.ToInt()
...
```

### rand
Random numbers and guids


### reflection
Currently modifying structures. Add items to slices etc..

