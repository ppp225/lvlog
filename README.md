# lvlog
Standard logger with levels. Uses log package internally, doesn't create its own logger instance, so it's possible to share flags etc. with std logger

# usage

see source code, it's simple enough

```go

import log "github.com/ppp225/lvlog"

func main() {
	log.Info("info msg")              // == std.Println(" [INFO]", "info msg")
	log.Infof("msg=%q", "info msg")   // == std.Printf("%s msg=%q", " [INFO]", "info msg")
	log.Debug("debug msg")            // won't be displayed
	log.SetLevel(log.ALL)             // default level is INFO and up
	log.Debug("debug msg")            // will be displayed now
	log.Fatalf("msg=%q", "fatal msg") // will be displayed and exit with status code 1
}
```

# benchmarks
```go
Benchmark_Off_baseline-12    	1000000000	     0.284 ns/op	       0 B/op	       0 allocs/op
Benchmark_Off_naive-12       	1000000000	     0.268 ns/op	       0 B/op	       0 allocs/op
Benchmark_Off_lvlog-12       	1000000000	     0.950 ns/op	       0 B/op	       0 allocs/op
Benchmark_Print_std-12       	 2201851	       526 ns/op	      48 B/op	       1 allocs/op
Benchmark_Print_lvlog-12     	 1646030	       826 ns/op	      96 B/op	       3 allocs/op
Benchmark_Printf_std-12      	 1120405	      1024 ns/op	      48 B/op	       1 allocs/op
Benchmark_Printf_lvlog-12    	 1132267	      1017 ns/op	      48 B/op	       1 allocs/op
```