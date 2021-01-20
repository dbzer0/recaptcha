package recaptcha

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

const recaptchaServerName = "https://www.google.com/recaptcha/api/siteverify"

type Recaptcha struct {
	privateKey string
}

func New(privateKey string) *Recaptcha {
	return &Recaptcha{privateKey: privateKey}
}

type recaptchaResponse struct {
	Success     bool      `json:"success"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

func (rc *Recaptcha) check(remoteIP, response string) (*recaptchaResponse, error) {
	resp, err := http.PostForm(recaptchaServerName, url.Values{
		"secret":   {rc.privateKey},
		"remoteip": {remoteIP},
		"response": {response}},
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var recapResp *recaptchaResponse
	if err := json.NewDecoder(resp.Body).Decode(&recapResp); err != nil {
		return recapResp, nil
	}

	return recapResp, nil
}

func (rc *Recaptcha) Confirm(remoteIP, response string) (bool, error) {
	resp, err := rc.check(remoteIP, response)
	if err != nil {
		return false, err
	}

	return resp.Success, nil
}
