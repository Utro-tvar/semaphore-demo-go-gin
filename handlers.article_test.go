package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/", showIndexPage)

	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := io.ReadAll(w.Body)

		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOK

	})
}

func TestArticleUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/article/view/:article_id", showArticle)

	route := fmt.Sprintf("/article/view/%d", articleList[0].ID)
	req, _ := http.NewRequest("GET", route, nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := io.ReadAll(w.Body)

		pageOK := err == nil && strings.Index(string(p), fmt.Sprintf("<title>%s</title>", articleList[0].Title)) > 0

		return statusOK && pageOK
	})
}
