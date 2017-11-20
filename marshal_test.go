package mappy

import (
	"fmt"
	"testing"
)

func TestMarshal(t *testing.T) {
	testCases := []struct {
		desc   string
		in     interface{}
		outMap map[string]string
	}{
		{
			desc: "marshal struct",
			in: struct {
				Name string `map:"name"`
			}{
				Name: "slomek",
			},
			outMap: map[string]string{
				"name": "slomek",
			},
		}, {
			desc: "marshal multiple-field struct",
			in: struct {
				FirstName string `map:"first_name"`
				LastName  string `map:"last_name"`
			}{
				FirstName: "Tim",
				LastName:  "Duncan",
			},
			outMap: map[string]string{
				"first_name": "Tim",
				"last_name":  "Duncan",
			},
		}, {
			desc: "marshal fields with tags only",
			in: struct {
				Username string `map:"username"`
				Password string
			}{
				Username: "slomek",
			},
			outMap: map[string]string{
				"username": "slomek",
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			outMap, err := Marshal(tC.in)
			if err != nil {
				t.Fatalf("Failed to marshal struct to map: %v", err)
			}

			for key, expVal := range tC.outMap {
				val, ok := outMap[key]
				if !ok {
					t.Errorf("Expected key '%s' in the map, but not found", key)
				} else if val != expVal {
					t.Errorf("Expected key '%s' value to be '%s', got: '%s'", key, expVal, val)
				}
			}

			for key, val := range outMap {
				expVal, ok := tC.outMap[key]
				if !ok {
					t.Errorf("Unexpected key '%s' in the map", key)
				} else if val != expVal {
					t.Errorf("Expected key '%s' value to be '%s', got: '%s'", key, expVal, val)
				}
			}
		})
	}
}

func ExampleMarshal() {
	type Person struct {
		FirstName string `map:"first_name"`
		LastName  string `map:"last_name"`
	}

	p := Person{FirstName: "Tim", LastName: "Duncan"}

	pMap, _ := Marshal(p)

	fmt.Println(pMap["first_name"])
	fmt.Println(pMap["last_name"])

	// Output:
	// Tim
	// Duncan
}

func BenchmarkMarshal(b *testing.B) {
	type Person struct {
		FirstName string `map:"first_name"`
		LastName  string `map:"last_name"`
	}

	p := Person{FirstName: "Tim", LastName: "Duncan"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pMap, _ := Marshal(p)
		if len(pMap) != 2 {
			b.Fail()
		}
	}
}

func BenchmarkStructToMap(b *testing.B) {
	type Person struct {
		FirstName string
		LastName  string
	}
	const (
		keyFirstName = "first_name"
		keyLastName  = "last_name"
	)

	p := Person{FirstName: "Tim", LastName: "Duncan"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pMap := map[string]string{
			keyFirstName: p.FirstName,
			keyLastName:  p.LastName,
		}
		if len(pMap) != 2 {
			b.Fail()
		}
	}
}
