package contype

import (
  "testing"
)

func Benchmark_contype(b *testing.B) {
  for i := 0; i < b.N; i++ {
    FileType("https://kampsy.com/apps")
  }
}

func Benchmark_timers(b *testing.B) {
  b.StopTimer()

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    FileType("build.golang.org/log.json")
  }
}
