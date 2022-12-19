package models

import "sync"

type Cart struct {
	sync.RWMutex
	ID string 
	Items map[string]CartItem
}

type CartItem struct {
	IDProduct string
	Count int64
}


func (c *Cart) Add(ci CartItem){
	c.Lock()
	defer c.Unlock()
	c.Items[ci.IDProduct] = ci
}