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
	serverName string
}

func New(privateKey string, opts ...Option) *Recaptcha {
	rc := &Recaptcha{privateKey: privateKey}

	for _, enhance := range opts {
		enhance(rc)
	}
	rc.EnsureDefaults()

	return rc
}

func (rc *Recaptcha) EnsureDefaults() {
	if rc.serverName == "" {
		rc.serverName = recaptchaServerName
	}
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
	resp, err := http.PostForm(rc.serverName, url.Values{
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
		return recapResp, err
	}

	return recapResp, rc.ErrorWrap(recapResp.ErrorCodes)
}

func (rc *Recaptcha) ErrorWrap(errorCodes []string) error {
	for _, errCode := range errorCodes {
		switch errCode {
		case "missing-input-secret":
			return ErrMissingInputSecret
		case "invalid-input-secret":
			return ErrMissingInputSecret
		case "invalid-input-response":
			return ErrInvalidInputResponse
		case "bad-request":
			return ErrBadRequest
		case "timeout-or-duplicate":
			return ErrTimeoutOrDuplicate
		default:
			return ErrUnknown
		}
	}

	return nil
}

func (rc *Recaptcha) Confirm(remoteIP, response string) (bool, error) {
	resp, err := rc.check(remoteIP, response)
	if err != nil {
		return false, err
	}

	return resp.Success, nil
}

