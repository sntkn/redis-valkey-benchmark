# redis-valkey-benchmark

Benchmark testing was conducted using both Redis and Valkey libraries.

## Setup

```bash
docker compose up -d
```

## Benchmark

```bash

❯ go test -bench=. -benchmem
goos: darwin
goarch: arm64
pkg: sntkn/redis-valkey-benchmark
cpu: Apple M2
BenchmarkRedisSetValue-8            5006            238893 ns/op             279 B/op          9 allocs/op
BenchmarkRedisGetValue-8            4923            239439 ns/op             220 B/op          8 allocs/op
BenchmarkValkeySetValue-8           5202            216464 ns/op              24 B/op          1 allocs/op
BenchmarkValkeyGetValue-8           4987            235406 ns/op              16 B/op          1 allocs/op
PASS
ok      sntkn/redis-valkey-benchmark    5.776s
```
