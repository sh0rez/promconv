package http

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func ExampleInstrument() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello %s!", r.PathValue("name"))
	})
	mux.Handle("/metrics", promhttp.Handler())

	handler := Instrument(Register(prometheus.DefaultRegisterer), mux)
	go func() {
		log.Fatalln(http.ListenAndServe(":4242", handler))
	}()

	if _, err := http.Get("http://localhost:4242/hello/prometheus"); err != nil {
		log.Fatalln(err)
	}

	res, err := http.Get("http://localhost:4242/metrics")
	if err != nil {
		log.Fatalln(err)
	}

	sc := bufio.NewScanner(res.Body)
	for sc.Scan() {
		if strings.HasPrefix(sc.Text(), "http_") {
			fmt.Println(sc.Text())
		}
	}

	// Output:
	// http_server_active_requests{http_request_method="GET",server_address="127.0.0.1",server_port="4242",url_scheme=""} 1
	// http_server_request_body_size_bucket{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type="",le="0.005"} 1
	// http_server_request_body_size_bucket{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type="",le="0.01"} 1
	// http_server_request_body_size_bucket{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type="",le="0.025"} 1
	// http_server_request_body_size_bucket{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type="",le="0.05"} 1
	// http_server_request_body_size_bucket{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type="",le="0.1"} 1
	// http_server_request_body_size_bucket{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type="",le="0.25"} 1
	// http_server_request_body_size_bucket{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type="",le="0.5"} 1
	// http_server_request_body_size_bucket{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type="",le="1"} 1
	// http_server_request_body_size_bucket{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type="",le="2.5"} 1
	// http_server_request_body_size_bucket{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type="",le="5"} 1
	// http_server_request_body_size_bucket{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type="",le="10"} 1
	// http_server_request_body_size_bucket{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type="",le="+Inf"} 1
	// http_server_request_body_size_sum{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type=""} 0
	// http_server_request_body_size_count{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type=""} 1
	// http_server_request_duration_bucket{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type="",le="0.005"} 1
	// http_server_request_duration_bucket{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type="",le="0.01"} 1
	// http_server_request_duration_bucket{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type="",le="0.025"} 1
	// http_server_request_duration_bucket{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type="",le="0.05"} 1
	// http_server_request_duration_bucket{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type="",le="0.1"} 1
	// http_server_request_duration_bucket{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type="",le="0.25"} 1
	// http_server_request_duration_bucket{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type="",le="0.5"} 1
	// http_server_request_duration_bucket{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type="",le="1"} 1
	// http_server_request_duration_bucket{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type="",le="2.5"} 1
	// http_server_request_duration_bucket{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type="",le="5"} 1
	// http_server_request_duration_bucket{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type="",le="10"} 1
	// http_server_request_duration_bucket{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type="",le="+Inf"} 1
	// http_server_request_duration_sum{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type=""} 1.4625e-05
	// http_server_request_duration_count{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type=""} 1
	// http_server_response_body_size_bucket{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type="",le="0.005"} 0
	// http_server_response_body_size_bucket{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type="",le="0.01"} 0
	// http_server_response_body_size_bucket{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type="",le="0.025"} 0
	// http_server_response_body_size_bucket{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type="",le="0.05"} 0
	// http_server_response_body_size_bucket{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type="",le="0.1"} 0
	// http_server_response_body_size_bucket{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type="",le="0.25"} 0
	// http_server_response_body_size_bucket{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type="",le="0.5"} 0
	// http_server_response_body_size_bucket{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type="",le="1"} 0
	// http_server_response_body_size_bucket{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type="",le="2.5"} 0
	// http_server_response_body_size_bucket{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type="",le="5"} 0
	// http_server_response_body_size_bucket{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type="",le="10"} 0
	// http_server_response_body_size_bucket{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type="",le="+Inf"} 1
	// http_server_response_body_size_sum{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type=""} 17
	// http_server_response_body_size_count{error_type="",http_request_method="GET",http_response_status_code="200",http_route="/hello/{name}",network_protocol_name="http",network_protocol_version="1.1",server_address="127.0.0.1",server_port="4242",url_scheme="",user_agent_synthetic_type=""} 1
}
