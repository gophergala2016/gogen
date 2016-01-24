package gogen

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ResourceSuite struct {
	suite.Suite
}

func (s *ResourceSuite) SetupTest() {}


func (s *ResourceSuite) TestResourceAdd() {
	rc := ResourceContainer{}
	rc.Add(42)
	assert.Equal(s.T(), 42, rc[0])
	rc.Add(1.25)
	rc.Add("wololo")
	assert.Equal(s.T(), "wololo", rc[2])
}


func TestResourceSuite(t *testing.T) {
	suite.Run(t, &ResourceSuite{})
}


