package logger

import (
	"io"
	"log"
	"os"

	"github.com/lucasbonilla/freterapido-api/internal/ports"
)

type Adapter struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewAdapter(config ports.Config) *Adapter {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, config.RunType(), log.Ldate|log.Ltime)

	return &Adapter{
		debug:   log.New(writer, "DEBUG: ", logger.Flags()),
		info:    log.New(writer, "INFO: ", logger.Flags()),
		warning: log.New(writer, "WARNING: ", logger.Flags()),
		err:     log.New(writer, "ERROR: ", logger.Flags()),
		writer:  writer,
	}
}

func (lA *Adapter) Debug(value ...interface{}) {
	lA.debug.Println(value...)
}

func (lA *Adapter) Info(value ...interface{}) {
	lA.info.Println(value...)
}

func (lA *Adapter) Warning(value ...interface{}) {
	lA.warning.Println(value...)
}

func (lA *Adapter) Error(value ...interface{}) {
	lA.err.Println(value...)
}

func (lA *Adapter) Debugf(format string, v ...interface{}) {
	lA.debug.Printf(format, v...)
}

func (lA *Adapter) Infof(format string, v ...interface{}) {
	lA.info.Printf(format, v...)
}

func (lA *Adapter) Warningf(format string, v ...interface{}) {
	lA.warning.Printf(format, v...)
}

func (lA *Adapter) Errorf(format string, v ...interface{}) {
	lA.err.Printf(format, v...)
}
