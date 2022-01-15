package logger

import "time"

type Logger interface {
	Debug(msg string, params map[string]interface{})
	Info(msg string, params map[string]interface{})
	Warn(msg string, params map[string]interface{})
	Error(msg string, params map[string]interface{})
	KafkaProcessMessage(topic string, partition int, message string, workerID int, offset int64, time time.Time)
	KafkaLogCommittedMessage(topic string, partition int, offset int64)
}
