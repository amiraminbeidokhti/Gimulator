package object

import (
	"testing"
)

func TestKeyString(t *testing.T) {

	var tests = []struct {
		key  Key
		want string
	}{
		{Key{"Type", "Namespace", "Name"}, "{Type: Type, Namespace: Namespace, Name: Name}"},
		{Key{"", "", ""}, "{Type: , Namespace: , Name: }"},
		{Key{}, "{Type: , Namespace: , Name: }"},
		{Key{Type: "t"}, "{Type: t, Namespace: , Name: }"},
	}

	for _, test := range tests {

		got := test.key.String()

		if got != test.want {
			t.Errorf("got %s, want %s", got, test.want)
		}
	}
}

func TestEqual(t *testing.T) {

	var tests = []struct {
		key1 Key
		key2 *Key
		want bool
	}{
		{Key{"Type", "Namespace", "Name"}, &Key{"Type", "Namespace", "Name"}, true},
		{Key{"", "", ""}, &Key{}, true},
		{Key{Type: "t"}, &Key{Type: "t"}, true},
		{Key{"type", "namespace", "name"}, &Key{"Type", "Namespace", "Name"}, false},
		{Key{"tpe", "namespace", "name"}, &Key{"Type", "Namespace", "Name"}, false},
		{Key{"t", "ns", "name"}, &Key{"t", "Namespace", "Name"}, false},
		{Key{"t", "ns", "n"}, &Key{"t", "ns", "Name"}, false},
	}

	for _, test := range tests {

		got := test.key1.Equal(test.key2)

		if got != test.want {
			t.Errorf("got %v, want %v", got, test.want)
		}
	}
}

func TestMatch(t *testing.T) {

	var tests = []struct {
		key1 Key
		key2 *Key
		want bool
	}{
		{Key{"Type", "Namespace", "Name"}, &Key{"Type", "Namespace", "Name"}, true},
		{Key{"", "", ""}, &Key{}, true},
		{Key{Type: "t"}, &Key{Type: "t"}, true},
		{Key{"type", "namespace", "name"}, &Key{"Type", "Namespace", "Name"}, false},
		{Key{"tpe", "namespace", "name"}, &Key{"Type", "Namespace", "Name"}, false},
		{Key{Type: "t"}, &Key{"t", "a", "b"}, true},
		{Key{"Type", "Namespace", "Name"}, &Key{Type: "Type"}, false},
		{Key{}, &Key{}, true},
		{Key{}, &Key{"a", "b", "c"}, true},
		{Key{Type: "type"}, &Key{Type: "type", Namespace: "ns"}, true},
		{Key{"Type", "ns", "n"}, &Key{"Type", "ns", "name"}, false},
	}	

	for _, test := range tests {

		got := test.key1.Match(test.key2)

		if got != test.want {
			t.Errorf("got %v, want %v", got, test.want)
		}
	}
}

func TestObjectString(t *testing.T) {

	var tests = []struct {
		o    Object
		want string
	}{
		{Object{Key: &Key{Type: "Type"}, Value: "Value"}, "{Key: {Type: Type, Namespace: , Name: }, Value: 'Value'}"},
		{Object{Key: &Key{"t", "ns", "n"}, Value: "thisisatestvalue"}, "{Key: {Type: t, Namespace: ns, Name: n}, Value: 'thisisates...'}"},
		{Object{}, "{Key: <nil>, Value: ''}"},
		{Object{Key: &Key{Name: "n"}, Value: ""}, "{Key: {Type: , Namespace: , Name: n}, Value: ''}"},
	}

	for _, test := range tests {

		got := test.o.String()

		if got != test.want {
			t.Errorf("got %v, want %v", got, test.want)
		}
	}
}
