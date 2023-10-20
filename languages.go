package tmdb

import (
	"fmt"
)

// Language struct
type Language struct {
	Iso6391     string
	EnglishName string
	Name        string
}

// GetLanguageList from configuration
// https://developers.themoviedb.org/3/configuration/languages
func (tmdb *TMDb) GetLanguageList() (*[]Language, error) {
	var languageList []Language
	uri := fmt.Sprintf("%s/configuration/languagest?api_key=%s", baseURL, tmdb.apiKey)
	result, err := getTmdb(uri, &languageList)
	return result.(*[]Language), err
}
