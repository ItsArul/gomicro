package web

type FailFindId struct {
	Id      uint   `json:"id" form:"id"`
	Message string `json:"message" form:"message"`
}

type FailUpdate struct {
	Id      uint   `json:"id" form:"id"`
	Message string `json:"message" form:"message"`
}

type FailFindCategory struct {
	Category string `json:"category" form:"category"`
	Message  string `json:"message" form:"message"`
}

type FailOrder struct {
	Id      uint   `json:"id" form:"id"`
	Message string `json:"message" form:"message"`
}

type FailUpload struct {
	Id      uint   `json:"id" form:"id"`
	Name    string `json:"name" form:"name"`
	Message string `json:"message" form:"message"`
}

type SuccessFindCategory struct {
	Category string `json:"category" form:"category"`
	Product  string `json:"product" form:"product"`
}

type SuccessUpdate struct {
	Id      uint        `json:"id" form:"id"`
	Product interface{} `json:"product" form:"product"`
}

type SuccessFindId struct {
	NameProduct string `json:"name_product" form:"name_product"`
	Description string `json:"description" form:"description"`
	Price       int    `json:"price" form:"price"`
}

type SuccessUpload struct {
	NameProduct string `json:"name_product" form:"name_product"`
	Description string `json:"description" form:"description"`
	Price       int    `json:"price" form:"price"`
	Quantity    int    `json:"quantity" form:"quantity"`
}
