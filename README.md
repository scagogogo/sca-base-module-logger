# SCA Base Module Logger

```bash
go get -u github.com/scagogogo/sca-base-module-logger
```

配置文件中设置：

```yaml
# 日志相关参数配置
logger:
  # 是否自动初始化日志模块
  auto-init: true
  # 日志的开启级别
  level: info
  # 日志文件的输出目录
  directory: "./logs"
  # 日志开启哪些端
  enable:
    # 是否开启标准输出
    stdout: true
    # 是否开启文件输出
    file: true
```

对于同一个选项在配置文件和环境变量同时设置的话，以环境变量优先



