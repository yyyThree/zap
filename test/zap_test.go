package test

import (
	"github.com/yyyThree/zap"
	"testing"
)

// 测试标准输出
func TestLogStdout(t *testing.T) {
	logger := zap.New(zap.Config{
		Env:    zap.EnvDebug,
		Writer: zap.WriterStdout,
	})
	logger.Debug("test_msg", zap.BaseMap{"test": 2})
}

// 测试文件日志写入
func TestLogFile(t *testing.T) {
	logger := zap.New()
	logger.Error("test_msg", zap.BaseMap{"test": 1})
}

// 测试redis日志写入
func TestLogRedis(t *testing.T) {
	logger := zap.New(zap.Config{
		Env:    zap.EnvDebug,
		Writer: zap.WriterRedis,
		Redis: zap.RedisConfig{
			Host:     "127.0.0.1",
			Port:     8003,
			RedisKey: "go-test",
		},
	})
	logger.Debug("test_msg", zap.BaseMap{"test": 2})
}
