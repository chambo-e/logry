package logry

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/evalphobia/logrus_sentry"
)

var hook *logrus_sentry.SentryHook

func init() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	if env != "production" {
		logrus.SetLevel(logrus.DebugLevel)
	}

	var err error
	hook, err = logrus_sentry.NewSentryHook(os.Getenv("SENTRY_DSN"), []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	})
	if err != nil {
		logrus.Errorf("logry: %s\n", err)
		return
	}

	hook.SetEnvironment(env)

	logrus.AddHook(hook)
}
