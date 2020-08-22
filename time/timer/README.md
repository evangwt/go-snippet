## timer复用


### 使用场景

高频场景下，每个循环中需要独立的time.After做超时控制，如果创建的timer未超时，则在堆上不会被GC回收，造成内存一路飙高，那么考虑两个问题:

1. 及时释放未超时的Timer

  直接使用NewTimer代替After（After也是调用NewTimer），通过Stop()来停止Timer，供GC回收。

2. 高频场景，复用Timer
  
  虽然Stop()后会被GC回收，但是未触发GC之前可以通过Reset()复用。


### 复用的问题

如果复用前timer已经超时，timer.C这个channel有值，因为其Len为1，所以要先读timer.C，保证后面定时依然能写入，不然会阻塞。
```go
if !t.Stop() {
  select {
  case <-t.C: // try to drain from the channel
  default:
  }
}
```


### Benchmark

```shell
goos: darwin
goarch: amd64
pkg: go-snippet/time/timer
BenchmarkPool-12    	10932666	       104 ns/op	       0 B/op	       0 allocs/op
BenchmarkStd-12     	 5109447	       234 ns/op	     208 B/op	       3 allocs/op
PASS
ok  	go-snippet/time/timer	2.697s
```
