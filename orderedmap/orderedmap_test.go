package orderedmap_test

import (
	"github.com/opslevel/opslevel-jq-parser/v2024/orderedmap"
	"reflect"
	"strings"
	"testing"
)

func TestOrderedMap(t *testing.T) {
	type TestOperation struct {
		params   string
		expected any
	}
	operations := []TestOperation{
		{"contains foo", false},
		{"keys", []string{}},
		{"len", 0},
		{"add foo bar", true},
		{"add foo bar2", false},
		{"contains foo", true},
		{"keys", []string{"bar"}},
		{"add bar foo", true},
		{"add hello world", true},
		{"contains world", false},
		{"keys", []string{"bar", "foo", "world"}},
	}

	oMap := orderedmap.New[string]()
	for i, op := range operations {
		split := strings.Split(op.params, " ")
		switch split[0] {
		case "contains":
			if oMap.Contains(split[1]) != op.expected {
				t.Errorf("op %d '%s' - Contains() is incorrect (expected %v)", i, op, op.expected)
			}
		case "keys":
			if !reflect.DeepEqual(oMap.Values(), op.expected) {
				t.Errorf("op %d '%s' - Values() is incorrect (expected %v)", i, op, op.expected)
			}
		case "len":
			if oMap.Len() != op.expected {
				t.Errorf("op %d '%s' - Len() is incorrect (expected %v)", i, op, op.expected)
			}
		case "add":
			if oMap.Add(split[1], split[2]) != op.expected {
				t.Errorf("op %d '%s' - Add() is incorrect (expected %v)", i, op, op.expected)
			}
		default:
			t.Errorf("op %d '%s' - unknown operation (broken test case)", i, op)
		}
	}
}
