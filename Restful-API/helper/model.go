package helper

import (
	"belajar-go-restful-api/model/domain"
	"belajar-go-restful-api/model/web"
)

func ToCategoryResponse(domain domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   domain.Id,
		Name: domain.Name,
	}
}

func ToCategoryResponses(domains []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, domain := range domains {
		categoryResponse := ToCategoryResponse(domain)
		categoryResponses = append(categoryResponses, categoryResponse)
	}
	return categoryResponses
}
