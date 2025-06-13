package container

// The command used to run the container (i.e. the command name).
//
// If using embedded credentials or sensitive data, it is recommended to remove them to prevent potential leakage
type AttrCommand string // container.command

func (AttrCommand) Development()    {}
func (AttrCommand) Recommended()    {}
func (AttrCommand) Key() string     { return "container_command" }
func (a AttrCommand) Value() string { return string(a) }

// All the command arguments (including the command/executable itself) run by the container
type AttrCommandArgs string // container.command_args

func (AttrCommandArgs) Development()    {}
func (AttrCommandArgs) Recommended()    {}
func (AttrCommandArgs) Key() string     { return "container_command_args" }
func (a AttrCommandArgs) Value() string { return string(a) }

// The full command run by the container as a single string representing the full command
type AttrCommandLine string // container.command_line

func (AttrCommandLine) Development()    {}
func (AttrCommandLine) Recommended()    {}
func (AttrCommandLine) Key() string     { return "container_command_line" }
func (a AttrCommandLine) Value() string { return string(a) }

// Deprecated, use `cpu.mode` instead
type AttrCpuState string // container.cpu.state

func (AttrCpuState) Development()    {}
func (AttrCpuState) Recommended()    {}
func (AttrCpuState) Key() string     { return "container_cpu_state" }
func (a AttrCpuState) Value() string { return string(a) }

const CpuStateUser AttrCpuState = "user"
const CpuStateSystem AttrCpuState = "system"
const CpuStateKernel AttrCpuState = "kernel"

// The name of the CSI ([Container Storage Interface]) plugin used by the volume.
//
// This can sometimes be referred to as a "driver" in CSI implementations. This should represent the `name` field of the GetPluginInfo RPC
//
// [Container Storage Interface]: https://github.com/container-storage-interface/spec
type AttrCsiPluginName string // container.csi.plugin.name

func (AttrCsiPluginName) Development()    {}
func (AttrCsiPluginName) Recommended()    {}
func (AttrCsiPluginName) Key() string     { return "container_csi_plugin_name" }
func (a AttrCsiPluginName) Value() string { return string(a) }

// The unique volume ID returned by the CSI ([Container Storage Interface]) plugin.
//
// This can sometimes be referred to as a "volume handle" in CSI implementations. This should represent the `Volume.volume_id` field in CSI spec
//
// [Container Storage Interface]: https://github.com/container-storage-interface/spec
type AttrCsiVolumeId string // container.csi.volume.id

func (AttrCsiVolumeId) Development()    {}
func (AttrCsiVolumeId) Recommended()    {}
func (AttrCsiVolumeId) Key() string     { return "container_csi_volume_id" }
func (a AttrCsiVolumeId) Value() string { return string(a) }

// Container ID. Usually a UUID, as for example used to [identify Docker containers]. The UUID might be abbreviated
//
// [identify Docker containers]: https://docs.docker.com/engine/containers/run/#container-identification
type AttrId string // container.id

func (AttrId) Development()    {}
func (AttrId) Recommended()    {}
func (AttrId) Key() string     { return "container_id" }
func (a AttrId) Value() string { return string(a) }

// Runtime specific image identifier. Usually a hash algorithm followed by a UUID.
//
// Docker defines a sha256 of the image id; `container.image.id` corresponds to the `Image` field from the Docker container inspect [API] endpoint.
// K8s defines a link to the container registry repository with digest `"imageID": "registry.azurecr.io /namespace/service/dockerfile@sha256:bdeabd40c3a8a492eaf9e8e44d0ebbb84bac7ee25ac0cf8a7159d25f62555625"`.
// The ID is assigned by the container runtime and can vary in different environments. Consider using `oci.manifest.digest` if it is important to identify the same image in different environments/runtimes
//
// [API]: https://docs.docker.com/engine/api/v1.43/#tag/Container/operation/ContainerInspect
type AttrImageId string // container.image.id

func (AttrImageId) Development()    {}
func (AttrImageId) Recommended()    {}
func (AttrImageId) Key() string     { return "container_image_id" }
func (a AttrImageId) Value() string { return string(a) }

// Name of the image the container was built on
type AttrImageName string // container.image.name

func (AttrImageName) Development()    {}
func (AttrImageName) Recommended()    {}
func (AttrImageName) Key() string     { return "container_image_name" }
func (a AttrImageName) Value() string { return string(a) }

// Repo digests of the container image as provided by the container runtime.
//
// [Docker] and [CRI] report those under the `RepoDigests` field
//
// [Docker]: https://docs.docker.com/engine/api/v1.43/#tag/Image/operation/ImageInspect
// [CRI]: https://github.com/kubernetes/cri-api/blob/c75ef5b473bbe2d0a4fc92f82235efd665ea8e9f/pkg/apis/runtime/v1/api.proto#L1237-L1238
type AttrImageRepoDigests string // container.image.repo_digests

func (AttrImageRepoDigests) Development()    {}
func (AttrImageRepoDigests) Recommended()    {}
func (AttrImageRepoDigests) Key() string     { return "container_image_repo_digests" }
func (a AttrImageRepoDigests) Value() string { return string(a) }

// Container image tags. An example can be found in [Docker Image Inspect]. Should be only the `<tag>` section of the full name for example from `registry.example.com/my-org/my-image:<tag>`
//
// [Docker Image Inspect]: https://docs.docker.com/engine/api/v1.43/#tag/Image/operation/ImageInspect
type AttrImageTags string // container.image.tags

func (AttrImageTags) Development()    {}
func (AttrImageTags) Recommended()    {}
func (AttrImageTags) Key() string     { return "container_image_tags" }
func (a AttrImageTags) Value() string { return string(a) }

// Container labels, `<key>` being the label name, the value being the label value.
//
// For example, a docker container label `app` with value `nginx` SHOULD be recorded as the `container.label.app` attribute with value `"nginx"`
type AttrLabel string // container.label

func (AttrLabel) Development()    {}
func (AttrLabel) Recommended()    {}
func (AttrLabel) Key() string     { return "container_label" }
func (a AttrLabel) Value() string { return string(a) }

// Deprecated, use `container.label` instead
type AttrLabels string // container.labels

func (AttrLabels) Development()    {}
func (AttrLabels) Recommended()    {}
func (AttrLabels) Key() string     { return "container_labels" }
func (a AttrLabels) Value() string { return string(a) }

// Container name used by container runtime
type AttrName string // container.name

func (AttrName) Development()    {}
func (AttrName) Recommended()    {}
func (AttrName) Key() string     { return "container_name" }
func (a AttrName) Value() string { return string(a) }

// The container runtime managing this container
type AttrRuntime string // container.runtime

func (AttrRuntime) Development()    {}
func (AttrRuntime) Recommended()    {}
func (AttrRuntime) Key() string     { return "container_runtime" }
func (a AttrRuntime) Value() string { return string(a) }

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "The command used to run the container (i.e. the command name).\n",
                    "examples": [
                        "otelcontribcol",
                    ],
                    "name": "container.command",
                    "note": "If using embedded credentials or sensitive data, it is recommended to remove them to prevent potential leakage.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "container",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "All the command arguments (including the command/executable itself) run by the container.\n",
                    "examples": [
                        [
                            "otelcontribcol",
                            "--config",
                            "config.yaml",
                        ],
                    ],
                    "name": "container.command_args",
                    "requirement_level": "recommended",
                    "root_namespace": "container",
                    "stability": "development",
                    "type": "string[]",
                },
                {
                    "brief": "The full command run by the container as a single string representing the full command.\n",
                    "examples": [
                        "otelcontribcol --config config.yaml",
                    ],
                    "name": "container.command_line",
                    "requirement_level": "recommended",
                    "root_namespace": "container",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Deprecated, use `cpu.mode` instead.",
                    "deprecated": {
                        "note": "Replaced by `cpu.mode`.",
                        "reason": "renamed",
                        "renamed_to": "cpu.mode",
                    },
                    "examples": [
                        "user",
                        "kernel",
                    ],
                    "name": "container.cpu.state",
                    "requirement_level": "recommended",
                    "root_namespace": "container",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "When tasks of the cgroup are in user mode (Linux). When all container processes are in user mode (Windows).",
                                "deprecated": none,
                                "id": "user",
                                "note": none,
                                "stability": "development",
                                "value": "user",
                            },
                            {
                                "brief": "When CPU is used by the system (host OS)",
                                "deprecated": none,
                                "id": "system",
                                "note": none,
                                "stability": "development",
                                "value": "system",
                            },
                            {
                                "brief": "When tasks of the cgroup are in kernel mode (Linux). When all container processes are in kernel mode (Windows).",
                                "deprecated": none,
                                "id": "kernel",
                                "note": none,
                                "stability": "development",
                                "value": "kernel",
                            },
                        ],
                    },
                },
                {
                    "brief": "The name of the CSI ([Container Storage Interface](https://github.com/container-storage-interface/spec)) plugin used by the volume.\n",
                    "examples": [
                        "pd.csi.storage.gke.io",
                    ],
                    "name": "container.csi.plugin.name",
                    "note": "This can sometimes be referred to as a \"driver\" in CSI implementations. This should represent the `name` field of the GetPluginInfo RPC.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "container",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The unique volume ID returned by the CSI ([Container Storage Interface](https://github.com/container-storage-interface/spec)) plugin.\n",
                    "examples": [
                        "projects/my-gcp-project/zones/my-gcp-zone/disks/my-gcp-disk",
                    ],
                    "name": "container.csi.volume.id",
                    "note": "This can sometimes be referred to as a \"volume handle\" in CSI implementations. This should represent the `Volume.volume_id` field in CSI spec.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "container",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Container ID. Usually a UUID, as for example used to [identify Docker containers](https://docs.docker.com/engine/containers/run/#container-identification). The UUID might be abbreviated.\n",
                    "examples": [
                        "a3bf90e006b2",
                    ],
                    "name": "container.id",
                    "requirement_level": "recommended",
                    "root_namespace": "container",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Runtime specific image identifier. Usually a hash algorithm followed by a UUID.\n",
                    "examples": [
                        "sha256:19c92d0a00d1b66d897bceaa7319bee0dd38a10a851c60bcec9474aa3f01e50f",
                    ],
                    "name": "container.image.id",
                    "note": "Docker defines a sha256 of the image id; `container.image.id` corresponds to the `Image` field from the Docker container inspect [API](https://docs.docker.com/engine/api/v1.43/#tag/Container/operation/ContainerInspect) endpoint.\nK8s defines a link to the container registry repository with digest `\"imageID\": \"registry.azurecr.io /namespace/service/dockerfile@sha256:bdeabd40c3a8a492eaf9e8e44d0ebbb84bac7ee25ac0cf8a7159d25f62555625\"`.\nThe ID is assigned by the container runtime and can vary in different environments. Consider using `oci.manifest.digest` if it is important to identify the same image in different environments/runtimes.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "container",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Name of the image the container was built on.\n",
                    "examples": [
                        "gcr.io/opentelemetry/operator",
                    ],
                    "name": "container.image.name",
                    "requirement_level": "recommended",
                    "root_namespace": "container",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Repo digests of the container image as provided by the container runtime.\n",
                    "examples": [
                        [
                            "example@sha256:afcc7f1ac1b49db317a7196c902e61c6c3c4607d63599ee1a82d702d249a0ccb",
                            "internal.registry.example.com:5000/example@sha256:b69959407d21e8a062e0416bf13405bb2b71ed7a84dde4158ebafacfa06f5578",
                        ],
                    ],
                    "name": "container.image.repo_digests",
                    "note": "[Docker](https://docs.docker.com/engine/api/v1.43/#tag/Image/operation/ImageInspect) and [CRI](https://github.com/kubernetes/cri-api/blob/c75ef5b473bbe2d0a4fc92f82235efd665ea8e9f/pkg/apis/runtime/v1/api.proto#L1237-L1238) report those under the `RepoDigests` field.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "container",
                    "stability": "development",
                    "type": "string[]",
                },
                {
                    "brief": "Container image tags. An example can be found in [Docker Image Inspect](https://docs.docker.com/engine/api/v1.43/#tag/Image/operation/ImageInspect). Should be only the `<tag>` section of the full name for example from `registry.example.com/my-org/my-image:<tag>`.\n",
                    "examples": [
                        [
                            "v1.27.1",
                            "3.5.7-0",
                        ],
                    ],
                    "name": "container.image.tags",
                    "requirement_level": "recommended",
                    "root_namespace": "container",
                    "stability": "development",
                    "type": "string[]",
                },
                {
                    "brief": "Container labels, `<key>` being the label name, the value being the label value.\n",
                    "examples": [
                        "nginx",
                    ],
                    "name": "container.label",
                    "note": "For example, a docker container label `app` with value `nginx` SHOULD be recorded as the `container.label.app` attribute with value `\"nginx\"`.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "container",
                    "stability": "development",
                    "type": "template[string]",
                },
                {
                    "brief": "Deprecated, use `container.label` instead.",
                    "deprecated": {
                        "note": "Replaced by `container.label`.",
                        "reason": "renamed",
                        "renamed_to": "container.label",
                    },
                    "examples": [
                        "nginx",
                    ],
                    "name": "container.labels",
                    "requirement_level": "recommended",
                    "root_namespace": "container",
                    "stability": "development",
                    "type": "template[string]",
                },
                {
                    "brief": "Container name used by container runtime.\n",
                    "examples": [
                        "opentelemetry-autoconf",
                    ],
                    "name": "container.name",
                    "requirement_level": "recommended",
                    "root_namespace": "container",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The container runtime managing this container.\n",
                    "examples": [
                        "docker",
                        "containerd",
                        "rkt",
                    ],
                    "name": "container.runtime",
                    "requirement_level": "recommended",
                    "root_namespace": "container",
                    "stability": "development",
                    "type": "string",
                },
            ],
            "root_namespace": "container",
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
