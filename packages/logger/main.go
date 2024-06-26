package logger

import (
	"fmt"
	"os"
	"time"
)

func getFileName(channel string) string {
	return fmt.Sprintf("logs/%s-%d-%d-%d.txt", channel, time.Now().Year(), time.Now().Month(), time.Now().Day())
}

func openFile(fileName string) (*os.File, error) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return file, nil
}

type EventLogger struct {
	file *os.File
}

func CreateEventLogger(channel string) (*EventLogger, error) {
	fileName := getFileName(channel)
	file, err := openFile(fileName)

	if err != nil {
		return nil, err
	}

	eventLogger := &EventLogger{
		file: file,
	}

	return eventLogger, nil
}

func (eventLogger *EventLogger) LogEvent(channel, message string) error {
	_, err := eventLogger.file.WriteString(fmt.Sprintf("%s: %s\n", time.Now().Format(time.Kitchen), message))
	if err != nil {
		return err
	}

	return nil
}
