# sentroni
--
    import "github.com/bradleyg/go-sentroni"

Recovery is a Negroni middleware that recovers from any panics, writes a 500 and
passes the error to Sentry.

Usage:

    n := negroni.New()
    n.Use(sentroni.NewRecovery(sentryDsn))
    n.Run()

## Usage

#### func  NewRecovery

```go
func NewRecovery(dsn string) *recovery
```
NewRecovery returns a *http.Handler interface. Pass your Sentry dsn as an
argument.
