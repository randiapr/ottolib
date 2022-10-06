package model

type ReqAuthApim struct {
	GrantType      string      `json:"grant_type"`
	AdditionalInfo interface{} `json:"additionalInfo"`
}

type ResAuthApim struct {
	ResponseCode    string      `json:"responseCode"`
	ResponseMessage string      `json:"responseMessage"`
	AccessToken     string      `json:"accessToken"`
	TokenType       string      `json:"tokenType"`
	ExpiresIn       string      `json:"expiredIn"`
	AdditionalInfo  interface{} `json:"additionalInfo"`
}

type HeaderTokenB2B struct {
	XTimestamp string `reqHeader:"X-TIMESTAMP" validate:"required"`
	XClientKey string `reqHeader:"X-CLIENT-KEY" validate:"required"`
	XSignature string `reqHeader:"X-SIGNATURE" validate:"required"`
}

type HeaderTokenApim struct {
}
