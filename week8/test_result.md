* 使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。

## 测试10
* ./redis-benchmark  -t set,get -n 10000 -d 10 -q
* SET: 142857.14 requests per second
* GET: 149253.73 requests per second

## 测试20
*  ./redis-benchmark  -t set,get -n 10000 -d 20 -q
* SET: 109890.11 requests per second
* GET: 138888.89 requests per second


## 测试30
* ./redis-benchmark  -t set,get -n 10000 -d 30 -q
* SET: 117647.05 requests per second
* GET: 100000.00 requests per second


## 测试50
* ./redis-benchmark  -t set,get -n 10000 -d 50 -q
* SET: 133333.33 requests per second
* GET: 144927.55 requests per second

## 测试100
* ./redis-benchmark  -t set,get -n 10000 -d 100 -q
* SET: 133333.33 requests per second
* GET: 140845.06 requests per second

## 测试200
* ./redis-benchmark  -t set,get -n 10000 -d 200 -q
* SET: 107526.88 requests per second
* GET: 108695.65 requests per second


## 测试1k
* ./redis-benchmark  -t set,get -n 10000 -d 1000 -q
* SET: 128205.13 requests per second
* GET: 113636.37 requests per second


## 测试5k
* ./redis-benchmark  -t set,get -n 10000 -d 5000 -q
* SET: 105263.16 requests per second
* GET: 108695.65 requests per second

## 测试10k
* ./redis-benchmark  -t set,get -n 10000 -d 10000 -q
* SET: 101010.10 requests per second
* GET: 104166.66 requests per second

## 测试50k
* ./redis-benchmark  -t set,get -n 10000 -d 50000 -q
* SET: 40983.61 requests per second
* GET: 42553.19 requests per second

## 测试100k
* ./redis-benchmark  -t set,get -n 10000 -d 100000 -q
* SET: 19762.85 requests per second
* GET: 15174.51 requests per second


# 结论：key在一定范围内变化的话 性能影响不大，当key 变大，超过某个值，性能会开始下降
# 5k下降了 10-20%左右  继续增大，性能下降更明显