package benchmark

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/ppp225/lvlog"
)

const (
	msg = "example, not too short log message"
)

func Benchmark_Off_baseline(b *testing.B) {
	for n := 0; n < b.N; n++ {
	}
}

func Benchmark_Off_naive(b *testing.B) {
	debug := false

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if debug {
			log.Println(" [INFO]", msg)
		}
	}
	b.StopTimer()
}

func Benchmark_Off_lvlog(b *testing.B) {
	lvlog.SetLevel(lvlog.NONE)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		lvlog.Info(msg)
	}
	b.StopTimer()
}

func Benchmark_Print_std(b *testing.B) {
	log.SetOutput(ioutil.Discard)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		log.Println(" [INFO]", msg)
	}
	b.StopTimer()
}

func Benchmark_Print_lvlog(b *testing.B) {
	lvlog.SetOutput(ioutil.Discard)
	lvlog.SetLevel(lvlog.ALL)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		lvlog.Info(msg)
	}
	b.StopTimer()
}

func Benchmark_Printf_std(b *testing.B) {
	log.SetOutput(ioutil.Discard)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		log.Printf("%s msg=%q", " [INFO]", msg)
	}
	b.StopTimer()
}

func Benchmark_Printf_lvlog(b *testing.B) {
	lvlog.SetOutput(ioutil.Discard)
	lvlog.SetLevel(lvlog.ALL)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		lvlog.Infof("msg=%q", msg)

	}
	b.StopTimer()
}
