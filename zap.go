package zap

import (
	"github.com/yyyThree/zap/helper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

const (
	EnvDebug      Env = "debug"
	EnvProduction Env = "production"
)

const (
	WriterFile   Writer = "file"
	WriterStdout Writer = "stdout"
	WriterRedis  Writer = "redis"
)

var loggers = make(map[string]*zap.SugaredLogger)

type Env string

type Writer string

type Config struct {
	Env    Env      // 运行环境 debug/production，默认为 production
	Writer Writer      // 日志写入方式 stdout/file/redis 默认为 file
	Redis  RedisConfig // redis 配置，Writer = redis 时必传
	LogDir string      // 日志存储文件夹，默认为 log
}

type Logger struct {
	logger *zap.SugaredLogger
}

func New(config ...Config) *Logger {
	c := Config{}
	if len(config) > 0 {
		c = config[0]
	}
	return &Logger{
		logger: getLogger(c),
	}
}

// 获取logger处理器
func getLogger(config Config) *zap.SugaredLogger {
	configMd5 := helper.StructToMd5(config)
	if logger, ok := loggers[configMd5]; ok {
		return logger
	}
	encoder := getEncoder()
	writeSyncer, err := getLogWriter(config) // 日志写入处理器
	if err != nil {
		panic("zap.getLogger error:" + err.Error())
	}
	core := zapcore.NewCore(encoder, writeSyncer, getLogLevel(config))
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(getStackLevel(config))).Sugar()
	loggers[configMd5] = logger
	return logger
}

// 编码器(如何写入日志)
func getEncoder() zapcore.Encoder {
	zapConfig := zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder, //将级别转换成大写
		TimeKey:     "time",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		CallerKey:     "caller",
		StacktraceKey: "stack",
		EncodeCaller:  zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1e6)
		},
	}
	return zapcore.NewJSONEncoder(zapConfig)
}

func isDebug(config Config) bool {
	return config.Env == EnvDebug
}

// 哪种级别及以上的日志会被写入
func getLogLevel(config Config) zapcore.Level {
	if isDebug(config) {
		return zapcore.DebugLevel
	}
	return zapcore.InfoLevel
}

// 哪种级别及以上的日志会写入堆栈信息
func getStackLevel(config Config) zapcore.Level {
	if isDebug(config) {
		return zapcore.DebugLevel
	}
	return zapcore.ErrorLevel
}