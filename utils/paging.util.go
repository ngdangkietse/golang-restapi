package utils

import (
	"golang-rest-api/constants"
	"golang-rest-api/models"
)

func PageDirection(direction string) string {
	if direction != constants.Asc && direction != constants.Desc {
		return constants.Asc
	}
	return direction
}

func PageIndex(index int) int {
	if index <= 0 {
		return constants.PageIndexDefault
	}
	return index
}

func PageSize(size int) int {
	if size <= 0 {
		return constants.PageSizeDefault
	}
	return size
}

func PageSort(sort string) string {
	if sort == "" {
		return constants.PageSortDefault
	}
	return sort
}

func PageOffset(index int, size int) int {
	return (index - 1) * size
}

func Paging(paging models.Paging) models.Paging {
	return models.Paging{
		PageIndex:     PageIndex(paging.PageIndex),
		PageSize:      PageSize(paging.PageSize),
		PageSort:      PageSort(paging.PageSort),
		PageDirection: PageDirection(paging.PageDirection)}
}
