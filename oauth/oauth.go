package oauth

import (
    "golang.org/x/oauth2"    
)

var Config = &oauth2.Config{
	ClientID:     "8d404f907fb04aeda0d36d236fac0828",
	ClientSecret: "asaodUU0CLffNucp14KV3ymIdWhL97DPyXqevvBi",
	Scopes:       []string{"publicData", "esi-skills.read_skills.v1", "esi-universe.read_structures.v1"},
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://login.eveonline.com/v2/oauth/authorize",
		TokenURL: "https://login.eveonline.com/v2/oauth/token",
	},
	RedirectURL: "https://arak1x9nlf.execute-api.us-west-1.amazonaws.com/dev/callback",
}

var AuthCodeURL string = Config.AuthCodeURL("state")