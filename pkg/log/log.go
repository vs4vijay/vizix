package log

import (
	"github.com/sirupsen/logrus"
)

func Init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func Log(data interface{}) {
	logrus.Info(data)
}
