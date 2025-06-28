package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is johnny"}
	diff2 := &String{Value: "My name is johnny"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with the same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with the same content have different hash keys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with different content have the same hash keys")
	}
}

func TestIntegerHashKey(t *testing.T) {
	n1 := &Integer{Value: 5}
	n2 := &Integer{Value: 5}
	diff1 := &Integer{Value: 1}
	diff2 := &Integer{Value: 1}

	if n1.HashKey() != n2.HashKey() {
		t.Errorf("integers with the same value have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("integers with the same value have different hash keys")
	}

	if n1.HashKey() == diff1.HashKey() {
		t.Errorf("integers with different values have the same hash keys")
	}
}

func TestBooleanHashKey(t *testing.T) {
	b1 := &Boolean{Value: true}
	b2 := &Boolean{Value: true}
	diff1 := &Boolean{Value: false}
	diff2 := &Boolean{Value: false}

	if b1.HashKey() != b2.HashKey() {
		t.Errorf("booleans with the same value have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("booleans with the same value have different hash keys")
	}

	if b1.HashKey() == diff1.HashKey() {
		t.Errorf("booleans with different values have the same hash keys")
	}
}
