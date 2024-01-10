package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Общий handler для всех тестов
var handler = http.HandlerFunc(mainHandle)

func performRequest(req *http.Request) *httptest.ResponseRecorder {
	// Здесь будет храниться объект для ответа
	responseRecorder := httptest.NewRecorder()
	// Вызов обработчика с тестовым запросом
	handler.ServeHTTP(responseRecorder, req)
	return responseRecorder
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil) // здесь нужно создать запрос к сервису

	response := performRequest(req)

	require.Equal(t, http.StatusOK, response.Code) // проверка, что код ответа 200
	cafes := strings.Split(response.Body.String(), ",")
	assert.Len(t, cafes, totalCount)
}

func TestMainHandlerSuccess(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	response := performRequest(req)

	require.Equal(t, http.StatusOK, response.Code)
	assert.NotEmptyf(t, response.Body.String(), "Тело ответа не должно быть пустым")
}

func TestMainHandWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=unknown", nil)

	response := performRequest(req)

	assert.Equal(t, http.StatusBadRequest, response.Code)
	assert.Contains(t, response.Body.String(), "wrong city value")
}

func TestMainHandNoCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2", nil)

	response := performRequest(req)

	assert.Equal(t, http.StatusBadRequest, response.Code)
	assert.Contains(t, response.Body.String(), "wrong city value")
}

func TestMainHandNoCount(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=moscow", nil)

	response := performRequest(req)

	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestMainHandlerNegativeCount(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=-2&city=moscow", nil)

	response := performRequest(req)

	require.Equal(t, http.StatusBadRequest, response.Code)
	assert.NotEmptyf(t, response.Body.String(), "Тело ответа не должно быть пустым")
	assert.Contains(t, response.Body.String(), "wrong count value")
}
