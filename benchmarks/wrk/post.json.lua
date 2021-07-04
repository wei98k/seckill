-- 设置请求方式
wrk.method = "POST"
-- json string
wrk.body = '{"gid": 1}'
-- 设置content-type
wrk.headers["Content-Type"] = "application/json"