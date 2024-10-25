package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/javadkavossi/Golange_Clean_webApi/src/api/helper"
)

type header struct {
	UserId  string
	Browser string
}

type personData struct {
	Name        string `json:"name" binding:"required,alpha,min=4,max=10"`
	Family      string `json:"family"`
	MobilNumber string `json:"mobil_number" binding:"required,mobil,min=11,max=11"`
}

type TestHandler struct {
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

// =====================

func (h *TestHandler) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "ok Test",
	})

}

func (h *TestHandler) HeaderBinder1(c *gin.Context) {
	userId := c.GetHeader("userId")

	c.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder1",
		"userId": userId,
	})
}

func (h *TestHandler) HeaderBinder2(c *gin.Context) {
	header := header{}
	c.BindHeader(&header)
	c.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder2",
		"header": header,
	})

}

func (h *TestHandler) QueryBinder1(c *gin.Context) {

	id := c.Query("id")
	name := c.Query("name")
	family := c.Query("family")
	fullName := name + " " + family
	c.JSON(http.StatusOK, gin.H{
		"result":    "HeaderBinder2",
		"id":        id,
		"full name": fullName,
	})

}

func (h *TestHandler) QueryBinder2(c *gin.Context) {

	id := c.QueryArray("id")
	name := c.Query("name")
	family := c.Query("family")
	fullName := name + " " + family
	c.JSON(http.StatusOK, gin.H{
		"result":    "HeaderBinder2",
		"ids":       id,
		"full name": fullName,
	})

}

func (h *TestHandler) ParamsBinder1(c *gin.Context) {

	id := c.Param("id")
	name := c.Param("name")
	family := c.Param("family")
	fullName := name + " " + family
	c.JSON(http.StatusOK, gin.H{
		"result":    "ParamsBinder1",
		"ids":       id,
		"full name": fullName,
	})

}

func (h *TestHandler) BodyBinder(c *gin.Context) {
	p := personData{}
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil,
				false, helper.ValidationError, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "BodyBinder",
		"person": p,
	}, true, 0))
}

// func (h *TestHandler) BodyBinder(c *gin.Context) {

// 	p := personData{}
// 	err := c.ShouldBindJSON(&p)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadGateway,
// 			helper.GenerateBaseResponseWithError(nil, false, -1, err))
// 		return
// 	}

// 	fullName := p.Name + " " + p.Family

// 	// c.JSON(http.StatusOK, gin.H{
// 	// 	"result":    "BodyBinder",
// 	// 	"name":      p,
// 	// 	"full name": fullName,
// 	// })
// 	c.JSON(http.StatusOK, helper.GenerateBaseResponse(
// 		gin.H{
// 			"result":    "BodyBinder",
// 			"name":      p,
// 			"full name": fullName,
// 		}, true, 0))

// }



func (h *TestHandler) FormBinder(c *gin.Context) {
	p := personData{}
	c.Bind(&p)
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "FormBinder",
		"person": p,
	}, true, 0))
}

func (h *TestHandler) FileBinder(c *gin.Context) {
	file, _ := c.FormFile("file")
	err := c.SaveUploadedFile(file, "file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result":    "FileBinder",
		"file Name": file.Filename,
		"size":      file.Size,
	})

}
