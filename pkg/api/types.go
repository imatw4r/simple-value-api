package api

type GetValueIndexRequest struct {
	Value string `json:"value" uri:"value" binding:"required"`
}

type GetValueIndexResponse struct {
	Index        int    `json:"index"`
	ErrorMessage string `json:"errorMessage"`
	Value        int    `json:"value"`
}
