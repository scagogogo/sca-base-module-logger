package sca_base_module_logger

import (
	"github.com/golang-infrastructure/go-pointer"
	project_root_directory "github.com/golang-infrastructure/go-project-root-directory"
	sca_base_module_config "github.com/scagogogo/sca-base-module-config"
	"github.com/scagogogo/sca-base-module-logger/rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// ScaAutoInitLoggerEnvName 此环境变量控制是否自动初始化logger模块，true自动初始化，false不自动初始化需要手动调用init初始化
var ScaAutoInitLoggerEnvName = "SCA_AUTO_INIT_LOGGER"

// ScaLoggerDirectoryEnvName 此环境变量控制日志文件存储的位置，注意优先级如果在配置文件中指定了的话则此处无效
var ScaLoggerDirectoryEnvName = "SCA_LOGGER_DIRECTORY"

// DefaultLoggerDirectory 默认情况下日志文件夹放在项目根目录下的这个文件夹中
const DefaultLoggerDirectory = "./logs"

var Log *zap.SugaredLogger

func init() {

	// 配置文件明确指定的优先级更高
	if !(sca_base_module_config.Config != nil && pointer.FromPointerOrDefault(sca_base_module_config.Config.Logger.AutoInit, true)) {
		return
	}

	// 然后是看环境变量
	if !sca_base_module_config.GetEnvBoolOrDefault(ScaAutoInitLoggerEnvName, true) {
		return
	}

	// 解析存放日志的目录
	var logDirectory string
	// 配置文件的优先级最高
	if sca_base_module_config.Config != nil && sca_base_module_config.Config.Logger.Directory != "" {
		logDirectory = sca_base_module_config.Config.Logger.Directory
	} else if p, b := os.LookupEnv(ScaLoggerDirectoryEnvName); b {
		// 然后是环境变量
		logDirectory = p
	} else {
		// 最后实在不行，就还是默认的路径
		logDirectory = DefaultLoggerDirectory
	}

	// 如果是相对路径的话，认为是相对于项目根目录的路径
	if !filepath.IsAbs(logDirectory) {
		projectRootDirectory, err := project_root_directory.GetRootDirectory()
		if err != nil {
			panic(err)
		}
		logDirectory = filepath.Join(projectRootDirectory, logDirectory)
	}
	err := Init(logDirectory)
	if err != nil {
		panic(err)
	}

}

// Init 调用初始化日志方法
func Init(logDirectory string) error {
	// 设置一些基本日志格式 具体含义还比较好理解，直接看zap源码也不难懂
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey: "msg",
		LevelKey:   "level",
		//EncodeLevel: zapcore.CapitalLevelEncoder,
		EncodeLevel: zapcore.LowercaseColorLevelEncoder,
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("[2006-01-02 15:04:05]"))
		},
		CallerKey:    "fileUtil",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})

	// 实现两个判断日志等级的interface
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel
	})

	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	infoWriterSlice := make([]zapcore.WriteSyncer, 0)
	errorWriterSlice := make([]zapcore.WriteSyncer, 0)

	// 标准输出
	if isEnableStdout() {
		infoWriterSlice = append(infoWriterSlice, zapcore.AddSync(os.Stdout))
		//errorWriterSlice = append(errorWriterSlice, zapcore.AddSync(os.Stdout))
	}

	// 文件输出
	if isEnableFileOutput() {

		// 获取 info、error日志文件的io.Writer 抽象 getWriter() 在下方实现
		infoWriter, err := getWriter(filepath.Join(logDirectory, "/info.log"))
		if err != nil {
			return err
		}
		infoWriterSlice = append(infoWriterSlice, zapcore.AddSync(infoWriter))

		errorWriter, err := getWriter(filepath.Join(logDirectory, "/error.log"))
		if err != nil {
			return err
		}
		errorWriterSlice = append(errorWriterSlice, zapcore.AddSync(errorWriter))
	}

	// 最后创建具体的Logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(infoWriterSlice...), infoLevel),
		zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(errorWriterSlice...), errorLevel),
	)

	log := zap.New(core, zap.AddCaller()) // 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数, 有点小坑
	Log = log.Sugar()
	return nil
}

const ScaLoggerEnableStdoutEnvName = "SCA_LOGGER_ENABLE_STDOUT"

func isEnableStdout() bool {

	// 配置文件
	if sca_base_module_config.Config != nil && sca_base_module_config.Config.Logger.Enable.Stdout != nil {
		return pointer.FromPointer(sca_base_module_config.Config.Logger.Enable.Stdout)
	}

	// 环境变量
	b := sca_base_module_config.GetEnvBool(ScaLoggerEnableStdoutEnvName)
	if b != nil {
		return pointer.FromPointer(b)
	}

	return true
}

const ScaLoggerEnableFileEnvName = "SCA_LOGGER_ENABLE_FILE"

func isEnableFileOutput() bool {

	// 配置文件
	if sca_base_module_config.Config != nil && sca_base_module_config.Config.Logger.Enable.File != nil {
		return pointer.FromPointer(sca_base_module_config.Config.Logger.Enable.File)
	}

	// 环境变量
	b := sca_base_module_config.GetEnvBool(ScaLoggerEnableFileEnvName)
	if b != nil {
		return pointer.FromPointer(b)
	}

	return true
}

func getWriter(filename string) (io.Writer, error) {
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// info.log是指向最新日志的链接
	return rotatelogs.New(
		strings.Replace(filename, ".log", "", -1)+"-%Y%m%d.log", // 没有使用go风格反人类的format格式
		//rotatelogs.WithLinkName(filename),
		// 每天分割一次日志
		rotatelogs.WithRotationTime(time.Hour*24),
		// 每个日志最打512M
		rotatelogs.WithRotationSize(1024*1024*512),
		// 保存7天内的日志
		rotatelogs.WithMaxAge(-1),
		// 保留最近多少个日志
		//rotatelogs.WithRotationCount(uint(config.GetIntFromEnvOrConfigFile("LOGGER_KEEP_LATEST_N_LOG_FILE", "logger.keep-latest-n-log-file", 3))),
		rotatelogs.WithRotationCount(10),
	)
}
