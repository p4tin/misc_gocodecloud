package mockery_testify

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"misc/mockery-testify/mocks"
)

func TestBar(t *testing.T) {
	mockThing := Thing{}

	mockDAL := &mocks.DataAccessLayer{}
	mockDAL.On("Insert", "foo", mockThing).Return(nil)

	actual := Bar(mockDAL)

	mockDAL.AssertExpectations(t)

	assert.Equal(t, mockThing, actual, "should return a Thing")
}
