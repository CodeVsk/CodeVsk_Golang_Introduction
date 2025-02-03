package controller

import (
	"net/http"

	"github.com/codevsk/golang/hexagonal/adapter/input/converter"
	"github.com/codevsk/golang/hexagonal/adapter/input/model/request"
	"github.com/codevsk/golang/hexagonal/application/domain"
	"github.com/codevsk/golang/hexagonal/application/port/input"
	"github.com/codevsk/golang/hexagonal/configuration/result"
	"github.com/gin-gonic/gin"
)

type productController struct {
	productsUseCase input.ProductsUseCase
}

func NewProductController(productsUseCase input.ProductsUseCase) *productController {
	return &productController{productsUseCase}
}

func (p *productController) GetProducts(c *gin.Context){
	productsRequest := request.ProductsRequest{}

	if err := c.ShouldBindQuery(&productsRequest); err != nil {

		errResult := result.NewBadRequestResult(err.Error())

		c.JSON(int(errResult.StatusCode), errResult)

		return
	}

	productsDomain := domain.ProductsReqDomain{
		Type: productsRequest.Type,
	}

	productsResponseDomain, err := p.productsUseCase.GetProducts(productsDomain)
	if err != nil {
		errResult := result.NewInternalResult(err)

		c.JSON(int(errResult.StatusCode), errResult)
		return
	}

	c.JSON(http.StatusOK, converter.ConvertDomainToResponse(productsResponseDomain))
}