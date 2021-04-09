# zap

*基于zap日志库编写的日志系统，支持标准输出、文件写入、redis写入，日志将以年/月/日自动做切割存储。*

## 一、安装与引用

```go
go get github.com/yyyThree/zap@v1.0.1
import "github.com/yyyThree/zap"
```

## 二、使用说明

1. 快速使用

    ```go
    logger := zap.New()
    logger.Warn("test_msg", zap.BaseMap{"test": 1})
    ```

2. 标准输出
    ```go
    logger := zap.New(zap.Config{
        Env:    zap.EnvDebug,
        Writer: zap.WriterStdout,
    })
    logger.Debug("test_msg", zap.BaseMap{"test": 2})
    ```

3. 文件写入
    ```go
    logger := zap.New(zap.Config{
        Env:    zap.EnvDebug,
        Writer: zap.WriterFile,
        LogDir: "log/"
    })
    logger.Debug("test_msg", zap.BaseMap{"test": 2})
    ```

4. `redis`写入
    ```go
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
    ```

5. 配置项（`zap.Config`）说明
    1. Env：运行环境，debug/production，默认为 production。
       
        *debug模式运行时所有日志及堆栈信息均会被写入，production模式运行时仅Info及以上级别的日志会被写入，仅Error及以上的日志会记录堆栈信息。*
    2. Writer：日志写入方式 stdout/file/redis 默认为 file
    3. Redis：redis 配置，Writer = redis 时必传
        1. Host
        2. Port
        3. User（非必须）
        4. Password（非必须）
        5. Db（非必须）
        6. RedisKey：redis列表写入的key前缀
    4. LogDir：日志存储文件夹地址，默认为 log
    
6. 多种日志记录方法
    1. Debug
    2. Info
    3. Warn
    4. Error
    5. Panic
    6. Fatal（会导致os.Exit）
    
