package main

import "testing"

func TestBadLength(t *testing.T) {
	s := "())"
	if isValid(s) {
		t.Errorf("The string %s is not valid.", s)
	}
}

func TestLongBadLength(t *testing.T) {
	s := "({{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}))"
	if isValid(s) {
		t.Errorf("The string %s is not valid.", s)
	}
}

func TestLength1(t *testing.T) {
	s := "("
	if isValid(s) {
		t.Errorf("The string %s is not valid.", s)
	}
}

func TestWrongStart(t *testing.T) {
	s := "]["
	if isValid(s) {
		t.Errorf("The string %s is not valid.", s)
	}
}

func TestBadOrder(t *testing.T) {
	s := "([)]"
	if isValid(s) {
		t.Errorf("The string %s is not valid.", s)
	}
}

func TestDuplicate(t *testing.T) {
	s := "(("
	if isValid(s) {
		t.Errorf("The string %s is not valid.", s)
	}
}

func TestLongDuplicate(t *testing.T) {
	s := "(((("
	if isValid(s) {
		t.Errorf("The string %s is not valid.", s)
	}
}

func TestDiffDuplicates(t *testing.T) {
	s := "[[(("
	if isValid(s) {
		t.Errorf("The string %s is not valid.", s)
	}
}

func TestGoodMix(t *testing.T) {
	s := "([{()}])"
	if !isValid(s) {
		t.Errorf("The string %s is valid.", s)
	}
}

func TestComplexGoodMix(t *testing.T) {
	s := "([{()[]({[]}[()])}])"
	if !isValid(s) {
		t.Errorf("The string %s is valid.", s)
	}
}

func TestSimple(t *testing.T) {
	s := "(){}[]"
	if !isValid(s) {
		t.Errorf("The string %s is valid.", s)
	}
}

func TestBaseCase(t *testing.T) {
	s := "{}"
	if !isValid(s) {
		t.Errorf("The string %s is valid.", s)
	}
}
