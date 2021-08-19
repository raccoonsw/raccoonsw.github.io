package models

type Paging struct {
	Type   string `form:"type" binding:"omitempty,oneof=virtual_good virtual_currency bundle"`
	Limit  int    `form:"limit" binding:"omitempty,min=1,max=100"`
	Offset int    `form:"offset" binding:"omitempty,min=0"`
}
