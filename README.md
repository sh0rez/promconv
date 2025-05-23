# PromConv

[![Go Reference](https://pkg.go.dev/badge/shorez.de/promconv.svg)](https://pkg.go.dev/shorez.de/promconv)

_[Semantic Conventions](https://opentelemetry.io/docs/specs/semconv/) for Prometheus_

*PromConv* has two goals:
- Bring semantic convetions to Prometheus
- Make the experience easier and safer than writing metrics by hand

To achieve those, *PromConv*:

- Offers a strongly typed API, making it hard to represent invalid states
- Heavily relies on code-generation using [Weaver](https://github.com/open-telemetry/weaver) based on official OpenTelemetry models.
