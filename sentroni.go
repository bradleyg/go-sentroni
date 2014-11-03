// Recovery is a Negroni middleware that recovers from any panics, writes a
// 500 and passes the error to Sentry.
//
// Usage:
//   n := negroni.New()
//   n.Use(sentroni.NewRecovery(sentryDsn))
//   n.Run()
package sentroni

import (
	"errors"
	"log"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/getsentry/raven-go"
)

type recovery struct {
	Logger *log.Logger
	Client *raven.Client
}

// NewRecovery returns a *http.Handler interface. Pass your Sentry dsn as an argument.
func NewRecovery(dsn string) *recovery {
	logger := log.New(os.Stdout, "[sentroni] ", 0)

	client, err := raven.NewClient(dsn, nil)
	if err != nil {
		logger.Fatal("FATAL: ", err)
	}

	return &recovery{
		Logger: logger,
		Client: client,
	}
}

func (rec *recovery) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer func() {
		if err := recover(); err != nil {
			rec.Logger.Printf("PANIC: %s\n%s", err, debug.Stack())
			rec.Client.CaptureError(errors.New(err.(string)), nil)
			rw.WriteHeader(http.StatusInternalServerError)
		}
	}()

	next(rw, r)
}
