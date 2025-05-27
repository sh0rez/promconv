# PromConv

[![Go Reference](https://pkg.go.dev/badge/shorez.de/promconv.svg)](https://pkg.go.dev/shorez.de/promconv)

_[Semantic Conventions](https://opentelemetry.io/docs/specs/semconv/) for Prometheus_

**PromConv** has two goals:
- Bring semantic convetions to Prometheus
- Make the experience easier and safer than writing metrics by hand

To achieve those, **PromConv**:

- Offers a strongly typed API, making it hard to represent invalid states
- Heavily relies on code-generation using [Weaver](https://github.com/open-telemetry/weaver) based on official OpenTelemetry models.

## Examples

| Link | Description |
|-|-|
| [HTTP Server](https://pkg.go.dev/shorez.de/promconv/otel/http#example-Instrument) | Go HTTP ServeMux Instrumentation, according to https://opentelemetry.io/docs/specs/semconv/http/http-metrics/#http-server |

## License

Apache v2
