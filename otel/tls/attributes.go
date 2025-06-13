package tls

// String indicating the [cipher] used during the current connection.
//
// The values allowed for `tls.cipher` MUST be one of the `Descriptions` of the [registered TLS Cipher Suits]
//
// [cipher]: https://datatracker.ietf.org/doc/html/rfc5246#appendix-A.5
// [registered TLS Cipher Suits]: https://www.iana.org/assignments/tls-parameters/tls-parameters.xhtml#table-tls-parameters-4
type AttrCipher string // tls.cipher

func (AttrCipher) Development()    {}
func (AttrCipher) Recommended()    {}
func (AttrCipher) Key() string     { return "tls_cipher" }
func (a AttrCipher) Value() string { return string(a) }

// PEM-encoded stand-alone certificate offered by the client. This is usually mutually-exclusive of `client.certificate_chain` since this value also exists in that list
type AttrClientCertificate string // tls.client.certificate

func (AttrClientCertificate) Development()    {}
func (AttrClientCertificate) Recommended()    {}
func (AttrClientCertificate) Key() string     { return "tls_client_certificate" }
func (a AttrClientCertificate) Value() string { return string(a) }

// Array of PEM-encoded certificates that make up the certificate chain offered by the client. This is usually mutually-exclusive of `client.certificate` since that value should be the first certificate in the chain
type AttrClientCertificateChain string // tls.client.certificate_chain

func (AttrClientCertificateChain) Development()    {}
func (AttrClientCertificateChain) Recommended()    {}
func (AttrClientCertificateChain) Key() string     { return "tls_client_certificate_chain" }
func (a AttrClientCertificateChain) Value() string { return string(a) }

// Certificate fingerprint using the MD5 digest of DER-encoded version of certificate offered by the client. For consistency with other hash values, this value should be formatted as an uppercase hash
type AttrClientHashMd5 string // tls.client.hash.md5

func (AttrClientHashMd5) Development()    {}
func (AttrClientHashMd5) Recommended()    {}
func (AttrClientHashMd5) Key() string     { return "tls_client_hash_md5" }
func (a AttrClientHashMd5) Value() string { return string(a) }

// Certificate fingerprint using the SHA1 digest of DER-encoded version of certificate offered by the client. For consistency with other hash values, this value should be formatted as an uppercase hash
type AttrClientHashSha1 string // tls.client.hash.sha1

func (AttrClientHashSha1) Development()    {}
func (AttrClientHashSha1) Recommended()    {}
func (AttrClientHashSha1) Key() string     { return "tls_client_hash_sha1" }
func (a AttrClientHashSha1) Value() string { return string(a) }

// Certificate fingerprint using the SHA256 digest of DER-encoded version of certificate offered by the client. For consistency with other hash values, this value should be formatted as an uppercase hash
type AttrClientHashSha256 string // tls.client.hash.sha256

func (AttrClientHashSha256) Development()    {}
func (AttrClientHashSha256) Recommended()    {}
func (AttrClientHashSha256) Key() string     { return "tls_client_hash_sha256" }
func (a AttrClientHashSha256) Value() string { return string(a) }

// Distinguished name of [subject] of the issuer of the x.509 certificate presented by the client
//
// [subject]: https://datatracker.ietf.org/doc/html/rfc5280#section-4.1.2.6
type AttrClientIssuer string // tls.client.issuer

func (AttrClientIssuer) Development()    {}
func (AttrClientIssuer) Recommended()    {}
func (AttrClientIssuer) Key() string     { return "tls_client_issuer" }
func (a AttrClientIssuer) Value() string { return string(a) }

// A hash that identifies clients based on how they perform an SSL/TLS handshake
type AttrClientJa3 string // tls.client.ja3

func (AttrClientJa3) Development()    {}
func (AttrClientJa3) Recommended()    {}
func (AttrClientJa3) Key() string     { return "tls_client_ja3" }
func (a AttrClientJa3) Value() string { return string(a) }

// Date/Time indicating when client certificate is no longer considered valid
type AttrClientNotAfter string // tls.client.not_after

func (AttrClientNotAfter) Development()    {}
func (AttrClientNotAfter) Recommended()    {}
func (AttrClientNotAfter) Key() string     { return "tls_client_not_after" }
func (a AttrClientNotAfter) Value() string { return string(a) }

// Date/Time indicating when client certificate is first considered valid
type AttrClientNotBefore string // tls.client.not_before

func (AttrClientNotBefore) Development()    {}
func (AttrClientNotBefore) Recommended()    {}
func (AttrClientNotBefore) Key() string     { return "tls_client_not_before" }
func (a AttrClientNotBefore) Value() string { return string(a) }

// Deprecated, use `server.address` instead
type AttrClientServerName string // tls.client.server_name

func (AttrClientServerName) Development()    {}
func (AttrClientServerName) Recommended()    {}
func (AttrClientServerName) Key() string     { return "tls_client_server_name" }
func (a AttrClientServerName) Value() string { return string(a) }

// Distinguished name of subject of the x.509 certificate presented by the client
type AttrClientSubject string // tls.client.subject

func (AttrClientSubject) Development()    {}
func (AttrClientSubject) Recommended()    {}
func (AttrClientSubject) Key() string     { return "tls_client_subject" }
func (a AttrClientSubject) Value() string { return string(a) }

// Array of ciphers offered by the client during the client hello
type AttrClientSupportedCiphers string // tls.client.supported_ciphers

func (AttrClientSupportedCiphers) Development()    {}
func (AttrClientSupportedCiphers) Recommended()    {}
func (AttrClientSupportedCiphers) Key() string     { return "tls_client_supported_ciphers" }
func (a AttrClientSupportedCiphers) Value() string { return string(a) }

// String indicating the curve used for the given cipher, when applicable
type AttrCurve string // tls.curve

func (AttrCurve) Development()    {}
func (AttrCurve) Recommended()    {}
func (AttrCurve) Key() string     { return "tls_curve" }
func (a AttrCurve) Value() string { return string(a) }

// Boolean flag indicating if the TLS negotiation was successful and transitioned to an encrypted tunnel
type AttrEstablished string // tls.established

func (AttrEstablished) Development()    {}
func (AttrEstablished) Recommended()    {}
func (AttrEstablished) Key() string     { return "tls_established" }
func (a AttrEstablished) Value() string { return string(a) }

// String indicating the protocol being tunneled. Per the values in the [IANA registry], this string should be lower case
//
// [IANA registry]: https://www.iana.org/assignments/tls-extensiontype-values/tls-extensiontype-values.xhtml#alpn-protocol-ids
type AttrNextProtocol string // tls.next_protocol

func (AttrNextProtocol) Development()    {}
func (AttrNextProtocol) Recommended()    {}
func (AttrNextProtocol) Key() string     { return "tls_next_protocol" }
func (a AttrNextProtocol) Value() string { return string(a) }

// Normalized lowercase protocol name parsed from original string of the negotiated [SSL/TLS protocol version]
//
// [SSL/TLS protocol version]: https://docs.openssl.org/1.1.1/man3/SSL_get_version/#return-values
type AttrProtocolName string // tls.protocol.name

func (AttrProtocolName) Development()    {}
func (AttrProtocolName) Recommended()    {}
func (AttrProtocolName) Key() string     { return "tls_protocol_name" }
func (a AttrProtocolName) Value() string { return string(a) }

const ProtocolNameSsl AttrProtocolName = "ssl"
const ProtocolNameTLS AttrProtocolName = "tls"

// Numeric part of the version parsed from the original string of the negotiated [SSL/TLS protocol version]
//
// [SSL/TLS protocol version]: https://docs.openssl.org/1.1.1/man3/SSL_get_version/#return-values
type AttrProtocolVersion string // tls.protocol.version

func (AttrProtocolVersion) Development()    {}
func (AttrProtocolVersion) Recommended()    {}
func (AttrProtocolVersion) Key() string     { return "tls_protocol_version" }
func (a AttrProtocolVersion) Value() string { return string(a) }

// Boolean flag indicating if this TLS connection was resumed from an existing TLS negotiation
type AttrResumed string // tls.resumed

func (AttrResumed) Development()    {}
func (AttrResumed) Recommended()    {}
func (AttrResumed) Key() string     { return "tls_resumed" }
func (a AttrResumed) Value() string { return string(a) }

// PEM-encoded stand-alone certificate offered by the server. This is usually mutually-exclusive of `server.certificate_chain` since this value also exists in that list
type AttrServerCertificate string // tls.server.certificate

func (AttrServerCertificate) Development()    {}
func (AttrServerCertificate) Recommended()    {}
func (AttrServerCertificate) Key() string     { return "tls_server_certificate" }
func (a AttrServerCertificate) Value() string { return string(a) }

// Array of PEM-encoded certificates that make up the certificate chain offered by the server. This is usually mutually-exclusive of `server.certificate` since that value should be the first certificate in the chain
type AttrServerCertificateChain string // tls.server.certificate_chain

func (AttrServerCertificateChain) Development()    {}
func (AttrServerCertificateChain) Recommended()    {}
func (AttrServerCertificateChain) Key() string     { return "tls_server_certificate_chain" }
func (a AttrServerCertificateChain) Value() string { return string(a) }

// Certificate fingerprint using the MD5 digest of DER-encoded version of certificate offered by the server. For consistency with other hash values, this value should be formatted as an uppercase hash
type AttrServerHashMd5 string // tls.server.hash.md5

func (AttrServerHashMd5) Development()    {}
func (AttrServerHashMd5) Recommended()    {}
func (AttrServerHashMd5) Key() string     { return "tls_server_hash_md5" }
func (a AttrServerHashMd5) Value() string { return string(a) }

// Certificate fingerprint using the SHA1 digest of DER-encoded version of certificate offered by the server. For consistency with other hash values, this value should be formatted as an uppercase hash
type AttrServerHashSha1 string // tls.server.hash.sha1

func (AttrServerHashSha1) Development()    {}
func (AttrServerHashSha1) Recommended()    {}
func (AttrServerHashSha1) Key() string     { return "tls_server_hash_sha1" }
func (a AttrServerHashSha1) Value() string { return string(a) }

// Certificate fingerprint using the SHA256 digest of DER-encoded version of certificate offered by the server. For consistency with other hash values, this value should be formatted as an uppercase hash
type AttrServerHashSha256 string // tls.server.hash.sha256

func (AttrServerHashSha256) Development()    {}
func (AttrServerHashSha256) Recommended()    {}
func (AttrServerHashSha256) Key() string     { return "tls_server_hash_sha256" }
func (a AttrServerHashSha256) Value() string { return string(a) }

// Distinguished name of [subject] of the issuer of the x.509 certificate presented by the client
//
// [subject]: https://datatracker.ietf.org/doc/html/rfc5280#section-4.1.2.6
type AttrServerIssuer string // tls.server.issuer

func (AttrServerIssuer) Development()    {}
func (AttrServerIssuer) Recommended()    {}
func (AttrServerIssuer) Key() string     { return "tls_server_issuer" }
func (a AttrServerIssuer) Value() string { return string(a) }

// A hash that identifies servers based on how they perform an SSL/TLS handshake
type AttrServerJa3s string // tls.server.ja3s

func (AttrServerJa3s) Development()    {}
func (AttrServerJa3s) Recommended()    {}
func (AttrServerJa3s) Key() string     { return "tls_server_ja3s" }
func (a AttrServerJa3s) Value() string { return string(a) }

// Date/Time indicating when server certificate is no longer considered valid
type AttrServerNotAfter string // tls.server.not_after

func (AttrServerNotAfter) Development()    {}
func (AttrServerNotAfter) Recommended()    {}
func (AttrServerNotAfter) Key() string     { return "tls_server_not_after" }
func (a AttrServerNotAfter) Value() string { return string(a) }

// Date/Time indicating when server certificate is first considered valid
type AttrServerNotBefore string // tls.server.not_before

func (AttrServerNotBefore) Development()    {}
func (AttrServerNotBefore) Recommended()    {}
func (AttrServerNotBefore) Key() string     { return "tls_server_not_before" }
func (a AttrServerNotBefore) Value() string { return string(a) }

// Distinguished name of subject of the x.509 certificate presented by the server
type AttrServerSubject string // tls.server.subject

func (AttrServerSubject) Development()    {}
func (AttrServerSubject) Recommended()    {}
func (AttrServerSubject) Key() string     { return "tls_server_subject" }
func (a AttrServerSubject) Value() string { return string(a) }

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "String indicating the [cipher](https://datatracker.ietf.org/doc/html/rfc5246#appendix-A.5) used during the current connection.\n",
                    "examples": [
                        "TLS_RSA_WITH_3DES_EDE_CBC_SHA",
                        "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256",
                    ],
                    "name": "tls.cipher",
                    "note": "The values allowed for `tls.cipher` MUST be one of the `Descriptions` of the [registered TLS Cipher Suits](https://www.iana.org/assignments/tls-parameters/tls-parameters.xhtml#table-tls-parameters-4).\n",
                    "requirement_level": "recommended",
                    "root_namespace": "tls",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "PEM-encoded stand-alone certificate offered by the client. This is usually mutually-exclusive of `client.certificate_chain` since this value also exists in that list.\n",
                    "examples": [
                        "MII...",
                    ],
                    "name": "tls.client.certificate",
                    "requirement_level": "recommended",
                    "root_namespace": "tls",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Array of PEM-encoded certificates that make up the certificate chain offered by the client. This is usually mutually-exclusive of `client.certificate` since that value should be the first certificate in the chain.\n",
                    "examples": [
                        [
                            "MII...",
                            "MI...",
                        ],
                    ],
                    "name": "tls.client.certificate_chain",
                    "requirement_level": "recommended",
                    "root_namespace": "tls",
                    "stability": "development",
                    "type": "string[]",
                },
                {
                    "brief": "Certificate fingerprint using the MD5 digest of DER-encoded version of certificate offered by the client. For consistency with other hash values, this value should be formatted as an uppercase hash.\n",
                    "examples": [
                        "0F76C7F2C55BFD7D8E8B8F4BFBF0C9EC",
                    ],
                    "name": "tls.client.hash.md5",
                    "requirement_level": "recommended",
                    "root_namespace": "tls",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Certificate fingerprint using the SHA1 digest of DER-encoded version of certificate offered by the client. For consistency with other hash values, this value should be formatted as an uppercase hash.\n",
                    "examples": [
                        "9E393D93138888D288266C2D915214D1D1CCEB2A",
                    ],
                    "name": "tls.client.hash.sha1",
                    "requirement_level": "recommended",
                    "root_namespace": "tls",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Certificate fingerprint using the SHA256 digest of DER-encoded version of certificate offered by the client. For consistency with other hash values, this value should be formatted as an uppercase hash.\n",
                    "examples": [
                        "0687F666A054EF17A08E2F2162EAB4CBC0D265E1D7875BE74BF3C712CA92DAF0",
                    ],
                    "name": "tls.client.hash.sha256",
                    "requirement_level": "recommended",
                    "root_namespace": "tls",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Distinguished name of [subject](https://datatracker.ietf.org/doc/html/rfc5280#section-4.1.2.6) of the issuer of the x.509 certificate presented by the client.",
                    "examples": [
                        "CN=Example Root CA, OU=Infrastructure Team, DC=example, DC=com",
                    ],
                    "name": "tls.client.issuer",
                    "requirement_level": "recommended",
                    "root_namespace": "tls",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "A hash that identifies clients based on how they perform an SSL/TLS handshake.",
                    "examples": [
                        "d4e5b18d6b55c71272893221c96ba240",
                    ],
                    "name": "tls.client.ja3",
                    "requirement_level": "recommended",
                    "root_namespace": "tls",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Date/Time indicating when client certificate is no longer considered valid.",
                    "examples": [
                        "2021-01-01T00:00:00.000Z",
                    ],
                    "name": "tls.client.not_after",
                    "requirement_level": "recommended",
                    "root_namespace": "tls",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Date/Time indicating when client certificate is first considered valid.",
                    "examples": [
                        "1970-01-01T00:00:00.000Z",
                    ],
                    "name": "tls.client.not_before",
                    "requirement_level": "recommended",
                    "root_namespace": "tls",
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
                        "opentelemetry.io",
                    ],
                    "name": "tls.client.server_name",
                    "requirement_level": "recommended",
                    "root_namespace": "tls",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Distinguished name of subject of the x.509 certificate presented by the client.",
                    "examples": [
                        "CN=myclient, OU=Documentation Team, DC=example, DC=com",
                    ],
                    "name": "tls.client.subject",
                    "requirement_level": "recommended",
                    "root_namespace": "tls",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Array of ciphers offered by the client during the client hello.",
                    "examples": [
                        [
                            "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
                            "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
                        ],
                    ],
                    "name": "tls.client.supported_ciphers",
                    "requirement_level": "recommended",
                    "root_namespace": "tls",
                    "stability": "development",
                    "type": "string[]",
                },
                {
                    "brief": "String indicating the curve used for the given cipher, when applicable",
                    "examples": [
                        "secp256r1",
                    ],
                    "name": "tls.curve",
                    "requirement_level": "recommended",
                    "root_namespace": "tls",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Boolean flag indicating if the TLS negotiation was successful and transitioned to an encrypted tunnel.",
                    "examples": [
                        true,
                    ],
                    "name": "tls.established",
                    "requirement_level": "recommended",
                    "root_namespace": "tls",
                    "stability": "development",
                    "type": "boolean",
                },
                {
                    "brief": "String indicating the protocol being tunneled. Per the values in the [IANA registry](https://www.iana.org/assignments/tls-extensiontype-values/tls-extensiontype-values.xhtml#alpn-protocol-ids), this string should be lower case.\n",
                    "examples": [
                        "http/1.1",
                    ],
                    "name": "tls.next_protocol",
                    "requirement_level": "recommended",
                    "root_namespace": "tls",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Normalized lowercase protocol name parsed from original string of the negotiated [SSL/TLS protocol version](https://docs.openssl.org/1.1.1/man3/SSL_get_version/#return-values)\n",
                    "name": "tls.protocol.name",
                    "requirement_level": "recommended",
                    "root_namespace": "tls",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "ssl",
                                "note": none,
                                "stability": "development",
                                "value": "ssl",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "tls",
                                "note": none,
                                "stability": "development",
                                "value": "tls",
                            },
                        ],
                    },
                },
                {
                    "brief": "Numeric part of the version parsed from the original string of the negotiated [SSL/TLS protocol version](https://docs.openssl.org/1.1.1/man3/SSL_get_version/#return-values)\n",
                    "examples": [
                        "1.2",
                        "3",
                    ],
                    "name": "tls.protocol.version",
                    "requirement_level": "recommended",
                    "root_namespace": "tls",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Boolean flag indicating if this TLS connection was resumed from an existing TLS negotiation.",
                    "examples": [
                        true,
                    ],
                    "name": "tls.resumed",
                    "requirement_level": "recommended",
                    "root_namespace": "tls",
                    "stability": "development",
                    "type": "boolean",
                },
                {
                    "brief": "PEM-encoded stand-alone certificate offered by the server. This is usually mutually-exclusive of `server.certificate_chain` since this value also exists in that list.\n",
                    "examples": [
                        "MII...",
                    ],
                    "name": "tls.server.certificate",
                    "requirement_level": "recommended",
                    "root_namespace": "tls",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Array of PEM-encoded certificates that make up the certificate chain offered by the server. This is usually mutually-exclusive of `server.certificate` since that value should be the first certificate in the chain.\n",
                    "examples": [
                        [
                            "MII...",
                            "MI...",
                        ],
                    ],
                    "name": "tls.server.certificate_chain",
                    "requirement_level": "recommended",
                    "root_namespace": "tls",
                    "stability": "development",
                    "type": "string[]",
                },
                {
                    "brief": "Certificate fingerprint using the MD5 digest of DER-encoded version of certificate offered by the server. For consistency with other hash values, this value should be formatted as an uppercase hash.\n",
                    "examples": [
                        "0F76C7F2C55BFD7D8E8B8F4BFBF0C9EC",
                    ],
                    "name": "tls.server.hash.md5",
                    "requirement_level": "recommended",
                    "root_namespace": "tls",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Certificate fingerprint using the SHA1 digest of DER-encoded version of certificate offered by the server. For consistency with other hash values, this value should be formatted as an uppercase hash.\n",
                    "examples": [
                        "9E393D93138888D288266C2D915214D1D1CCEB2A",
                    ],
                    "name": "tls.server.hash.sha1",
                    "requirement_level": "recommended",
                    "root_namespace": "tls",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Certificate fingerprint using the SHA256 digest of DER-encoded version of certificate offered by the server. For consistency with other hash values, this value should be formatted as an uppercase hash.\n",
                    "examples": [
                        "0687F666A054EF17A08E2F2162EAB4CBC0D265E1D7875BE74BF3C712CA92DAF0",
                    ],
                    "name": "tls.server.hash.sha256",
                    "requirement_level": "recommended",
                    "root_namespace": "tls",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Distinguished name of [subject](https://datatracker.ietf.org/doc/html/rfc5280#section-4.1.2.6) of the issuer of the x.509 certificate presented by the client.",
                    "examples": [
                        "CN=Example Root CA, OU=Infrastructure Team, DC=example, DC=com",
                    ],
                    "name": "tls.server.issuer",
                    "requirement_level": "recommended",
                    "root_namespace": "tls",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "A hash that identifies servers based on how they perform an SSL/TLS handshake.",
                    "examples": [
                        "d4e5b18d6b55c71272893221c96ba240",
                    ],
                    "name": "tls.server.ja3s",
                    "requirement_level": "recommended",
                    "root_namespace": "tls",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Date/Time indicating when server certificate is no longer considered valid.",
                    "examples": [
                        "2021-01-01T00:00:00.000Z",
                    ],
                    "name": "tls.server.not_after",
                    "requirement_level": "recommended",
                    "root_namespace": "tls",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Date/Time indicating when server certificate is first considered valid.",
                    "examples": [
                        "1970-01-01T00:00:00.000Z",
                    ],
                    "name": "tls.server.not_before",
                    "requirement_level": "recommended",
                    "root_namespace": "tls",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Distinguished name of subject of the x.509 certificate presented by the server.",
                    "examples": [
                        "CN=myserver, OU=Documentation Team, DC=example, DC=com",
                    ],
                    "name": "tls.server.subject",
                    "requirement_level": "recommended",
                    "root_namespace": "tls",
                    "stability": "development",
                    "type": "string",
                },
            ],
            "root_namespace": "tls",
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
            "attr.go.j2",
        ],
    },
} */
