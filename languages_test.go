package tmdb

import (
	. "gopkg.in/check.v1"
)

func (s *TmdbSuite) TestGetLanguageList(c *C) {
	languageResult, err := s.tmdb.GetLanguageList()
	s.baseTest(&languageResult, err, c)
	firstLanguage := languageResult[0]
	c.Assert(firstLanguage, NotNil)
}
