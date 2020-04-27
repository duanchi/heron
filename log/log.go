package log

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.heurd.com/heron-go/heron/types/config"
	"io"
	"os"
	"strings"
	"time"
)

type Logrus struct {
	std *logrus.Logger
	enabled bool
}

func (this *Logrus) Enabled(status bool) {
	this.enabled = status
}

func (this *Logrus) GetEnabled() bool {
	return this.enabled
}

func (this *Logrus) StandardLogger() *logrus.Logger {
	return this.std
}

// SetOutput sets the standard logger output.
func (this *Logrus) SetOutput(out io.Writer) {
	this.std.SetOutput(out)
}

// SetFormatter sets the standard logger formatter.
func (this *Logrus) SetFormatter(formatter logrus.Formatter) {
	this.std.SetFormatter(formatter)
}

// SetReportCaller sets whether the standard logger will include the calling
// method as a field.
func (this *Logrus) SetReportCaller(include bool) {
	this.std.SetReportCaller(include)
}

// SetLevel sets the standard logger level.
func (this *Logrus) SetLevel(level logrus.Level) {
	this.std.SetLevel(level)
}

// GetLevel returns the standard logger level.
func (this *Logrus) GetLevel() logrus.Level {
	return this.std.GetLevel()
}

// IsLevelEnabled checks if the log level of the standard logger is greater than the level param
func (this *Logrus) IsLevelEnabled(level logrus.Level) bool {
	return this.std.IsLevelEnabled(level)
}

// AddHook adds a hook to the standard logger hooks.
func (this *Logrus) AddHook(hook logrus.Hook) {
	this.std.AddHook(hook)
}

// WithError creates an entry from the standard logger and adds an error to it, using the value defined in ErrorKey as key.
func (this *Logrus) WithError(err error) *logrus.Entry {
	return this.std.WithField(logrus.ErrorKey, err)
}

// WithContext creates an entry from the standard logger and adds a context to it.
func (this *Logrus) WithContext(ctx context.Context) *logrus.Entry {
	return this.std.WithContext(ctx)
}

// WithField creates an entry from the standard logger and adds a field to
// it. If you want multiple fields, use `WithFields`.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func (this *Logrus) WithField(key string, value interface{}) *logrus.Entry {
	return this.std.WithField(key, value)
}

// WithFields creates an entry from the standard logger and adds multiple
// fields to it. This is simply a helper for `WithField`, invoking it
// once for each field.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func (this *Logrus) WithFields(fields map[string]interface{}) *logrus.Entry {
	return this.std.WithFields(fields)
}

// WithTime creats an entry from the standard logger and overrides the time of
// logs generated with it.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func (this *Logrus) WithTime(t time.Time) *logrus.Entry {
	return this.std.WithTime(t)
}

// Trace logs a message at level Trace on the standard logger.
func (this *Logrus) Trace(args ...interface{}) {
	if this.enabled {
		this.std.Trace(args...)
	}
}

// Debug logs a message at level Debug on the standard logger.
func (this *Logrus) Debug(args ...interface{}) {
	if this.enabled {
		this.std.Debug(args...)
	}
}

// Print logs a message at level Info on the standard logger.
func (this *Logrus) Print(args ...interface{}) {
	if this.enabled {
		this.std.Print(args...)
	}
}

// Info logs a message at level Info on the standard logger.
func (this *Logrus) Info(args ...interface{}) {
	if this.enabled {
		this.std.Info(args...)
	}
}

// Warn logs a message at level Warn on the standard logger.
func (this *Logrus) Warn(args ...interface{}) {
	if this.enabled {
		this.std.Warn(args...)
	}
}

// Warning logs a message at level Warn on the standard logger.
func (this *Logrus) Warning(args ...interface{}) {
	if this.enabled {
		this.std.Warning(args...)
	}
}

// Error logs a message at level Error on the standard logger.
func (this *Logrus) Error(args ...interface{}) {
	if this.enabled {
		this.std.Error(args...)
	}
}

// Panic logs a message at level Panic on the standard logger.
func (this *Logrus) Panic(args ...interface{}) {
	if this.enabled {
		this.std.Panic(args...)
	}
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func (this *Logrus) Fatal(args ...interface{}) {
	if this.enabled {
		this.std.Fatal(args...)
	}
}

// Tracef logs a message at level Trace on the standard logger.
func (this *Logrus) Tracef(format string, args ...interface{}) {
	if this.enabled {
		this.std.Tracef(format, args...)
	}
}

// Debugf logs a message at level Debug on the standard logger.
func (this *Logrus) Debugf(format string, args ...interface{}) {
	if this.enabled {
		this.std.Debugf(format, args...)
	}
}

// Printf logs a message at level Info on the standard logger.
func (this *Logrus) Printf(format string, args ...interface{}) {
	if this.enabled {
		this.std.Printf(format, args...)
	}
}

// Infof logs a message at level Info on the standard logger.
func (this *Logrus) Infof(format string, args ...interface{}) {
	if this.enabled {
		this.std.Infof(format, args...)
	}
}

// Warnf logs a message at level Warn on the standard logger.
func (this *Logrus) Warnf(format string, args ...interface{}) {
	if this.enabled {
		this.std.Warnf(format, args...)
	}
}

// Warningf logs a message at level Warn on the standard logger.
func (this *Logrus) Warningf(format string, args ...interface{}) {
	if this.enabled {
		this.std.Warningf(format, args...)
	}
}

// Errorf logs a message at level Error on the standard logger.
func (this *Logrus) Errorf(format string, args ...interface{}) {
	if this.enabled {
		this.std.Errorf(format, args...)
	}
}

// Panicf logs a message at level Panic on the standard logger.
func (this *Logrus) Panicf(format string, args ...interface{}) {
	if this.enabled {
		this.std.Panicf(format, args...)
	}
}

// Fatalf logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func (this *Logrus) Fatalf(format string, args ...interface{}) {
	if this.enabled {
		this.std.Fatalf(format, args...)
	}
}

// Traceln logs a message at level Trace on the standard logger.
func (this *Logrus) Traceln(args ...interface{}) {
	if this.enabled {
		this.std.Traceln(args...)
	}
}

// Debugln logs a message at level Debug on the standard logger.
func (this *Logrus) Debugln(args ...interface{}) {
	if this.enabled {
		this.std.Debugln(args...)
	}
}

// Println logs a message at level Info on the standard logger.
func (this *Logrus) Println(args ...interface{}) {
	if this.enabled {
		this.std.Println(args...)
	}
}

// Infoln logs a message at level Info on the standard logger.
func (this *Logrus) Infoln(args ...interface{}) {
	if this.enabled {
		this.std.Infoln(args...)
	}
}

// Warnln logs a message at level Warn on the standard logger.
func (this *Logrus) Warnln(args ...interface{}) {
	if this.enabled {
		this.std.Warnln(args...)
	}
}

// Warningln logs a message at level Warn on the standard logger.
func (this *Logrus) Warningln(args ...interface{}) {
	if this.enabled {
		this.std.Warningln(args...)
	}
}

// Errorln logs a message at level Error on the standard logger.
func (this *Logrus) Errorln(args ...interface{}) {
	if this.enabled {
		this.std.Errorln(args...)
	}
}

// Panicln logs a message at level Panic on the standard logger.
func (this *Logrus) Panicln(args ...interface{}) {
	if this.enabled {
		this.std.Panicln(args...)
	}
}

// Fatalln logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func (this *Logrus) Fatalln(args ...interface{}) {
	if this.enabled {
		this.std.Fatalln(args...)
	}
}

type Fields map[string]interface{}

var Log = Logrus{
	std: logrus.New(),
	enabled: true,
}

func Init (config config.Log) {

	if config.Output == "" {
		config.Output = "stdout://"
	}

	outputStack := strings.Split(config.Output, "://")

	switch outputStack[0] {
	case "stdout":
	default:
		Log.std.Out = os.Stdout
	}

	switch config.Level {
	case "debug":
		Log.SetLevel(logrus.DebugLevel)
	case "warn":
		Log.SetLevel(logrus.WarnLevel)
	case "error":
		Log.SetLevel(logrus.ErrorLevel)
	case "fatal":
		Log.SetLevel(logrus.FatalLevel)
	case "panic":
		Log.SetLevel(logrus.PanicLevel)
	case "info":
	default:
		Log.SetLevel(logrus.InfoLevel)
	}

	switch config.Format.Type {
	case "json":
		Log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: config.Format.Timestamp,
			DisableTimestamp: config.Timestamp,
			PrettyPrint: config.Format.Json.Pretty,
		})
	case "text":
	case "plain":
		Log.SetFormatter(&logrus.TextFormatter{
			DisableColors: config.Format.Text.Colors,
			TimestampFormat: config.Format.Timestamp,
			FullTimestamp: config.Format.Text.FullTimestamp,
			DisableTimestamp: config.Timestamp,
		})
	}


}