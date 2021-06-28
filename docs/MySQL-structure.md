# mysql结构设计

## Users (用户表)

| 名称 | 类型 | 描述 |
|:---- |:-----|:----|
| id  | bigint(20)| |
| username | varchar(20) | 用户名 |
| phone | char(11) | 手机号 |
| email | varchar(50) | 邮箱 |
| nickname | varchar(20) | 昵称 |
| password_hash | char(32) | 密码 |
| user_type | int(1) | 1 普通用户 |
| created_at | date | 创建时间 |
| updated_at | date | 更新时间 |

## Goods (商品表)

| 名称 | 类型 | 描述 |
|:---- |:-----|:----|
| id  | bigint(20)| |
| title | varchar(200) | 商品标题 |
| intro | varchar(500) | 商品简介 |
| details | bigtext | 商品详情内容 |
| category | bigint(20) | 商品分类 |
| brand | bigint(20) | 商品品牌 |
| price | double(10, 2) | 商品价格 |
| stock | int(10) | 库存 |
| created_at | date | 创建时间 |
| updated_at | date | 更新时间 |

## Orders (订单表)

| 名称 | 类型 | 描述 |
|:---- |:-----|:----|
| id  | bigint(20)| |
| sn | char(50) | 单号 eee202101221232222 |
| gid | bigint(20) | 商品ID goods_id |
| uid | bigint(20) | 用户ID users_i |
| stat | int(1) | 订单状态 1 创建成功 2 支付成功 3 申请取消 | 
| created_at | date | 创建时间 |
| updated_at | date | 更新时间 |

## OrderGoods (订单冗余表-订单详细表)

## OrderSeckill (秒杀订单表)

## GoodsCategory (商品分类表)

## Category (分类表)

## log (日志表)
