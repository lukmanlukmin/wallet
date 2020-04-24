package request

type Limitofset struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}
