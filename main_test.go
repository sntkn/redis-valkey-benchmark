package main

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/valkey-io/valkey-go"
)

var redisCli *redis.Client
var valkeyCli valkey.Client

func TestMain(m *testing.M) {
	// 共通のセットアップ
	ctx := context.Background()
	redisCli = NewRedisClient(ctx)
	valkeyCli = NewValkeyClient()

	// テスト実行
	code := m.Run()

	// 共通のクリーンアップ
	redisCli.Close()
	valkeyCli.Close()

	// 終了コードを返す
	os.Exit(code)
}

func BenchmarkRedisSetValue(b *testing.B) {
	ctx := context.Background()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("%s-%d", "bench_key1", i)
		if err := redisCli.Set(ctx, key, "bench_value", 0).Err(); err != nil {
			b.Fatalf("Failed to set value: %v", err)
		}
	}
}

func BenchmarkRedisGetValue(b *testing.B) {
	ctx := context.Background()
	if err := redisCli.Set(ctx, "bench_key1", "bench_value", 0).Err(); err != nil {
		b.Fatalf("Failed to set value: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if _, err := redisCli.Get(ctx, "bench_key1").Result(); err != nil {
			b.Fatalf("Failed to get value: %v", err)
		}
	}
}

func BenchmarkValkeySetValue(b *testing.B) {
	ctx := context.Background()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("%s-%d", "bench_key2", i)
		if err := valkeyCli.Do(ctx, valkeyCli.B().Set().Key(key).Value("bench_value").Build()).Error(); err != nil {
			b.Fatalf("Failed to set value: %v", err)
		}
	}
}

func BenchmarkValkeyGetValue(b *testing.B) {
	ctx := context.Background()
	if err := valkeyCli.Do(ctx, valkeyCli.B().Set().Key("bench_key2").Value("bench_value").Build()).Error(); err != nil {
		b.Fatalf("Failed to set value: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if _, err := valkeyCli.Do(ctx, valkeyCli.B().Get().Key("bench_key2").Build()).ToString(); err != nil {
			b.Fatalf("Failed to get value: %v", err)
		}
	}
}
