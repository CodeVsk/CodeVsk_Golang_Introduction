package converter

import (
	"fmt"

	"github.com/codevsk/golang/hexagonal/adapter/input/model/response"
	"github.com/codevsk/golang/hexagonal/application/domain"
	"github.com/jinzhu/copier"
)

func ConvertDomainToResponse(productDomain *domain.ProductsDomain) *response.ProductsResponse {
	productsResponse := &response.ProductsResponse{}
	copier.Copy(productsResponse, productDomain)
	
	fmt.Println(productsResponse)

	return productsResponse
}