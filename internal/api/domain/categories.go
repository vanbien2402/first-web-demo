package domain

//CategoryController Category Controller interface
type CategoryController interface {
}

//CategoryService Category Service interface
type CategoryService interface {
}

//CategoryRepository Category Repository interface
type CategoryRepository interface {
}

type CategoryCreateParams struct {
	Name          string `json:"name" binding:"required,max=50"`
	ExpiredInDays int    `json:"expiredInDays" binding:"required"`
	Remark        string `json:"remark" binding:"max=1000"`
}
