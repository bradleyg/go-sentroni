# sentroni
--
    import "github.com/bradleyg/go-sentroni"


## Usage

#### type Recovery

```go
type Recovery struct {
	Logger *log.Logger
	Client *raven.Client
}
```

Recovery is a Negroni middleware that recovers from any panics, writes a 500 and
passes the error to Sentry.

#### func  NewRecovery

```go
func NewRecovery(dsn string) *Recovery
```
NewRecovery returns a new instance of Recovery. Pass your Sentry dsn as an
argument.

#### func (*Recovery) ServeHTTP

```go
func (rec *Recovery) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)
```
