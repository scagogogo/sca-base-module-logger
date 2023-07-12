package logger

import sca_base_module_logger "github.com/scagogogo/sca-base-module-logger"

// Debug uses fmt.Sprint to construct and log a message.
func Debug(args ...interface{}) {
	sca_base_module_logger.Log.Debug(args...)
}

// Info uses fmt.Sprint to construct and log a message.
func Info(args ...interface{}) {
	sca_base_module_logger.Log.Info(args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func Warn(args ...interface{}) {
	sca_base_module_logger.Log.Warn(args...)
}

// Error uses fmt.Sprint to construct and log a message.
func Error(args ...interface{}) {
	sca_base_module_logger.Log.Error(args...)
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanic(args ...interface{}) {
	sca_base_module_logger.Log.DPanic(args...)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func Panic(args ...interface{}) {
	sca_base_module_logger.Log.Panic(args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func Fatal(args ...interface{}) {
	sca_base_module_logger.Log.Fatal(args...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(template string, args ...interface{}) {
	sca_base_module_logger.Log.Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...interface{}) {
	sca_base_module_logger.Log.Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...interface{}) {
	sca_base_module_logger.Log.Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
	sca_base_module_logger.Log.Errorf(template, args...)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanicf(template string, args ...interface{}) {
	sca_base_module_logger.Log.DPanicf(template, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func Panicf(template string, args ...interface{}) {
	sca_base_module_logger.Log.Panicf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func Fatalf(template string, args ...interface{}) {
	sca_base_module_logger.Log.Fatalf(template, args...)
}

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//
//	s.With(keysAndValues).Debug(msg)
func Debugw(msg string, keysAndValues ...interface{}) {
	sca_base_module_logger.Log.Debugw(msg, keysAndValues...)
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Infow(msg string, keysAndValues ...interface{}) {
	sca_base_module_logger.Log.Infow(msg, keysAndValues...)
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Warnw(msg string, keysAndValues ...interface{}) {
	sca_base_module_logger.Log.Warnw(msg, keysAndValues...)
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Errorw(msg string, keysAndValues ...interface{}) {
	sca_base_module_logger.Log.Errorw(msg, keysAndValues...)
}

// DPanicw logs a message with some additional context. In development, the
// logger then panics. (See DPanicLevel for details.) The variadic key-value
// pairs are treated as they are in With.
func DPanicw(msg string, keysAndValues ...interface{}) {
	sca_base_module_logger.Log.DPanicw(msg, keysAndValues...)
}

// Panicw logs a message with some additional context, then panics. The
// variadic key-value pairs are treated as they are in With.
func Panicw(msg string, keysAndValues ...interface{}) {
	sca_base_module_logger.Log.Panicw(msg, keysAndValues...)
}

// Fatalw logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs are treated as they are in With.
func Fatalw(msg string, keysAndValues ...interface{}) {
	sca_base_module_logger.Log.Fatalw(msg, keysAndValues...)
}

// Debugln uses fmt.Sprintln to construct and log a message.
func Debugln(args ...interface{}) {
	sca_base_module_logger.Log.Debugln(args...)
}

// Infoln uses fmt.Sprintln to construct and log a message.
func Infoln(args ...interface{}) {
	sca_base_module_logger.Log.Infoln(args...)
}

// Warnln uses fmt.Sprintln to construct and log a message.
func Warnln(args ...interface{}) {
	sca_base_module_logger.Log.Warnln(args...)
}

// Errorln uses fmt.Sprintln to construct and log a message.
func Errorln(args ...interface{}) {
	sca_base_module_logger.Log.Errorln(args...)
}

// DPanicln uses fmt.Sprintln to construct and log a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanicln(args ...interface{}) {
	sca_base_module_logger.Log.DPanicln(args...)
}

// Panicln uses fmt.Sprintln to construct and log a message, then panics.
func Panicln(args ...interface{}) {
	sca_base_module_logger.Log.Panicln(args...)
}

// Fatalln uses fmt.Sprintln to construct and log a message, then calls os.Exit.
func Fatalln(args ...interface{}) {
	sca_base_module_logger.Log.Fatalln(args...)
}

// Sync flushes any buffered log entries.
func Sync() error {
	return sca_base_module_logger.Log.Sync()
}
