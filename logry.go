// Package logry is a simple wrapper around logrus and logrus_sentry
//
// It requires no configuration to be used, only an env var SENTRY_DSN
//
//     import (
//         "github.com/chambo-e/logry"
//     )
//
//     func main() {
//         logry.Debugln("Will be printed on stdout except if APP_ENV=production")
//         logry.Println("Will be printed on stdout")
//         logry.Errorln("Will be printed on stderr and sent to Sentry")
//         logry.Fatalln("Will be printed on stderr and sent to Sentry")
//         logry.Panicln("Will be printed on stderr and sent to Sentry")
//     }
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
