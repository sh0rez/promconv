package http

import (
	"github.com/prometheus/client_golang/prometheus"
)

import (
	"shorez.de/promconv/otel/error"
	"shorez.de/promconv/otel/network"
	"shorez.de/promconv/otel/server"
	"shorez.de/promconv/otel/url"
	"shorez.de/promconv/otel/user_agent"
)

// Size of HTTP server response bodies.
type ServerResponseBodySize struct {
	*prometheus.HistogramVec
	extra ServerResponseBodySizeExtra
}

func NewServerResponseBodySize() ServerResponseBodySize {
	labels := []string{AttrRequestMethod("").Key(), url.AttrScheme("").Key(), error.AttrType("").Key(), AttrResponseStatusCode("").Key(), AttrRoute("").Key(), network.AttrProtocolName("").Key(), network.AttrProtocolVersion("").Key(), server.AttrAddress("").Key(), server.AttrPort("").Key(), user_agent.AttrSyntheticType("").Key()}
	return ServerResponseBodySize{HistogramVec: prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_server_response_body_size",
		Help: "Size of HTTP server response bodies.",
	}, labels)}
}

func (m ServerResponseBodySize) With(requestMethod AttrRequestMethod, scheme url.AttrScheme, extras ...interface {
	ErrorType() error.AttrType
	HttpResponseStatusCode() AttrResponseStatusCode
	HttpRoute() AttrRoute
	NetworkProtocolName() network.AttrProtocolName
	NetworkProtocolVersion() network.AttrProtocolVersion
	ServerAddress() server.AttrAddress
	ServerPort() server.AttrPort
	UserAgentSyntheticType() user_agent.AttrSyntheticType
}) prometheus.Observer {
	if extras == nil {
		extras = append(extras, m.extra)
	}
	extra := extras[0]

	return m.HistogramVec.WithLabelValues(requestMethod.Value(), scheme.Value(), extra.ErrorType().Value(), extra.HttpResponseStatusCode().Value(), extra.HttpRoute().Value(), extra.NetworkProtocolName().Value(), extra.NetworkProtocolVersion().Value(), extra.ServerAddress().Value(), extra.ServerPort().Value(), extra.UserAgentSyntheticType().Value())
}

// Deprecated: Use [ServerResponseBodySize.With] instead
func (m ServerResponseBodySize) WithLabelValues(lvs ...string) prometheus.Observer {
	return m.HistogramVec.WithLabelValues(lvs...)
}

func (a ServerResponseBodySize) WithErrorType(attr interface{ ErrorType() error.AttrType }) ServerResponseBodySize {
	a.extra.AttrErrorType = attr.ErrorType()
	return a
}
func (a ServerResponseBodySize) WithResponseStatusCode(attr interface{ HttpResponseStatusCode() AttrResponseStatusCode }) ServerResponseBodySize {
	a.extra.AttrResponseStatusCode = attr.HttpResponseStatusCode()
	return a
}
func (a ServerResponseBodySize) WithRoute(attr interface{ HttpRoute() AttrRoute }) ServerResponseBodySize {
	a.extra.AttrRoute = attr.HttpRoute()
	return a
}
func (a ServerResponseBodySize) WithNetworkProtocolName(attr interface {
	NetworkProtocolName() network.AttrProtocolName
}) ServerResponseBodySize {
	a.extra.AttrNetworkProtocolName = attr.NetworkProtocolName()
	return a
}
func (a ServerResponseBodySize) WithNetworkProtocolVersion(attr interface {
	NetworkProtocolVersion() network.AttrProtocolVersion
}) ServerResponseBodySize {
	a.extra.AttrNetworkProtocolVersion = attr.NetworkProtocolVersion()
	return a
}
func (a ServerResponseBodySize) WithServerAddress(attr interface{ ServerAddress() server.AttrAddress }) ServerResponseBodySize {
	a.extra.AttrServerAddress = attr.ServerAddress()
	return a
}
func (a ServerResponseBodySize) WithServerPort(attr interface{ ServerPort() server.AttrPort }) ServerResponseBodySize {
	a.extra.AttrServerPort = attr.ServerPort()
	return a
}
func (a ServerResponseBodySize) WithUserAgentSyntheticType(attr interface {
	UserAgentSyntheticType() user_agent.AttrSyntheticType
}) ServerResponseBodySize {
	a.extra.AttrUserAgentSyntheticType = attr.UserAgentSyntheticType()
	return a
}

type ServerResponseBodySizeExtra struct {
	// Describes a class of error the operation ended with
	AttrErrorType error.AttrType `otel:"error.type"` // [HTTP response status code]
	//
	// [HTTP response status code]: https://tools.ietf.org/html/rfc7231#section-6
	AttrResponseStatusCode AttrResponseStatusCode `otel:"http.response.status_code"` // The matched route, that is, the path template in the format used by the respective server framework
	AttrRoute              AttrRoute              `otel:"http.route"`                // [OSI application layer] or non-OSI equivalent
	//
	// [OSI application layer]: https://wikipedia.org/wiki/Application_layer
	AttrNetworkProtocolName    network.AttrProtocolName     `otel:"network.protocol.name"`    // The actual version of the protocol used for network communication
	AttrNetworkProtocolVersion network.AttrProtocolVersion  `otel:"network.protocol.version"` // Name of the local HTTP server that received the request
	AttrServerAddress          server.AttrAddress           `otel:"server.address"`           // Port of the local HTTP server that received the request
	AttrServerPort             server.AttrPort              `otel:"server.port"`              // Specifies the category of synthetic traffic, such as tests or bots
	AttrUserAgentSyntheticType user_agent.AttrSyntheticType `otel:"user_agent.synthetic.type"`
}

func (a ServerResponseBodySizeExtra) ErrorType() error.AttrType { return a.AttrErrorType }
func (a ServerResponseBodySizeExtra) HttpResponseStatusCode() AttrResponseStatusCode {
	return a.AttrResponseStatusCode
}
func (a ServerResponseBodySizeExtra) HttpRoute() AttrRoute { return a.AttrRoute }
func (a ServerResponseBodySizeExtra) NetworkProtocolName() network.AttrProtocolName {
	return a.AttrNetworkProtocolName
}
func (a ServerResponseBodySizeExtra) NetworkProtocolVersion() network.AttrProtocolVersion {
	return a.AttrNetworkProtocolVersion
}
func (a ServerResponseBodySizeExtra) ServerAddress() server.AttrAddress { return a.AttrServerAddress }
func (a ServerResponseBodySizeExtra) ServerPort() server.AttrPort       { return a.AttrServerPort }
func (a ServerResponseBodySizeExtra) UserAgentSyntheticType() user_agent.AttrSyntheticType {
	return a.AttrUserAgentSyntheticType
}

/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "ServerResponseBodySizeExtra",
        "Instr": "Histogram",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "server.response.body.size",
        "Type": "ServerResponseBodySize",
        "attributes": [
            {
                "brief": "HTTP request method.",
                "examples": [
                    "GET",
                    "POST",
                    "HEAD",
                ],
                "name": "http.request.method",
                "note": "HTTP request method value SHOULD be \"known\" to the instrumentation.\nBy default, this convention defines \"known\" methods as the ones listed in [RFC9110](https://www.rfc-editor.org/rfc/rfc9110.html#name-methods)\nand the PATCH method defined in [RFC5789](https://www.rfc-editor.org/rfc/rfc5789.html).\n\nIf the HTTP request method is not known to instrumentation, it MUST set the `http.request.method` attribute to `_OTHER`.\n\nIf the HTTP instrumentation could end up converting valid HTTP request methods to `_OTHER`, then it MUST provide a way to override\nthe list of known HTTP methods. If this override is done via environment variable, then the environment variable MUST be named\nOTEL_INSTRUMENTATION_HTTP_KNOWN_METHODS and support a comma-separated list of case-sensitive known HTTP methods\n(this list MUST be a full override of the default known method, it is not a list of known methods in addition to the defaults).\n\nHTTP method names are case-sensitive and `http.request.method` attribute value MUST match a known HTTP method name exactly.\nInstrumentations for specific web frameworks that consider HTTP methods to be case insensitive, SHOULD populate a canonical equivalent.\nTracing instrumentations that do so, MUST also set `http.request.method_original` to the original value.\n",
                "requirement_level": "required",
                "stability": "stable",
                "type": {
                    "members": [
                        {
                            "brief": "CONNECT method.",
                            "deprecated": none,
                            "id": "connect",
                            "note": none,
                            "stability": "stable",
                            "value": "CONNECT",
                        },
                        {
                            "brief": "DELETE method.",
                            "deprecated": none,
                            "id": "delete",
                            "note": none,
                            "stability": "stable",
                            "value": "DELETE",
                        },
                        {
                            "brief": "GET method.",
                            "deprecated": none,
                            "id": "get",
                            "note": none,
                            "stability": "stable",
                            "value": "GET",
                        },
                        {
                            "brief": "HEAD method.",
                            "deprecated": none,
                            "id": "head",
                            "note": none,
                            "stability": "stable",
                            "value": "HEAD",
                        },
                        {
                            "brief": "OPTIONS method.",
                            "deprecated": none,
                            "id": "options",
                            "note": none,
                            "stability": "stable",
                            "value": "OPTIONS",
                        },
                        {
                            "brief": "PATCH method.",
                            "deprecated": none,
                            "id": "patch",
                            "note": none,
                            "stability": "stable",
                            "value": "PATCH",
                        },
                        {
                            "brief": "POST method.",
                            "deprecated": none,
                            "id": "post",
                            "note": none,
                            "stability": "stable",
                            "value": "POST",
                        },
                        {
                            "brief": "PUT method.",
                            "deprecated": none,
                            "id": "put",
                            "note": none,
                            "stability": "stable",
                            "value": "PUT",
                        },
                        {
                            "brief": "TRACE method.",
                            "deprecated": none,
                            "id": "trace",
                            "note": none,
                            "stability": "stable",
                            "value": "TRACE",
                        },
                        {
                            "brief": "Any HTTP method that the instrumentation has no prior knowledge of.",
                            "deprecated": none,
                            "id": "other",
                            "note": none,
                            "stability": "stable",
                            "value": "_OTHER",
                        },
                    ],
                },
            },
            {
                "brief": "The [URI scheme](https://www.rfc-editor.org/rfc/rfc3986#section-3.1) component identifying the used protocol.\n",
                "examples": [
                    "http",
                    "https",
                ],
                "name": "url.scheme",
                "note": "The scheme of the original client request, if known (e.g. from [Forwarded#proto](https://developer.mozilla.org/docs/Web/HTTP/Headers/Forwarded#proto), [X-Forwarded-Proto](https://developer.mozilla.org/docs/Web/HTTP/Headers/X-Forwarded-Proto), or a similar header). Otherwise, the scheme of the immediate peer request.\n",
                "requirement_level": "required",
                "stability": "stable",
                "type": "string",
            },
            {
                "brief": "Describes a class of error the operation ended with.\n",
                "examples": [
                    "timeout",
                    "java.net.UnknownHostException",
                    "server_certificate_invalid",
                    "500",
                ],
                "name": "error.type",
                "note": "If the request fails with an error before response status code was sent or received,\n`error.type` SHOULD be set to exception type (its fully-qualified class name, if applicable)\nor a component-specific low cardinality error identifier.\n\nIf response status code was sent or received and status indicates an error according to [HTTP span status definition](/docs/http/http-spans.md),\n`error.type` SHOULD be set to the status code number (represented as a string), an exception type (if thrown) or a component-specific error identifier.\n\nThe `error.type` value SHOULD be predictable and SHOULD have low cardinality.\nInstrumentations SHOULD document the list of errors they report.\n\nThe cardinality of `error.type` within one instrumentation library SHOULD be low, but\ntelemetry consumers that aggregate data from multiple instrumentation libraries and applications\nshould be prepared for `error.type` to have high cardinality at query time, when no\nadditional filters are applied.\n\nIf the request has completed successfully, instrumentations SHOULD NOT set `error.type`.\n",
                "requirement_level": {
                    "conditionally_required": "If request has ended with an error.",
                },
                "stability": "stable",
                "type": {
                    "members": [
                        {
                            "brief": "A fallback error value to be used when the instrumentation doesn't define a custom value.\n",
                            "deprecated": none,
                            "id": "other",
                            "note": none,
                            "stability": "stable",
                            "value": "_OTHER",
                        },
                    ],
                },
            },
            {
                "brief": "[HTTP response status code](https://tools.ietf.org/html/rfc7231#section-6).",
                "examples": [
                    200,
                ],
                "name": "http.response.status_code",
                "requirement_level": {
                    "conditionally_required": "If and only if one was received/sent.",
                },
                "stability": "stable",
                "type": "int",
            },
            {
                "brief": "The matched route, that is, the path template in the format used by the respective server framework.\n",
                "examples": [
                    "/users/:userID?",
                    "{controller}/{action}/{id?}",
                ],
                "name": "http.route",
                "note": "MUST NOT be populated when this is not supported by the HTTP server framework as the route attribute should have low-cardinality and the URI path can NOT substitute it.\nSHOULD include the [application root](/docs/http/http-spans.md#http-server-definitions) if there is one.\n",
                "requirement_level": {
                    "conditionally_required": "If and only if it's available",
                },
                "stability": "stable",
                "type": "string",
            },
            {
                "brief": "[OSI application layer](https://wikipedia.org/wiki/Application_layer) or non-OSI equivalent.",
                "examples": [
                    "http",
                    "spdy",
                ],
                "name": "network.protocol.name",
                "note": "The value SHOULD be normalized to lowercase.",
                "requirement_level": {
                    "conditionally_required": "If not `http` and `network.protocol.version` is set.",
                },
                "stability": "stable",
                "type": "string",
            },
            {
                "brief": "The actual version of the protocol used for network communication.",
                "examples": [
                    "1.0",
                    "1.1",
                    "2",
                    "3",
                ],
                "name": "network.protocol.version",
                "note": "If protocol version is subject to negotiation (for example using [ALPN](https://www.rfc-editor.org/rfc/rfc7301.html)), this attribute SHOULD be set to the negotiated version. If the actual protocol version is not known, this attribute SHOULD NOT be set.\n",
                "requirement_level": "recommended",
                "stability": "stable",
                "type": "string",
            },
            {
                "brief": "Name of the local HTTP server that received the request.\n",
                "examples": [
                    "example.com",
                    "10.1.2.80",
                    "/tmp/my.sock",
                ],
                "name": "server.address",
                "note": "See [Setting `server.address` and `server.port` attributes](/docs/http/http-spans.md#setting-serveraddress-and-serverport-attributes).\n> **Warning**\n> Since this attribute is based on HTTP headers, opting in to it may allow an attacker\n> to trigger cardinality limits, degrading the usefulness of the metric.\n",
                "requirement_level": "opt_in",
                "stability": "stable",
                "type": "string",
            },
            {
                "brief": "Port of the local HTTP server that received the request.\n",
                "examples": [
                    80,
                    8080,
                    443,
                ],
                "name": "server.port",
                "note": "See [Setting `server.address` and `server.port` attributes](/docs/http/http-spans.md#setting-serveraddress-and-serverport-attributes).\n> **Warning**\n> Since this attribute is based on HTTP headers, opting in to it may allow an attacker\n> to trigger cardinality limits, degrading the usefulness of the metric.\n",
                "requirement_level": "opt_in",
                "stability": "stable",
                "type": "int",
            },
            {
                "brief": "Specifies the category of synthetic traffic, such as tests or bots.\n",
                "name": "user_agent.synthetic.type",
                "note": "This attribute MAY be derived from the contents of the `user_agent.original` attribute. Components that populate the attribute are responsible for determining what they consider to be synthetic bot or test traffic. This attribute can either be set for self-identification purposes, or on telemetry detected to be generated as a result of a synthetic request. This attribute is useful for distinguishing between genuine client traffic and synthetic traffic generated by bots or tests.\n",
                "requirement_level": "opt_in",
                "stability": "development",
                "type": {
                    "members": [
                        {
                            "brief": "Bot source.",
                            "deprecated": none,
                            "id": "bot",
                            "note": none,
                            "stability": "development",
                            "value": "bot",
                        },
                        {
                            "brief": "Synthetic test source.",
                            "deprecated": none,
                            "id": "test",
                            "note": none,
                            "stability": "development",
                            "value": "test",
                        },
                    ],
                },
            },
        ],
        "ctx": {
            "attributes": [
                {
                    "brief": "Describes a class of error the operation ended with.\n",
                    "examples": [
                        "timeout",
                        "java.net.UnknownHostException",
                        "server_certificate_invalid",
                        "500",
                    ],
                    "name": "error.type",
                    "note": "If the request fails with an error before response status code was sent or received,\n`error.type` SHOULD be set to exception type (its fully-qualified class name, if applicable)\nor a component-specific low cardinality error identifier.\n\nIf response status code was sent or received and status indicates an error according to [HTTP span status definition](/docs/http/http-spans.md),\n`error.type` SHOULD be set to the status code number (represented as a string), an exception type (if thrown) or a component-specific error identifier.\n\nThe `error.type` value SHOULD be predictable and SHOULD have low cardinality.\nInstrumentations SHOULD document the list of errors they report.\n\nThe cardinality of `error.type` within one instrumentation library SHOULD be low, but\ntelemetry consumers that aggregate data from multiple instrumentation libraries and applications\nshould be prepared for `error.type` to have high cardinality at query time, when no\nadditional filters are applied.\n\nIf the request has completed successfully, instrumentations SHOULD NOT set `error.type`.\n",
                    "requirement_level": {
                        "conditionally_required": "If request has ended with an error.",
                    },
                    "stability": "stable",
                    "type": {
                        "members": [
                            {
                                "brief": "A fallback error value to be used when the instrumentation doesn't define a custom value.\n",
                                "deprecated": none,
                                "id": "other",
                                "note": none,
                                "stability": "stable",
                                "value": "_OTHER",
                            },
                        ],
                    },
                },
                {
                    "brief": "[HTTP response status code](https://tools.ietf.org/html/rfc7231#section-6).",
                    "examples": [
                        200,
                    ],
                    "name": "http.response.status_code",
                    "requirement_level": {
                        "conditionally_required": "If and only if one was received/sent.",
                    },
                    "stability": "stable",
                    "type": "int",
                },
                {
                    "brief": "[OSI application layer](https://wikipedia.org/wiki/Application_layer) or non-OSI equivalent.",
                    "examples": [
                        "http",
                        "spdy",
                    ],
                    "name": "network.protocol.name",
                    "note": "The value SHOULD be normalized to lowercase.",
                    "requirement_level": {
                        "conditionally_required": "If not `http` and `network.protocol.version` is set.",
                    },
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "The actual version of the protocol used for network communication.",
                    "examples": [
                        "1.0",
                        "1.1",
                        "2",
                        "3",
                    ],
                    "name": "network.protocol.version",
                    "note": "If protocol version is subject to negotiation (for example using [ALPN](https://www.rfc-editor.org/rfc/rfc7301.html)), this attribute SHOULD be set to the negotiated version. If the actual protocol version is not known, this attribute SHOULD NOT be set.\n",
                    "requirement_level": "recommended",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "The matched route, that is, the path template in the format used by the respective server framework.\n",
                    "examples": [
                        "/users/:userID?",
                        "{controller}/{action}/{id?}",
                    ],
                    "name": "http.route",
                    "note": "MUST NOT be populated when this is not supported by the HTTP server framework as the route attribute should have low-cardinality and the URI path can NOT substitute it.\nSHOULD include the [application root](/docs/http/http-spans.md#http-server-definitions) if there is one.\n",
                    "requirement_level": {
                        "conditionally_required": "If and only if it's available",
                    },
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "HTTP request method.",
                    "examples": [
                        "GET",
                        "POST",
                        "HEAD",
                    ],
                    "name": "http.request.method",
                    "note": "HTTP request method value SHOULD be \"known\" to the instrumentation.\nBy default, this convention defines \"known\" methods as the ones listed in [RFC9110](https://www.rfc-editor.org/rfc/rfc9110.html#name-methods)\nand the PATCH method defined in [RFC5789](https://www.rfc-editor.org/rfc/rfc5789.html).\n\nIf the HTTP request method is not known to instrumentation, it MUST set the `http.request.method` attribute to `_OTHER`.\n\nIf the HTTP instrumentation could end up converting valid HTTP request methods to `_OTHER`, then it MUST provide a way to override\nthe list of known HTTP methods. If this override is done via environment variable, then the environment variable MUST be named\nOTEL_INSTRUMENTATION_HTTP_KNOWN_METHODS and support a comma-separated list of case-sensitive known HTTP methods\n(this list MUST be a full override of the default known method, it is not a list of known methods in addition to the defaults).\n\nHTTP method names are case-sensitive and `http.request.method` attribute value MUST match a known HTTP method name exactly.\nInstrumentations for specific web frameworks that consider HTTP methods to be case insensitive, SHOULD populate a canonical equivalent.\nTracing instrumentations that do so, MUST also set `http.request.method_original` to the original value.\n",
                    "requirement_level": "required",
                    "stability": "stable",
                    "type": {
                        "members": [
                            {
                                "brief": "CONNECT method.",
                                "deprecated": none,
                                "id": "connect",
                                "note": none,
                                "stability": "stable",
                                "value": "CONNECT",
                            },
                            {
                                "brief": "DELETE method.",
                                "deprecated": none,
                                "id": "delete",
                                "note": none,
                                "stability": "stable",
                                "value": "DELETE",
                            },
                            {
                                "brief": "GET method.",
                                "deprecated": none,
                                "id": "get",
                                "note": none,
                                "stability": "stable",
                                "value": "GET",
                            },
                            {
                                "brief": "HEAD method.",
                                "deprecated": none,
                                "id": "head",
                                "note": none,
                                "stability": "stable",
                                "value": "HEAD",
                            },
                            {
                                "brief": "OPTIONS method.",
                                "deprecated": none,
                                "id": "options",
                                "note": none,
                                "stability": "stable",
                                "value": "OPTIONS",
                            },
                            {
                                "brief": "PATCH method.",
                                "deprecated": none,
                                "id": "patch",
                                "note": none,
                                "stability": "stable",
                                "value": "PATCH",
                            },
                            {
                                "brief": "POST method.",
                                "deprecated": none,
                                "id": "post",
                                "note": none,
                                "stability": "stable",
                                "value": "POST",
                            },
                            {
                                "brief": "PUT method.",
                                "deprecated": none,
                                "id": "put",
                                "note": none,
                                "stability": "stable",
                                "value": "PUT",
                            },
                            {
                                "brief": "TRACE method.",
                                "deprecated": none,
                                "id": "trace",
                                "note": none,
                                "stability": "stable",
                                "value": "TRACE",
                            },
                            {
                                "brief": "Any HTTP method that the instrumentation has no prior knowledge of.",
                                "deprecated": none,
                                "id": "other",
                                "note": none,
                                "stability": "stable",
                                "value": "_OTHER",
                            },
                        ],
                    },
                },
                {
                    "brief": "The [URI scheme](https://www.rfc-editor.org/rfc/rfc3986#section-3.1) component identifying the used protocol.\n",
                    "examples": [
                        "http",
                        "https",
                    ],
                    "name": "url.scheme",
                    "note": "The scheme of the original client request, if known (e.g. from [Forwarded#proto](https://developer.mozilla.org/docs/Web/HTTP/Headers/Forwarded#proto), [X-Forwarded-Proto](https://developer.mozilla.org/docs/Web/HTTP/Headers/X-Forwarded-Proto), or a similar header). Otherwise, the scheme of the immediate peer request.\n",
                    "requirement_level": "required",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Specifies the category of synthetic traffic, such as tests or bots.\n",
                    "name": "user_agent.synthetic.type",
                    "note": "This attribute MAY be derived from the contents of the `user_agent.original` attribute. Components that populate the attribute are responsible for determining what they consider to be synthetic bot or test traffic. This attribute can either be set for self-identification purposes, or on telemetry detected to be generated as a result of a synthetic request. This attribute is useful for distinguishing between genuine client traffic and synthetic traffic generated by bots or tests.\n",
                    "requirement_level": "opt_in",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "Bot source.",
                                "deprecated": none,
                                "id": "bot",
                                "note": none,
                                "stability": "development",
                                "value": "bot",
                            },
                            {
                                "brief": "Synthetic test source.",
                                "deprecated": none,
                                "id": "test",
                                "note": none,
                                "stability": "development",
                                "value": "test",
                            },
                        ],
                    },
                },
                {
                    "brief": "Name of the local HTTP server that received the request.\n",
                    "examples": [
                        "example.com",
                        "10.1.2.80",
                        "/tmp/my.sock",
                    ],
                    "name": "server.address",
                    "note": "See [Setting `server.address` and `server.port` attributes](/docs/http/http-spans.md#setting-serveraddress-and-serverport-attributes).\n> **Warning**\n> Since this attribute is based on HTTP headers, opting in to it may allow an attacker\n> to trigger cardinality limits, degrading the usefulness of the metric.\n",
                    "requirement_level": "opt_in",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Port of the local HTTP server that received the request.\n",
                    "examples": [
                        80,
                        8080,
                        443,
                    ],
                    "name": "server.port",
                    "note": "See [Setting `server.address` and `server.port` attributes](/docs/http/http-spans.md#setting-serveraddress-and-serverport-attributes).\n> **Warning**\n> Since this attribute is based on HTTP headers, opting in to it may allow an attacker\n> to trigger cardinality limits, degrading the usefulness of the metric.\n",
                    "requirement_level": "opt_in",
                    "stability": "stable",
                    "type": "int",
                },
            ],
            "brief": "Size of HTTP server response bodies.",
            "events": [],
            "id": "metric.http.server.response.body.size",
            "instrument": "histogram",
            "lineage": {
                "attributes": {
                    "error.type": {
                        "inherited_fields": [
                            "brief",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "examples",
                            "note",
                            "requirement_level",
                        ],
                        "source_group": "registry.error",
                    },
                    "http.request.method": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.http",
                    },
                    "http.response.status_code": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.http",
                    },
                    "http.route": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.http",
                    },
                    "network.protocol.name": {
                        "inherited_fields": [
                            "brief",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "examples",
                            "requirement_level",
                        ],
                        "source_group": "registry.network",
                    },
                    "network.protocol.version": {
                        "inherited_fields": [
                            "brief",
                            "note",
                            "requirement_level",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "examples",
                        ],
                        "source_group": "registry.network",
                    },
                    "server.address": {
                        "inherited_fields": [
                            "examples",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "brief",
                            "note",
                            "requirement_level",
                        ],
                        "source_group": "registry.server",
                    },
                    "server.port": {
                        "inherited_fields": [
                            "examples",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "brief",
                            "note",
                            "requirement_level",
                        ],
                        "source_group": "registry.server",
                    },
                    "url.scheme": {
                        "inherited_fields": [
                            "brief",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "examples",
                            "note",
                            "requirement_level",
                        ],
                        "source_group": "registry.url",
                    },
                    "user_agent.synthetic.type": {
                        "inherited_fields": [
                            "brief",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.user_agent.os",
                    },
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/http/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "http.server.response.body.size",
            "name": none,
            "note": "The size of the response payload body in bytes. This is the number of bytes transferred excluding headers and is often, but not always, present as the [Content-Length](https://www.rfc-editor.org/rfc/rfc9110.html#field.content-length) header. For requests using transport encoding, this should be the compressed size.\n",
            "root_namespace": "http",
            "span_kind": none,
            "stability": "development",
            "type": "metric",
            "unit": "By",
        },
        "for_each_attr": <macro for_each_attr>,
        "module": "shorez.de/promconv/otel",
    },
    env: Environment {
        globals: {
            "concat_if": weaver_forge::extensions::util::concat_if,
            "cycler": minijinja_contrib::globals::cycler,
            "debug": minijinja::functions::builtins::debug,
            "dict": minijinja::functions::builtins::dict,
            "joiner": minijinja_contrib::globals::joiner,
            "namespace": minijinja::functions::builtins::namespace,
            "params": {
                "params": {},
            },
            "range": minijinja::functions::builtins::range,
            "template": {},
        },
        tests: [
            "!=",
            "<",
            "<=",
            "==",
            ">",
            ">=",
            "array",
            "boolean",
            "defined",
            "deprecated",
            "divisibleby",
            "endingwith",
            "enum",
            "enum_type",
            "eq",
            "equalto",
            "escaped",
            "even",
            "experimental",
            "false",
            "filter",
            "float",
            "ge",
            "greaterthan",
            "gt",
            "in",
            "int",
            "integer",
            "iterable",
            "le",
            "lessthan",
            "lower",
            "lt",
            "mapping",
            "ne",
            "none",
            "number",
            "odd",
            "safe",
            "sameas",
            "sequence",
            "simple_type",
            "stable",
            "startingwith",
            "string",
            "template_type",
            "test",
            "true",
            "undefined",
            "upper",
        ],
        filters: [
            "abs",
            "acronym",
            "ansi_bg_black",
            "ansi_bg_blue",
            "ansi_bg_bright_black",
            "ansi_bg_bright_blue",
            "ansi_bg_bright_cyan",
            "ansi_bg_bright_green",
            "ansi_bg_bright_magenta",
            "ansi_bg_bright_red",
            "ansi_bg_bright_white",
            "ansi_bg_bright_yellow",
            "ansi_bg_cyan",
            "ansi_bg_green",
            "ansi_bg_magenta",
            "ansi_bg_red",
            "ansi_bg_white",
            "ansi_bg_yellow",
            "ansi_black",
            "ansi_blue",
            "ansi_bold",
            "ansi_bright_black",
            "ansi_bright_blue",
            "ansi_bright_cyan",
            "ansi_bright_green",
            "ansi_bright_magenta",
            "ansi_bright_red",
            "ansi_bright_white",
            "ansi_bright_yellow",
            "ansi_cyan",
            "ansi_green",
            "ansi_italic",
            "ansi_magenta",
            "ansi_red",
            "ansi_strikethrough",
            "ansi_underline",
            "ansi_white",
            "ansi_yellow",
            "attr",
            "attribute_id",
            "attribute_namespace",
            "attribute_registry_file",
            "attribute_registry_namespace",
            "attribute_registry_title",
            "attribute_sort",
            "batch",
            "body_fields",
            "bool",
            "camel_case",
            "camel_case_const",
            "capitalize",
            "capitalize_first",
            "comment",
            "comment_with_prefix",
            "count",
            "d",
            "default",
            "dictsort",
            "e",
            "enum_type",
            "escape",
            "filesizeformat",
            "first",
            "flatten",
            "float",
            "groupby",
            "indent",
            "instantiated_type",
            "int",
            "items",
            "join",
            "kebab_case",
            "kebab_case_const",
            "last",
            "length",
            "lines",
            "list",
            "lower",
            "lower_case",
            "map",
            "map_text",
            "markdown_to_html",
            "max",
            "metric_namespace",
            "min",
            "not_required",
            "pascal_case",
            "pascal_case_const",
            "pluralize",
            "pprint",
            "print_member_value",
            "regex_replace",
            "reject",
            "rejectattr",
            "replace",
            "required",
            "reverse",
            "round",
            "safe",
            "screaming_kebab_case",
            "screaming_snake_case",
            "screaming_snake_case_const",
            "select",
            "selectattr",
            "slice",
            "snake_case",
            "snake_case_const",
            "sort",
            "split",
            "split_id",
            "string",
            "striptags",
            "sum",
            "title",
            "title_case",
            "tojson",
            "toyaml",
            "trim",
            "truncate",
            "type_mapping",
            "unique",
            "upper",
            "upper_case",
            "urlencode",
        ],
        templates: [
            "vec.go.j2",
        ],
    },
}
*/
