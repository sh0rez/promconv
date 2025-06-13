package geo

// Two-letter code representing continent’s name
type AttrContinentCode string // geo.continent.code

func (AttrContinentCode) Development()    {}
func (AttrContinentCode) Recommended()    {}
func (AttrContinentCode) Key() string     { return "geo_continent_code" }
func (a AttrContinentCode) Value() string { return string(a) }

const ContinentCodeAf AttrContinentCode = "AF"
const ContinentCodeAn AttrContinentCode = "AN"
const ContinentCodeAs AttrContinentCode = "AS"
const ContinentCodeEu AttrContinentCode = "EU"
const ContinentCodeNa AttrContinentCode = "NA"
const ContinentCodeOc AttrContinentCode = "OC"
const ContinentCodeSa AttrContinentCode = "SA"

// Two-letter ISO Country Code ([ISO 3166-1 alpha2])
//
// [ISO 3166-1 alpha2]: https://wikipedia.org/wiki/ISO_3166-1#Codes
type AttrCountryIsoCode string // geo.country.iso_code

func (AttrCountryIsoCode) Development()    {}
func (AttrCountryIsoCode) Recommended()    {}
func (AttrCountryIsoCode) Key() string     { return "geo_country_iso_code" }
func (a AttrCountryIsoCode) Value() string { return string(a) }

// Locality name. Represents the name of a city, town, village, or similar populated place
type AttrLocalityName string // geo.locality.name

func (AttrLocalityName) Development()    {}
func (AttrLocalityName) Recommended()    {}
func (AttrLocalityName) Key() string     { return "geo_locality_name" }
func (a AttrLocalityName) Value() string { return string(a) }

// Latitude of the geo location in [WGS84]
//
// [WGS84]: https://wikipedia.org/wiki/World_Geodetic_System#WGS84
type AttrLocationLat string // geo.location.lat

func (AttrLocationLat) Development()    {}
func (AttrLocationLat) Recommended()    {}
func (AttrLocationLat) Key() string     { return "geo_location_lat" }
func (a AttrLocationLat) Value() string { return string(a) }

// Longitude of the geo location in [WGS84]
//
// [WGS84]: https://wikipedia.org/wiki/World_Geodetic_System#WGS84
type AttrLocationLon string // geo.location.lon

func (AttrLocationLon) Development()    {}
func (AttrLocationLon) Recommended()    {}
func (AttrLocationLon) Key() string     { return "geo_location_lon" }
func (a AttrLocationLon) Value() string { return string(a) }

// Postal code associated with the location. Values appropriate for this field may also be known as a postcode or ZIP code and will vary widely from country to country
type AttrPostalCode string // geo.postal_code

func (AttrPostalCode) Development()    {}
func (AttrPostalCode) Recommended()    {}
func (AttrPostalCode) Key() string     { return "geo_postal_code" }
func (a AttrPostalCode) Value() string { return string(a) }

// Region ISO code ([ISO 3166-2])
//
// [ISO 3166-2]: https://wikipedia.org/wiki/ISO_3166-2
type AttrRegionIsoCode string // geo.region.iso_code

func (AttrRegionIsoCode) Development()    {}
func (AttrRegionIsoCode) Recommended()    {}
func (AttrRegionIsoCode) Key() string     { return "geo_region_iso_code" }
func (a AttrRegionIsoCode) Value() string { return string(a) }

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "Two-letter code representing continent’s name.\n",
                    "name": "geo.continent.code",
                    "requirement_level": "recommended",
                    "root_namespace": "geo",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": "Africa",
                                "deprecated": none,
                                "id": "af",
                                "note": none,
                                "stability": "development",
                                "value": "AF",
                            },
                            {
                                "brief": "Antarctica",
                                "deprecated": none,
                                "id": "an",
                                "note": none,
                                "stability": "development",
                                "value": "AN",
                            },
                            {
                                "brief": "Asia",
                                "deprecated": none,
                                "id": "as",
                                "note": none,
                                "stability": "development",
                                "value": "AS",
                            },
                            {
                                "brief": "Europe",
                                "deprecated": none,
                                "id": "eu",
                                "note": none,
                                "stability": "development",
                                "value": "EU",
                            },
                            {
                                "brief": "North America",
                                "deprecated": none,
                                "id": "na",
                                "note": none,
                                "stability": "development",
                                "value": "NA",
                            },
                            {
                                "brief": "Oceania",
                                "deprecated": none,
                                "id": "oc",
                                "note": none,
                                "stability": "development",
                                "value": "OC",
                            },
                            {
                                "brief": "South America",
                                "deprecated": none,
                                "id": "sa",
                                "note": none,
                                "stability": "development",
                                "value": "SA",
                            },
                        ],
                    },
                },
                {
                    "brief": "Two-letter ISO Country Code ([ISO 3166-1 alpha2](https://wikipedia.org/wiki/ISO_3166-1#Codes)).\n",
                    "examples": [
                        "CA",
                    ],
                    "name": "geo.country.iso_code",
                    "requirement_level": "recommended",
                    "root_namespace": "geo",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Locality name. Represents the name of a city, town, village, or similar populated place.\n",
                    "examples": [
                        "Montreal",
                        "Berlin",
                    ],
                    "name": "geo.locality.name",
                    "requirement_level": "recommended",
                    "root_namespace": "geo",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Latitude of the geo location in [WGS84](https://wikipedia.org/wiki/World_Geodetic_System#WGS84).\n",
                    "examples": [
                        45.505918,
                    ],
                    "name": "geo.location.lat",
                    "requirement_level": "recommended",
                    "root_namespace": "geo",
                    "stability": "development",
                    "type": "double",
                },
                {
                    "brief": "Longitude of the geo location in [WGS84](https://wikipedia.org/wiki/World_Geodetic_System#WGS84).\n",
                    "examples": [
                        -73.61483,
                    ],
                    "name": "geo.location.lon",
                    "requirement_level": "recommended",
                    "root_namespace": "geo",
                    "stability": "development",
                    "type": "double",
                },
                {
                    "brief": "Postal code associated with the location. Values appropriate for this field may also be known as a postcode or ZIP code and will vary widely from country to country.\n",
                    "examples": [
                        "94040",
                    ],
                    "name": "geo.postal_code",
                    "requirement_level": "recommended",
                    "root_namespace": "geo",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Region ISO code ([ISO 3166-2](https://wikipedia.org/wiki/ISO_3166-2)).\n",
                    "examples": [
                        "CA-QC",
                    ],
                    "name": "geo.region.iso_code",
                    "requirement_level": "recommended",
                    "root_namespace": "geo",
                    "stability": "development",
                    "type": "string",
                },
            ],
            "root_namespace": "geo",
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
