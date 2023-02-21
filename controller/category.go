package controller

import (
	"app/models"
	"fmt"
	"log"
)

func (c *Controller) CreateCategory(req *models.CreateCategory) (string, error) {
	id, err := c.store.Category().Create(req)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (c *Controller) DeleteCategory(req *models.CategoryPrimaryKey) error {
	err := c.store.Category().Delete(req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) UpdateCategory(req *models.UpdateCategory, categoryId string) error {
	err := c.store.Category().Update(req, categoryId)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) GetByIdCategory(req *models.CategoryPrimaryKey) (models.Category, error) {
	category, err := c.store.Category().GetByID(req)
	if err != nil {
		return models.Category{}, err
	}
	return category, nil
}

func (c *Controller) GetAllCategory(req *models.GetListCategoryRequest) (models.GetListCategoryResponse, error) {
	categories, err := c.store.Category().GetAll(req)
	if err != nil {
		return models.GetListCategoryResponse{}, err
	}
	return categories, nil
}

func (c *Controller) Statistics() {
	m := map[string]int{}
	products, err := c.store.Product().GetAllProductWithCategory()
	if err != nil {
		log.Println(err)
		return
	}
	for _, v := range products {
		category, err := c.store.Category().GetByID(&models.CategoryPrimaryKey{Id: v.CategoryID})
		if err != nil {
			log.Println(err)
			return
		}
		m[category.Name]++
	}
	for k, v := range m {
		fmt.Printf("%v: %v\n", k, v)
	}
}
