//go:build integration
// +build integration

// run integration tests command
// go test ./test -v -tags=integration
package test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
