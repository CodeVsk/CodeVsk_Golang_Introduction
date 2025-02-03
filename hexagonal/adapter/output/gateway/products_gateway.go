package gateway

import (
	"github.com/codevsk/golang/hexagonal/adapter/output/model/response"
	"github.com/codevsk/golang/hexagonal/application/domain"
	"github.com/go-resty/resty/v2"
	"github.com/jinzhu/copier"
)

type productsGateway struct{}

func NewProductsGateway() *productsGateway {
	client = resty.New().SetBaseURL("https://dummyjson.com")

	return &productsGateway{}
}

var (
	client *resty.Client
)

func (pg *productsGateway) GetProductsPort(productsDomain domain.ProductsReqDomain) (*domain.ProductsDomain, error) {
	productsResponse := &response.ProductsClientResponse{}

	_, err := client.R().
		SetQueryParams(map[string]string{
			"q": productsDomain.Type,
		}).
		SetResult(productsResponse).
		Get("/products/search")

	if err != nil {
		return nil, err
	}

	productsResponseDomain := &domain.ProductsDomain{}
	copier.Copy(productsResponseDomain, productsResponse)

	return productsResponseDomain, nil
}