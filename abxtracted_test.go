package abxtracted_test

import (
	"testing"

	abxtracted "github.com/abxtracted/go-abxtracted"
	"github.com/stretchr/testify/assert"
)

const (
	testProject    = "bc378dd9-1da5-4365-8f6b-1071587d1ab5"
	testExperiment = "integration_tests"
)

func TestNewClient(t *testing.T) {
	var assert = assert.New(t)
	var abx, err = abxtracted.New()
	assert.NoError(err)
	assert.NotNil(abx)
}

func TestGetExperiment(t *testing.T) {
	var assert = assert.New(t)
	var abx, _ = abxtracted.New()
	var customer = "TestGetExperiment"
	exp, err := abx.Get(testProject, customer, testExperiment)
	assert.NoError(err)
	assert.Equal(testExperiment, exp.Experiment)
	assert.Equal(customer, exp.CustomerIdentity)
}

func TestGetInvalidExperiment(t *testing.T) {
	var assert = assert.New(t)
	var abx, _ = abxtracted.New()
	var customer = "TestGetInvalidExperiment"
	exp, err := abx.Get(testProject, customer, "invalid-exp")
	assert.Error(err)
	assert.Equal(abxtracted.Experiment{}, exp)
}

func TestCompleteExperiment(t *testing.T) {
	var assert = assert.New(t)
	var abx, _ = abxtracted.New()
	var customer = "TestCompleteExperiment"
	exp, err := abx.Get(testProject, customer, testExperiment)
	assert.NoError(err)
	assert.NoError(abx.Complete(testProject, exp))
}

func TestCompletedExperiment(t *testing.T) {
	var assert = assert.New(t)
	const experiment = "complete_integration_test"
	var abx, _ = abxtracted.New()
	var customer = "TestCompletedExperiment"
	exp, err := abx.Get(testProject, customer, experiment)
	assert.NoError(err)
	assert.NoError(abx.Complete(testProject, exp))
	assert.Equal(experiment, exp.Experiment)
	assert.Equal(customer, exp.CustomerIdentity)
	assert.Equal("variation", exp.Scenario)
}
