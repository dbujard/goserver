package main

import "testing"

func TestStubForTesting(t *testing.T){
  expected := "of the buildnight"
  actual := StubForTesting()
  if actual != expected {
    t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
  }
}