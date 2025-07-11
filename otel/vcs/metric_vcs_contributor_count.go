package vcs

import (
	"github.com/prometheus/client_golang/prometheus"
)

// The number of unique contributors to a repository
type ContributorCount struct {
	*prometheus.GaugeVec
	extra ContributorCountExtra
}

func NewContributorCount() ContributorCount {
	labels := []string{AttrRepositoryUrlFull("").Key(), AttrOwnerName("").Key(), AttrRepositoryName("").Key(), AttrProviderName("").Key()}
	return ContributorCount{GaugeVec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "vcs_contributor_count",
		Help: "The number of unique contributors to a repository",
	}, labels)}
}

func (m ContributorCount) With(repositoryUrlFull AttrRepositoryUrlFull, extras ...interface {
	VcsOwnerName() AttrOwnerName
	VcsRepositoryName() AttrRepositoryName
	VcsProviderName() AttrProviderName
}) prometheus.Gauge {
	if extras == nil {
		extras = append(extras, m.extra)
	}
	extra := extras[0]

	return m.GaugeVec.WithLabelValues(repositoryUrlFull.Value(), extra.VcsOwnerName().Value(), extra.VcsRepositoryName().Value(), extra.VcsProviderName().Value())
}

// Deprecated: Use [ContributorCount.With] instead
func (m ContributorCount) WithLabelValues(lvs ...string) prometheus.Gauge {
	return m.GaugeVec.WithLabelValues(lvs...)
}

func (a ContributorCount) WithOwnerName(attr interface{ VcsOwnerName() AttrOwnerName }) ContributorCount {
	a.extra.AttrOwnerName = attr.VcsOwnerName()
	return a
}
func (a ContributorCount) WithRepositoryName(attr interface{ VcsRepositoryName() AttrRepositoryName }) ContributorCount {
	a.extra.AttrRepositoryName = attr.VcsRepositoryName()
	return a
}
func (a ContributorCount) WithProviderName(attr interface{ VcsProviderName() AttrProviderName }) ContributorCount {
	a.extra.AttrProviderName = attr.VcsProviderName()
	return a
}

type ContributorCountExtra struct {
	// The group owner within the version control system
	AttrOwnerName      AttrOwnerName      `otel:"vcs.owner.name"`      // The human readable name of the repository. It SHOULD NOT include any additional identifier like Group/SubGroup in GitLab or organization in GitHub
	AttrRepositoryName AttrRepositoryName `otel:"vcs.repository.name"` // The name of the version control system provider
	AttrProviderName   AttrProviderName   `otel:"vcs.provider.name"`
}

func (a ContributorCountExtra) VcsOwnerName() AttrOwnerName           { return a.AttrOwnerName }
func (a ContributorCountExtra) VcsRepositoryName() AttrRepositoryName { return a.AttrRepositoryName }
func (a ContributorCountExtra) VcsProviderName() AttrProviderName     { return a.AttrProviderName }

/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "ContributorCountExtra",
        "Instr": "Gauge",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "contributor.count",
        "Type": "ContributorCount",
        "attributes": [
            {
                "brief": "The [canonical URL](https://support.google.com/webmasters/answer/10347851?hl=en#:~:text=A%20canonical%20URL%20is%20the,Google%20chooses%20one%20as%20canonical.) of the repository providing the complete HTTP(S) address in order to locate and identify the repository through a browser.\n",
                "examples": [
                    "https://github.com/opentelemetry/open-telemetry-collector-contrib",
                    "https://gitlab.com/my-org/my-project/my-projects-project/repo",
                ],
                "name": "vcs.repository.url.full",
                "note": "In Git Version Control Systems, the canonical URL SHOULD NOT include\nthe `.git` extension.\n",
                "requirement_level": "required",
                "stability": "development",
                "type": "string",
            },
            {
                "brief": "The group owner within the version control system.\n",
                "examples": [
                    "my-org",
                    "myteam",
                    "business-unit",
                ],
                "name": "vcs.owner.name",
                "requirement_level": "recommended",
                "stability": "development",
                "type": "string",
            },
            {
                "brief": "The human readable name of the repository. It SHOULD NOT include any additional identifier like Group/SubGroup in GitLab or organization in GitHub.\n",
                "examples": [
                    "semantic-conventions",
                    "my-cool-repo",
                ],
                "name": "vcs.repository.name",
                "note": "Due to it only being the name, it can clash with forks of the same\nrepository if collecting telemetry across multiple orgs or groups in\nthe same backends.\n",
                "requirement_level": "recommended",
                "stability": "development",
                "type": "string",
            },
            {
                "brief": "The name of the version control system provider.\n",
                "examples": [
                    "github",
                    "gitlab",
                    "gitea",
                    "bitbucket",
                ],
                "name": "vcs.provider.name",
                "requirement_level": "opt_in",
                "stability": "development",
                "type": {
                    "members": [
                        {
                            "brief": "[GitHub](https://github.com)",
                            "deprecated": none,
                            "id": "github",
                            "note": none,
                            "stability": "development",
                            "value": "github",
                        },
                        {
                            "brief": "[GitLab](https://gitlab.com)",
                            "deprecated": none,
                            "id": "gitlab",
                            "note": none,
                            "stability": "development",
                            "value": "gitlab",
                        },
                        {
                            "brief": "Deprecated, use `gitea` instead.",
                            "deprecated": "Replaced by `gitea`.",
                            "id": "gittea",
                            "note": none,
                            "stability": "development",
                            "value": "gittea",
                        },
                        {
                            "brief": "[Gitea](https://gitea.io)",
                            "deprecated": none,
                            "id": "gitea",
                            "note": none,
                            "stability": "development",
                            "value": "gitea",
                        },
                        {
                            "brief": "[Bitbucket](https://bitbucket.org)",
                            "deprecated": none,
                            "id": "bitbucket",
                            "note": none,
                            "stability": "development",
                            "value": "bitbucket",
                        },
                    ],
                },
            },
        ],
        "ctx": {
            "attributes": [
                {
                    "brief": "The human readable name of the repository. It SHOULD NOT include any additional identifier like Group/SubGroup in GitLab or organization in GitHub.\n",
                    "examples": [
                        "semantic-conventions",
                        "my-cool-repo",
                    ],
                    "name": "vcs.repository.name",
                    "note": "Due to it only being the name, it can clash with forks of the same\nrepository if collecting telemetry across multiple orgs or groups in\nthe same backends.\n",
                    "requirement_level": "recommended",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The group owner within the version control system.\n",
                    "examples": [
                        "my-org",
                        "myteam",
                        "business-unit",
                    ],
                    "name": "vcs.owner.name",
                    "requirement_level": "recommended",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The [canonical URL](https://support.google.com/webmasters/answer/10347851?hl=en#:~:text=A%20canonical%20URL%20is%20the,Google%20chooses%20one%20as%20canonical.) of the repository providing the complete HTTP(S) address in order to locate and identify the repository through a browser.\n",
                    "examples": [
                        "https://github.com/opentelemetry/open-telemetry-collector-contrib",
                        "https://gitlab.com/my-org/my-project/my-projects-project/repo",
                    ],
                    "name": "vcs.repository.url.full",
                    "note": "In Git Version Control Systems, the canonical URL SHOULD NOT include\nthe `.git` extension.\n",
                    "requirement_level": "required",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The name of the version control system provider.\n",
                    "examples": [
                        "github",
                        "gitlab",
                        "gitea",
                        "bitbucket",
                    ],
                    "name": "vcs.provider.name",
                    "requirement_level": "opt_in",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "[GitHub](https://github.com)",
                                "deprecated": none,
                                "id": "github",
                                "note": none,
                                "stability": "development",
                                "value": "github",
                            },
                            {
                                "brief": "[GitLab](https://gitlab.com)",
                                "deprecated": none,
                                "id": "gitlab",
                                "note": none,
                                "stability": "development",
                                "value": "gitlab",
                            },
                            {
                                "brief": "Deprecated, use `gitea` instead.",
                                "deprecated": "Replaced by `gitea`.",
                                "id": "gittea",
                                "note": none,
                                "stability": "development",
                                "value": "gittea",
                            },
                            {
                                "brief": "[Gitea](https://gitea.io)",
                                "deprecated": none,
                                "id": "gitea",
                                "note": none,
                                "stability": "development",
                                "value": "gitea",
                            },
                            {
                                "brief": "[Bitbucket](https://bitbucket.org)",
                                "deprecated": none,
                                "id": "bitbucket",
                                "note": none,
                                "stability": "development",
                                "value": "bitbucket",
                            },
                        ],
                    },
                },
            ],
            "brief": "The number of unique contributors to a repository",
            "entity_associations": [
                "vcs.repo",
            ],
            "events": [],
            "id": "metric.vcs.contributor.count",
            "instrument": "gauge",
            "lineage": {
                "attributes": {
                    "vcs.owner.name": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.vcs.repository",
                    },
                    "vcs.provider.name": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.vcs.repository",
                    },
                    "vcs.repository.name": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.vcs.repository",
                    },
                    "vcs.repository.url.full": {
                        "inherited_fields": [
                            "brief",
                            "examples",
                            "note",
                            "stability",
                        ],
                        "locally_overridden_fields": [
                            "requirement_level",
                        ],
                        "source_group": "registry.vcs.repository",
                    },
                },
                "provenance": {
                    "path": "https://github.com/open-telemetry/semantic-conventions.git[model]/vcs/metrics.yaml",
                    "registry_id": "main",
                },
            },
            "metric_name": "vcs.contributor.count",
            "name": none,
            "root_namespace": "vcs",
            "span_kind": none,
            "stability": "development",
            "type": "metric",
            "unit": "{contributor}",
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
