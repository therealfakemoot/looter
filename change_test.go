package genloot

import (
	"reflect"
	"testing"
)

func Test_ReduceCoins(t *testing.T) {
	wallet := CashValue{CP: 300, SP: 150, EP: 3, GP: 7, PP: 2}

	expected := CashValue{
		PP: 4,
		GP: 6,
		EP: 1,
	}

	actual := reduceCoins(wallet)

	if !reflect.DeepEqual(expected, actual) {
		t.Logf("actual: %#+v\n", actual)
		t.Log("you made change incorrectly")
		t.Fail()

	}
}
