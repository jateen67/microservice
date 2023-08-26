package db

type LogEntryDB interface {
	InsertLogEntry(logEntry LogEntry) error
}
