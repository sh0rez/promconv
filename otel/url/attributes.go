package url

// Domain extracted from the `url.full`, such as "opentelemetry.io".
//
// In some cases a URL may refer to an IP and/or port directly, without a domain name. In this case, the IP address would go to the domain field. If the URL contains a [literal IPv6 address] enclosed by `[` and `]`, the `[` and `]` characters should also be captured in the domain field
//
// [literal IPv6 address]: https://www.rfc-editor.org/rfc/rfc2732#section-2
type AttrDomain string // url.domain

func (AttrDomain) Development() {}
func (AttrDomain) Recommended() {}

// The file extension extracted from the `url.full`, excluding the leading dot.
//
// The file extension is only set if it exists, as not every url has a file extension. When the file name has multiple extensions `example.tar.gz`, only the last one should be captured `gz`, not `tar.gz`
type AttrExtension string // url.extension

func (AttrExtension) Development() {}
func (AttrExtension) Recommended() {}

// The [URI fragment] component
//
// [URI fragment]: https://www.rfc-editor.org/rfc/rfc3986#section-3.5
type AttrFragment string // url.fragment

func (AttrFragment) Stable()      {}
func (AttrFragment) Recommended() {}

// Absolute URL describing a network resource according to [RFC3986]
// For network calls, URL usually has `scheme://host[:port][path][?query][#fragment]` format, where the fragment
// is not transmitted over HTTP, but if it is known, it SHOULD be included nevertheless.
//
// `url.full` MUST NOT contain credentials passed via URL in form of `https://username:password@www.example.com/`.
// In such case username and password SHOULD be redacted and attribute's value SHOULD be `https://REDACTED:REDACTED@www.example.com/`.
//
// `url.full` SHOULD capture the absolute URL when it is available (or can be reconstructed).
//
// Sensitive content provided in `url.full` SHOULD be scrubbed when instrumentations can identify it.
//
// Query string values for the following keys SHOULD be redacted by default and replaced by the
// value `REDACTED`:
//
//   - [`AWSAccessKeyId`]
//   - [`Signature`]
//   - [`sig`]
//   - [`X-Goog-Signature`]
//
// This list is subject to change over time.
//
// When a query string value is redacted, the query string key SHOULD still be preserved, e.g.
// `https://www.example.com/path?color=blue&sig=REDACTED`
//
// [RFC3986]: https://www.rfc-editor.org/rfc/rfc3986
// [`AWSAccessKeyId`]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/RESTAuthentication.html#RESTAuthenticationQueryStringAuth
// [`Signature`]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/RESTAuthentication.html#RESTAuthenticationQueryStringAuth
// [`sig`]: https://learn.microsoft.com/azure/storage/common/storage-sas-overview#sas-token
// [`X-Goog-Signature`]: https://cloud.google.com/storage/docs/access-control/signed-urls
type AttrFull string // url.full

func (AttrFull) Stable()      {}
func (AttrFull) Recommended() {}

// Unmodified original URL as seen in the event source.
//
// In network monitoring, the observed URL may be a full URL, whereas in access logs, the URL is often just represented as a path. This field is meant to represent the URL as it was observed, complete or not.
// `url.original` might contain credentials passed via URL in form of `https://username:password@www.example.com/`. In such case password and username SHOULD NOT be redacted and attribute's value SHOULD remain the same
type AttrOriginal string // url.original

func (AttrOriginal) Development() {}
func (AttrOriginal) Recommended() {}

// The [URI path] component
//
// # Sensitive content provided in `url.path` SHOULD be scrubbed when instrumentations can identify it
//
// [URI path]: https://www.rfc-editor.org/rfc/rfc3986#section-3.3
type AttrPath string // url.path

func (AttrPath) Stable()      {}
func (AttrPath) Recommended() {}

// Port extracted from the `url.full`
type AttrPort string // url.port

func (AttrPort) Development() {}
func (AttrPort) Recommended() {}

// The [URI query] component
//
// Sensitive content provided in `url.query` SHOULD be scrubbed when instrumentations can identify it.
//
// Query string values for the following keys SHOULD be redacted by default and replaced by the value `REDACTED`:
//
//   - [`AWSAccessKeyId`]
//   - [`Signature`]
//   - [`sig`]
//   - [`X-Goog-Signature`]
//
// This list is subject to change over time.
//
// When a query string value is redacted, the query string key SHOULD still be preserved, e.g.
// `q=OpenTelemetry&sig=REDACTED`
//
// [URI query]: https://www.rfc-editor.org/rfc/rfc3986#section-3.4
// [`AWSAccessKeyId`]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/RESTAuthentication.html#RESTAuthenticationQueryStringAuth
// [`Signature`]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/RESTAuthentication.html#RESTAuthenticationQueryStringAuth
// [`sig`]: https://learn.microsoft.com/azure/storage/common/storage-sas-overview#sas-token
// [`X-Goog-Signature`]: https://cloud.google.com/storage/docs/access-control/signed-urls
type AttrQuery string // url.query

func (AttrQuery) Stable()      {}
func (AttrQuery) Recommended() {}

// The highest registered url domain, stripped of the subdomain.
//
// This value can be determined precisely with the [public suffix list]. For example, the registered domain for `foo.example.com` is `example.com`. Trying to approximate this by simply taking the last two labels will not work well for TLDs such as `co.uk`
//
// [public suffix list]: https://publicsuffix.org/
type AttrRegisteredDomain string // url.registered_domain

func (AttrRegisteredDomain) Development() {}
func (AttrRegisteredDomain) Recommended() {}

// The [URI scheme] component identifying the used protocol
//
// [URI scheme]: https://www.rfc-editor.org/rfc/rfc3986#section-3.1
type AttrScheme string // url.scheme

func (AttrScheme) Stable()      {}
func (AttrScheme) Recommended() {}

// The subdomain portion of a fully qualified domain name includes all of the names except the host name under the registered_domain. In a partially qualified domain, or if the qualification level of the full name cannot be determined, subdomain contains all of the names below the registered domain.
//
// The subdomain portion of `www.east.mydomain.co.uk` is `east`. If the domain has multiple levels of subdomain, such as `sub2.sub1.example.com`, the subdomain field should contain `sub2.sub1`, with no trailing period
type AttrSubdomain string // url.subdomain

func (AttrSubdomain) Development() {}
func (AttrSubdomain) Recommended() {}

// The low-cardinality template of an [absolute path reference]
//
// [absolute path reference]: https://www.rfc-editor.org/rfc/rfc3986#section-4.2
type AttrTemplate string // url.template

func (AttrTemplate) Development() {}
func (AttrTemplate) Recommended() {}

// The effective top level domain (eTLD), also known as the domain suffix, is the last part of the domain name. For example, the top level domain for example.com is `com`.
//
// This value can be determined precisely with the [public suffix list]
//
// [public suffix list]: https://publicsuffix.org/
type AttrTopLevelDomain string // url.top_level_domain

func (AttrTopLevelDomain) Development() {}
func (AttrTopLevelDomain) Recommended() {}

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "Domain extracted from the `url.full`, such as \"opentelemetry.io\".\n",
                    "examples": [
                        "www.foo.bar",
                        "opentelemetry.io",
                        "3.12.167.2",
                        "[1080:0:0:0:8:800:200C:417A]",
                    ],
                    "name": "url.domain",
                    "note": "In some cases a URL may refer to an IP and/or port directly, without a domain name. In this case, the IP address would go to the domain field. If the URL contains a [literal IPv6 address](https://www.rfc-editor.org/rfc/rfc2732#section-2) enclosed by `[` and `]`, the `[` and `]` characters should also be captured in the domain field.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "url",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The file extension extracted from the `url.full`, excluding the leading dot.\n",
                    "examples": [
                        "png",
                        "gz",
                    ],
                    "name": "url.extension",
                    "note": "The file extension is only set if it exists, as not every url has a file extension. When the file name has multiple extensions `example.tar.gz`, only the last one should be captured `gz`, not `tar.gz`.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "url",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The [URI fragment](https://www.rfc-editor.org/rfc/rfc3986#section-3.5) component\n",
                    "examples": [
                        "SemConv",
                    ],
                    "name": "url.fragment",
                    "requirement_level": "recommended",
                    "root_namespace": "url",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Absolute URL describing a network resource according to [RFC3986](https://www.rfc-editor.org/rfc/rfc3986)",
                    "examples": [
                        "https://www.foo.bar/search?q=OpenTelemetry#SemConv",
                        "//localhost",
                    ],
                    "name": "url.full",
                    "note": "For network calls, URL usually has `scheme://host[:port][path][?query][#fragment]` format, where the fragment\nis not transmitted over HTTP, but if it is known, it SHOULD be included nevertheless.\n\n`url.full` MUST NOT contain credentials passed via URL in form of `https://username:password@www.example.com/`.\nIn such case username and password SHOULD be redacted and attribute's value SHOULD be `https://REDACTED:REDACTED@www.example.com/`.\n\n`url.full` SHOULD capture the absolute URL when it is available (or can be reconstructed).\n\nSensitive content provided in `url.full` SHOULD be scrubbed when instrumentations can identify it.\n\n![Development](https://img.shields.io/badge/-development-blue)\nQuery string values for the following keys SHOULD be redacted by default and replaced by the\nvalue `REDACTED`:\n\n* [`AWSAccessKeyId`](https://docs.aws.amazon.com/AmazonS3/latest/userguide/RESTAuthentication.html#RESTAuthenticationQueryStringAuth)\n* [`Signature`](https://docs.aws.amazon.com/AmazonS3/latest/userguide/RESTAuthentication.html#RESTAuthenticationQueryStringAuth)\n* [`sig`](https://learn.microsoft.com/azure/storage/common/storage-sas-overview#sas-token)\n* [`X-Goog-Signature`](https://cloud.google.com/storage/docs/access-control/signed-urls)\n\nThis list is subject to change over time.\n\nWhen a query string value is redacted, the query string key SHOULD still be preserved, e.g.\n`https://www.example.com/path?color=blue&sig=REDACTED`.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "url",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Unmodified original URL as seen in the event source.\n",
                    "examples": [
                        "https://www.foo.bar/search?q=OpenTelemetry#SemConv",
                        "search?q=OpenTelemetry",
                    ],
                    "name": "url.original",
                    "note": "In network monitoring, the observed URL may be a full URL, whereas in access logs, the URL is often just represented as a path. This field is meant to represent the URL as it was observed, complete or not.\n`url.original` might contain credentials passed via URL in form of `https://username:password@www.example.com/`. In such case password and username SHOULD NOT be redacted and attribute's value SHOULD remain the same.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "url",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The [URI path](https://www.rfc-editor.org/rfc/rfc3986#section-3.3) component\n",
                    "examples": [
                        "/search",
                    ],
                    "name": "url.path",
                    "note": "Sensitive content provided in `url.path` SHOULD be scrubbed when instrumentations can identify it.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "url",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "Port extracted from the `url.full`\n",
                    "examples": [
                        443,
                    ],
                    "name": "url.port",
                    "requirement_level": "recommended",
                    "root_namespace": "url",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "The [URI query](https://www.rfc-editor.org/rfc/rfc3986#section-3.4) component\n",
                    "examples": [
                        "q=OpenTelemetry",
                    ],
                    "name": "url.query",
                    "note": "Sensitive content provided in `url.query` SHOULD be scrubbed when instrumentations can identify it.\n\n![Development](https://img.shields.io/badge/-development-blue)\nQuery string values for the following keys SHOULD be redacted by default and replaced by the value `REDACTED`:\n\n* [`AWSAccessKeyId`](https://docs.aws.amazon.com/AmazonS3/latest/userguide/RESTAuthentication.html#RESTAuthenticationQueryStringAuth)\n* [`Signature`](https://docs.aws.amazon.com/AmazonS3/latest/userguide/RESTAuthentication.html#RESTAuthenticationQueryStringAuth)\n* [`sig`](https://learn.microsoft.com/azure/storage/common/storage-sas-overview#sas-token)\n* [`X-Goog-Signature`](https://cloud.google.com/storage/docs/access-control/signed-urls)\n\nThis list is subject to change over time.\n\nWhen a query string value is redacted, the query string key SHOULD still be preserved, e.g.\n`q=OpenTelemetry&sig=REDACTED`.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "url",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "The highest registered url domain, stripped of the subdomain.\n",
                    "examples": [
                        "example.com",
                        "foo.co.uk",
                    ],
                    "name": "url.registered_domain",
                    "note": "This value can be determined precisely with the [public suffix list](https://publicsuffix.org/). For example, the registered domain for `foo.example.com` is `example.com`. Trying to approximate this by simply taking the last two labels will not work well for TLDs such as `co.uk`.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "url",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The [URI scheme](https://www.rfc-editor.org/rfc/rfc3986#section-3.1) component identifying the used protocol.\n",
                    "examples": [
                        "https",
                        "ftp",
                        "telnet",
                    ],
                    "name": "url.scheme",
                    "requirement_level": "recommended",
                    "root_namespace": "url",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "The subdomain portion of a fully qualified domain name includes all of the names except the host name under the registered_domain. In a partially qualified domain, or if the qualification level of the full name cannot be determined, subdomain contains all of the names below the registered domain.\n",
                    "examples": [
                        "east",
                        "sub2.sub1",
                    ],
                    "name": "url.subdomain",
                    "note": "The subdomain portion of `www.east.mydomain.co.uk` is `east`. If the domain has multiple levels of subdomain, such as `sub2.sub1.example.com`, the subdomain field should contain `sub2.sub1`, with no trailing period.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "url",
                    "stability": "development",
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
                    "requirement_level": "recommended",
                    "root_namespace": "url",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The effective top level domain (eTLD), also known as the domain suffix, is the last part of the domain name. For example, the top level domain for example.com is `com`.\n",
                    "examples": [
                        "com",
                        "co.uk",
                    ],
                    "name": "url.top_level_domain",
                    "note": "This value can be determined precisely with the [public suffix list](https://publicsuffix.org/).\n",
                    "requirement_level": "recommended",
                    "root_namespace": "url",
                    "stability": "development",
                    "type": "string",
                },
            ],
            "root_namespace": "url",
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
