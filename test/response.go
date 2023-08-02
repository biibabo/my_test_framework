// test/response.go

package test

type ResponseData struct {
	RetCode int         `json:"RetCode"`
	Message string      `json:"Message"`
	Action  string      `json:"Action"`
	Data    CaptchaData `json:"Data"`
}

type CaptchaData struct {
	CaptchaId string `json:"CaptchaId"`
}
