package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
)

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() *Logger {
	return &Logger{e}
}

func init() {
	log.Println("Init Logger")

	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
		DisableColors: false,
		FullTimestamp: true,
	}

	l.SetOutput(ioutil.Discard)

	l.AddHook(&writer.Hook{
		Writer:    os.Stdout,
		LogLevels: logrus.AllLevels,
	})

	l.SetLevel(logrus.TraceLevel)

	e = logrus.NewEntry(l)
}
