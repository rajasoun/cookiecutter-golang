package testutils_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/testutils"
)

func TestIsRunningAsGoTest(t *testing.T) {
	t.Parallel()
	// Set up assertions using the testify/assert package.
	assert := assert.New(t)

	// Define test cases.
	testCases := []struct {
		name    string
		args    []string
		wantRes bool
	}{
		{
			name:    "test flag present",
			args:    []string{"-test.run", "TestIsRunningAsGoTest"},
			wantRes: true,
		},
		{
			name:    "test flag not present",
			args:    []string{"foo", "bar"},
			wantRes: false,
		},
		{
			name:    "test flag  present with verbose flag",
			args:    []string{"-test.v", "TestIsRunningAsGoTest"},
			wantRes: true,
		},
	}

	// Test each test case.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Replace os.Args with the test case's arguments.
			origArgs := os.Args
			defer func() { os.Args = origArgs }()
			os.Args = tc.args
			// Call the function and check the result using assertions.
			assert.Equal(tc.wantRes, testutils.IsRunningAsGoTest())
		})
	}
}
