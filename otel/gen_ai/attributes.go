package gen_ai

// Free-form description of the GenAI agent provided by the application
type AttrAgentDescription string // gen_ai.agent.description

func (AttrAgentDescription) Development() {}
func (AttrAgentDescription) Recommended() {}

// The unique identifier of the GenAI agent
type AttrAgentId string // gen_ai.agent.id

func (AttrAgentId) Development() {}
func (AttrAgentId) Recommended() {}

// Human-readable name of the GenAI agent provided by the application
type AttrAgentName string // gen_ai.agent.name

func (AttrAgentName) Development() {}
func (AttrAgentName) Recommended() {}

// Deprecated, use Event API to report completions contents
type AttrCompletion string // gen_ai.completion

func (AttrCompletion) Development() {}
func (AttrCompletion) Recommended() {}

// The unique identifier for a conversation (session, thread), used to store and correlate messages within this conversation
type AttrConversationId string // gen_ai.conversation.id

func (AttrConversationId) Development() {}
func (AttrConversationId) Recommended() {}

// The data source identifier.
// Data sources are used by AI agents and RAG applications to store grounding data. A data source may be an external database, object store, document collection, website, or any other storage system used by the GenAI agent or application. The `gen_ai.data_source.id` SHOULD match the identifier used by the GenAI system rather than a name specific to the external storage, such as a database or object store. Semantic conventions referencing `gen_ai.data_source.id` MAY also leverage additional attributes, such as `db.*`, to further identify and describe the data source
type AttrDataSourceId string // gen_ai.data_source.id

func (AttrDataSourceId) Development() {}
func (AttrDataSourceId) Recommended() {}

// Deprecated, use `gen_ai.output.type`
type AttrOpenaiRequestResponseFormat string // gen_ai.openai.request.response_format

func (AttrOpenaiRequestResponseFormat) Development() {}
func (AttrOpenaiRequestResponseFormat) Recommended() {}

// Deprecated, use `gen_ai.request.seed`
type AttrOpenaiRequestSeed string // gen_ai.openai.request.seed

func (AttrOpenaiRequestSeed) Development() {}
func (AttrOpenaiRequestSeed) Recommended() {}

// The service tier requested. May be a specific tier, default, or auto
type AttrOpenaiRequestServiceTier string // gen_ai.openai.request.service_tier

func (AttrOpenaiRequestServiceTier) Development() {}
func (AttrOpenaiRequestServiceTier) Recommended() {}

// The service tier used for the response
type AttrOpenaiResponseServiceTier string // gen_ai.openai.response.service_tier

func (AttrOpenaiResponseServiceTier) Development() {}
func (AttrOpenaiResponseServiceTier) Recommended() {}

// A fingerprint to track any eventual change in the Generative AI environment
type AttrOpenaiResponseSystemFingerprint string // gen_ai.openai.response.system_fingerprint

func (AttrOpenaiResponseSystemFingerprint) Development() {}
func (AttrOpenaiResponseSystemFingerprint) Recommended() {}

// The name of the operation being performed.
// If one of the predefined values applies, but specific system uses a different name it's RECOMMENDED to document it in the semantic conventions for specific GenAI system and use system-specific name in the instrumentation. If a different name is not documented, instrumentation libraries SHOULD use applicable predefined value
type AttrOperationName string // gen_ai.operation.name

func (AttrOperationName) Development() {}
func (AttrOperationName) Recommended() {}

// Represents the content type requested by the client.
// This attribute SHOULD be used when the client requests output of a specific type. The model may return zero or more outputs of this type.
// This attribute specifies the output modality and not the actual output format. For example, if an image is requested, the actual output could be a URL pointing to an image file.
// Additional output format details may be recorded in the future in the `gen_ai.output.{type}.*` attributes
type AttrOutputType string // gen_ai.output.type

func (AttrOutputType) Development() {}
func (AttrOutputType) Recommended() {}

// Deprecated, use Event API to report prompt contents
type AttrPrompt string // gen_ai.prompt

func (AttrPrompt) Development() {}
func (AttrPrompt) Recommended() {}

// The target number of candidate completions to return
type AttrRequestChoiceCount string // gen_ai.request.choice.count

func (AttrRequestChoiceCount) Development() {}
func (AttrRequestChoiceCount) Recommended() {}

// The encoding formats requested in an embeddings operation, if specified.
// In some GenAI systems the encoding formats are called embedding types. Also, some GenAI systems only accept a single format per request
type AttrRequestEncodingFormats string // gen_ai.request.encoding_formats

func (AttrRequestEncodingFormats) Development() {}
func (AttrRequestEncodingFormats) Recommended() {}

// The frequency penalty setting for the GenAI request
type AttrRequestFrequencyPenalty string // gen_ai.request.frequency_penalty

func (AttrRequestFrequencyPenalty) Development() {}
func (AttrRequestFrequencyPenalty) Recommended() {}

// The maximum number of tokens the model generates for a request
type AttrRequestMaxTokens string // gen_ai.request.max_tokens

func (AttrRequestMaxTokens) Development() {}
func (AttrRequestMaxTokens) Recommended() {}

// The name of the GenAI model a request is being made to
type AttrRequestModel string // gen_ai.request.model

func (AttrRequestModel) Development() {}
func (AttrRequestModel) Recommended() {}

// The presence penalty setting for the GenAI request
type AttrRequestPresencePenalty string // gen_ai.request.presence_penalty

func (AttrRequestPresencePenalty) Development() {}
func (AttrRequestPresencePenalty) Recommended() {}

// Requests with same seed value more likely to return same result
type AttrRequestSeed string // gen_ai.request.seed

func (AttrRequestSeed) Development() {}
func (AttrRequestSeed) Recommended() {}

// List of sequences that the model will use to stop generating further tokens
type AttrRequestStopSequences string // gen_ai.request.stop_sequences

func (AttrRequestStopSequences) Development() {}
func (AttrRequestStopSequences) Recommended() {}

// The temperature setting for the GenAI request
type AttrRequestTemperature string // gen_ai.request.temperature

func (AttrRequestTemperature) Development() {}
func (AttrRequestTemperature) Recommended() {}

// The top_k sampling setting for the GenAI request
type AttrRequestTopK string // gen_ai.request.top_k

func (AttrRequestTopK) Development() {}
func (AttrRequestTopK) Recommended() {}

// The top_p sampling setting for the GenAI request
type AttrRequestTopP string // gen_ai.request.top_p

func (AttrRequestTopP) Development() {}
func (AttrRequestTopP) Recommended() {}

// Array of reasons the model stopped generating tokens, corresponding to each generation received
type AttrResponseFinishReasons string // gen_ai.response.finish_reasons

func (AttrResponseFinishReasons) Development() {}
func (AttrResponseFinishReasons) Recommended() {}

// The unique identifier for the completion
type AttrResponseId string // gen_ai.response.id

func (AttrResponseId) Development() {}
func (AttrResponseId) Recommended() {}

// The name of the model that generated the response
type AttrResponseModel string // gen_ai.response.model

func (AttrResponseModel) Development() {}
func (AttrResponseModel) Recommended() {}

// The Generative AI product as identified by the client or server instrumentation.
// The `gen_ai.system` describes a family of GenAI models with specific model identified
// by `gen_ai.request.model` and `gen_ai.response.model` attributes.
//
// The actual GenAI product may differ from the one identified by the client.
// Multiple systems, including Azure OpenAI and Gemini, are accessible by OpenAI client
// libraries. In such cases, the `gen_ai.system` is set to `openai` based on the
// instrumentation's best knowledge, instead of the actual system. The `server.address`
// attribute may help identify the actual system in use for `openai`.
//
// For custom model, a custom friendly name SHOULD be used.
// If none of these options apply, the `gen_ai.system` SHOULD be set to `_OTHER`
type AttrSystem string // gen_ai.system

func (AttrSystem) Development() {}
func (AttrSystem) Recommended() {}

// The type of token being counted
type AttrTokenType string // gen_ai.token.type

func (AttrTokenType) Development() {}
func (AttrTokenType) Recommended() {}

// The tool call identifier
type AttrToolCallId string // gen_ai.tool.call.id

func (AttrToolCallId) Development() {}
func (AttrToolCallId) Recommended() {}

// The tool description
type AttrToolDescription string // gen_ai.tool.description

func (AttrToolDescription) Development() {}
func (AttrToolDescription) Recommended() {}

// Name of the tool utilized by the agent
type AttrToolName string // gen_ai.tool.name

func (AttrToolName) Development() {}
func (AttrToolName) Recommended() {}

// Type of the tool utilized by the agent
// Extension: A tool executed on the agent-side to directly call external APIs, bridging the gap between the agent and real-world systems.
// Agent-side operations involve actions that are performed by the agent on the server or within the agent's controlled environment.
// Function: A tool executed on the client-side, where the agent generates parameters for a predefined function, and the client executes the logic.
// Client-side operations are actions taken on the user's end or within the client application.
// Datastore: A tool used by the agent to access and query structured or unstructured external data for retrieval-augmented tasks or knowledge updates
type AttrToolType string // gen_ai.tool.type

func (AttrToolType) Development() {}
func (AttrToolType) Recommended() {}

// Deprecated, use `gen_ai.usage.output_tokens` instead
type AttrUsageCompletionTokens string // gen_ai.usage.completion_tokens

func (AttrUsageCompletionTokens) Development() {}
func (AttrUsageCompletionTokens) Recommended() {}

// The number of tokens used in the GenAI input (prompt)
type AttrUsageInputTokens string // gen_ai.usage.input_tokens

func (AttrUsageInputTokens) Development() {}
func (AttrUsageInputTokens) Recommended() {}

// The number of tokens used in the GenAI response (completion)
type AttrUsageOutputTokens string // gen_ai.usage.output_tokens

func (AttrUsageOutputTokens) Development() {}
func (AttrUsageOutputTokens) Recommended() {}

// Deprecated, use `gen_ai.usage.input_tokens` instead
type AttrUsagePromptTokens string // gen_ai.usage.prompt_tokens

func (AttrUsagePromptTokens) Development() {}
func (AttrUsagePromptTokens) Recommended() {}

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "Free-form description of the GenAI agent provided by the application.",
                    "examples": [
                        "Helps with math problems",
                        "Generates fiction stories",
                    ],
                    "name": "gen_ai.agent.description",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The unique identifier of the GenAI agent.",
                    "examples": [
                        "asst_5j66UpCpwteGg4YSxUnt7lPY",
                    ],
                    "name": "gen_ai.agent.id",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Human-readable name of the GenAI agent provided by the application.",
                    "examples": [
                        "Math Tutor",
                        "Fiction Writer",
                    ],
                    "name": "gen_ai.agent.name",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use Event API to report completions contents.",
                    "deprecated": {
                        "note": "Removed, no replacement at this time.",
                        "reason": "obsoleted",
                    },
                    "examples": [
                        "[{'role': 'assistant', 'content': 'The capital of France is Paris.'}]",
                    ],
                    "name": "gen_ai.completion",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The unique identifier for a conversation (session, thread), used to store and correlate messages within this conversation.",
                    "examples": [
                        "conv_5j66UpCpwteGg4YSxUnt7lPY",
                    ],
                    "name": "gen_ai.conversation.id",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The data source identifier.",
                    "examples": [
                        "H7STPQYOND",
                    ],
                    "name": "gen_ai.data_source.id",
                    "note": "Data sources are used by AI agents and RAG applications to store grounding data. A data source may be an external database, object store, document collection, website, or any other storage system used by the GenAI agent or application. The `gen_ai.data_source.id` SHOULD match the identifier used by the GenAI system rather than a name specific to the external storage, such as a database or object store. Semantic conventions referencing `gen_ai.data_source.id` MAY also leverage additional attributes, such as `db.*`, to further identify and describe the data source.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `gen_ai.output.type`.\n",
                    "deprecated": {
                        "note": "Replaced by `gen_ai.output.type`.",
                        "reason": "renamed",
                        "renamed_to": "gen_ai.output.type",
                    },
                    "name": "gen_ai.openai.request.response_format",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "Text response format",
                                "deprecated": none,
                                "id": "text",
                                "note": none,
                                "stability": "development",
                                "value": "text",
                            },
                            {
                                "brief": "JSON object response format",
                                "deprecated": none,
                                "id": "json_object",
                                "note": none,
                                "stability": "development",
                                "value": "json_object",
                            },
                            {
                                "brief": "JSON schema response format",
                                "deprecated": none,
                                "id": "json_schema",
                                "note": none,
                                "stability": "development",
                                "value": "json_schema",
                            },
                        ],
                    },
                },
                {
                    "brief": "Deprecated, use `gen_ai.request.seed`.",
                    "deprecated": {
                        "note": "Replaced by `gen_ai.request.seed`.",
                        "reason": "renamed",
                        "renamed_to": "gen_ai.request.seed",
                    },
                    "examples": [
                        100,
                    ],
                    "name": "gen_ai.openai.request.seed",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "The service tier requested. May be a specific tier, default, or auto.",
                    "examples": [
                        "auto",
                        "default",
                    ],
                    "name": "gen_ai.openai.request.service_tier",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "The system will utilize scale tier credits until they are exhausted.",
                                "deprecated": none,
                                "id": "auto",
                                "note": none,
                                "stability": "development",
                                "value": "auto",
                            },
                            {
                                "brief": "The system will utilize the default scale tier.",
                                "deprecated": none,
                                "id": "default",
                                "note": none,
                                "stability": "development",
                                "value": "default",
                            },
                        ],
                    },
                },
                {
                    "brief": "The service tier used for the response.",
                    "examples": [
                        "scale",
                        "default",
                    ],
                    "name": "gen_ai.openai.response.service_tier",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "A fingerprint to track any eventual change in the Generative AI environment.",
                    "examples": [
                        "fp_44709d6fcb",
                    ],
                    "name": "gen_ai.openai.response.system_fingerprint",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The name of the operation being performed.",
                    "name": "gen_ai.operation.name",
                    "note": "If one of the predefined values applies, but specific system uses a different name it's RECOMMENDED to document it in the semantic conventions for specific GenAI system and use system-specific name in the instrumentation. If a different name is not documented, instrumentation libraries SHOULD use applicable predefined value.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
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
                    "brief": "Represents the content type requested by the client.",
                    "name": "gen_ai.output.type",
                    "note": "This attribute SHOULD be used when the client requests output of a specific type. The model may return zero or more outputs of this type.\nThis attribute specifies the output modality and not the actual output format. For example, if an image is requested, the actual output could be a URL pointing to an image file.\nAdditional output format details may be recorded in the future in the `gen_ai.output.{type}.*` attributes.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "Plain text",
                                "deprecated": none,
                                "id": "text",
                                "note": none,
                                "stability": "development",
                                "value": "text",
                            },
                            {
                                "brief": "JSON object with known or unknown schema",
                                "deprecated": none,
                                "id": "json",
                                "note": none,
                                "stability": "development",
                                "value": "json",
                            },
                            {
                                "brief": "Image",
                                "deprecated": none,
                                "id": "image",
                                "note": none,
                                "stability": "development",
                                "value": "image",
                            },
                            {
                                "brief": "Speech",
                                "deprecated": none,
                                "id": "speech",
                                "note": none,
                                "stability": "development",
                                "value": "speech",
                            },
                        ],
                    },
                },
                {
                    "brief": "Deprecated, use Event API to report prompt contents.",
                    "deprecated": {
                        "note": "Removed, no replacement at this time.",
                        "reason": "obsoleted",
                    },
                    "examples": [
                        "[{'role': 'user', 'content': 'What is the capital of France?'}]",
                    ],
                    "name": "gen_ai.prompt",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The target number of candidate completions to return.",
                    "examples": [
                        3,
                    ],
                    "name": "gen_ai.request.choice.count",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "The encoding formats requested in an embeddings operation, if specified.",
                    "examples": [
                        [
                            "base64",
                        ],
                        [
                            "float",
                            "binary",
                        ],
                    ],
                    "name": "gen_ai.request.encoding_formats",
                    "note": "In some GenAI systems the encoding formats are called embedding types. Also, some GenAI systems only accept a single format per request.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": "string[]",
                },
                {
                    "brief": "The frequency penalty setting for the GenAI request.",
                    "examples": [
                        0.1,
                    ],
                    "name": "gen_ai.request.frequency_penalty",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": "double",
                },
                {
                    "brief": "The maximum number of tokens the model generates for a request.",
                    "examples": [
                        100,
                    ],
                    "name": "gen_ai.request.max_tokens",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "The name of the GenAI model a request is being made to.",
                    "examples": "gpt-4",
                    "name": "gen_ai.request.model",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The presence penalty setting for the GenAI request.",
                    "examples": [
                        0.1,
                    ],
                    "name": "gen_ai.request.presence_penalty",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": "double",
                },
                {
                    "brief": "Requests with same seed value more likely to return same result.",
                    "examples": [
                        100,
                    ],
                    "name": "gen_ai.request.seed",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "List of sequences that the model will use to stop generating further tokens.",
                    "examples": [
                        [
                            "forest",
                            "lived",
                        ],
                    ],
                    "name": "gen_ai.request.stop_sequences",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": "string[]",
                },
                {
                    "brief": "The temperature setting for the GenAI request.",
                    "examples": [
                        0.0,
                    ],
                    "name": "gen_ai.request.temperature",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": "double",
                },
                {
                    "brief": "The top_k sampling setting for the GenAI request.",
                    "examples": [
                        1.0,
                    ],
                    "name": "gen_ai.request.top_k",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": "double",
                },
                {
                    "brief": "The top_p sampling setting for the GenAI request.",
                    "examples": [
                        1.0,
                    ],
                    "name": "gen_ai.request.top_p",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": "double",
                },
                {
                    "brief": "Array of reasons the model stopped generating tokens, corresponding to each generation received.",
                    "examples": [
                        [
                            "stop",
                        ],
                        [
                            "stop",
                            "length",
                        ],
                    ],
                    "name": "gen_ai.response.finish_reasons",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": "string[]",
                },
                {
                    "brief": "The unique identifier for the completion.",
                    "examples": [
                        "chatcmpl-123",
                    ],
                    "name": "gen_ai.response.id",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The name of the model that generated the response.",
                    "examples": [
                        "gpt-4-0613",
                    ],
                    "name": "gen_ai.response.model",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The Generative AI product as identified by the client or server instrumentation.",
                    "examples": "openai",
                    "name": "gen_ai.system",
                    "note": "The `gen_ai.system` describes a family of GenAI models with specific model identified\nby `gen_ai.request.model` and `gen_ai.response.model` attributes.\n\nThe actual GenAI product may differ from the one identified by the client.\nMultiple systems, including Azure OpenAI and Gemini, are accessible by OpenAI client\nlibraries. In such cases, the `gen_ai.system` is set to `openai` based on the\ninstrumentation's best knowledge, instead of the actual system. The `server.address`\nattribute may help identify the actual system in use for `openai`.\n\nFor custom model, a custom friendly name SHOULD be used.\nIf none of these options apply, the `gen_ai.system` SHOULD be set to `_OTHER`.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
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
                    "brief": "The type of token being counted.",
                    "examples": [
                        "input",
                        "output",
                    ],
                    "name": "gen_ai.token.type",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": {
                        "allow_custom_values": none,
                        "members": [
                            {
                                "brief": "Input tokens (prompt, input, etc.)",
                                "deprecated": none,
                                "id": "input",
                                "note": none,
                                "stability": "development",
                                "value": "input",
                            },
                            {
                                "brief": "Output tokens (completion, response, etc.)",
                                "deprecated": "Replaced by `output`.",
                                "id": "completion",
                                "note": none,
                                "stability": "development",
                                "value": "output",
                            },
                            {
                                "brief": "Output tokens (completion, response, etc.)",
                                "deprecated": none,
                                "id": "output",
                                "note": none,
                                "stability": "development",
                                "value": "output",
                            },
                        ],
                    },
                },
                {
                    "brief": "The tool call identifier.",
                    "examples": [
                        "call_mszuSIzqtI65i1wAUOE8w5H4",
                    ],
                    "name": "gen_ai.tool.call.id",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The tool description.",
                    "examples": [
                        "Multiply two numbers",
                    ],
                    "name": "gen_ai.tool.description",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Name of the tool utilized by the agent.",
                    "examples": [
                        "Flights",
                    ],
                    "name": "gen_ai.tool.name",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Type of the tool utilized by the agent",
                    "examples": [
                        "function",
                        "extension",
                        "datastore",
                    ],
                    "name": "gen_ai.tool.type",
                    "note": "Extension: A tool executed on the agent-side to directly call external APIs, bridging the gap between the agent and real-world systems.\n  Agent-side operations involve actions that are performed by the agent on the server or within the agent's controlled environment.\nFunction: A tool executed on the client-side, where the agent generates parameters for a predefined function, and the client executes the logic.\n  Client-side operations are actions taken on the user's end or within the client application.\nDatastore: A tool used by the agent to access and query structured or unstructured external data for retrieval-augmented tasks or knowledge updates.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `gen_ai.usage.output_tokens` instead.",
                    "deprecated": {
                        "note": "Replaced by `gen_ai.usage.output_tokens`.",
                        "reason": "renamed",
                        "renamed_to": "gen_ai.usage.output_tokens",
                    },
                    "examples": [
                        42,
                    ],
                    "name": "gen_ai.usage.completion_tokens",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "The number of tokens used in the GenAI input (prompt).",
                    "examples": [
                        100,
                    ],
                    "name": "gen_ai.usage.input_tokens",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "The number of tokens used in the GenAI response (completion).",
                    "examples": [
                        180,
                    ],
                    "name": "gen_ai.usage.output_tokens",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "Deprecated, use `gen_ai.usage.input_tokens` instead.",
                    "deprecated": {
                        "note": "Replaced by `gen_ai.usage.input_tokens`.",
                        "reason": "renamed",
                        "renamed_to": "gen_ai.usage.input_tokens",
                    },
                    "examples": [
                        42,
                    ],
                    "name": "gen_ai.usage.prompt_tokens",
                    "requirement_level": "recommended",
                    "root_namespace": "gen_ai",
                    "stability": "development",
                    "type": "int",
                },
            ],
            "root_namespace": "gen_ai",
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
