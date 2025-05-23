package http

// Deprecated, use `client.address` instead
type AttrClientIp string // http.client_ip

func (AttrClientIp) Development() {}
func (AttrClientIp) Recommended() {}

// State of the HTTP connection in the HTTP connection pool
type AttrConnectionState string // http.connection.state

func (AttrConnectionState) Development() {}
func (AttrConnectionState) Recommended() {}

// Deprecated, use `network.protocol.name` instead
type AttrFlavor string // http.flavor

func (AttrFlavor) Development() {}
func (AttrFlavor) Recommended() {}

// Deprecated, use one of `server.address`, `client.address` or `http.request.header.host` instead, depending on the usage
type AttrHost string // http.host

func (AttrHost) Development() {}
func (AttrHost) Recommended() {}

// Deprecated, use `http.request.method` instead
type AttrMethod string // http.method

func (AttrMethod) Development() {}
func (AttrMethod) Recommended() {}

// The size of the request payload body in bytes. This is the number of bytes transferred excluding headers and is often, but not always, present as the [Content-Length] header. For requests using transport encoding, this should be the compressed size
//
// [Content-Length]: https://www.rfc-editor.org/rfc/rfc9110.html#field.content-length
type AttrRequestBodySize string // http.request.body.size

func (AttrRequestBodySize) Development() {}
func (AttrRequestBodySize) Recommended() {}

// HTTP request headers, `<key>` being the normalized HTTP Header name (lowercase), the value being the header values.
//
// Instrumentations SHOULD require an explicit configuration of which headers are to be captured.
// Including all request headers can be a security risk - explicit configuration helps avoid leaking sensitive information.
//
// The `User-Agent` header is already captured in the `user_agent.original` attribute.
// Users MAY explicitly configure instrumentations to capture them even though it is not recommended.
//
// The attribute value MUST consist of either multiple header values as an array of strings
// or a single-item array containing a possibly comma-concatenated string, depending on the way
// the HTTP library provides access to headers.
//
// Examples:
//
//   - A header `Content-Type: application/json` SHOULD be recorded as the `http.request.header.content-type`
//     attribute with value `["application/json"]`.
//   - A header `X-Forwarded-For: 1.2.3.4, 1.2.3.5` SHOULD be recorded as the `http.request.header.x-forwarded-for`
//     attribute with value `["1.2.3.4", "1.2.3.5"]` or `["1.2.3.4, 1.2.3.5"]` depending on the HTTP library
type AttrRequestHeader string // http.request.header

func (AttrRequestHeader) Stable()      {}
func (AttrRequestHeader) Recommended() {}

// HTTP request method.
// HTTP request method value SHOULD be "known" to the instrumentation.
// By default, this convention defines "known" methods as the ones listed in [RFC9110]
// and the PATCH method defined in [RFC5789].
//
// If the HTTP request method is not known to instrumentation, it MUST set the `http.request.method` attribute to `_OTHER`.
//
// If the HTTP instrumentation could end up converting valid HTTP request methods to `_OTHER`, then it MUST provide a way to override
// the list of known HTTP methods. If this override is done via environment variable, then the environment variable MUST be named
// OTEL_INSTRUMENTATION_HTTP_KNOWN_METHODS and support a comma-separated list of case-sensitive known HTTP methods
// (this list MUST be a full override of the default known method, it is not a list of known methods in addition to the defaults).
//
// HTTP method names are case-sensitive and `http.request.method` attribute value MUST match a known HTTP method name exactly.
// Instrumentations for specific web frameworks that consider HTTP methods to be case insensitive, SHOULD populate a canonical equivalent.
// Tracing instrumentations that do so, MUST also set `http.request.method_original` to the original value
//
// [RFC9110]: https://www.rfc-editor.org/rfc/rfc9110.html#name-methods
// [RFC5789]: https://www.rfc-editor.org/rfc/rfc5789.html
type AttrRequestMethod string // http.request.method

func (AttrRequestMethod) Stable()      {}
func (AttrRequestMethod) Recommended() {}

// Original HTTP method sent by the client in the request line
type AttrRequestMethodOriginal string // http.request.method_original

func (AttrRequestMethodOriginal) Stable()      {}
func (AttrRequestMethodOriginal) Recommended() {}

// The ordinal number of request resending attempt (for any reason, including redirects).
//
// The resend count SHOULD be updated each time an HTTP request gets resent by the client, regardless of what was the cause of the resending (e.g. redirection, authorization failure, 503 Server Unavailable, network issues, or any other)
type AttrRequestResendCount string // http.request.resend_count

func (AttrRequestResendCount) Stable()      {}
func (AttrRequestResendCount) Recommended() {}

// The total size of the request in bytes. This should be the total number of bytes sent over the wire, including the request line (HTTP/1.1), framing (HTTP/2 and HTTP/3), headers, and request body if any
type AttrRequestSize string // http.request.size

func (AttrRequestSize) Development() {}
func (AttrRequestSize) Recommended() {}

// Deprecated, use `http.request.header.content-length` instead
type AttrRequestContentLength string // http.request_content_length

func (AttrRequestContentLength) Development() {}
func (AttrRequestContentLength) Recommended() {}

// Deprecated, use `http.request.body.size` instead
type AttrRequestContentLengthUncompressed string // http.request_content_length_uncompressed

func (AttrRequestContentLengthUncompressed) Development() {}
func (AttrRequestContentLengthUncompressed) Recommended() {}

// The size of the response payload body in bytes. This is the number of bytes transferred excluding headers and is often, but not always, present as the [Content-Length] header. For requests using transport encoding, this should be the compressed size
//
// [Content-Length]: https://www.rfc-editor.org/rfc/rfc9110.html#field.content-length
type AttrResponseBodySize string // http.response.body.size

func (AttrResponseBodySize) Development() {}
func (AttrResponseBodySize) Recommended() {}

// HTTP response headers, `<key>` being the normalized HTTP Header name (lowercase), the value being the header values.
//
// Instrumentations SHOULD require an explicit configuration of which headers are to be captured.
// Including all response headers can be a security risk - explicit configuration helps avoid leaking sensitive information.
//
// Users MAY explicitly configure instrumentations to capture them even though it is not recommended.
//
// The attribute value MUST consist of either multiple header values as an array of strings
// or a single-item array containing a possibly comma-concatenated string, depending on the way
// the HTTP library provides access to headers.
//
// Examples:
//
//   - A header `Content-Type: application/json` header SHOULD be recorded as the `http.request.response.content-type`
//     attribute with value `["application/json"]`.
//   - A header `My-custom-header: abc, def` header SHOULD be recorded as the `http.response.header.my-custom-header`
//     attribute with value `["abc", "def"]` or `["abc, def"]` depending on the HTTP library
type AttrResponseHeader string // http.response.header

func (AttrResponseHeader) Stable()      {}
func (AttrResponseHeader) Recommended() {}

// The total size of the response in bytes. This should be the total number of bytes sent over the wire, including the status line (HTTP/1.1), framing (HTTP/2 and HTTP/3), headers, and response body and trailers if any
type AttrResponseSize string // http.response.size

func (AttrResponseSize) Development() {}
func (AttrResponseSize) Recommended() {}

// [HTTP response status code]
//
// [HTTP response status code]: https://tools.ietf.org/html/rfc7231#section-6
type AttrResponseStatusCode string // http.response.status_code

func (AttrResponseStatusCode) Stable()      {}
func (AttrResponseStatusCode) Recommended() {}

// Deprecated, use `http.response.header.content-length` instead
type AttrResponseContentLength string // http.response_content_length

func (AttrResponseContentLength) Development() {}
func (AttrResponseContentLength) Recommended() {}

// Deprecated, use `http.response.body.size` instead
type AttrResponseContentLengthUncompressed string // http.response_content_length_uncompressed

func (AttrResponseContentLengthUncompressed) Development() {}
func (AttrResponseContentLengthUncompressed) Recommended() {}

// The matched route, that is, the path template in the format used by the respective server framework.
//
// MUST NOT be populated when this is not supported by the HTTP server framework as the route attribute should have low-cardinality and the URI path can NOT substitute it.
// SHOULD include the [application root] if there is one
//
// [application root]: /docs/http/http-spans.md#http-server-definitions
type AttrRoute string // http.route

func (AttrRoute) Stable()      {}
func (AttrRoute) Recommended() {}

// Deprecated, use `url.scheme` instead
type AttrScheme string // http.scheme

func (AttrScheme) Development() {}
func (AttrScheme) Recommended() {}

// Deprecated, use `server.address` instead
type AttrServerName string // http.server_name

func (AttrServerName) Development() {}
func (AttrServerName) Recommended() {}

// Deprecated, use `http.response.status_code` instead
type AttrStatusCode string // http.status_code

func (AttrStatusCode) Development() {}
func (AttrStatusCode) Recommended() {}

// Deprecated, use `url.path` and `url.query` instead
type AttrTarget string // http.target

func (AttrTarget) Development() {}
func (AttrTarget) Recommended() {}

// Deprecated, use `url.full` instead
type AttrUrl string // http.url

func (AttrUrl) Development() {}
func (AttrUrl) Recommended() {}

// Deprecated, use `user_agent.original` instead
type AttrUserAgent string // http.user_agent

func (AttrUserAgent) Development() {}
func (AttrUserAgent) Recommended() {}

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "Deprecated, use `client.address` instead.",
                    "deprecated": {
                        "note": "Replaced by `client.address`.",
                        "reason": "renamed",
                        "renamed_to": "client.address",
                    },
                    "examples": "83.164.160.102",
                    "name": "http.client_ip",
                    "requirement_level": "recommended",
                    "root_namespace": "http",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "State of the HTTP connection in the HTTP connection pool.",
                    "examples": [
                        "active",
                        "idle",
                    ],
                    "name": "http.connection.state",
                    "requirement_level": "recommended",
                    "root_namespace": "http",
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "active state.",
                                "deprecated": none,
                                "id": "active",
                                "note": none,
                                "stability": "development",
                                "value": "active",
                            },
                            {
                                "brief": "idle state.",
                                "deprecated": none,
                                "id": "idle",
                                "note": none,
                                "stability": "development",
                                "value": "idle",
                            },
                        ],
                    },
                },
                {
                    "brief": "Deprecated, use `network.protocol.name` instead.",
                    "deprecated": {
                        "note": "Replaced by `network.protocol.name`.",
                        "reason": "renamed",
                        "renamed_to": "network.protocol.name",
                    },
                    "name": "http.flavor",
                    "requirement_level": "recommended",
                    "root_namespace": "http",
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "HTTP/1.0",
                                "deprecated": none,
                                "id": "http_1_0",
                                "note": none,
                                "stability": "development",
                                "value": "1.0",
                            },
                            {
                                "brief": "HTTP/1.1",
                                "deprecated": none,
                                "id": "http_1_1",
                                "note": none,
                                "stability": "development",
                                "value": "1.1",
                            },
                            {
                                "brief": "HTTP/2",
                                "deprecated": none,
                                "id": "http_2_0",
                                "note": none,
                                "stability": "development",
                                "value": "2.0",
                            },
                            {
                                "brief": "HTTP/3",
                                "deprecated": none,
                                "id": "http_3_0",
                                "note": none,
                                "stability": "development",
                                "value": "3.0",
                            },
                            {
                                "brief": "SPDY protocol.",
                                "deprecated": none,
                                "id": "spdy",
                                "note": none,
                                "stability": "development",
                                "value": "SPDY",
                            },
                            {
                                "brief": "QUIC protocol.",
                                "deprecated": none,
                                "id": "quic",
                                "note": none,
                                "stability": "development",
                                "value": "QUIC",
                            },
                        ],
                    },
                },
                {
                    "brief": "Deprecated, use one of `server.address`, `client.address` or `http.request.header.host` instead, depending on the usage.",
                    "deprecated": {
                        "note": "Replaced by one of `server.address`, `client.address` or `http.request.header.host`, depending on the usage.\n",
                        "reason": "uncategorized",
                    },
                    "examples": [
                        "www.example.org",
                    ],
                    "name": "http.host",
                    "requirement_level": "recommended",
                    "root_namespace": "http",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `http.request.method` instead.",
                    "deprecated": {
                        "note": "Replaced by `http.request.method`.",
                        "reason": "renamed",
                        "renamed_to": "http.request.method",
                    },
                    "examples": [
                        "GET",
                        "POST",
                        "HEAD",
                    ],
                    "name": "http.method",
                    "requirement_level": "recommended",
                    "root_namespace": "http",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The size of the request payload body in bytes. This is the number of bytes transferred excluding headers and is often, but not always, present as the [Content-Length](https://www.rfc-editor.org/rfc/rfc9110.html#field.content-length) header. For requests using transport encoding, this should be the compressed size.\n",
                    "examples": 3495,
                    "name": "http.request.body.size",
                    "requirement_level": "recommended",
                    "root_namespace": "http",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "HTTP request headers, `<key>` being the normalized HTTP Header name (lowercase), the value being the header values.\n",
                    "examples": [
                        [
                            "application/json",
                        ],
                        [
                            "1.2.3.4",
                            "1.2.3.5",
                        ],
                    ],
                    "name": "http.request.header",
                    "note": "Instrumentations SHOULD require an explicit configuration of which headers are to be captured.\nIncluding all request headers can be a security risk - explicit configuration helps avoid leaking sensitive information.\n\nThe `User-Agent` header is already captured in the `user_agent.original` attribute.\nUsers MAY explicitly configure instrumentations to capture them even though it is not recommended.\n\nThe attribute value MUST consist of either multiple header values as an array of strings\nor a single-item array containing a possibly comma-concatenated string, depending on the way\nthe HTTP library provides access to headers.\n\nExamples:\n\n- A header `Content-Type: application/json` SHOULD be recorded as the `http.request.header.content-type`\n  attribute with value `[\"application/json\"]`.\n- A header `X-Forwarded-For: 1.2.3.4, 1.2.3.5` SHOULD be recorded as the `http.request.header.x-forwarded-for`\n  attribute with value `[\"1.2.3.4\", \"1.2.3.5\"]` or `[\"1.2.3.4, 1.2.3.5\"]` depending on the HTTP library.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "http",
                    "stability": "stable",
                    "type": "template[string[]]",
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
                    "requirement_level": "recommended",
                    "root_namespace": "http",
                    "stability": "stable",
                    "type": {
                        "allow_custom_values": none,
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
                    "brief": "Original HTTP method sent by the client in the request line.",
                    "examples": [
                        "GeT",
                        "ACL",
                        "foo",
                    ],
                    "name": "http.request.method_original",
                    "requirement_level": "recommended",
                    "root_namespace": "http",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "The ordinal number of request resending attempt (for any reason, including redirects).\n",
                    "examples": 3,
                    "name": "http.request.resend_count",
                    "note": "The resend count SHOULD be updated each time an HTTP request gets resent by the client, regardless of what was the cause of the resending (e.g. redirection, authorization failure, 503 Server Unavailable, network issues, or any other).\n",
                    "requirement_level": "recommended",
                    "root_namespace": "http",
                    "stability": "stable",
                    "type": "int",
                },
                {
                    "brief": "The total size of the request in bytes. This should be the total number of bytes sent over the wire, including the request line (HTTP/1.1), framing (HTTP/2 and HTTP/3), headers, and request body if any.\n",
                    "examples": 1437,
                    "name": "http.request.size",
                    "requirement_level": "recommended",
                    "root_namespace": "http",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "Deprecated, use `http.request.header.content-length` instead.",
                    "deprecated": {
                        "note": "Replaced by `http.request.header.content-length`.",
                        "reason": "uncategorized",
                    },
                    "examples": 3495,
                    "name": "http.request_content_length",
                    "requirement_level": "recommended",
                    "root_namespace": "http",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "Deprecated, use `http.request.body.size` instead.",
                    "deprecated": {
                        "note": "Replaced by `http.request.body.size`.",
                        "reason": "renamed",
                        "renamed_to": "http.request.body.size",
                    },
                    "examples": 5493,
                    "name": "http.request_content_length_uncompressed",
                    "requirement_level": "recommended",
                    "root_namespace": "http",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "The size of the response payload body in bytes. This is the number of bytes transferred excluding headers and is often, but not always, present as the [Content-Length](https://www.rfc-editor.org/rfc/rfc9110.html#field.content-length) header. For requests using transport encoding, this should be the compressed size.\n",
                    "examples": 3495,
                    "name": "http.response.body.size",
                    "requirement_level": "recommended",
                    "root_namespace": "http",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "HTTP response headers, `<key>` being the normalized HTTP Header name (lowercase), the value being the header values.\n",
                    "examples": [
                        [
                            "application/json",
                        ],
                        [
                            "abc",
                            "def",
                        ],
                    ],
                    "name": "http.response.header",
                    "note": "Instrumentations SHOULD require an explicit configuration of which headers are to be captured.\nIncluding all response headers can be a security risk - explicit configuration helps avoid leaking sensitive information.\n\nUsers MAY explicitly configure instrumentations to capture them even though it is not recommended.\n\nThe attribute value MUST consist of either multiple header values as an array of strings\nor a single-item array containing a possibly comma-concatenated string, depending on the way\nthe HTTP library provides access to headers.\n\nExamples:\n\n- A header `Content-Type: application/json` header SHOULD be recorded as the `http.request.response.content-type`\n  attribute with value `[\"application/json\"]`.\n- A header `My-custom-header: abc, def` header SHOULD be recorded as the `http.response.header.my-custom-header`\n  attribute with value `[\"abc\", \"def\"]` or `[\"abc, def\"]` depending on the HTTP library.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "http",
                    "stability": "stable",
                    "type": "template[string[]]",
                },
                {
                    "brief": "The total size of the response in bytes. This should be the total number of bytes sent over the wire, including the status line (HTTP/1.1), framing (HTTP/2 and HTTP/3), headers, and response body and trailers if any.\n",
                    "examples": 1437,
                    "name": "http.response.size",
                    "requirement_level": "recommended",
                    "root_namespace": "http",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "[HTTP response status code](https://tools.ietf.org/html/rfc7231#section-6).",
                    "examples": [
                        200,
                    ],
                    "name": "http.response.status_code",
                    "requirement_level": "recommended",
                    "root_namespace": "http",
                    "stability": "stable",
                    "type": "int",
                },
                {
                    "brief": "Deprecated, use `http.response.header.content-length` instead.",
                    "deprecated": {
                        "note": "hp.response.header.content-length",
                        "reason": "uncategorized",
                    },
                    "examples": 3495,
                    "name": "http.response_content_length",
                    "requirement_level": "recommended",
                    "root_namespace": "http",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "Deprecated, use `http.response.body.size` instead.",
                    "deprecated": {
                        "note": "Replaced by `http.response.body.size`.",
                        "reason": "renamed",
                        "renamed_to": "http.response.body.size",
                    },
                    "examples": 5493,
                    "name": "http.response_content_length_uncompressed",
                    "requirement_level": "recommended",
                    "root_namespace": "http",
                    "stability": "development",
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
                    "requirement_level": "recommended",
                    "root_namespace": "http",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `url.scheme` instead.",
                    "deprecated": {
                        "note": "Replaced by `url.scheme`.",
                        "reason": "renamed",
                        "renamed_to": "url.scheme",
                    },
                    "examples": [
                        "http",
                        "https",
                    ],
                    "name": "http.scheme",
                    "requirement_level": "recommended",
                    "root_namespace": "http",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `server.address` instead.",
                    "deprecated": {
                        "note": "Replaced by `server.address`.",
                        "reason": "renamed",
                        "renamed_to": "server.address",
                    },
                    "examples": [
                        "example.com",
                    ],
                    "name": "http.server_name",
                    "requirement_level": "recommended",
                    "root_namespace": "http",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `http.response.status_code` instead.",
                    "deprecated": {
                        "note": "Replaced by `http.response.status_code`.",
                        "reason": "renamed",
                        "renamed_to": "http.response.status_code",
                    },
                    "examples": [
                        200,
                    ],
                    "name": "http.status_code",
                    "requirement_level": "recommended",
                    "root_namespace": "http",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "Deprecated, use `url.path` and `url.query` instead.",
                    "deprecated": {
                        "note": "Split to `url.path` and `url.query`.",
                        "reason": "obsoleted",
                    },
                    "examples": [
                        "/search?q=OpenTelemetry#SemConv",
                    ],
                    "name": "http.target",
                    "requirement_level": "recommended",
                    "root_namespace": "http",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `url.full` instead.",
                    "deprecated": {
                        "note": "Replaced by `url.full`.",
                        "reason": "renamed",
                        "renamed_to": "url.full",
                    },
                    "examples": [
                        "https://www.foo.bar/search?q=OpenTelemetry#SemConv",
                    ],
                    "name": "http.url",
                    "requirement_level": "recommended",
                    "root_namespace": "http",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `user_agent.original` instead.",
                    "deprecated": {
                        "note": "Replaced by `user_agent.original`.",
                        "reason": "renamed",
                        "renamed_to": "user_agent.original",
                    },
                    "examples": [
                        "CERN-LineMode/2.15 libwww/2.17b3",
                        "Mozilla/5.0 (iPhone; CPU iPhone OS 14_7_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1.2 Mobile/15E148 Safari/604.1",
                    ],
                    "name": "http.user_agent",
                    "requirement_level": "recommended",
                    "root_namespace": "http",
                    "stability": "development",
                    "type": "string",
                },
            ],
            "root_namespace": "http",
        },
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
            "attr.go.j2",
        ],
    },
} */
