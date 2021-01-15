package parser

import (
	"concurrent-crawler/model"
	"regexp"

	"concurrent-crawler/engine"
)

var userInfoRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`)

func parseProfile(contents []byte, name string) engine.ParseResult {
	match := userInfoRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	profile := model.Profile{}
	profile.Name = name
	for i, user := range match {
		switch i {
		case 0:
			profile.Marriage = string(user[1])
			break
		case 1:
			profile.Age = string(user[1])
			break
		case 2:
			profile.Xinzuo = string(user[1])
			break
		case 3:
			profile.Height = string(user[1])
			break
		case 4:
			break
		case 5:
			profile.Income = string(user[1])
			break
		case 6:
			profile.Education = string(user[1])
			break

		}
	}
	result.Items = append(result.Items, profile)
	return result
}
