package regex

import (
	"fmt"
	"regexp"
)

var b bool
var KeySecret = map[string]string{
	"BasicAuthCredentials":      `"(?://)[a-zA-Z0-9]+:[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z]+"`,
	"FacebookOauth":             `"[f|F][a|A][c|C][e|E][b|B][o|O][o|O][k|K].*[)'|"][0-9a-f]{32}['|"]"`,
	"FacebookSecretKey":         `"(?i)(facebook|fb)(.{0,20})?(?-i)['"][0-9a-f]{32}"`,
	"GoogleAPIKey":              `AIza[0-9A-Za-z\\-_]{35}`,
	"ArtifactoryAPIToken":       `(?:\s|=|:|"|^)AKC[a-zA-Z0-9]{10,}`,
	"AuthorizationBasic":        `basic [a-zA-Z0-9_\-\\.:=]+`,
	"ArtifactoryPassword":       `(?:\s|=|:|"|^)AP[\dABCDEF][a-zA-Z0-9]{8,}`,
	"AuthorizationBearer":       `bearer [a-zA-Z0-9_\-.=]+`,
	"CloudinaryBasicAuth":       `"cloudinary://[0-9]{15}:[0-9A-Za-z]+@[a-z]+"`,
	"FacebookAccessToken":       `"EAACEdEose0cBA[0-9A-Za-z]+"`,
	"FacebookAPIKey":            `"(?i)(facebook|fb)(.{0,20})?['"][0-9]{13,17}"`,
	"AWSClientID":               `(A3T[A-Z0-9]|AKIA|AGPA|AIDA|AROA|AIPA|ANPA|ANVA|ASIA)[A-Z0-9]{16}`,
	"AWSMWSKey":                 `amzn\.mws\.[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`,
	"AWSSecretKey":              `(?i)aws(?:.{0,20})?["'][0-9a-zA-Z/+]{40}["']`,
	"Github":                    `"(?i)github(.{0,20})?(?-i)['"][0-9a-zA-Z]{35,40}"`,
	"GoogleCloudPlatformAPIKey": `"(?i)(google|gcp|youtube|drive|yt)(.{0,20})?['\"][AIza[)0-9a-z\\-_]{35}]['"]"`,
	"GoogleOauth":               `"[0-9]+-[0-9A-Za-z_]{32}\.apps\.googleusercontent\.com"`,
	"HerokuAPIKey":              `"[h|H][e|E][r|R][o|O][k|K][u|U].{0,30}[0-9A-F)]{8}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{12}"`,
	"LinkedInSecretKey":         `"(?i)linkedin(.{0,20})?['"][0-9a-z]{16}['"]"`,
	"MailchampAPIKey":           `"[0-9a-f]{32}-us[0-9]{1,2}"`,
	"MailgunAPIKey":             `"key-[0-9a-zA-Z]{32}"`,
	"PicaticAPIKey":             `"sk_live_[0-9a-z]{32}"`,
	"SlackToken":                `"xox[baprs]-([0-9a-zA-Z]{10,48})?"`,
	"SlackWebhook":              `"https://hooks.slack.com/services/T[)a-zA-Z0-9_]{8}/B[a-zA-Z0-9_]{8}/[a-zA-Z0-9_]{24}"`,
	"StripeAPIKey":              `"[rs]k_live_[0-9a-zA-Z]{24}"`,
	"Firebase":                  `"AAAA[A-Za-z0-9_-]{7}:[A-Za-z0-9_-]{140}"`,
	"Google_captcha":            `"6L[0-9A-Za-z-_]{38}|^6[0-9a-zA-Z_-]{39}$"`,
	"Google_oauth":              `"ya29\.[0-9A-Za-z\-_]+"`,
	"Amazon_aws_access_key_id":  `"A[SK]IA[0-9A-Z]{16}"`,
	"Amazon_aws_url":            `"s3\.amazonaws.com[]+|[a-zA-Z0-9_-]*\.s3\.amazonaws.com"`,
	"Amazon_mws_auth_toke":      `"amzn\\.mws\\.[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`,
	"SquareAccessToken":         `"sqOatp-[0-9A-Za-z\\-_]{22}"`,
	"SquareOauthSecret":         `"sq0csp-[ 0-9A-Za-z\\-_]{43}"`,
	"TwilioAPIKey":              `"SK[0-9a-fA-F]{32}"`,
	"TwitterOauth":              `"[t|T][w|W][i|I][t|T][t|T][e|E][r|R].{0,30}[')"\\s][0-9a-zA-Z]{35,44}['"\\s]"`,
	"TwitterSecretKey":          `"(?i)twitter(.{0,20})?['"][0-9a-z]{35,44}"`,
}

func Regex(response string) {
	b = false
	for key, pattern := range KeySecret {
		match := regexp.MustCompile(pattern)
		data := match.FindAllString(response, -1)
		if len(data) > 0 {
			b = true
			fmt.Println(key+": ", data)

		}

	}

	if b != true {
		fmt.Println("No match found.")
	}
}
