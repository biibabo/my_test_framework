// assertion/assertion.go

package assertion

import (
	"reflect"
	"testing"
)

// AssertJSONStringEquals 断言两个JSON字符串是否相等
func AssertJSONStringEquals(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("Expected JSON: %s, but got: %s", expected, actual)
	}
}
