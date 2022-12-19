package storage

import (
	"sync"

	"github.com/google/uuid"
	"github.com/vslitvinov/entebot/internal/models"
)

type Cart struct {
	cache sync.Map
}

func NewCart() *Cart {
	return &Cart{
		cache: sync.Map{},
	}
}


func (c *Cart) create() models.Cart {

	return models.Cart{
		ID: uuid.NewString(),
		Items: make(map[string]models.CartItem),
	} 

}


func (c *Cart) Add(cid,pid string, count int64) (bool, error){

	// var cart models.Cart 

	// cart, ok := c.cache.Load(cid)
	// if !ok {
	// 	cart = c.create() 
	// } 

	// cart.(models.Cart).Add(models.CartItem{
	// 	IDProduct: pid,
	// 	Count: count,
	// })

	return false, nil

}

func (c *Cart) Delete(pid string){

}

func (c *Cart) UpdateCount(pid string, count int64){

}

// AddProduct 
// DeleteProduct 
// UpdateCountProduct 
