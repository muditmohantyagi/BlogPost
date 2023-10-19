package dto

type CreatePost struct {
	Title string `binding:"required"`
	Post  string `binding:"required"`
}
type EditPost struct {
	ID    uint   `binding:"required,number"`
	Title string `binding:"required"`
	Post  string `binding:"required"`
}
type DeletePost struct {
	ID uint `binding:"required,number"`
}
type ViewPost struct {
	ID uint `binding:"required,number"`
}
