package promconv

//go:generate weaver registry generate go otel
//go:generate go run ./examples/generate.go
//go:generate gofmt -s -w otel examples
