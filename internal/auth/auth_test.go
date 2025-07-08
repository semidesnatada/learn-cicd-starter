package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	testHeader := http.Header{
		"Authorization": []string{"ApiKey 1283e874874"},
	}

	expected := "1283e874874"
	out, err := GetAPIKey(testHeader)
	if err != nil {
		t.Fatalf("could not access API key for header: %v", testHeader)
	}

	if !reflect.DeepEqual(expected, out) {
		t.Fatalf("Error in test. Expected: %v, Got: %v.", expected, out)
	}

}
