package suites

import (
	"github.com/stretchr/testify/suite"
)

type BaseSuite struct {
	suite.Suite
}

func (bs *BaseSuite) ErrorContainsAll(err error, contains ...string) bool {
	result := bs.Error(err)
	if !result {
		return result
	}

	for _, contain := range contains {
		resultSub := bs.ErrorContains(err, contain, "should have '%s' contained in error message", contain)
		if !resultSub {
			result = false
		}
	}

	return result
}
