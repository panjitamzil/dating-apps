package repository

import (
	"context"
	"fmt"
	"reflect"
	"time"
)

func (r *Repo) SetKey(key, value string) error {
	timeDuration := time.Duration(86400) * time.Second
	err := r.redis.SetEX(context.Background(), key, value, timeDuration).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) Exist(key string) (int64, error) {
	result, err := r.redis.Exists(context.Background(), key).Result()
	if err != nil {
		return -1, err
	}

	return result, nil
}

func (r *Repo) GetKeys(key, email string) (int, error) {
	keyword := fmt.Sprintf("*%s*", email)

	result, err := r.redis.Do(context.Background(), "KEYS", keyword).Result()
	if err != nil {
		return -1, err
	}

	v := reflect.ValueOf(result)

	return v.Len(), nil
}
