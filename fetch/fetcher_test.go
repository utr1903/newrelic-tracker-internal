package internal

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type graphqlClientMock struct {
	failRequest bool
}

func (c *graphqlClientMock) Execute(
	queryVariables any,
	result any,
) error {
	if c.failRequest {
		return errors.New("error")
	}

	// Create mock response to convert into bytes
	responseMock := map[string]string{
		"test": "test",
	}
	bytes, err := json.Marshal(responseMock)
	if err != nil {
		panic(err)
	}

	// Put the responseMock into result
	err = json.Unmarshal(bytes, result)
	if err != nil {
		panic(err)
	}

	return nil
}

func Test_GraphQlRequestFails(t *testing.T) {
	gqlc := &graphqlClientMock{
		failRequest: true,
	}

	res := map[string]string{}
	err := Fetch(gqlc, "qv", &res)

	assert.NotNil(t, err)
}

func Test_GraphQlRequestSucceeds(t *testing.T) {
	gqlc := &graphqlClientMock{
		failRequest: false,
	}

	res := map[string]string{}
	err := Fetch(gqlc, "qv", &res)
	assert.NotNil(t, res)
	assert.Nil(t, err)
}
