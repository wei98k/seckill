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

----------