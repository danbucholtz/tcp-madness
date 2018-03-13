package main

import (
	"testing"
)

func TestRequestStringtoResponseString_basicIndex(t *testing.T) {
	WipeDataStore()
	response := RequestStringtoResponseString("INDEX|ls|\n")
	expectedResponse := "OK\n"
	if response != expectedResponse {
		t.Errorf("First Expected response: %s Received response: %s", expectedResponse, response)
	}

	response = RequestStringtoResponseString("INDEX|pwd|\n")
	expectedResponse = "OK\n"
	if response != expectedResponse {
		t.Errorf("Second Expected response: %s Received response: %s", expectedResponse, response)
	}

	response = RequestStringtoResponseString("INDEX|cat|ls,pwd\n")
	expectedResponse = "OK\n"
	if response != expectedResponse {
		t.Errorf("Third Expected response: %s Received response: %s", expectedResponse, response)
	}

	response = RequestStringtoResponseString("INDEX|taco|banana,pwd\n")
	expectedResponse = "FAIL\n"
	if response != expectedResponse {
		t.Errorf("Third Expected response: %s Received response: %s", expectedResponse, response)
	}
}

func TestRequestStringtoResponseString_removeIndex(t *testing.T) {
	WipeDataStore()
	response := RequestStringtoResponseString("REMOVE|ls|\n")
	expectedResponse := "OK\n"
	if response != expectedResponse {
		t.Errorf("First Expected response: %s Received response: %s", expectedResponse, response)
	}

	response = RequestStringtoResponseString("INDEX|ls|\n")
	expectedResponse = "OK\n"
	if response != expectedResponse {
		t.Errorf("Second Expected response: %s Received response: %s", expectedResponse, response)
	}

	response = RequestStringtoResponseString("REMOVE|ls|\n")
	expectedResponse = "OK\n"
	if response != expectedResponse {
		t.Errorf("Third Expected response: %s Received response: %s", expectedResponse, response)
	}
}

func TestRequestStringtoResponseString_tryToRemoveSomethingDependedOnIndex(t *testing.T) {
	WipeDataStore()

	response := RequestStringtoResponseString("INDEX|ls|\n")
	expectedResponse := "OK\n"
	if response != expectedResponse {
		t.Errorf("First Expected response: %s Received response: %s", expectedResponse, response)
	}

	response = RequestStringtoResponseString("INDEX|taco|ls\n")
	expectedResponse = "OK\n"
	if response != expectedResponse {
		t.Errorf("Second Expected response: %s Received response: %s", expectedResponse, response)
	}

	response = RequestStringtoResponseString("REMOVE|ls|\n")
	expectedResponse = "FAIL\n"
	if response != expectedResponse {
		t.Errorf("Third Expected response: %s Received response: %s", expectedResponse, response)
	}
}

func TestRequestStringtoResponseString_queryBasic(t *testing.T) {
	WipeDataStore()

	response := RequestStringtoResponseString("QUERY|ls|\n")
	expectedResponse := "FAIL\n"
	if response != expectedResponse {
		t.Errorf("First Expected response: %s Received response: %s", expectedResponse, response)
	}

	response = RequestStringtoResponseString("INDEX|ls|\n")
	expectedResponse = "OK\n"
	if response != expectedResponse {
		t.Errorf("Second Expected response: %s Received response: %s", expectedResponse, response)
	}

	response = RequestStringtoResponseString("QUERY|ls|\n")
	expectedResponse = "OK\n"
	if response != expectedResponse {
		t.Errorf("Third Expected response: %s Received response: %s", expectedResponse, response)
	}

}
