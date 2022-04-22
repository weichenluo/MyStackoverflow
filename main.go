package main

import (
	"MyStackoverflow/cache"
	"MyStackoverflow/handler/answer"
	"MyStackoverflow/handler/question"
	"MyStackoverflow/handler/topic"
	"MyStackoverflow/handler/user"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	groupUser := r.Group("/user")
	{
		groupUser.POST("/add", func(c *gin.Context) {
			user.AddUser(c)
		})
	}

	groupTopic := r.Group("/topic")
	{
		groupTopic.POST("/add", func(c *gin.Context) {
			topic.AddTopic(c)
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
		groupQuestion.GET("/list", func(c *gin.Context) {
			question.ListQuestion(c)
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
		groupAnswer.GET("/list", func(c *gin.Context) {
			answer.ListAnswer(c)
		})
	}

	// pre-computed cache
	cache.Init()
	// listen and serve on 0.0.0.0:8080
	err := r.Run()
	if err != nil {
		return
	}
}
