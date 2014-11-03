package sentroni

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/codegangsta/negroni"
)

func handler(res http.ResponseWriter, req *http.Request) {
	panic("oops!")
}

func TestRecovery(t *testing.T) {
	recorder := httptest.NewRecorder()

	rec := NewRecovery("")
	rec.Logger = log.New(bytes.NewBufferString(""), "[sentroni] ", 0)

	n := negroni.New()
	n.Use(rec)
	n.UseHandler(http.HandlerFunc(handler))
	n.ServeHTTP(recorder, (*http.Request)(nil))

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("Expected %d, Actual %d", recorder.Code, http.StatusInternalServerError)
	}
}
