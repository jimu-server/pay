package control

type PayArgs struct {
	Value string `json:"value" form:"value" binding:"required"`
}
