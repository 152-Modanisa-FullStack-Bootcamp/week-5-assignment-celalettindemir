package assignment

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUint32(t *testing.T) {
	/*
		Sum uint32 numbers, return uint32 sum value and boolean overflow flag
		cases need to pass:
			math.MaxUint32, 1 => 0, true
			1, 1 => 2, false
			42, 2701 => 2743, false
			42, math.MaxUint32 => 41, true
			4294967290, 5 => 4294967295, false
			4294967290, 6 => 0, true
			4294967290, 10 => 4, true
	*/
	type input struct {
		x, y uint32
	}
	type result struct {
		sum      uint32
		overflow bool
	}
	cases := []struct {
		input
		result
	}{
		{input{math.MaxUint32, 1}, result{0, true}}, // other cases
		{input{1, 1}, result{2, false}},
		{input{42, 2701}, result{2743, false}},
		{input{4294967290, 5}, result{4294967295, false}},
		{input{4294967290, 6}, result{0, true}},
		{input{4294967290, 10}, result{4, true}},
	}

	for _, c := range cases {
		sum, overflow := AddUint32(c.x, c.y)

		assert.Equal(t, c.sum, sum)
		assert.Equal(t, c.overflow, overflow)
	}

}

func TestCeilNumber(t *testing.T) {
	/*
		Ceil the number within 0.25
		cases need to pass:
			42.42 => 42.50
			42 => 42
			42.01 => 42.25
			42.24 => 42.25
			42.25 => 42.25
			42.26 => 42.50
			42.55 => 42.75
			42.75 => 42.75
			42.76 => 43
			42.99 => 43
			43.13 => 43.25
	*/
	cases := []struct {
		input, result float64
	}{
		{42.42, 42.50}, // other cases
		{42, 42},
		{42.01, 42.25},
		{42.24, 42.25},
		{42.25, 42.25},
		{42.26, 42.50},
		{42.55, 42.75},
		{42.75, 42.75},
		{42.76, 43},
		{42.99, 43},
		{43.13, 43.25},
	}

	for _, c := range cases {
		point := CeilNumber(c.input)

		assert.Equal(t, c.result, point)
	}

}

func TestAlphabetSoup(t *testing.T) {
	/*
		String with the letters in alphabetical order.
		cases need to pass:
		 	"hello" => "ehllo"
			"" => ""
			"h" => "h"
			"ab" => "ab"
			"ba" => "ab"
			"bac" => "abc"
			"cba" => "abc"
	*/
	cases := []struct {
		input, result string
	}{
		{"hello", "ehllo"}, // other cases
		{"", ""},
		{"h", "h"},
		{"ab", "ab"},
		{"ba", "ab"},
		{"bac", "abc"},
		{"cba", "abc"},
	}

	for _, c := range cases {
		result := AlphabetSoup(c.input)
		assert.Equal(t, c.result, result)
	}

}

func TestStringMask(t *testing.T) {
	/*
		Replace after n(uint) character of string with '*' character.
		cases need to pass:
			"!mysecret*", 2 => "!m********"
			"", n(any positive number) => "*"
			"a", 1 => "*"
			"string", 0 => "******"
			"string", 3 => "str***"
			"string", 5 => "strin*"
			"string", 6 => "******"
			"string", 7(bigger than len of "string") => "******"
			"s*r*n*", 3 => "s*r***"
	*/

	type input struct {
		str string
		n   int32
	}
	cases := []struct {
		input
		result string
	}{
		{input{"!mysecret*", 2}, "!m********"}, // other cases
		{input{"", 5}, "*"},
		{input{"a", 1}, "*"},
		{input{"string", 0}, "******"},
		{input{"string", 3}, "str***"},
		{input{"string", 5}, "strin*"},
		{input{"string", 6}, "******"},
		{input{"string", 10}, "******"},
		{input{"s*r*n*", 3}, "s*r***"},
	}

	for _, c := range cases {
		result := StringMask(c.str, uint(c.n))
		assert.Equal(t, c.result, result)
	}

}

func TestWordSplit(t *testing.T) {
	words := "apple,bat,cat,goodbye,hello,yellow,why"
	/*
		Your goal is to determine if the first element in the array can be split into two words,
		where both words exist in the dictionary(words variable) that is provided in the second element of array.

		cases need to pass:
			[2]string{"hellocat",words} => hello,cat
			[2]string{"catbat",words} => cat,bat
			[2]string{"yellowapple",words} => yellow,apple
			[2]string{"",words} => not possible
			[2]string{"notcat",words} => not possible
			[2]string{"bootcamprocks!",words} => not possible
	*/
	cases := []struct {
		input  [2]string
		result string
	}{
		{[2]string{"hellocat", words}, "hello,cat"}, // other cases
		{[2]string{"catbat", words}, "cat,bat"},
		{[2]string{"yellowapple", words}, "yellow,apple"},
		{[2]string{"", words}, "not possible"},
		{[2]string{"notcat", words}, "not possible"},
		{[2]string{"bootcamprocks!", words}, "not possible"},
	}

	for _, c := range cases {
		result := WordSplit(c.input)
		assert.Equal(t, c.result, result)
	}
}

func TestVariadicSet(t *testing.T) {
	/*
		FINAL BOSS ALERT :)
		Tip: Learn and apply golang variadic functions(search engine -> "golang variadic function" -> WOW You can really dance! )

		Convert inputs to set(no duplicate element)
		cases need to pass:
			4,2,5,4,2,4 => []interface{4,2,5}
			"bootcamp","rocks!","really","rocks!" => []interface{"bootcamp","rocks!","really"}
			1,uint32(1),"first",2,uint32(2),"second",1,uint32(2),"first" => []interface{1,uint32(1),"first",2,uint32(2),"second"}
	*/

	cases := []struct {
		input  []interface{}
		result []interface{}
	}{
		{[]interface{}{4, 2, 5, 4, 2, 4}, []interface{}{4, 2, 5}}, // other cases
		{[]interface{}{"bootcamp", "rocks!", "really", "rocks!"}, []interface{}{"bootcamp", "rocks!", "really"}},
		{[]interface{}{1, uint32(1), "first", 2, uint32(2), "second", 1, uint32(2), "first"}, []interface{}{1, uint32(1), "first", 2, uint32(2), "second"}},
	}

	for _, c := range cases {
		set := VariadicSet(c.input...)
		assert.Equal(t, c.result, set)
	}
}
