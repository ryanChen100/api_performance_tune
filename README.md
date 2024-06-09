# api_performance_tune
 ```
 API優化測試
 ```

### 1.modidfy : cmd >>  config
### 2.run : cmd 

----

## 一般map存取
api_performance_tune\pkg\router\normal\normal_test.go

-  go test -run=Bench -bench=Benchmark
![alt text](result_pin/normal.png)

## 使用cache
api_performance_tune\pkg\router\gocache\gocache_test.go

-  go test -run=Bench -bench=Benchmark
![alt text](image.png)
![alt text](result_pin/gocache.png)


## 使用sync.Pool
api_performance_tune\pkg\router\syncpool\syncpool_test.go

- go test -run=Bench -bench=Benchmark
![alt text](result_pin/sync.Pool.png)

## 使用fasthttp
api_performance_tune\pkg\router\fast\fast_test.go

- go test -run=Bench -bench=Benchmark
![alt text](result_pin/fasthttp.png)

## 使用gzip
api_performance_tune\pkg\router\gzip\gzip_test.go

- go test -run=Bench -bench=Benchmark
![alt text](result_pin/gzip.png)