package entity

type Person struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Age       int    `json:"age" binding:"gte=5,lte=1000"`
	Email     string `json:"email" binding:"required,email"`
}

type Video struct {
	Title       string  `json:"title" binding:"min=2,max=20"`
	Description string  `json:"description" binding:"max=30"`
	URL         string  `json:"url" binding:"required,url"`
	Person      *Person `json:"person" binding:"required"`
}
