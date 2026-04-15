package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"sort"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestAsyncAPISpecEndpoint(t *testing.T) {
	state := newFakeState(t)
	srv := New(state)
	ts := httptest.NewServer(srv.handler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/v0/asyncapi.yaml")
	if err != nil {
		t.Fatalf("GET /v0/asyncapi.yaml: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want 200", resp.StatusCode)
	}
	ct := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(ct, "text/yaml") {
		t.Errorf("Content-Type = %q, want text/yaml", ct)
	}
	body, _ := io.ReadAll(resp.Body)
	if !strings.Contains(string(body), "asyncapi:") {
		t.Error("response body does not contain AsyncAPI spec")
	}
}

func TestOpenAPISpecEndpoint(t *testing.T) {
	state := newFakeState(t)
	srv := New(state)
	ts := httptest.NewServer(srv.handler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/v0/openapi.yaml")
	if err != nil {
		t.Fatalf("GET /v0/openapi.yaml: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want 200", resp.StatusCode)
	}
	ct := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(ct, "text/yaml") {
		t.Errorf("Content-Type = %q, want text/yaml", ct)
	}
	body, _ := io.ReadAll(resp.Body)
	if !strings.Contains(string(body), "openapi:") {
		t.Error("response body does not contain OpenAPI spec")
	}
}

func TestOpenAPISpecContainsOnlyHTTPSurvivorPaths(t *testing.T) {
	state := newFakeState(t)
	srv := New(state)
	ts := httptest.NewServer(srv.handler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/v0/openapi.yaml")
	if err != nil {
		t.Fatalf("GET /v0/openapi.yaml: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}

	var doc struct {
		Paths map[string]any `yaml:"paths"`
	}
	if err := yaml.Unmarshal(body, &doc); err != nil {
		t.Fatalf("unmarshal openapi yaml: %v", err)
	}

	var got []string
	for path := range doc.Paths {
		got = append(got, path)
	}
	sort.Strings(got)

	want := []string{
		"/health",
		"/v0/asyncapi.yaml",
		"/v0/city",
		"/v0/openapi.yaml",
		"/v0/provider-readiness",
		"/v0/readiness",
		"/v0/ws",
	}
	if strings.Join(got, "\n") != strings.Join(want, "\n") {
		t.Fatalf("openapi paths = %v, want %v", got, want)
	}
}

func TestAsyncAPISpecContainsExpectedProtocolChannels(t *testing.T) {
	state := newFakeState(t)
	srv := New(state)
	ts := httptest.NewServer(srv.handler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/v0/asyncapi.yaml")
	if err != nil {
		t.Fatalf("GET /v0/asyncapi.yaml: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}

	var doc struct {
		Channels map[string]any `yaml:"channels"`
	}
	if err := yaml.Unmarshal(body, &doc); err != nil {
		t.Fatalf("unmarshal asyncapi yaml: %v", err)
	}

	var gotProtocol []string
	for path := range doc.Channels {
		if strings.HasPrefix(path, "protocol/") {
			gotProtocol = append(gotProtocol, path)
		}
	}
	sort.Strings(gotProtocol)

	wantProtocol := []string{
		"protocol/error",
		"protocol/event",
		"protocol/hello",
		"protocol/subscription-start",
		"protocol/subscription-stop",
	}
	if strings.Join(gotProtocol, "\n") != strings.Join(wantProtocol, "\n") {
		t.Fatalf("asyncapi protocol channels = %v, want %v", gotProtocol, wantProtocol)
	}

	if _, ok := doc.Channels["actions/cities.list/request"]; !ok {
		t.Fatal("asyncapi missing supervisor action channel actions/cities.list/request")
	}
	if _, ok := doc.Channels["actions/health.get/request"]; ok {
		t.Fatal("asyncapi advertised HTTP-only action channel actions/health.get/request")
	}
}
