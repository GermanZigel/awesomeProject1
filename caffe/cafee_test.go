package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("expected status code: %d, got %d", http.StatusOK, status)
	}
}

func TestMainHandlerWhenOkFy(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	assert.Equal(t, http.StatusOK, status)
}
func TestMainHandlerWhenMissingCount(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusBadRequest {
		t.Errorf("expected status code: %d, got %d", http.StatusBadRequest, status)
	}

	expected := `count missing`
	if responseRecorder.Body.String() != expected {
		t.Errorf("expected body: %s, got %s", expected, responseRecorder.Body.String())
	}
}
func TestMainHandlerWhenMissingCountFy(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	assert.Equal(t, http.StatusBadRequest, status)
	assert.Equal(t, `count missing`, responseRecorder.Body.String())

}
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil) // здесь нужно создать запрос к сервису
	var cafeList = map[string][]string{
		"moscow": []string{"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"},
	}
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("expected status code: %d, got %d", http.StatusOK, status)
	}
	expectedCafee := strings.Join(cafeList["moscow"], ",")
	resCaff := responseRecorder.Body.String()
	if resCaff != expectedCafee {
		t.Errorf("expected cafee: %s, got %s", cafeList["moscow"], resCaff)
	}
	list := strings.Split(resCaff, ",")
	if len(list) != totalCount {
		t.Errorf("expected cafe count: %d, got %d", totalCount, len(list))
	}
}
func TestMainHandlerWhenCountMoreThanTotalFy(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil) // здесь нужно создать запрос к сервису
	var cafeList = map[string][]string{
		"moscow": []string{"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"},
	}
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	expectedCafee := strings.Join(cafeList["moscow"], ",")
	resCaff := responseRecorder.Body.String()
	list := strings.Split(resCaff, ",")
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, expectedCafee, resCaff)
	assert.Equal(t, totalCount, len(list))
}
