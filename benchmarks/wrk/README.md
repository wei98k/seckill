`wrk -t4 -c10 -d5s -T1s --script=post.json.lua --latency http://127.0.0.1:8001/seckill/order`

```bigquery
jianwei:wrk jw$ wrk -t4 -c10 -d5s -T1s --script=post.json.lua --latency http://127.0.0.1:8001/seckill/order
Running 5s test @ http://127.0.0.1:8001/seckill/order
  4 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    61.84ms   70.50ms 883.00ms   96.22%
    Req/Sec    37.34     11.03    70.00     69.23%
  Latency Distribution
     50%   48.13ms
     75%   62.52ms
     90%   82.04ms
     99%  409.75ms
  741 requests in 5.05s, 78.88KB read
Requests/sec:    146.79
Transfer/sec:     15.63KB
```

- 时间：2021年08月07日10:05:47
- 阶段：v1

```bigquery
jw@jianwei wrk$ wrk -t4 -c10 -d5s -T1s --script=post.json.lua --latency http://127.0.0.1:8001/seckill/order
Running 5s test @ http://127.0.0.1:8001/seckill/order
  4 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   506.40ms  526.09ms 999.36ms  100.00%
    Req/Sec     1.40      0.60     3.00     65.00%
  Latency Distribution
     50%  997.67ms
     75%  998.54ms
     90%  999.36ms
     99%  999.36ms
  44 requests in 5.08s, 4.68KB read
  Socket errors: connect 0, read 0, write 0, timeout 36
Requests/sec:      8.67
Transfer/sec:      0.92KB
```

- 时间：2021年08月07日10:06:29
- 阶段: v1

```bigquery
jw@jianwei wrk$ wrk -t4 -c10 -d5s -T1s --script=post.json.lua --latency http://127.0.0.1:8001/seckill/order
Running 5s test @ http://127.0.0.1:8001/seckill/order
  4 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     3.77ms    2.59ms  11.28ms   86.11%
    Req/Sec     3.38      5.97    22.00     90.48%
  Latency Distribution
     50%    2.79ms
     75%    3.55ms
     90%    9.73ms
     99%   11.28ms
  76 requests in 5.06s, 8.09KB read
  Socket errors: connect 0, read 0, write 0, timeout 40
Requests/sec:     15.03
Transfer/sec:      1.60KB
```