package main

import (
	"MyStackoverflow/handler/answer"
	"MyStackoverflow/handler/keyword_search"
	"MyStackoverflow/handler/question"
	"MyStackoverflow/handler/topic"
	"MyStackoverflow/handler/user"
	"github.com/gin-gonic/gin"
)

func RegisterRouter() *gin.Engine {

	r := gin.Default()
	groupUser := r.Group("/user")
	{
		groupUser.POST("/add", func(c *gin.Context) {
			user.AddUser(c)
		})
		groupUser.GET("/get", func(c *gin.Context) {
			user.GetUser(c)
		})
		groupUser.GET("/list", func(c *gin.Context) {
			user.ListUser(c)
		})
		groupUser.POST("/edit", func(c *gin.Context) {
			user.EditUser(c)
		})
	}

	groupTopic := r.Group("/topic")
	{
		groupTopic.POST("/add", func(c *gin.Context) {
			topic.AddTopic(c)
		})
		groupTopic.GET("/list", func(c *gin.Context) {
			topic.ListTopic(c)
		})
	}

	groupQuestion := r.Group("/question")
	{
		groupQuestion.POST("/add", func(c *gin.Context) {
			question.AddQuestion(c)
		})
		groupQuestion.POST("/like", func(c *gin.Context) {
			question.LikeQuestion(c)
		})
		groupQuestion.GET("/get", func(c *gin.Context) {
			question.GetQuestion(c)
		})
		groupQuestion.GET("/list", func(c *gin.Context) {
			question.ListQuestion(c)
		})
		groupQuestion.POST("/edit", func(c *gin.Context) {
			question.EditQuestion(c)
		})
	}

	groupAnswer := r.Group("/answer")
	{
		groupAnswer.POST("/add", func(c *gin.Context) {
			answer.AddAnswer(c)
		})
		groupAnswer.POST("/like", func(c *gin.Context) {
			answer.LikeAnswer(c)
		})
		groupAnswer.GET("/get", func(c *gin.Context) {
			answer.GetAnswer(c)
		})
		groupAnswer.GET("/list", func(c *gin.Context) {
			answer.ListAnswer(c)
		})
		groupAnswer.POST("/select", func(c *gin.Context) {
			answer.SelectBest(c)
		})
		groupAnswer.POST("/rate", func(c *gin.Context) {
			answer.RateAnswer(c)
		})
		groupAnswer.POST("/edit", func(c *gin.Context) {
			answer.EditAnswer(c)
		})
	}

	groupKeyword := r.Group("/keyword_search")
	{
		groupKeyword.GET("/list", func(c *gin.Context) {
			keyword_search.ListByKeyword(c)
		})
	}
	return r
}
