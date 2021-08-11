# FQA

## 2021年08月07日

#### 问题描述 http请求接口，发送消息队列写入kafka中报告错误: `context deadline exceeded`, 但是查看kafka数据是写入成功的。

http请求接口，发送消息队列写入kafka中报告错误: `context deadline exceeded`, 但是查看kafka数据是写入成功的。

#### 错误信息

```shell
INFO service.name= service.version= ts=2021-08-07T13:44:44+08:00 caller=order.go:11 msg=CreateOrder request: gid:1  amount:2
ERROR service.name= service.version= ts=2021-08-07T13:44:45+08:00 caller=order.go:37 msg=context deadline exceeded
INFO service.name= service.version= ts=2021-08-07T13:44:45+08:00 caller=tracing.go:44 kind=server component=http operation=/api.order.service.v1.Order/CreateOrder args=gid:1  amount:2 code=0 reason= stack= latency=1.004600849
```

#### 测试与分析

- 根据错误提示是context 超时了、
- 单独运行kafka-sdk是不会产生错误的， 那么就是框架里的使用方式不对了、

通过调整配置文件中http.timeout = 2s 的时候就不会出现超时, 说明kafka.send需要的发送时间比较长、需要排查golang服务器和kafka服务器之间的问题(kafka配置的问题或kratos接入的kafka-sdk有问题)

----------

[2021年08月08日15:05:42] 本地搭建kafka-qps显示正常那就与docker-kafka配置的关系

```shell
jw@jianwei wrk$ wrk -t4 -c10 -d5s -T2s --script=post.json.lua --latency http://127.0.0.1:8001/seckill/order
Running 5s test @ http://127.0.0.1:8001/seckill/order
  4 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     7.22ms   20.48ms 202.93ms   96.16%
    Req/Sec   560.23    135.44   717.00     87.82%
  Latency Distribution
     50%    3.10ms
     75%    3.75ms
     90%    5.61ms
     99%  133.76ms
  11077 requests in 5.02s, 1.15MB read
Requests/sec:   2208.23
Transfer/sec:    235.06KB
```