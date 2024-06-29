package response_test

import (
	"ZrifAppsCore/response"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBuildSuccessResponse(t *testing.T) {
	// Test success response
	w := httptest.NewRecorder()
	payload := map[string]string{"key": "value"}
	response.BuildResponse(http.StatusOK, "success", payload, nil, w)
	var actualResponse response.Response
	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
	if err != nil {
		t.Errorf("Error unmarshaling response: %v", err)
	}
	if actualResponse.Meta.Message != "success" {
		t.Errorf("Unexpected message: %s", actualResponse.Meta.Message)
	}
	if actualResponse.Meta.Code != http.StatusOK {
		t.Errorf("Unexpected code: %d", actualResponse.Meta.Code)
	}
	if actualResponse.Meta.Status != "success" {
		t.Errorf("Unexpected status: %s", actualResponse.Meta.Status)
	}

	actualData, err := json.Marshal(actualResponse.Data)
	if err != nil {
		t.Errorf("Error marshaling actual data: %v", err)
	}
	expectedData, err := json.Marshal(payload)
	if err != nil {
		t.Errorf("Error marshaling expected data: %v", err)
	}
	if string(actualData) != string(expectedData) {
		t.Errorf("Unexpected payload: got %s, want %s", string(actualData), string(expectedData))
	}
}

func TestBuildErrorResponse(t *testing.T) {
	// Test error response
	w := httptest.NewRecorder()
	errorPayload := map[string]string{"error": "error"}
	response.BuildResponse(http.StatusInternalServerError, "error", nil, errorPayload, w)
	var actualResponseError response.ResponseError
	err := json.Unmarshal(w.Body.Bytes(), &actualResponseError)
	if err != nil {
		t.Errorf("Error unmarshaling response: %v", err)
	}
	if actualResponseError.Meta.Message != "error" {
		t.Errorf("Unexpected message: %s", actualResponseError.Meta.Message)
	}
	if actualResponseError.Meta.Code != http.StatusInternalServerError {
		t.Errorf("Unexpected code: %d", actualResponseError.Meta.Code)
	}
	if actualResponseError.Meta.Status != "error" {
		t.Errorf("Unexpected status: %s", actualResponseError.Meta.Status)
	}

	actualErrorData, err := json.Marshal(actualResponseError.Error)
	if err != nil {
		t.Errorf("Error marshaling actual error data: %v", err)
	}
	expectedErrorData, err := json.Marshal(errorPayload)
	if err != nil {
		t.Errorf("Error marshaling expected error data: %v", err)
	}
	if string(actualErrorData) != string(expectedErrorData) {
		t.Errorf("Unexpected error payload: got %s, want %s", string(actualErrorData), string(expectedErrorData))
	}
}
