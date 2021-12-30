package mediator

import (
	"reflect"
	"testing"
)

func Test_Mediator(t *testing.T) {
	expected := &Three{}
	actual := Sum(One{}, Two{})
	t.Log(getType(actual))
	if getType(expected) != getType(actual) {
		t.Errorf("Sum not working. expected: %v\nactual: %v\n", getType(expected), getType(actual))
	}

	expectedN := 3
	actualN := Sum(1, 2)
	t.Log(actualN)
	if expectedN != actualN {
		t.Errorf("Sum not working. expected: %v\nactual: %v\n", expectedN, actualN)
	}

}

func getType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}
