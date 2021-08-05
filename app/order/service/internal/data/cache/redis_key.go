package cache

import "fmt"

// user:userid:9:username
// order:oid:12:sn
// article:aid:999:views

// KeyPrefix 前缀接口
type KeyPrefix interface {
	expire()
	pre()
}

// KeyPrefixRepo 前缀抽象类
type KeyPrefixRepo struct {
}

func (KeyPrefixRepo) expire() {
	fmt.Println("set redis key expire")
}

func (KeyPrefixRepo) pre()  {
}

func createKey(keyPrefix KeyPrefix)  {
	keyPrefix.expire()
	keyPrefix.pre()
}

type CreateRedisKey struct {
	KeyPrefixRepo
}

func (*CreateRedisKey) pre()  {
	fmt.Println("set order expire")
}