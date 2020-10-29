package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnectionString(t *testing.T) {
	testCases := []struct {
		User     string
		Pass     string
		Host     string
		Port     string
		Db       string
		Expected string
	}{
		{
			User:     "test",
			Pass:     "password",
			Host:     "host",
			Port:     "1001",
			Db:       "database",
			Expected: "test:password@tcp(host:1001)/database?parseTime=true",
		},
		{
			User:     "test",
			Pass:     "password@password",
			Host:     "host",
			Port:     "1002",
			Db:       "database_test",
			Expected: "test:password@password@tcp(host:1002)/database_test?parseTime=true",
		},
	}

	for _, testCase := range testCases {
		output := ConnectionString(
			testCase.User, testCase.Pass, testCase.Host, testCase.Port, testCase.Db)

		assert.Equal(t, output, testCase.Expected)
	}
}
