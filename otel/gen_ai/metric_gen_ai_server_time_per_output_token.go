package gen_ai

import (
	"github.com/prometheus/client_golang/prometheus"
)

import (
	"shorez.de/promconv/otel/server"
)

// Time per output token generated after the first token for successful responses
type ServerTimePerOutputToken struct {
	*prometheus.HistogramVec
	extra ServerTimePerOutputTokenExtra
}

func NewServerTimePerOutputToken() ServerTimePerOutputToken {
	labels := []string{AttrOperationName("").Key(), AttrSystem("").Key(), AttrRequestModel("").Key(), server.AttrPort("").Key(), AttrResponseModel("").Key(), server.AttrAddress("").Key()}
	return ServerTimePerOutputToken{HistogramVec: prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "gen_ai_server_time_per_output_token",
		Help: "Time per output token generated after the first token for successful responses",
	}, labels)}
}

func (m ServerTimePerOutputToken) With(operationName AttrOperationName, system AttrSystem, extras ...interface {
	GenAiRequestModel() AttrRequestModel
	ServerPort() server.AttrPort
	GenAiResponseModel() AttrResponseModel
	ServerAddress() server.AttrAddress
}) prometheus.Observer {
	if extras == nil {
		extras = append(extras, m.extra)
	}
	extra := extras[0]

	return m.HistogramVec.WithLabelValues(operationName.Value(), system.Value(), extra.GenAiRequestModel().Value(), extra.ServerPort().Value(), extra.GenAiResponseModel().Value(), extra.ServerAddress().Value())
}

// Deprecated: Use [ServerTimePerOutputToken.With] instead
func (m ServerTimePerOutputToken) WithLabelValues(lvs ...string) prometheus.Observer {
	return m.HistogramVec.WithLabelValues(lvs...)
}

func (a ServerTimePerOutputToken) WithRequestModel(attr interface{ GenAiRequestModel() AttrRequestModel }) ServerTimePerOutputToken {
	a.extra.AttrRequestModel = attr.GenAiRequestModel()
	return a
}
func (a ServerTimePerOutputToken) WithServerPort(attr interface{ ServerPort() server.AttrPort }) ServerTimePerOutputToken {
	a.extra.AttrServerPort = attr.ServerPort()
	return a
}
func (a ServerTimePerOutputToken) WithResponseModel(attr interface{ GenAiResponseModel() AttrResponseModel }) ServerTimePerOutputToken {
	a.extra.AttrResponseModel = attr.GenAiResponseModel()
	return a
}
func (a ServerTimePerOutputToken) WithServerAddress(attr interface{ ServerAddress() server.AttrAddress }) ServerTimePerOutputToken {
	a.extra.AttrServerAddress = attr.ServerAddress()
	return a
}

type ServerTimePerOutputTokenExtra struct {
	// The name of the GenAI model a request is being made to
	AttrRequestModel  AttrRequestModel   `otel:"gen_ai.request.model"`  // GenAI server port
	AttrServerPort    server.AttrPort    `otel:"server.port"`           // The name of the model that generated the response
	AttrResponseModel AttrResponseModel  `otel:"gen_ai.response.model"` // GenAI server address
	AttrServerAddress server.AttrAddress `otel:"server.address"`
}

func (a ServerTimePerOutputTokenExtra) GenAiRequestModel() AttrRequestModel {
	return a.AttrRequestModel
}
func (a ServerTimePerOutputTokenExtra) ServerPort() server.AttrPort { return a.AttrServerPort }
func (a ServerTimePerOutputTokenExtra) GenAiResponseModel() AttrResponseModel {
	return a.AttrResponseModel
}
func (a ServerTimePerOutputTokenExtra) ServerAddress() server.AttrAddress { return a.AttrServerAddress }

/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "ServerTimePerOutputTokenExtra",
        "Instr": "Histogram",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "server.time_per_output_token",
        "Type": "ServerTimePerOutputToken",
        "attributes": [
            {
                "brief": "The name of the operation being performed.",
                "name": "gen_ai.operation.name",
                "note": "If one of the predefined values applies, but specific system uses a different name it's RECOMMENDED to document it in the semantic conventions for specific GenAI system and use system-specific name in the instrumentation. If a different name is not documented, instrumentation libraries SHOULD use applicable predefined value.\n",
                "requirement_level": "required",
                "stability": "development",
                "type": {
                    "members": [
                        {
                            "brief": "Chat completion operation such as [OpenAI Chat API](https://platform.openai.com/docs/api-reference/chat)",
                            "deprecated": none,
                            "id": "chat",
                            "note": none,
                            "stability": "development",
                            "value": "chat",
                        },
                        {
                            "brief": "Multimodal content generation operation such as [Gemini Generate Content](https://ai.google.dev/api/generate-content)",
                            "deprecated": none,
                            "id": "generate_content",
                            "note": none,
                            "stability": "development",
                            "value": "generate_content",
                        },
                        {
                            "brief": "Text completions operation such as [OpenAI Completions API (Legacy)](https://platform.openai.com/docs/api-reference/completions)",
                            "deprecated": none,
                            "id": "text_completion",
                            "note": none,
                            "stability": "development",
                            "value": "text_completion",
                        },
                        {
                            "brief": "Embeddings operation such as [OpenAI Create embeddings API](https://platform.openai.com/docs/api-reference/embeddings/create)",
                            "deprecated": none,
                            "id": "embeddings",
                            "note": none,
                            "stability": "development",
                            "value": "embeddings",
                        },
                        {
                            "brief": "Create GenAI agent",
                            "deprecated": none,
                            "id": "create_agent",
                            "note": none,
                            "stability": "development",
                            "value": "create_agent",
                        },
                        {
                            "brief": "Invoke GenAI agent",
                            "deprecated": none,
                            "id": "invoke_agent",
                            "note": none,
                            "stability": "development",
                            "value": "invoke_agent",
                        },
                        {
                            "brief": "Execute a tool",
                            "deprecated": none,
                            "id": "execute_tool",
                            "note": none,
                            "stability": "development",
                            "value": "execute_tool",
                        },
                    ],
                },
            },
            {
                "brief": "The Generative AI product as identified by the client or server instrumentation.",
                "examples": "openai",
                "name": "gen_ai.system",
                "note": "The `gen_ai.system` describes a family of GenAI models with specific model identified\nby `gen_ai.request.model` and `gen_ai.response.model` attributes.\n\nThe actual GenAI product may differ from the one identified by the client.\nMultiple systems, including Azure OpenAI and Gemini, are accessible by OpenAI client\nlibraries. In such cases, the `gen_ai.system` is set to `openai` based on the\ninstrumentation's best knowledge, instead of the actual system. The `server.address`\nattribute may help identify the actual system in use for `openai`.\n\nFor custom model, a custom friendly name SHOULD be used.\nIf none of these options apply, the `gen_ai.system` SHOULD be set to `_OTHER`.\n",
                "requirement_level": "required",
                "stability": "development",
                "type": {
                    "members": [
                        {
                            "brief": "OpenAI",
                            "deprecated": none,
                            "id": "openai",
                            "note": none,
                            "stability": "development",
                            "value": "openai",
                        },
                        {
                            "brief": "Any Google generative AI endpoint",
                            "deprecated": none,
                            "id": "gcp.gen_ai",
                            "note": "May be used when specific backend is unknown. May use common attributes prefixed with 'gcp.gen_ai.'.\n",
                            "stability": "development",
                            "value": "gcp.gen_ai",
                        },
                        {
                            "brief": "Vertex AI",
                            "deprecated": none,
                            "id": "gcp.vertex_ai",
                            "note": "This refers to the 'aiplatform.googleapis.com' endpoint. May use common attributes prefixed with 'gcp.gen_ai.'.\n",
                            "stability": "development",
                            "value": "gcp.vertex_ai",
                        },
                        {
                            "brief": "Gemini",
                            "deprecated": none,
                            "id": "gcp.gemini",
                            "note": "This refers to the 'generativelanguage.googleapis.com' endpoint. Also known as the AI Studio API. May use common attributes prefixed with 'gcp.gen_ai.'.\n",
                            "stability": "development",
                            "value": "gcp.gemini",
                        },
                        {
                            "brief": "Vertex AI",
                            "deprecated": "Use 'gcp.vertex_ai' instead.",
                            "id": "vertex_ai",
                            "note": none,
                            "stability": "development",
                            "value": "vertex_ai",
                        },
                        {
                            "brief": "Gemini",
                            "deprecated": "Use 'gcp.gemini' instead.",
                            "id": "gemini",
                            "note": none,
                            "stability": "development",
                            "value": "gemini",
                        },
                        {
                            "brief": "Anthropic",
                            "deprecated": none,
                            "id": "anthropic",
                            "note": none,
                            "stability": "development",
                            "value": "anthropic",
                        },
                        {
                            "brief": "Cohere",
                            "deprecated": none,
                            "id": "cohere",
                            "note": none,
                            "stability": "development",
                            "value": "cohere",
                        },
                        {
                            "brief": "Azure AI Inference",
                            "deprecated": none,
                            "id": "az.ai.inference",
                            "note": none,
                            "stability": "development",
                            "value": "az.ai.inference",
                        },
                        {
                            "brief": "Azure OpenAI",
                            "deprecated": none,
                            "id": "az.ai.openai",
                            "note": none,
                            "stability": "development",
                            "value": "az.ai.openai",
                        },
                        {
                            "brief": "IBM Watsonx AI",
                            "deprecated": none,
                            "id": "ibm.watsonx.ai",
                            "note": none,
                            "stability": "development",
                            "value": "ibm.watsonx.ai",
                        },
                        {
                            "brief": "AWS Bedrock",
                            "deprecated": none,
                            "id": "aws.bedrock",
                            "note": none,
                            "stability": "development",
                            "value": "aws.bedrock",
                        },
                        {
                            "brief": "Perplexity",
                            "deprecated": none,
                            "id": "perplexity",
                            "note": none,
                            "stability": "development",
                            "value": "perplexity",
                        },
                        {
                            "brief": "xAI",
                            "deprecated": none,
                            "id": "xai",
                            "note": none,
                            "stability": "development",
                            "value": "xai",
                        },
                        {
                            "brief": "DeepSeek",
                            "deprecated": none,
                            "id": "deepseek",
                            "note": none,
                            "stability": "development",
                            "value": "deepseek",
                        },
                        {
                            "brief": "Groq",
                            "deprecated": none,
                            "id": "groq",
                            "note": none,
                            "stability": "development",
                            "value": "groq",
                        },
                        {
                            "brief": "Mistral AI",
                            "deprecated": none,
                            "id": "mistral_ai",
                            "note": none,
                            "stability": "development",
                            "value": "mistral_ai",
                        },
                    ],
                },
            },
            {
                "brief": "The name of the GenAI model a request is being made to.",
                "examples": "gpt-4",
                "name": "gen_ai.request.model",
                "requirement_level": {
                    "conditionally_required": "If available.",
                },
                "stability": "development",
                "type": "string",
            },
            {
                "brief": "GenAI server port.",
                "examples": [
                    80,
                    8080,
                    443,
                ],
                "name": "server.port",
                "note": "When observed from the client side, and when communicating through an intermediary, `server.port` SHOULD represent the server port behind any intermediaries, for example proxies, if it's available.\n",
                "requirement_level": {
                    "conditionally_required": "If `server.address` is set.",
                },
                "stability": "stable",
                "type": "int",
            },
            {
                "brief": "The name of the model that generated the response.",
                "examples": [
                    "gpt-4-0613",
                ],
                "name": "gen_ai.response.model",
                "requirement_level": "recommended",
                "stability": "development",
                "type": "string",
            },
            {
                "brief": "GenAI server address.",
                "examples": [
                    "example.com",
                    "10.1.2.80",
                    "/tmp/my.sock",
                ],
                "name": "server.address",
                "note": "When observed from the client side, and when communicating through an intermediary, `server.address` SHOULD represent the server address behind any intermediaries, for example proxies, if it's available.\n",
                "requirement_level": "recommended",
                "stability": "stable",
                "type": "string",
            },
        ],
        "ctx": {
            "attributes": [
                {
                    "brief": "The name of the model that generated the response.",
                    "examples": [
                        "gpt-4-0613",
                    ],
                    "name": "gen_ai.response.model",
                    "requirement_level": "recommended",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The name of the operation being performed.",
                    "name": "gen_ai.operation.name",
                    "note": "If one of the predefined values applies, but specific system uses a different name it's RECOMMENDED to document it in the semantic conventions for specific GenAI system and use system-specific name in the instrumentation. If a different name is not documented, instrumentation libraries SHOULD use applicable predefined value.\n",
                    "requirement_level": "required",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "Chat completion operation such as [OpenAI Chat API](https://platform.openai.com/docs/api-reference/chat)",
                                "deprecated": none,
                                "id": "chat",
                                "note": none,
                                "stability": "development",
                                "value": "chat",
                            },
                            {
                                "brief": "Multimodal content generation operation such as [Gemini Generate Content](https://ai.google.dev/api/generate-content)",
                                "deprecated": none,
                                "id": "generate_content",
                                "note": none,
                                "stability": "development",
                                "value": "generate_content",
                            },
                            {
                                "brief": "Text completions operation such as [OpenAI Completions API (Legacy)](https://platform.openai.com/docs/api-reference/completions)",
                                "deprecated": none,
                                "id": "text_completion",
                                "note": none,
                                "stability": "development",
                                "value": "text_completion",
                            },
                            {
                                "brief": "Embeddings operation such as [OpenAI Create embeddings API](https://platform.openai.com/docs/api-reference/embeddings/create)",
                                "deprecated": none,
                                "id": "embeddings",
                                "note": none,
                                "stability": "development",
                                "value": "embeddings",
                            },
                            {
                                "brief": "Create GenAI agent",
                                "deprecated": none,
                                "id": "create_agent",
                                "note": none,
                                "stability": "development",
                                "value": "create_agent",
                            },
                            {
                                "brief": "Invoke GenAI agent",
                                "deprecated": none,
                                "id": "invoke_agent",
                                "note": none,
                                "stability": "development",
                                "value": "invoke_agent",
                            },
                            {
                                "brief": "Execute a tool",
                                "deprecated": none,
                                "id": "execute_tool",
                                "note": none,
                                "stability": "development",
                                "value": "execute_tool",
                            },
                        ],
                    },
                },
                {
                    "brief": "The Generative AI product as identified by the client or server instrumentation.",
                    "examples": "openai",
                    "name": "gen_ai.system",
                    "note": "The `gen_ai.system` describes a family of GenAI models with specific model identified\nby `gen_ai.request.model` and `gen_ai.response.model` attributes.\n\nThe actual GenAI product may differ from the one identified by the client.\nMultiple systems, including Azure OpenAI and Gemini, are accessible by OpenAI client\nlibraries. In such cases, the `gen_ai.system` is set to `openai` based on the\ninstrumentation's best knowledge, instead of the actual system. The `server.address`\nattribute may help identify the actual system in use for `openai`.\n\nFor custom model, a custom friendly name SHOULD be used.\nIf none of these options apply, the `gen_ai.system` SHOULD be set to `_OTHER`.\n",
                    "requirement_level": "required",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "OpenAI",
                                "deprecated": none,
                                "id": "openai",
                                "note": none,
                                "stability": "development",
                                "value": "openai",
                            },
                            {
                                "brief": "Any Google generative AI endpoint",
                                "deprecated": none,
                                "id": "gcp.gen_ai",
                                "note": "May be used when specific backend is unknown. May use common attributes prefixed with 'gcp.gen_ai.'.\n",
                                "stability": "development",
                                "value": "gcp.gen_ai",
                            },
                            {
                                "brief": "Vertex AI",
                                "deprecated": none,
                                "id": "gcp.vertex_ai",
                                "note": "This refers to the 'aiplatform.googleapis.com' endpoint. May use common attributes prefixed with 'gcp.gen_ai.'.\n",
                                "stability": "development",
                                "value": "gcp.vertex_ai",
                            },
                            {
                                "brief": "Gemini",
                                "deprecated": none,
                                "id": "gcp.gemini",
                                "note": "This refers to the 'generativelanguage.googleapis.com' endpoint. Also known as the AI Studio API. May use common attributes prefixed with 'gcp.gen_ai.'.\n",
                                "stability": "development",
                                "value": "gcp.gemini",
                            },
                            {
                                "brief": "Vertex AI",
                                "deprecated": "Use 'gcp.vertex_ai' instead.",
                                "id": "vertex_ai",
                                "note": none,
                                "stability": "development",
                                "value": "vertex_ai",
                            },
                            {
                                "brief": "Gemini",
                                "deprecated": "Use 'gcp.gemini' instead.",
                                "id": "gemini",
                                "note": none,
                                "stability": "development",
                                "value": "gemini",
                            },
                            {
                                "brief": "Anthropic",
                                "deprecated": none,
                                "id": "anthropic",
                                "note": none,
                                "stability": "development",
                                "value": "anthropic",
                            },
                            {
                                "brief": "Cohere",
                                "deprecated": none,
                                "id": "cohere",
                                "note": none,
                                "stability": "development",
                                "value": "cohere",
                            },
                            {
                                "brief": "Azure AI Inference",
                                "deprecated": none,
                                "id": "az.ai.inference",
                                "note": none,
                                "stability": "development",
                                "value": "az.ai.inference",
                            },
                            {
                                "brief": "Azure OpenAI",
                                "deprecated": none,
                                "id": "az.ai.openai",
                                "note": none,
                                "stability": "development",
                                "value": "az.ai.openai",
                            },
                            {
                                "brief": "IBM Watsonx AI",
                                "deprecated": none,
                                "id": "ibm.watsonx.ai",
                                "note": none,
                                "stability": "development",
                                "value": "ibm.watsonx.ai",
                            },
                            {
                                "brief": "AWS Bedrock",
                                "deprecated": none,
                                "id": "aws.bedrock",
                                "note": none,
                                "stability": "development",
                                "value": "aws.bedrock",
                            },
                            {
                                "brief": "Perplexity",
                                "deprecated": none,
                                "id": "perplexity",
                                "note": none,
                                "stability": "development",
                                "value": "perplexity",
                            },
                            {
                                "brief": "xAI",
                                "deprecated": none,
                                "id": "xai",
                                "note": none,
                                "stability": "development",
                                "value": "xai",
                            },
                            {
                                "brief": "DeepSeek",
                                "deprecated": none,
                                "id": "deepseek",
                                "note": none,
                                "stability": "development",
                                "value": "deepseek",
                            },
                            {
                                "brief": "Groq",
                                "deprecated": none,
                                "id": "groq",
                                "note": none,
                                "stability": "development",
                                "value": "groq",
                            },
                            {
                                "brief": "Mistral AI",
                                "deprecated": none,
                                "id": "mistral_ai",
                                "note": none,
                                "stability": "development",
                                "value": "mistral_ai",
                            },
                        ],
                    },
                },
                {
                    "brief": "The name of the GenAI model a request is being made to.",
                    "examples": "gpt-4",
                    "name": "gen_ai.request.model",
                    "requirement_level": {
                        "conditionally_required": "If available.",
                    },
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "GenAI server address.",
                    "examples": [
                        "example.com",
                        "10.1.2.80",
                        "/tmp/my.sock",
                    ],
                    "name": "server.address",
                    "note": "When observed from the client side, and when communicating through an intermediary, `server.address` SHOULD represent the server address behind any intermediaries, for example proxies, if it's available.\n",
                    "requirement_level": "recommended",
                    "stability": "stable",
                    "type": "string",
                },
                {
                    "brief": "GenAI server port.",
                    "examples": [
                        80,
                        8080,
                        443,
                    ],
                    "name": "server.port",
                    "note": "When observed from the client side, and when communicating through an intermediary, `server.port` SHOULD represent the server port behind any intermediaries, for example proxies, if it's available.\n",
                    "requirement_level": {
                        "conditionally_required": "If `server.address` is set.",
                    },
                    "stability": "stable",
                    "type": "int",
                },
            ],
            "brief": "Time per output token generated after the first token for successful responses",
            "events": [],
            "id": "metric.gen_ai.server.time_per_output_token",
            "instrument": "histogram",
            "lineage": {
                "attributes": {
                    "gen_ai.operation.name": {
                        "inherited_fields": [
                            "brief",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.gen_ai",
                    },
                    "gen_ai.request.model": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.gen_ai",
                    },
                    "gen_ai.response.model": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.gen_ai",
                    },
                    "gen_ai.system": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.gen_ai",
                    },
                    "server.address": {
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
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/gen-ai/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "gen_ai.server.time_per_output_token",
            "name": none,
            "root_namespace": "gen_ai",
            "span_kind": none,
            "stability": "development",
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
