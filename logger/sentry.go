package reporting

import (
	"fmt"
	"log"
	"sync"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/hardikbansal/gpt-enterprise-ui/config"
)

type Logger struct {
}

func GetGinSupport() gin.HandlerFunc {
	return sentrygin.New(sentrygin.Options{
		Repanic: true,
	})
}

var lock = &sync.Mutex{}

var logger *Logger

func InitiateLogger() {
	if logger == nil {
		defer lock.Unlock()
		lock.Lock()
		if logger == nil {
			LogMessage("Creating single log instance now.")
			logger = initReporting()
			if logger == nil {
				LogPanic("connection not initialized")
			}
		}
	}
}

func initReporting() *Logger {
	config := config.GetInstance()
	if config.IsSentryActivated {
		err := sentry.Init(sentry.ClientOptions{
			Dsn: fmt.Sprintf("http://%v@%v", config.SentryKey, config.SentryUrl),
			// Enable printing of SDK debug messages.
			// Useful when getting started or trying to figure something out.
			Debug:            config.IsDebug,
			AttachStacktrace: true,
			EnableTracing:    true,
			Environment:      config.SentryEnvironment,
		})
		if err != nil {
			LogPanic("sentry.Init: %s", err)
			return nil
		}
	}

	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	return &Logger{}
}

func LogProblem(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	LogMessage(msg)
	sentry.CaptureMessage(msg)
}

func LogError(err error) {
	LogMessage(err.Error())
	sentry.CaptureException(err)
}

func LogPanic(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	panic(msg)
}

func LogMessage(format string, args ...any) {
	log.Printf(format, args...)
}
