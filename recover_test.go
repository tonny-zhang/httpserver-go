package httpserver

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecoverWithWriter(t *testing.T) {
	buf := new(bytes.Buffer)

	router := NewRouter()
	router.Use(RecoverWithWriter(buf, func(ctx *Context, err interface{}) {
		ctx.writer.WriteHeader(http.StatusOK)
		ctx.writer.Write([]byte("[RECOVER]" + err.(string)))
	}))
	router.Get("/panic", func(c *Context) {
		panic("test")
	})

	w := doRequest(&router, http.MethodGet, "/panic")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "[RECOVER]test", w.Body.String())
	assert.Contains(t, buf.String(), "panic(\"test\")")
}