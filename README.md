# BlogPost
This is an example of blog post creation using Golang, GIN, GORM(MySql), and RestAPI.
PostMan collection is included in the root folder.

HOW TO USE:
1) create the database manually with the name user_blog
2)  go to the root folder from CLI and run the migration command (To know all available commands:go run main.go help). it will create database tables:
go run main.go migrat
3) run the application:
go run main.go
