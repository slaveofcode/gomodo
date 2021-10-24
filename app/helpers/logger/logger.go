package logger

import "github.com/sirupsen/logrus"

func Init(level ...interface{}) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	if len(level) > 0 {
		switch l := level[0].(type) {
		case logrus.Level:
			logrus.SetLevel(l)
		}
	}
}
