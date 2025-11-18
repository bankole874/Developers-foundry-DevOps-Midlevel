package logging


import (
"strings"


"github.com/sirupsen/logrus"
)


func Setup(level string) {
l := strings.ToLower(level)
switch l {
case "debug":
logrus.SetLevel(logrus.DebugLevel)
case "warn":
logrus.SetLevel(logrus.WarnLevel)
case "error":
logrus.SetLevel(logrus.ErrorLevel)
default:
logrus.SetLevel(logrus.InfoLevel)
}
logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
}
