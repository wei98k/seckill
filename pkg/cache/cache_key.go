package cache

// KeyPrefix 前缀接口
type KeyPrefix interface {
	expire(t int) int
	pre(s []interface{}) string
}

// KeyPrefixRepo 前缀抽象类
type KeyPrefixRepo struct {
}

func (KeyPrefixRepo) expire(t int) int {
	return t
}

func (KeyPrefixRepo) pre(s []interface{}) string {
	return "default"
}
