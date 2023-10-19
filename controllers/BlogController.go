package controllers

import (
	"net/http"

	"blog.com/dto"
	"blog.com/models"
	"blog.com/pkg/handle"
	"blog.com/pkg/helper"
	"github.com/gin-gonic/gin"
)

type BlogController struct{}

// creation
func (con BlogController) CreateBlog(c *gin.Context) {
	var InputDTO dto.CreatePost

	if errDTO := c.ShouldBindJSON(&InputDTO); errDTO != nil {
		msg := handle.Error(errDTO)
		c.AbortWithStatusJSON(http.StatusBadRequest, msg)
		return
	}
	var post models.Post
	post.Title = InputDTO.Title
	post.Post = InputDTO.Post

	_, err := models.CreatePost(post)
	if err != nil {
		helper.ELog.Error(err.Error())
		response := helper.Error("Error in post creation", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.Success(true, "ok", "Post created successfully")
	c.JSON(http.StatusOK, response)
}

// listing
func (con BlogController) ListBlog(c *gin.Context) {
	list, err := models.ListPost()
	if err != nil {
		helper.ELog.Error(err.Error())
		response := helper.Error("Error in post listing", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.Success(true, "ok", list)
	c.JSON(http.StatusOK, response)
}

// Edit
func (con BlogController) EditBlog(c *gin.Context) {
	var InputDTO dto.EditPost

	if errDTO := c.ShouldBindJSON(&InputDTO); errDTO != nil {
		msg := handle.Error(errDTO)
		c.AbortWithStatusJSON(http.StatusBadRequest, msg)
		return
	}
	var post models.Post
	post.ID = InputDTO.ID
	post.Title = InputDTO.Title
	post.Post = InputDTO.Post
	_, err := models.UpdatePost(post)
	if err != nil {
		helper.ELog.Error(err.Error())
		response := helper.Error("Error in edit post", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.Success(true, "ok", "Post updated successfully")
	c.JSON(http.StatusOK, response)
}
func (con BlogController) DeleteBlog(c *gin.Context) {
	var InputDTO dto.DeletePost

	if errDTO := c.ShouldBindJSON(&InputDTO); errDTO != nil {
		msg := handle.Error(errDTO)
		c.AbortWithStatusJSON(http.StatusBadRequest, msg)
		return
	}
	_, err := models.DeletePost(InputDTO.ID)
	if err != nil {
		helper.ELog.Error(err.Error())
		response := helper.Error("Error in deleting the post", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.Success(true, "ok", "Post deleted successfully")
	c.JSON(http.StatusOK, response)
}

// view blog
func (con BlogController) ViewBlog(c *gin.Context) {
	var InputDTO dto.ViewPost

	if errDTO := c.ShouldBindJSON(&InputDTO); errDTO != nil {
		msg := handle.Error(errDTO)
		c.AbortWithStatusJSON(http.StatusBadRequest, msg)
		return
	}
	post, err := models.ViewPost(InputDTO.ID)
	if err != nil {
		helper.ELog.Error(err.Error())
		response := helper.Error("Error in view post", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.Success(true, "ok", post)
	c.JSON(http.StatusOK, response)
}
