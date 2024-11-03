package helper

import (
	"testing"
	"github.com/stretchr/testify/suite"
	"github.com/stretchr/testify/assert"
)

type ExempleTestSuite struct {
	suite.Suite
	VariableHelloWorld string
}

func (suite *ExempleTestSuite) SetupTest() {
	suite.VariableHelloWorld = "hello world"
}

func (suite *ExempleTestSuite) TestExample() {
	assert.Equal(suite.T(), "Hello Raihan", suite.VariableHelloWorld, "result must be hello world")
}

func (suite *ExempleTestSuite) TestExampleMalay() {
	assert.Equal(suite.T(), "hello world", suite.VariableHelloWorld)
}

func TestHelloWorldSuite(t *testing.T) {
	suite.Run(t, new(ExempleTestSuite))
}

