package controllers

import (
	"Z02zz/internal/models"
	"Z02zz/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PostController defines a controller for managing blog posts
type PostController struct {
	PostService *services.PostService
}

// NewPostController creates a new instance of PostController
func NewPostController(service *services.PostService) *PostController {
	return &PostController{PostService: service}
}

// GetPosts retrieves all blog posts and renders them in the archives template
func (ctrl *PostController) GetPosts(c *gin.Context) {
	// 从服务层获取所有文章
	posts, err := ctrl.PostService.GetAllPosts()
	if err != nil {
		// 如果获取文章时出错，返回500错误
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 渲染archieves.html模板，并传递获取到的文章
	c.HTML(http.StatusOK, "archieves.html", gin.H{
		"posts": posts,
	})
}

// GetPost retrieves a specific post by ID
func (ctrl *PostController) GetPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	post, err := ctrl.PostService.GetPostByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if post == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, post)
}

// NewPostForm renders the form to create a new post
func (ctrl *PostController) NewPostForm(c *gin.Context) {
	c.HTML(http.StatusOK, "new_archieves.html", nil)
}

// CreatePost handles the creation of a new blog post
func (ctrl *PostController) CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.PostService.CreatePost(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post created successfully"})
}

// UpdatePost handles updating an existing blog post
func (ctrl *PostController) UpdatePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post.ID = uint(id)

	if err := ctrl.PostService.UpdatePost(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})
}

// DeletePost handles deleting a blog post by ID
func (ctrl *PostController) DeletePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := ctrl.PostService.DeletePost(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
