package main

import "testing"

func TestStubForTesting(t *testing.T){
  expected := "of the king"
  actual := StubForTesting()
  if actual != expected {
    t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
  }
}