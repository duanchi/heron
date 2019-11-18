package gateway

type Data struct {
	Data struct{
		Token string `json:"token"`
		User string `json:"user"`
		More string `json:"more"`
	} `json:"data"`
	Url string `json:"url"`
}
