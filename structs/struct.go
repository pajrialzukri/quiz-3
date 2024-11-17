package structs

type User struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	CreatedAt  string `json:"created_at"`
	CreatedBy  string `json:"created_by"`
	ModifiedAt string `json:"modified_at"`
	ModifiedBy string `json:"modified_by"`
}

type UserPayload struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Category struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	CreatedAt  string `json:"created_at"`
	CreatedBy  string `json:"created_by"`
	ModifiedAt string `json:"modified_at"`
	ModifiedBy string `json:"modified_by"`
}

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	ReleaseYear int    `json:"release_year"`
	Price       int    `json:"price"`
	TotalPage   int    `json:"total_page"`
	Thickness   string `json:"thickness"`
	CategoryID  int    `json:"category_id"`
	CreatedAt   string `json:"created_at"`
	CreatedBy   string `json:"created_by"`
	ModifiedAt  string `json:"modified_at"`
	ModifiedBy  string `json:"modified_by"`
}
