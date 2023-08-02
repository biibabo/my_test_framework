// runner/runner.go

package runner

import (
	"my_test_framework/test"
	"testing"
)

// RunAllTests 执行所有测试用例
func RunAllTests() {
	testing.Main(testing.Verbose(), test.TestUserManagement)
}
