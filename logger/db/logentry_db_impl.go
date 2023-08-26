package db

type LogEntryDBImpl struct{}

func NewUserDBImpl() *LogEntryDBImpl {
	return nil
}

func (l *LogEntryDBImpl) CreateLogEntry(logEntry LogEntry) error {
	return nil
}
