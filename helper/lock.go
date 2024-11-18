package helper

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// DistributedLock 结构体表示一个分布式锁
type DistributedLock struct {
	client *redis.Client
	key    string
	ttl    time.Duration
	ctx    context.Context
}

// NewDistributedLock 创建一个新的分布式锁实例
func NewDistributedLock(ctx context.Context, key string, ttl time.Duration) *DistributedLock {
	return &DistributedLock{
		client: RedisClient,
		key:    key,
		ttl:    ttl,
		ctx:    ctx,
	}
}

// Lock 尝试获取锁
func (dl *DistributedLock) Lock() (bool, error) {
	// 使用 SETNX 命令尝试设置锁
	locked, err := dl.client.SetNX(dl.ctx, dl.key, "1", dl.ttl).Result()
	if err != nil {
		return false, err
	}
	return locked, nil
}

// Unlock 释放锁
func (dl *DistributedLock) Unlock() error {
	// 使用 DEL 命令删除锁
	return dl.client.Del(dl.ctx, dl.key).Err()
}

// TryLockWithTimeout 尝试在指定时间内获取锁
func (dl *DistributedLock) TryLockWithTimeout(timeout time.Duration) (bool, error) {
	startTime := time.Now()
	for time.Since(startTime) < timeout {
		locked, err := dl.Lock()
		if err != nil {
			return false, err
		}
		if locked {
			return true, nil
		}
		time.Sleep(100 * time.Millisecond) // 等待一段时间后重试
	}
	return false, fmt.Errorf("timeout acquiring lock")
}
