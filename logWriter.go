package zap

import (
	redisClient "github.com/go-redis/redis/v8"
	"github.com/yyyThree/zap/helper"
	"github.com/yyyThree/zap/library/redis"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
)

type RedisConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Db       int
	RedisKey string
}

type redisWriter struct {
	config RedisConfig
	client *redisClient.Client
}

// 日志写入处理器
func getLogWriter(config Config) (zapcore.WriteSyncer, error) {
	switch config.Writer {
	case WriterStdout:
		return getStdoutLogWriter()
	case WriterRedis:
		return getRedisWriter(config.Redis)
	default: // 默认为文件存储
		return getFileLogWriter(config)
	}
}

// 屏幕输出
func getStdoutLogWriter() (zapcore.WriteSyncer, error) {
	return zapcore.AddSync(os.Stdout), nil
}

// redis 实现 o.Writer 接口
func (r *redisWriter) Write(b []byte) (int, error) {
	// 日志切割设置，按 年/月/日 切割
	k := r.config.RedisKey + "." + helper.FormatDateNow()
	s := string(b)
	s = strings.TrimRight(s, "\n")
	n, err := r.client.RPush(redis.GetCtx(), k, s).Result()
	return int(n), err
}

// redis写入
func getRedisWriter(config RedisConfig) (zapcore.WriteSyncer, error) {
	client, err := redis.GetConn(redis.Config{
		Host:     config.Host,
		Port:     config.Port,
		User:     config.User,
		Password: config.Password,
		Db:       config.Db,
	})
	if err != nil {
		return nil, err
	}
	return zapcore.AddSync(&redisWriter{
		config: config,
		client: client,
	}), nil
}

// 文件写入
func getFileLogWriter(config Config) (zapcore.WriteSyncer, error) {
	// 日志切割设置，按 年/月/日 切割
	if config.LogDir == "" {
		config.LogDir = "log"
	}
	file := strings.TrimRight(config.LogDir, "/") + "/" + helper.FormatDateNowBySlash() + ".log"
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file,  // 日志文件位置
		MaxSize:    500,   // 日志文件最大大小(MB)
		MaxBackups: 2,     // 保留旧文件最大数量
		MaxAge:     0,     // 保留旧文件最长天数
		Compress:   false, // 是否压缩旧文件
	}
	return zapcore.AddSync(lumberJackLogger), nil
}
