package vcs

import (
	"github.com/prometheus/client_golang/prometheus"
)

// The amount of time since its creation it took a change (pull request/merge request/changelist) to get merged into the target(base) ref.
type ChangeTimeToMerge struct {
	*prometheus.GaugeVec
	extra ChangeTimeToMergeExtra
}

func NewChangeTimeToMerge() ChangeTimeToMerge {
	labels := []string{AttrRefHeadName("").Key(), AttrRepositoryUrlFull("").Key(), AttrOwnerName("").Key(), AttrRefBaseName("").Key(), AttrRepositoryName("").Key(), AttrProviderName("").Key(), AttrRefBaseRevision("").Key(), AttrRefHeadRevision("").Key()}
	return ChangeTimeToMerge{GaugeVec: prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "vcs_change_time_to_merge",
		Help: "The amount of time since its creation it took a change (pull request/merge request/changelist) to get merged into the target(base) ref.",
	}, labels)}
}

func (m ChangeTimeToMerge) With(refHeadName AttrRefHeadName, repositoryUrlFull AttrRepositoryUrlFull, extras ...interface {
	VcsOwnerName() AttrOwnerName
	VcsRefBaseName() AttrRefBaseName
	VcsRepositoryName() AttrRepositoryName
	VcsProviderName() AttrProviderName
	VcsRefBaseRevision() AttrRefBaseRevision
	VcsRefHeadRevision() AttrRefHeadRevision
}) prometheus.Gauge {
	if extras == nil {
		extras = append(extras, m.extra)
	}
	extra := extras[0]

	return m.GaugeVec.WithLabelValues(refHeadName.Value(), repositoryUrlFull.Value(), extra.VcsOwnerName().Value(), extra.VcsRefBaseName().Value(), extra.VcsRepositoryName().Value(), extra.VcsProviderName().Value(), extra.VcsRefBaseRevision().Value(), extra.VcsRefHeadRevision().Value())
}

// Deprecated: Use [ChangeTimeToMerge.With] instead
func (m ChangeTimeToMerge) WithLabelValues(lvs ...string) prometheus.Gauge {
	return m.GaugeVec.WithLabelValues(lvs...)
}

func (a ChangeTimeToMerge) WithOwnerName(attr interface{ VcsOwnerName() AttrOwnerName }) ChangeTimeToMerge {
	a.extra.AttrOwnerName = attr.VcsOwnerName()
	return a
}
func (a ChangeTimeToMerge) WithRefBaseName(attr interface{ VcsRefBaseName() AttrRefBaseName }) ChangeTimeToMerge {
	a.extra.AttrRefBaseName = attr.VcsRefBaseName()
	return a
}
func (a ChangeTimeToMerge) WithRepositoryName(attr interface{ VcsRepositoryName() AttrRepositoryName }) ChangeTimeToMerge {
	a.extra.AttrRepositoryName = attr.VcsRepositoryName()
	return a
}
func (a ChangeTimeToMerge) WithProviderName(attr interface{ VcsProviderName() AttrProviderName }) ChangeTimeToMerge {
	a.extra.AttrProviderName = attr.VcsProviderName()
	return a
}
func (a ChangeTimeToMerge) WithRefBaseRevision(attr interface{ VcsRefBaseRevision() AttrRefBaseRevision }) ChangeTimeToMerge {
	a.extra.AttrRefBaseRevision = attr.VcsRefBaseRevision()
	return a
}
func (a ChangeTimeToMerge) WithRefHeadRevision(attr interface{ VcsRefHeadRevision() AttrRefHeadRevision }) ChangeTimeToMerge {
	a.extra.AttrRefHeadRevision = attr.VcsRefHeadRevision()
	return a
}

type ChangeTimeToMergeExtra struct {
	// The group owner within the version control system
	AttrOwnerName AttrOwnerName `otel:"vcs.owner.name"` // The name of the [reference] such as **branch** or **tag** in the repository
	//
	// [reference]: https://git-scm.com/docs/gitglossary#def_ref
	AttrRefBaseName    AttrRefBaseName    `otel:"vcs.ref.base.name"`   // The human readable name of the repository. It SHOULD NOT include any additional identifier like Group/SubGroup in GitLab or organization in GitHub
	AttrRepositoryName AttrRepositoryName `otel:"vcs.repository.name"` // The name of the version control system provider
	AttrProviderName   AttrProviderName   `otel:"vcs.provider.name"`   // The revision, literally [revised version], The revision most often refers to a commit object in Git, or a revision number in SVN
	//
	// [revised version]: https://www.merriam-webster.com/dictionary/revision
	AttrRefBaseRevision AttrRefBaseRevision `otel:"vcs.ref.base.revision"` // The revision, literally [revised version], The revision most often refers to a commit object in Git, or a revision number in SVN
	//
	// [revised version]: https://www.merriam-webster.com/dictionary/revision
	AttrRefHeadRevision AttrRefHeadRevision `otel:"vcs.ref.head.revision"`
}

func (a ChangeTimeToMergeExtra) VcsOwnerName() AttrOwnerName           { return a.AttrOwnerName }
func (a ChangeTimeToMergeExtra) VcsRefBaseName() AttrRefBaseName       { return a.AttrRefBaseName }
func (a ChangeTimeToMergeExtra) VcsRepositoryName() AttrRepositoryName { return a.AttrRepositoryName }
func (a ChangeTimeToMergeExtra) VcsProviderName() AttrProviderName     { return a.AttrProviderName }
func (a ChangeTimeToMergeExtra) VcsRefBaseRevision() AttrRefBaseRevision {
	return a.AttrRefBaseRevision
}
func (a ChangeTimeToMergeExtra) VcsRefHeadRevision() AttrRefHeadRevision {
	return a.AttrRefHeadRevision
}

/*
State {
    name: "vec.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "AttrExtra": "ChangeTimeToMergeExtra",
        "Instr": "Gauge",
        "InstrMap": {
            "counter": "Counter",
            "gauge": "Gauge",
            "histogram": "Histogram",
            "updowncounter": "Gauge",
        },
        "Name": "change.time_to_merge",
        "Type": "ChangeTimeToMerge",
        "attributes": [
            {
                "brief": "The name of the [reference](https://git-scm.com/docs/gitglossary#def_ref) such as **branch** or **tag** in the repository.\n",
                "examples": [
                    "my-feature-branch",
                    "tag-1-test",
                ],
                "name": "vcs.ref.head.name",
                "note": "`head` refers to where you are right now; the current reference at a\ngiven time.\n",
                "requirement_level": "required",
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
                "brief": "The name of the [reference](https://git-scm.com/docs/gitglossary#def_ref) such as **branch** or **tag** in the repository.\n",
                "examples": [
                    "my-feature-branch",
                    "tag-1-test",
                ],
                "name": "vcs.ref.base.name",
                "note": "`base` refers to the starting point of a change. For example, `main`\nwould be the base reference of type branch if you've created a new\nreference of type branch from it and created new commits.\n",
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
            {
                "brief": "The revision, literally [revised version](https://www.merriam-webster.com/dictionary/revision), The revision most often refers to a commit object in Git, or a revision number in SVN.\n",
                "examples": [
                    "9d59409acf479dfa0df1aa568182e43e43df8bbe28d60fcf2bc52e30068802cc",
                    "main",
                    "123",
                    "HEAD",
                ],
                "name": "vcs.ref.base.revision",
                "note": "`base` refers to the starting point of a change. For example, `main`\nwould be the base reference of type branch if you've created a new\nreference of type branch from it and created new commits. The\nrevision can be a full [hash value (see\nglossary)](https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.186-5.pdf),\nof the recorded change to a ref within a repository pointing to a\ncommit [commit](https://git-scm.com/docs/git-commit) object. It does\nnot necessarily have to be a hash; it can simply define a [revision\nnumber](https://svnbook.red-bean.com/en/1.7/svn.tour.revs.specifiers.html)\nwhich is an integer that is monotonically increasing. In cases where\nit is identical to the `ref.base.name`, it SHOULD still be included.\nIt is up to the implementer to decide which value to set as the\nrevision based on the VCS system and situational context.\n",
                "requirement_level": "opt_in",
                "stability": "development",
                "type": "string",
            },
            {
                "brief": "The revision, literally [revised version](https://www.merriam-webster.com/dictionary/revision), The revision most often refers to a commit object in Git, or a revision number in SVN.\n",
                "examples": [
                    "9d59409acf479dfa0df1aa568182e43e43df8bbe28d60fcf2bc52e30068802cc",
                    "main",
                    "123",
                    "HEAD",
                ],
                "name": "vcs.ref.head.revision",
                "note": "`head` refers to where you are right now; the current reference at a\ngiven time.The revision can be a full [hash value (see\nglossary)](https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.186-5.pdf),\nof the recorded change to a ref within a repository pointing to a\ncommit [commit](https://git-scm.com/docs/git-commit) object. It does\nnot necessarily have to be a hash; it can simply define a [revision\nnumber](https://svnbook.red-bean.com/en/1.7/svn.tour.revs.specifiers.html)\nwhich is an integer that is monotonically increasing. In cases where\nit is identical to the `ref.head.name`, it SHOULD still be included.\nIt is up to the implementer to decide which value to set as the\nrevision based on the VCS system and situational context.\n",
                "requirement_level": "opt_in",
                "stability": "development",
                "type": "string",
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
                    "brief": "The name of the [reference](https://git-scm.com/docs/gitglossary#def_ref) such as **branch** or **tag** in the repository.\n",
                    "examples": [
                        "my-feature-branch",
                        "tag-1-test",
                    ],
                    "name": "vcs.ref.base.name",
                    "note": "`base` refers to the starting point of a change. For example, `main`\nwould be the base reference of type branch if you've created a new\nreference of type branch from it and created new commits.\n",
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
                {
                    "brief": "The name of the [reference](https://git-scm.com/docs/gitglossary#def_ref) such as **branch** or **tag** in the repository.\n",
                    "examples": [
                        "my-feature-branch",
                        "tag-1-test",
                    ],
                    "name": "vcs.ref.head.name",
                    "note": "`head` refers to where you are right now; the current reference at a\ngiven time.\n",
                    "requirement_level": "required",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The revision, literally [revised version](https://www.merriam-webster.com/dictionary/revision), The revision most often refers to a commit object in Git, or a revision number in SVN.\n",
                    "examples": [
                        "9d59409acf479dfa0df1aa568182e43e43df8bbe28d60fcf2bc52e30068802cc",
                        "main",
                        "123",
                        "HEAD",
                    ],
                    "name": "vcs.ref.head.revision",
                    "note": "`head` refers to where you are right now; the current reference at a\ngiven time.The revision can be a full [hash value (see\nglossary)](https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.186-5.pdf),\nof the recorded change to a ref within a repository pointing to a\ncommit [commit](https://git-scm.com/docs/git-commit) object. It does\nnot necessarily have to be a hash; it can simply define a [revision\nnumber](https://svnbook.red-bean.com/en/1.7/svn.tour.revs.specifiers.html)\nwhich is an integer that is monotonically increasing. In cases where\nit is identical to the `ref.head.name`, it SHOULD still be included.\nIt is up to the implementer to decide which value to set as the\nrevision based on the VCS system and situational context.\n",
                    "requirement_level": "opt_in",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The revision, literally [revised version](https://www.merriam-webster.com/dictionary/revision), The revision most often refers to a commit object in Git, or a revision number in SVN.\n",
                    "examples": [
                        "9d59409acf479dfa0df1aa568182e43e43df8bbe28d60fcf2bc52e30068802cc",
                        "main",
                        "123",
                        "HEAD",
                    ],
                    "name": "vcs.ref.base.revision",
                    "note": "`base` refers to the starting point of a change. For example, `main`\nwould be the base reference of type branch if you've created a new\nreference of type branch from it and created new commits. The\nrevision can be a full [hash value (see\nglossary)](https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.186-5.pdf),\nof the recorded change to a ref within a repository pointing to a\ncommit [commit](https://git-scm.com/docs/git-commit) object. It does\nnot necessarily have to be a hash; it can simply define a [revision\nnumber](https://svnbook.red-bean.com/en/1.7/svn.tour.revs.specifiers.html)\nwhich is an integer that is monotonically increasing. In cases where\nit is identical to the `ref.base.name`, it SHOULD still be included.\nIt is up to the implementer to decide which value to set as the\nrevision based on the VCS system and situational context.\n",
                    "requirement_level": "opt_in",
                    "stability": "development",
                    "type": "string",
                },
            ],
            "brief": "The amount of time since its creation it took a change (pull request/merge request/changelist) to get merged into the target(base) ref.",
            "entity_associations": [
                "vcs.repo",
                "vcs.ref",
            ],
            "events": [],
            "id": "metric.vcs.change.time_to_merge",
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
                    "vcs.ref.base.name": {
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
                    "vcs.ref.base.revision": {
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
                    "vcs.ref.head.name": {
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
                    "vcs.ref.head.revision": {
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
            "metric_name": "vcs.change.time_to_merge",
            "name": none,
            "root_namespace": "vcs",
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
