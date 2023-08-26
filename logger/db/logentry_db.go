package db

type LogEntryDB interface {
	CreateLogEntry(logEntry LogEntry) error
}
