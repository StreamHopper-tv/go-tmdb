package tmdb

import (
	"fmt"
)

// Language struct
type Language struct {
	Languages []struct {
		Iso6391     string `json:"iso_639_1"`
		EnglishName string `json:"english_name"`
		Name        string `json:"name"`
	}
}

// GetLanguageList from configuration
// https://developers.themoviedb.org/3/configuration/languages
func (tmdb *TMDb) GetLanguageList() (*Language, error) {
	var languageList Language
	uri := fmt.Sprintf("%s/configuration/languagest?api_key=%s", baseURL, tmdb.apiKey)
	fmt.Printf("%s\n", uri)

	result, err := getTmdb(uri, &languageList)
	fmt.Printf("%+v\n", result)
	return result.(*Language), err
}