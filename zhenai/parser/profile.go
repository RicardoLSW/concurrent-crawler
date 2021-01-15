package parser

import (
	"regexp"

	"concurrent-crawler/engine"
)

var userInfoRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`)

func parseProfile(contents []byte, name string) engine.ParseResult {
	match := userInfoRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	var profile []interface{}
	profile = append(profile, name)
	for _, item := range match {
		profile = append(profile, item)
	}
	//profile.Marriage = string(match[0][1])
	//profile.Age = string(match[1][1])
	//profile.Xinzuo = string(match[2][1])
	//profile.Height = string(match[3][1])
	//profile.Income = string(match[5][1])
	//profile.Education = string(match[6][1])
	result.Items = append(result.Items, profile)
	return result
}
