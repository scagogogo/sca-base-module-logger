package sca_base_module_logger

// Debug uses fmt.Sprint to construct and log a message.
func Debug(args ...interface{}) {
	if Log != nil {
		Log.Debug(args...)
	}
}

// Info uses fmt.Sprint to construct and log a message.
func Info(args ...interface{}) {
	if Log != nil {
		Log.Info(args...)
	}
}

// Warn uses fmt.Sprint to construct and log a message.
func Warn(args ...interface{}) {
	if Log != nil {
		Log.Warn(args...)
	}
}

// Error uses fmt.Sprint to construct and log a message.
func Error(args ...interface{}) {
	if Log != nil {
		Log.Error(args...)
	}
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanic(args ...interface{}) {
	if Log != nil {
		Log.DPanic(args...)
	}
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func Panic(args ...interface{}) {
	if Log != nil {
		Log.Panic(args...)
	}
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func Fatal(args ...interface{}) {
	if Log != nil {
		Log.Fatal(args...)
	}
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(template string, args ...interface{}) {
	if Log != nil {
		Log.Debugf(template, args...)
	}
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...interface{}) {
	if Log != nil {
		Log.Infof(template, args...)
	}
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...interface{}) {
	if Log != nil {
		Log.Warnf(template, args...)
	}
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
	if Log != nil {
		Log.Errorf(template, args...)
	}
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanicf(template string, args ...interface{}) {
	if Log != nil {
		Log.DPanicf(template, args...)
	}
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func Panicf(template string, args ...interface{}) {
	if Log != nil {
		Log.Panicf(template, args...)
	}
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func Fatalf(template string, args ...interface{}) {
	if Log != nil {
		Log.Fatalf(template, args...)
	}
}

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//
//	s.With(keysAndValues).Debug(msg)
func Debugw(msg string, keysAndValues ...interface{}) {
	if Log != nil {
		Log.Debugw(msg, keysAndValues...)
	}
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Infow(msg string, keysAndValues ...interface{}) {
	if Log != nil {
		Log.Infow(msg, keysAndValues...)
	}
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Warnw(msg string, keysAndValues ...interface{}) {
	if Log != nil {
		Log.Warnw(msg, keysAndValues...)
	}
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Errorw(msg string, keysAndValues ...interface{}) {
	if Log != nil {
		Log.Errorw(msg, keysAndValues...)
	}
}

// DPanicw logs a message with some additional context. In development, the
// logger then panics. (See DPanicLevel for details.) The variadic key-value
// pairs are treated as they are in With.
func DPanicw(msg string, keysAndValues ...interface{}) {
	if Log != nil {
		Log.DPanicw(msg, keysAndValues...)
	}
}

// Panicw logs a message with some additional context, then panics. The
// variadic key-value pairs are treated as they are in With.
func Panicw(msg string, keysAndValues ...interface{}) {
	if Log != nil {
		Log.Panicw(msg, keysAndValues...)
	}
}

// Fatalw logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs are treated as they are in With.
func Fatalw(msg string, keysAndValues ...interface{}) {
	if Log != nil {
		Log.Fatalw(msg, keysAndValues...)
	}
}

// Debugln uses fmt.Sprintln to construct and log a message.
func Debugln(args ...interface{}) {
	if Log != nil {
		Log.Debugln(args...)
	}
}

// Infoln uses fmt.Sprintln to construct and log a message.
func Infoln(args ...interface{}) {
	if Log != nil {
		Log.Infoln(args...)
	}
}

// Warnln uses fmt.Sprintln to construct and log a message.
func Warnln(args ...interface{}) {
	if Log != nil {
		Log.Warnln(args...)
	}
}

// Errorln uses fmt.Sprintln to construct and log a message.
func Errorln(args ...interface{}) {
	if Log != nil {
		Log.Errorln(args...)
	}
}

// DPanicln uses fmt.Sprintln to construct and log a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanicln(args ...interface{}) {
	if Log != nil {
		Log.DPanicln(args...)
	}
}

// Panicln uses fmt.Sprintln to construct and log a message, then panics.
func Panicln(args ...interface{}) {
	if Log != nil {
		Log.Panicln(args...)
	}
}

// Fatalln uses fmt.Sprintln to construct and log a message, then calls os.Exit.
func Fatalln(args ...interface{}) {
	if Log != nil {
		Log.Fatalln(args...)
	}
}

// Sync flushes any buffered log entries.
func Sync() error {
	if Log != nil {
		return Log.Sync()
	} else {
		return nil
	}
}
