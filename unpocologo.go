package unpocologo

import (
	"fmt"
	"log"

	"github.com/fatih/color"
)

type level int

const (
	milestone level = iota
	info
	warn
	softError
	hardError
)

func (l level) color(event string) string {
	switch l {
	case milestone:
		return color.GreenString(event)
	case info:
		return color.CyanString(event)
	case warn:
		return color.YellowString(event)
	case softError:
		return color.MagentaString(event)
	case hardError:
		return color.RedString(event)
	default:
		return event
	}
}

func formatPrefix(prefix string) string {
	if prefix == "" {
		return ""
	}
	return fmt.Sprintf("[%s] ", prefix)
}

func write(prefix string, level level, event, detail string) {
	prefix = formatPrefix(prefix)
	event = level.color(event)

	if detail == "" {
		log.Print(prefix + event)
		return
	}
	log.Print(prefix + event + " | " + detail)
}

func SetFlags(flags int) {
	log.SetFlags(flags)
}

func Milestone(event, detail string) { write("", milestone, event, detail) }
func Info(event, detail string)      { write("", info, event, detail) }
func Warn(event, detail string)      { write("", warn, event, detail) }
func SoftError(event, detail string) { write("", softError, event, detail) }
func HardError(event, detail string) { write("", hardError, event, detail) }

func Milestonef(event, format string, v ...any) {
	write("", milestone, event, fmt.Sprintf(format, v...))
}
func Infof(event, format string, v ...any) { write("", info, event, fmt.Sprintf(format, v...)) }
func Warnf(event, format string, v ...any) { write("", warn, event, fmt.Sprintf(format, v...)) }
func SoftErrorf(event, format string, v ...any) {
	write("", softError, event, fmt.Sprintf(format, v...))
}
func HardErrorf(event, format string, v ...any) {
	write("", hardError, event, fmt.Sprintf(format, v...))
}

type Logger struct {
	prefix string
}

func New(prefix string) *Logger {
	return &Logger{prefix: prefix}
}

func (l *Logger) Milestone(event, detail string) { write(l.prefix, milestone, event, detail) }
func (l *Logger) Info(event, detail string)      { write(l.prefix, info, event, detail) }
func (l *Logger) Warn(event, detail string)      { write(l.prefix, warn, event, detail) }
func (l *Logger) SoftError(event, detail string) { write(l.prefix, softError, event, detail) }
func (l *Logger) HardError(event, detail string) { write(l.prefix, hardError, event, detail) }

func (l *Logger) Milestonef(event, format string, v ...any) {
	write(l.prefix, milestone, event, fmt.Sprintf(format, v...))
}
func (l *Logger) Infof(event, format string, v ...any) {
	write(l.prefix, info, event, fmt.Sprintf(format, v...))
}
func (l *Logger) Warnf(event, format string, v ...any) {
	write(l.prefix, warn, event, fmt.Sprintf(format, v...))
}
func (l *Logger) SoftErrorf(event, format string, v ...any) {
	write(l.prefix, softError, event, fmt.Sprintf(format, v...))
}
func (l *Logger) HardErrorf(event, format string, v ...any) {
	write(l.prefix, hardError, event, fmt.Sprintf(format, v...))
}
