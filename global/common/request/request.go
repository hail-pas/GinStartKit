package request

type PageRequestIn struct {
	PageNum  int64 `json:"pageNum" form:"PageNum" validate:"required,gt=0"`   // 页码
	PageSize int64 `json:"pageSize" form:"pageSize" validate:"required,gt=0"` // 页长
}

// IdRequestIn Find by id structure
type IdRequestIn struct {
	ID int `json:"id" form:"id" validate:"required,gt=0"`
}

func (r *IdRequestIn) Uint() uint {
	return uint(r.ID)
}

type IdsRequestIn struct {
	Ids []int `json:"ids" form:"ids"`
}

type Empty struct{}
