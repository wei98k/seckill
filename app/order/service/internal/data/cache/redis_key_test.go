package cache

import "testing"

func TestRedisKey(t *testing.T)  {
	redisKey := &CreateRedisKey{}
	createKey(redisKey)

}