package rest

import "net/http"

type CartService interface {}

type Cart struct {
	service CartService
}
 
func (c *Cart) Get (w http.ResponseWriter, r *http.Request) {}
func (c *Cart) AddItem (w http.ResponseWriter, r *http.Request) {}
func (c *Cart) DeleteItem (w http.ResponseWriter, r *http.Request) {}
func (c *Cart) EditCountItem (w http.ResponseWriter, r *http.Request) {}
func (c *Cart) CreateOrder (w http.ResponseWriter, r *http.Request) {} 