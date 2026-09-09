package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHomePage(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	homePage(rec, req)
	if rec.Code != http.StatusOK { t.Fatalf("expected 200, got %d", rec.Code) }
	if !strings.Contains(rec.Body.String(), "Tharun Kumaran") { t.Fatal("expected name in home page") }
}

func TestContactPage(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/contact", nil)
	rec := httptest.NewRecorder()
	contactPage(rec, req)
	if !strings.Contains(rec.Body.String(), "tharunkumaranm@outlook.com") { t.Fatal("expected contact email") }
}
