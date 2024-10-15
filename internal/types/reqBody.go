package types

type AuthReqBody struct {
	Login string `json:"login"`
	Pass  string `json:"pass"`
}

type ListReqBody struct {
	UserData string `json:"user_data"`
}
