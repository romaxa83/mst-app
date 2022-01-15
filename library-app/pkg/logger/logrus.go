package logger

import (
	"github.com/romaxa83/mst-app/pkg/constants"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"time"
)

// TODO create general interface with generic fields

func Debug(msg ...interface{}) {
	logrus.Debug(msg...)
}

func Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func Info(msg ...interface{}) {
	logrus.Info(msg...)
}

func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func Warn(msg ...interface{}) {
	logrus.Warn(msg...)
}

func Warnf(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

func Error(msg ...interface{}) {
	logrus.Error(msg...)
}

func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

// todo zap перевести на logrus
func KafkaProcessMessage(topic string, partition int, message string, workerID int, offset int64, time time.Time) {
	logrus.Debug(
		"Processing Kafka message",
		zap.String(constants.Topic, topic),
		zap.Int(constants.Partition, partition),
		zap.String(constants.Message, message),
		zap.Int(constants.WorkerID, workerID),
		zap.Int64(constants.Offset, offset),
		zap.Time(constants.Time, time),
	)
}

func KafkaLogCommittedMessage(topic string, partition int, offset int64) {
	logrus.Debug(
		"Committed Kafka message",
		zap.String(constants.Topic, topic),
		zap.Int(constants.Partition, partition),
		zap.Int64(constants.Offset, offset),
	)
}
