package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEveryPathIsThePageAs503(t *testing.T) {
	for _, path := range []string{"/", "/some/deep/path?x=1", "/favicon.ico"} {
		res := do(t, http.MethodGet, path)

		if res.StatusCode != http.StatusServiceUnavailable {
			t.Errorf("%s: status %d, want 503", path, res.StatusCode)
		}
		if got := res.Header.Get("Retry-After"); got != retryAfter {
			t.Errorf("%s: Retry-After %q, want %q", path, got, retryAfter)
		}
		if got := res.Header.Get("Cache-Control"); got != "no-store" {
			t.Errorf("%s: Cache-Control %q, want no-store", path, got)
		}
		if body, _ := io.ReadAll(res.Body); !bytes.Equal(body, page) {
			t.Errorf("%s: body is not the page", path)
		}
	}
}

func TestHeadHasNoBody(t *testing.T) {
	res := do(t, http.MethodHead, "/")

	if res.StatusCode != http.StatusServiceUnavailable {
		t.Errorf("status %d, want 503", res.StatusCode)
	}
	if body, _ := io.ReadAll(res.Body); len(body) != 0 {
		t.Errorf("HEAD returned %d bytes", len(body))
	}
}

func do(t *testing.T, method, path string) *http.Response {
	t.Helper()

	rec := httptest.NewRecorder()
	handler().ServeHTTP(rec, httptest.NewRequest(method, path, nil))

	return rec.Result()
}
