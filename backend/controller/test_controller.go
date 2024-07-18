package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"backend/data/request"
	"backend/data/response"
	"backend/helper"
	"backend/service"
	"strconv"
)

type TestController struct {
	testService service.TestService
}

func NewTestController(service service.TestService) *TestController {
	return &TestController{testService: service}
}

func (controller *TestController) Create(ctx *gin.Context) {
	createTestRequest := request.CreateTestRequest{}
	err := ctx.ShouldBindJSON(&createTestRequest)
	helper.ErrorPanic(err)

	controller.testService.Create(createTestRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TestController) Update(ctx *gin.Context) {
	updateTestRequest := request.UpdateTestRequest{}
	err := ctx.ShouldBindJSON(&updateTestRequest)
	helper.ErrorPanic(err)

	testId := ctx.Param("testId")
	id, err := strconv.Atoi(testId)
	helper.ErrorPanic(err)

	updateTestRequest.Id = id

	controller.testService.Update(updateTestRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TestController) Delete(ctx *gin.Context) {
	testId := ctx.Param("testId")
	id, err := strconv.Atoi(testId)
	helper.ErrorPanic(err)
	controller.testService.Delete(id)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *TestController) FindById(ctx *gin.Context) {
	testId := ctx.Param("testId")
	id, err := strconv.Atoi(testId)
	helper.ErrorPanic(err)

	testResponse := controller.testService.FindById(id)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   testResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TestController) FindAll(ctx *gin.Context) {
	testResponse := controller.testService.FindAll()

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   testResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}