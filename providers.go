package tmdb

import (
	"fmt"
)

// Region struct
type RegionResults struct {
	Regions []Region `json:"results"`
}
type Region struct {
	Iso31661    string `json:"iso_3166_1"`
	EnglishName string `json:"english_name"`
	Name        string `json:"native_name"`
}
type ProviderResults struct {
	Providers []Provider `json:"results"`
}
type Provider struct {
	DisplayPriorities map[string]string `json:"display_priorities"`
	LogoPath          string            `json:"logo_path"`
	ProviderName      string            `json:"provider_name"`
	ID                int               `json:"provider_id"`
}

// GetRegionList from watch providers
// https://developers.themoviedb.org/3/watch/providers/regions?language=en-US
func (tmdb *TMDb) GetRegionList() (*[]Region, error) {
	var regionList RegionResults
	uri := fmt.Sprintf("%s/watch/providers/regions?language=en-US&api_key=%s", baseURL, tmdb.apiKey)
	result, err := getTmdb(uri, &regionList)
	return &result.(*RegionResults).Regions, err
}

func (tmdb *TMDb) GetTVProviders() (*[]Provider, error) {
	var providerList ProviderResults
	uri := fmt.Sprintf("%s/watch/providers/tv?language=en-US&watch_region=US&api_key=%s", baseURL, tmdb.apiKey)
	result, err := getTmdb(uri, &providerList)
	return &result.(*ProviderResults).Providers, err
}

func (tmdb *TMDb) GetMovieProviders() (*[]Provider, error) {
	var providerList ProviderResults
	uri := fmt.Sprintf("%s/watch/providers/movie?language=en-US&watch_region=US&api_key=%s", baseURL, tmdb.apiKey)
	result, err := getTmdb(uri, &providerList)
	return &result.(*ProviderResults).Providers, err
}
