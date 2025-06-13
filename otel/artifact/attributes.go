package artifact

// The provenance filename of the built attestation which directly relates to the build artifact filename. This filename SHOULD accompany the artifact at publish time. See the [SLSA Relationship] specification for more information
//
// [SLSA Relationship]: https://slsa.dev/spec/v1.0/distributing-provenance#relationship-between-artifacts-and-attestations
type AttrAttestationFilename string // artifact.attestation.filename

func (AttrAttestationFilename) Development()    {}
func (AttrAttestationFilename) Recommended()    {}
func (AttrAttestationFilename) Key() string     { return "artifact_attestation_filename" }
func (a AttrAttestationFilename) Value() string { return string(a) }

// The full [hash value (see glossary)], of the built attestation. Some envelopes in the [software attestation space] also refer to this as the **digest**
//
// [hash value (see glossary)]: https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.186-5.pdf
// [software attestation space]: https://github.com/in-toto/attestation/tree/main/spec
type AttrAttestationHash string // artifact.attestation.hash

func (AttrAttestationHash) Development()    {}
func (AttrAttestationHash) Recommended()    {}
func (AttrAttestationHash) Key() string     { return "artifact_attestation_hash" }
func (a AttrAttestationHash) Value() string { return string(a) }

// The id of the build [software attestation]
//
// [software attestation]: https://slsa.dev/attestation-model
type AttrAttestationId string // artifact.attestation.id

func (AttrAttestationId) Development()    {}
func (AttrAttestationId) Recommended()    {}
func (AttrAttestationId) Key() string     { return "artifact_attestation_id" }
func (a AttrAttestationId) Value() string { return string(a) }

// The human readable file name of the artifact, typically generated during build and release processes. Often includes the package name and version in the file name.
//
// This file name can also act as the [Package Name]
// in cases where the package ecosystem maps accordingly.
// Additionally, the artifact [can be published]
// for others, but that is not a guarantee
//
// [Package Name]: https://slsa.dev/spec/v1.0/terminology#package-model
// [can be published]: https://slsa.dev/spec/v1.0/terminology#software-supply-chain
type AttrFilename string // artifact.filename

func (AttrFilename) Development()    {}
func (AttrFilename) Recommended()    {}
func (AttrFilename) Key() string     { return "artifact_filename" }
func (a AttrFilename) Value() string { return string(a) }

// The full [hash value (see glossary)], often found in checksum.txt on a release of the artifact and used to verify package integrity.
//
// The specific algorithm used to create the cryptographic hash value is
// not defined. In situations where an artifact has multiple
// cryptographic hashes, it is up to the implementer to choose which
// hash value to set here; this should be the most secure hash algorithm
// that is suitable for the situation and consistent with the
// corresponding attestation. The implementer can then provide the other
// hash values through an additional set of attribute extensions as they
// deem necessary
//
// [hash value (see glossary)]: https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.186-5.pdf
type AttrHash string // artifact.hash

func (AttrHash) Development()    {}
func (AttrHash) Recommended()    {}
func (AttrHash) Key() string     { return "artifact_hash" }
func (a AttrHash) Value() string { return string(a) }

// The [Package URL] of the [package artifact] provides a standard way to identify and locate the packaged artifact
//
// [Package URL]: https://github.com/package-url/purl-spec
// [package artifact]: https://slsa.dev/spec/v1.0/terminology#package-model
type AttrPurl string // artifact.purl

func (AttrPurl) Development()    {}
func (AttrPurl) Recommended()    {}
func (AttrPurl) Key() string     { return "artifact_purl" }
func (a AttrPurl) Value() string { return string(a) }

// The version of the artifact
type AttrVersion string // artifact.version

func (AttrVersion) Development()    {}
func (AttrVersion) Recommended()    {}
func (AttrVersion) Key() string     { return "artifact_version" }
func (a AttrVersion) Value() string { return string(a) }

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "The provenance filename of the built attestation which directly relates to the build artifact filename. This filename SHOULD accompany the artifact at publish time. See the [SLSA Relationship](https://slsa.dev/spec/v1.0/distributing-provenance#relationship-between-artifacts-and-attestations) specification for more information.\n",
                    "examples": [
                        "golang-binary-amd64-v0.1.0.attestation",
                        "docker-image-amd64-v0.1.0.intoto.json1",
                        "release-1.tar.gz.attestation",
                        "file-name-package.tar.gz.intoto.json1",
                    ],
                    "name": "artifact.attestation.filename",
                    "requirement_level": "recommended",
                    "root_namespace": "artifact",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The full [hash value (see glossary)](https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.186-5.pdf), of the built attestation. Some envelopes in the [software attestation space](https://github.com/in-toto/attestation/tree/main/spec) also refer to this as the **digest**.\n",
                    "examples": [
                        "1b31dfcd5b7f9267bf2ff47651df1cfb9147b9e4df1f335accf65b4cda498408",
                    ],
                    "name": "artifact.attestation.hash",
                    "requirement_level": "recommended",
                    "root_namespace": "artifact",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The id of the build [software attestation](https://slsa.dev/attestation-model).\n",
                    "examples": [
                        "123",
                    ],
                    "name": "artifact.attestation.id",
                    "requirement_level": "recommended",
                    "root_namespace": "artifact",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The human readable file name of the artifact, typically generated during build and release processes. Often includes the package name and version in the file name.\n",
                    "examples": [
                        "golang-binary-amd64-v0.1.0",
                        "docker-image-amd64-v0.1.0",
                        "release-1.tar.gz",
                        "file-name-package.tar.gz",
                    ],
                    "name": "artifact.filename",
                    "note": "This file name can also act as the [Package Name](https://slsa.dev/spec/v1.0/terminology#package-model)\nin cases where the package ecosystem maps accordingly.\nAdditionally, the artifact [can be published](https://slsa.dev/spec/v1.0/terminology#software-supply-chain)\nfor others, but that is not a guarantee.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "artifact",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The full [hash value (see glossary)](https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.186-5.pdf), often found in checksum.txt on a release of the artifact and used to verify package integrity.\n",
                    "examples": [
                        "9ff4c52759e2c4ac70b7d517bc7fcdc1cda631ca0045271ddd1b192544f8a3e9",
                    ],
                    "name": "artifact.hash",
                    "note": "The specific algorithm used to create the cryptographic hash value is\nnot defined. In situations where an artifact has multiple\ncryptographic hashes, it is up to the implementer to choose which\nhash value to set here; this should be the most secure hash algorithm\nthat is suitable for the situation and consistent with the\ncorresponding attestation. The implementer can then provide the other\nhash values through an additional set of attribute extensions as they\ndeem necessary.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "artifact",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The [Package URL](https://github.com/package-url/purl-spec) of the [package artifact](https://slsa.dev/spec/v1.0/terminology#package-model) provides a standard way to identify and locate the packaged artifact.\n",
                    "examples": [
                        "pkg:github/package-url/purl-spec@1209109710924",
                        "pkg:npm/foo@12.12.3",
                    ],
                    "name": "artifact.purl",
                    "requirement_level": "recommended",
                    "root_namespace": "artifact",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The version of the artifact.\n",
                    "examples": [
                        "v0.1.0",
                        "1.2.1",
                        "122691-build",
                    ],
                    "name": "artifact.version",
                    "requirement_level": "recommended",
                    "root_namespace": "artifact",
                    "stability": "development",
                    "type": "string",
                },
            ],
            "root_namespace": "artifact",
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
