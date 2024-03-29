package mappy

import (
	"errors"
	"fmt"
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

func TestUnmarshalError(t *testing.T) {
	type Person struct {
		Name string `map:"name"`
		Age  int    `map:"age"`
	}
	in := map[string]string{
		"name": "slomek",
		"age":  "99",
	}

	var s Person
	if err := Unmarshal(in, &s); err != nil {
		if errors.Unwrap(err) != ErrMapUnmarshal {
			t.Fatalf("Expected error returned to be %v, got: %v", ErrMapUnmarshal, err)
		}
		return
	}
	t.Fatalf("Expected error, not returned though")
}

func TestUnmarshalEmptyMap(t *testing.T) {
	type Person struct {
		Name string `map:"name"`
		Age  int    `map:"age"`
	}

	testCases := []struct {
		desc string
		in   map[string]string
	}{
		{
			desc: "nil map",
			in:   nil,
		}, {
			desc: "nil map",
			in:   map[string]string{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var s Person
			if err := Unmarshal(tC.in, &s); err != nil {
				t.Fatalf("Failed to unmarshal struct to map: %v", err)
			}
		})
	}
}

func ExampleUnmarshal() {
	type Person struct {
		FirstName string `map:"first_name"`
		LastName  string `map:"last_name"`
	}

	pMap := map[string]string{
		"first_name": "Shaquille",
		"last_name":  "O'Neal",
	}

	var p Person
	Unmarshal(pMap, &p)

	fmt.Println(p.FirstName)
	fmt.Println(p.LastName)

	// Output:
	// Shaquille
	// O'Neal
}

func BenchmarkUnarshal(b *testing.B) {
	type Person struct {
		FirstName string `map:"first_name"`
		LastName  string `map:"last_name"`
	}

	pMap := map[string]string{
		"first_name": "Shaquille",
		"last_name":  "O'Neal",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var p Person
		Unmarshal(pMap, &p)
		if p.FirstName == "" {
			b.Fail()
		}
	}
}

func BenchmarkMapToStruct(b *testing.B) {
	type Person struct {
		FirstName string `map:"first_name"`
		LastName  string `map:"last_name"`
	}
	const (
		keyFirstName = "first_name"
		keyLastName  = "last_name"
	)

	pMap := map[string]string{
		keyFirstName: "Shaquille",
		keyLastName:  "O'Neal",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p := Person{
			FirstName: pMap[keyFirstName],
			LastName:  pMap[keyLastName],
		}
		if p.FirstName == "" {
			b.Fail()
		}
	}
}
