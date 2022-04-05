package feature

type GetProductReq struct {
	Version string `form:"version"`
}
type GetProductRes struct {
	Message string `json:"message"`
}
