package engine

import (
	"learningnotes/reptile/crawler/fetcher"
)

func Worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(r.Url)
	//log.Warn("Fetching %s", r.Url)
	if err != nil {
		return ParseResult{}, err
	}
	return r.Parse.Parse(body, r.Url), nil
}
