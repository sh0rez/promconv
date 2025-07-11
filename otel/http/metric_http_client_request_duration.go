package http

import (
	"github.com/prometheus/client_golang/prometheus"
)

import (
	"shorez.de/promconv/otel/error"
	"shorez.de/promconv/otel/network"
	"shorez.de/promconv/otel/server"
	"shorez.de/promconv/otel/url"
)

// Duration of HTTP client requests.
type ClientRequestDuration struct {
	*prometheus.HistogramVec
	extra ClientRequestDurationExtra
}

func NewClientRequestDuration() ClientRequestDuration {
	labels := []string{AttrRequestMethod("").Key(), server.AttrAddress("").Key(), server.AttrPort("").Key(), error.AttrType("").Key(), AttrResponseStatusCode("").Key(), network.AttrProtocolName("").Key(), network.AttrProtocolVersion("").Key(), url.AttrScheme("").Key(), url.AttrTemplate("").Key()}
	return ClientRequestDuration{HistogramVec: prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_client_request_duration",
		Help: "Duration of HTTP client requests.",
	}, labels)}
}

func (m ClientRequestDuration) With(requestMethod AttrRequestMethod, address server.AttrAddress, port server.AttrPort, extras ...interface {
	ErrorType() error.AttrType
	HttpResponseStatusCode() AttrResponseStatusCode
	NetworkProtocolName() network.AttrProtocolName
	NetworkProtocolVersion() network.AttrProtocolVersion
	UrlScheme() url.AttrScheme
	UrlTemplate() url.AttrTemplate
}) prometheus.Observer {
	if extras == nil {
		extras = append(extras, m.extra)
	}
	extra := extras[0]

	return m.HistogramVec.WithLabelValues(requestMethod.Value(), address.Value(), port.Value(), extra.ErrorType().Value(), extra.HttpResponseStatusCode().Value(), extra.NetworkProtocolName().Value(), extra.NetworkProtocolVersion().Value(), extra.UrlScheme().Value(), extra.UrlTemplate().Value())
}

// Deprecated: Use [ClientRequestDuration.With] instead
func (m ClientRequestDuration) WithLabelValues(lvs ...string) prometheus.Observer {
	return m.HistogramVec.WithLabelValues(lvs...)
}

func (a ClientRequestDuration) WithErrorType(attr interface{ ErrorType() error.AttrType }) ClientRequestDuration {
	a.extra.AttrErrorType = attr.ErrorType()
	return a
}
func (a ClientRequestDuration) WithResponseStatusCode(attr interface{ HttpResponseStatusCode() AttrResponseStatusCode }) ClientRequestDuration {
	a.extra.AttrResponseStatusCode = attr.HttpResponseStatusCode()
	return a
}
func (a ClientRequestDuration) WithNetworkProtocolName(attr interface {
	NetworkProtocolName() network.AttrProtocolName
}) ClientRequestDuration {
	a.extra.AttrNetworkProtocolName = attr.NetworkProtocolName()
	return a
}
func (a ClientRequestDuration) WithNetworkProtocolVersion(attr interface {
	NetworkProtocolVersion() network.AttrProtocolVersion
}) ClientRequestDuration {
	a.extra.AttrNetworkProtocolVersion = attr.NetworkProtocolVersion()
	return a
}
func (a ClientRequestDuration) WithUrlScheme(attr interface{ UrlScheme() url.AttrScheme }) ClientRequestDuration {
	a.extra.AttrUrlScheme = attr.UrlScheme()
	return a
}
func (a ClientRequestDuration) WithUrlTemplate(attr interface{ UrlTemplate() url.AttrTemplate }) ClientRequestDuration {
	a.extra.AttrUrlTemplate = attr.UrlTemplate()
	return a
}

type ClientRequestDurationExtra struct {
	// Describes a class of error the operation ended with
	AttrErrorType error.AttrType `otel:"error.type"` // [HTTP response status code]
	//
	// [HTTP response status code]: https://tools.ietf.org/html/rfc7231#section-6
	AttrResponseStatusCode AttrResponseStatusCode `otel:"http.response.status_code"` // [OSI application layer] or non-OSI equivalent
	//
	// [OSI application layer]: https://wikipedia.org/wiki/Application_layer
	AttrNetworkProtocolName    network.AttrProtocolName    `otel:"network.protocol.name"`    // The actual version of the protocol used for network communication
	AttrNetworkProtocolVersion network.AttrProtocolVersion `otel:"network.protocol.version"` // The [URI scheme] component identifying the used protocol
	//
	// [URI scheme]: https://www.rfc-editor.org/rfc/rfc3986#section-3.1
	AttrUrlScheme url.AttrScheme `otel:"url.scheme"` // The low-cardinality template of an [absolute path reference]
	//
	// [absolute path reference]: https://www.rfc-editor.org/rfc/rfc3986#section-4.2
	AttrUrlTemplate url.AttrTemplate `otel:"url.template"`
}

func (a ClientRequestDurationExtra) ErrorType() error.AttrType { return a.AttrErrorType }
func (a ClientRequestDurationExtra) HttpResponseStatusCode() AttrResponseStatusCode {
	return a.AttrResponseStatusCode
}
func (a ClientRequestDurationExtra) NetworkProtocolName() network.AttrProtocolName {
	return a.AttrNetworkProtocolName
}
func (a ClientRequestDurationExtra) NetworkProtocolVersion() network.AttrProtocolVersion {
	return a.AttrNetworkProtocolVersion
}
func (a ClientRequestDurationExtra) UrlScheme() url.AttrScheme     { return a.AttrUrlScheme }
func (a ClientRequestDurationExtra) UrlTemplate() url.AttrTemplate { return a.AttrUrlTemplate }

/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "ClientRequestDurationExtra",
        "Instr": "Histogram",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "client.request.duration",
        "Type": "ClientRequestDuration",
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
                "brief": "Host identifier of the [\"URI origin\"](https://www.rfc-editor.org/rfc/rfc9110.html#name-uri-origin) HTTP request is sent to.\n",
                "examples": [
                    "example.com",
                    "10.1.2.80",
                    "/tmp/my.sock",
                ],
                "name": "server.address",
                "note": "If an HTTP client request is explicitly made to an IP address, e.g. `http://x.x.x.x:8080`, then `server.address` SHOULD be the IP address `x.x.x.x`. A DNS lookup SHOULD NOT be used.\n",
                "requirement_level": "required",
                "stability": "stable",
                "type": "string",
            },
            {
                "brief": "Port identifier of the [\"URI origin\"](https://www.rfc-editor.org/rfc/rfc9110.html#name-uri-origin) HTTP request is sent to.\n",
                "examples": [
                    80,
                    8080,
                    443,
                ],
                "name": "server.port",
                "note": "When observed from the client side, and when communicating through an intermediary, `server.port` SHOULD represent the server port behind any intermediaries, for example proxies, if it's available.\n",
                "requirement_level": "required",
                "stability": "stable",
                "type": "int",
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
                "brief": "The [URI scheme](https://www.rfc-editor.org/rfc/rfc3986#section-3.1) component identifying the used protocol.\n",
                "examples": [
                    "http",
                    "https",
                ],
                "name": "url.scheme",
                "requirement_level": "opt_in",
                "stability": "stable",
                "type": "string",
            },
            {
                "brief": "The low-cardinality template of an [absolute path reference](https://www.rfc-editor.org/rfc/rfc3986#section-4.2).\n",
                "examples": [
                    "/users/{id}",
                    "/users/:id",
                    "/users?id={id}",
                ],
                "name": "url.template",
                "note": "The `url.template` MUST have low cardinality. It is not usually available on HTTP clients, but may be known by the application or specialized HTTP instrumentation.\n",
                "requirement_level": "opt_in",
                "stability": "development",
                "type": "string",
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
                    "brief": "The [URI scheme](https://www.rfc-editor.org/rfc/rfc3986#section-3.1) component identifying the used protocol.\n",
                    "examples": [
                        "http",
                        "https",
                    ],
                    "name": "url.scheme",
                    "requirement_level": "opt_in",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "The low-cardinality template of an [absolute path reference](https://www.rfc-editor.org/rfc/rfc3986#section-4.2).\n",
                    "examples": [
                        "/users/{id}",
                        "/users/:id",
                        "/users?id={id}",
                    ],
                    "name": "url.template",
                    "note": "The `url.template` MUST have low cardinality. It is not usually available on HTTP clients, but may be known by the application or specialized HTTP instrumentation.\n",
                    "requirement_level": "opt_in",
                    "stability": "development",
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
                    "brief": "Host identifier of the [\"URI origin\"](https://www.rfc-editor.org/rfc/rfc9110.html#name-uri-origin) HTTP request is sent to.\n",
                    "examples": [
                        "example.com",
                        "10.1.2.80",
                        "/tmp/my.sock",
                    ],
                    "name": "server.address",
                    "note": "If an HTTP client request is explicitly made to an IP address, e.g. `http://x.x.x.x:8080`, then `server.address` SHOULD be the IP address `x.x.x.x`. A DNS lookup SHOULD NOT be used.\n",
                    "requirement_level": "required",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Port identifier of the [\"URI origin\"](https://www.rfc-editor.org/rfc/rfc9110.html#name-uri-origin) HTTP request is sent to.\n",
                    "examples": [
                        80,
                        8080,
                        443,
                    ],
                    "name": "server.port",
                    "note": "When observed from the client side, and when communicating through an intermediary, `server.port` SHOULD represent the server port behind any intermediaries, for example proxies, if it's available.\n",
                    "requirement_level": "required",
                    "stability": "stable",
                    "type": "int",
                },
            ],
            "brief": "Duration of HTTP client requests.",
            "events": [],
            "id": "metric.http.client.request.duration",
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
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "brief",
                            "requirement_level",
                        ],
                        "source_group": "registry.server",
                    },
                    "url.scheme": {
                        "inherited_fields": [
                            "brief",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "examples",
                            "requirement_level",
                        ],
                        "source_group": "registry.url",
                    },
                    "url.template": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "note",
                            "requirement_level",
                        ],
                        "source_group": "registry.url",
                    },
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/http/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "http.client.request.duration",
            "name": none,
            "root_namespace": "http",
            "span_kind": none,
            "stability": "stable",
            "type": "metric",
            "unit": "s",
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
