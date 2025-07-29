package entities

import "github.com/marcosfrancomarinho/go-first-project/src/domain/valuesobject"

type ContainerPage struct {
	limit *valuesobject.Limit
	page  *valuesobject.Page
}

func NewContainerPage(limit *valuesobject.Limit, page *valuesobject.Page) *ContainerPage {
	return &ContainerPage{limit, page}
}

func (c *ContainerPage) GetOffset() int {
	return (c.page.GetValue() - 1) * c.limit.GetValue()
}

func (c *ContainerPage) GetQuantityPage() int {
	return c.page.GetValue()
}

func (c *ContainerPage) GetLimitOfPage() int {
	return c.limit.GetValue()
}
