# mappy

[![Build Status](https://travis-ci.org/slomek/mappy.svg?branch=master)](https://travis-ci.org/slomek/mappy)
[![Documentation](https://godoc.org/github.com/slomek/mappy?status.svg)](https://godoc.org/github.com/slomek/mappy?status.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/slomek/mappy)](https://goreportcard.com/report/github.com/slomek/mappy)

The purpose of this tiny library is to easily parse `string`->`string` maps to typed structures in Go. I've faced a real problem where I needed to define a list of keys that may appear in the map, then have each key associated with `struct`'s field and they had to be manually transformed. `mappy` tackles the problem using _struct tags_, where each field is represented with the key that is to be found in the map.

## Usage

In order to transform a struct instance into a map, use `Marshal(..)` function:

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

In order to transform a map into a struct, use `Unmarshal(..)` function:

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

## Performance

The idea behind `mappy` is to enhance readability of the code, not to improve its performance. If you value the speed of execution above anything else, this will not be a good choice for you. Manual mapping is much faster for both marshaling:

    BenchmarkMarshal-8               3000000               575 ns/op
    BenchmarkStructToMap-8           5000000               240 ns/op

and unmarshaling:

    BenchmarkUnarshal-8              5000000               368 ns/op
    BenchmarkMapToStruct-8         100000000              13.3 ns/op

If, however, you find such solution good enough and useful, enjoy!
