package common

import (
	"context"
	"github.com/Woringsuhang/user/global"

	"fmt"
	"github.com/redis/go-redis/v9"

	"time"
)

func withClint(hand func(cli *redis.Client) error) error {

	cli := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.ConfigAll.Redis.Address, global.ConfigAll.Redis.Ip),
		DB:   0,
	})
	defer cli.Close()

	err := hand(cli)
	if err != nil {
		return err
	}

	return nil
}

func GetByKey(ctx context.Context, key string) (string, error) {
	var data string
	var err error

	err = withClint(func(cli *redis.Client) error {
		data, err = cli.Get(ctx, key).Result()
		return err
	})
	if err != nil {
		return "", err
	}
	return data, nil
}

func ExistKey(ctx context.Context, key string) (bool, error) {
	var data int64
	var err error

	err = withClint(func(cli *redis.Client) error {
		data, err = cli.Exists(ctx, key).Result()
		return err
	})
	if err != nil {
		return false, err
	}
	if data > 0 {
		return true, nil
	}
	return false, nil
}

func SetKey(ctx context.Context, key string, val interface{}, duration time.Duration) error {
	return withClint(func(cli *redis.Client) error {
		return cli.Set(ctx, key, val, duration).Err()
	})
}
