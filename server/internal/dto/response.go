package dto

type CommonResponse struct {
	Data interface{} `json:"data"`
}

func NewCommonResponse(data interface{}) *CommonResponse {
	return &CommonResponse{
		Data: data,
	}
}

type CrimeData struct {
	Image       string `json:"image"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
