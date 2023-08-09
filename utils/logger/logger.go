package logger

import (
	"github.com/sirupsen/logrus"
)

type requestLog struct {
	requestID string
}

func NewLogger(reqID string) logrus.FieldLogger {

	log := logrus.StandardLogger()
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.TextFormatter{})
	log.SetReportCaller(true)

	var retLogger logrus.FieldLogger = log
	return retLogger.WithField("req ID:", reqID)

}
