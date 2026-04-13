package testing

import (
	"fmt"

	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
)

// GenericMatcher is a helper function to create a GomegaMatcher for a specific
// type T.  It takes a check function that performs the actual matching logic,
// and two functions that generate the failure messages for both positive and
// negative cases.
func GenericMatcher[T any](
	checkFunc func(actualValue T) bool,
	positiveFailureMessageFunc func(actualValue T) string,
	negativeFailureMessageFunc func(actualValue T) string,
) types.GomegaMatcher {
	return &concreteMatcher[T]{
		checkFunc:                  checkFunc,
		positiveFailureMessageFunc: positiveFailureMessageFunc,
		negativeFailureMessageFunc: negativeFailureMessageFunc,
	}
}

type concreteMatcher[T any] struct {
	checkFunc                  func(T) bool
	positiveFailureMessageFunc func(T) string
	negativeFailureMessageFunc func(T) string
}

var _ types.GomegaMatcher = (*concreteMatcher[any])(nil)

func (m *concreteMatcher[T]) Match(actualValue any) (bool, error) {
	castActualValue, ok := actualValue.(T)
	if !ok {
		var zeroValue T // used just to get the type for error message formatting
		return false, fmt.Errorf("Matcher expected actual of type <%T>.  Got:\n%s", zeroValue, format.Object(actualValue, 1))
	}
	return m.checkFunc(castActualValue), nil
}

func (m *concreteMatcher[T]) FailureMessage(actualValue any) string {
	castActualValue, ok := actualValue.(T)
	if !ok {
		panic("method called on incorrect type")
	}
	return m.positiveFailureMessageFunc(castActualValue)
}

func (m *concreteMatcher[T]) NegatedFailureMessage(actualValue any) string {
	castActualValue, ok := actualValue.(T)
	if !ok {
		panic("method called on incorrect type")
	}
	return m.negativeFailureMessageFunc(castActualValue)
}
