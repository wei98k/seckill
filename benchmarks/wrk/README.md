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