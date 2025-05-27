package http

import (
	"io"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"shorez.de/promconv/otel/network"
	"shorez.de/promconv/otel/server"
	"shorez.de/promconv/otel/url"
)

func Register(reg prometheus.Registerer) Metrics {
	m := Metrics{
		Active:       NewServerActiveRequests(),
		ResponseSize: NewServerResponseBodySize(),
		RequestSize:  NewServerRequestBodySize(),
		Duration:     NewServerRequestDuration(),
	}
	reg.MustRegister(m.Active, m.ResponseSize, m.RequestSize, m.Duration)
	return m
}

type Metrics struct {
	Active       ServerActiveRequests
	ResponseSize ServerResponseBodySize
	RequestSize  ServerRequestBodySize
	Duration     ServerRequestDuration
}

type base = ServerRequestDurationExtra

// TODO: generate from yaml
type Attributes struct {
	base
}

func Method(r *http.Request) AttrRequestMethod {
	return AttrRequestMethod(r.Method)
}

func Scheme(r *http.Request) url.AttrScheme {
	return url.AttrScheme(r.URL.Scheme)
}

func Instrument(m Metrics, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var attrs Attributes
		addrport := r.Context().Value(http.LocalAddrContextKey).(net.Addr)
		attrs.ServerAddress, attrs.ServerPort = server.AddrPort(addrport)

		attrs.HttpRoute = AttrRoute(r.URL.Path)
		attrs.NetworkProtocolName = "http"
		if _, v, ok := strings.Cut(r.Proto, "/"); ok {
			attrs.NetworkProtocolVersion = network.AttrProtocolVersion(v)
		}

		if mux, ok := h.(interface {
			Handler(r *http.Request) (http.Handler, string)
		}); ok {
			_, route := mux.Handler(r)
			attrs.HttpRoute = AttrRoute(route)
		}

		active := m.Active.With(Method(r), Scheme(r), attrs)
		active.Inc()
		defer active.Dec()

		start := time.Now()

		body := countReader{ReadCloser: r.Body}
		r.Body = &body

		res := countWriter{ResponseWriter: w, code: http.StatusOK}
		h.ServeHTTP(&res, r)

		attrs.HttpResponseStatusCode = AttrResponseStatusCode(strconv.Itoa(res.code))

		m.RequestSize.With(Method(r), Scheme(r), attrs).Observe(float64(body.size))
		m.ResponseSize.With(Method(r), Scheme(r), attrs).Observe(float64(res.size))
		m.Duration.With(Method(r), Scheme(r), attrs).Observe(float64(time.Since(start).Seconds()))
	})
}

var _ http.ResponseWriter = (*countWriter)(nil)

type countWriter struct {
	http.ResponseWriter
	code int
	size uint64
}

func (cw *countWriter) WriteHeader(code int) {
	cw.code = code
	cw.ResponseWriter.WriteHeader(code)
}

func (cw *countWriter) Write(data []byte) (int, error) {
	n, err := cw.ResponseWriter.Write(data)
	cw.size += uint64(n)
	return n, err
}

var _ io.ReadCloser = (*countReader)(nil)

type countReader struct {
	io.ReadCloser
	size uint64
}

func (cr *countReader) Read(into []byte) (int, error) {
	n, err := cr.ReadCloser.Read(into)
	cr.size += uint64(n)
	return n, err
}
