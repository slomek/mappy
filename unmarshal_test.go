package mappy

import (
	"reflect"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	type Some struct {
		Name string `map:"name"`
	}
	in := map[string]string{
		"name": "slomek",
	}
	out := Some{Name: "slomek"}

	var s Some
	if err := Unmarshal(in, &s); err != nil {
		t.Fatalf("Failed to unmarshal struct to map: %v", err)
	}

	if !reflect.DeepEqual(s, out) {
		t.Errorf("Expected true, got false")
	}
}

func TestUnmarshalMultiField(t *testing.T) {
	type Person struct {
		FirstName string `map:"first_name"`
		LastName  string `map:"last_name"`
	}
	in := map[string]string{
		"first_name": "Tim",
		"last_name":  "Duncan",
	}
	out := Person{FirstName: "Tim", LastName: "Duncan"}

	var s Person
	if err := Unmarshal(in, &s); err != nil {
		t.Fatalf("Failed to unmarshal struct to map: %v", err)
	}

	if !reflect.DeepEqual(s, out) {
		t.Errorf("Expected true, got false")
	}
}
func TestUnmarshalTaggedFieldsOnly(t *testing.T) {
	type Person struct {
		Username string `map:"username"`
		Password string
	}
	in := map[string]string{
		"username": "slomek",
		"password": "qwerty",
	}
	out := Person{Username: "slomek"}

	var s Person
	if err := Unmarshal(in, &s); err != nil {
		t.Fatalf("Failed to unmarshal struct to map: %v", err)
	}

	if !reflect.DeepEqual(s, out) {
		t.Errorf("Expected true, got false")
	}
}
