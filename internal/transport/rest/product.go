package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/vslitvinov/entebot/internal/models"
)

type ProductService interface {
	Create(ctx context.Context, mp models.Product) (string, error)
	DeleteByID(ctx context.Context, pid string) error
	GetByID(ctx context.Context, pid string) (models.Product, error)
	ArchiveByID(ctx context.Context, pid string, archive bool) error
	GetByCategory(ctx context.Context, category string) ([]models.Product, error) 
}

type Product struct {
	service ProductService
}

func NewProduct(s ProductService) *Product {
	return &Product{
		service:s,
	}
}

func (p *Product) RegisterHandler(router *http.ServeMux) {

	router.HandleFunc("getbyid", p.GetByID)
	router.HandleFunc("getbycategory", p.GetByCategory)

}


type ProductGetByIDResponse struct {
	Status string `json:"status"`
	Data models.Product `json:"data"`
}
 
func (p *Product) GetByID(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		r.ParseForm()                    
    	pid := r.Form.Get("pid")

     	res := ProductGetByIDResponse{}

    	mp, err := p.service.GetByID(context.TODO(),pid)
    	if err != nil {
    		log.Println(err)
    		res.Status = "error get product by id"
    	} else {
    		res.Status = "ok"
    		res.Data = mp
    	}

    	resJson, err := json.Marshal(res)
		if err != nil {
			log.Println("Marshal json req")
		}
		fmt.Fprintf(w,string(resJson))
	}

}

type ProductGetByCategoryResponse struct {
	Status string `json:"status"`
	Data []models.Product `json:"data"`
}

func (p *Product) GetByCategory(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		r.ParseForm()                    
    	category := r.Form.Get("category")

     	res := ProductGetByCategoryResponse{}

    	mps, err := p.service.GetByCategory(context.TODO(),category)
    	if err != nil {
    		log.Println(err)
    		res.Status = "error get product by id"
    	} else {
    		res.Status = "ok"
    		res.Data = mps
    	}

    	resJson, err := json.Marshal(res)
		if err != nil {
			log.Println("Marshal json req")
		}
		fmt.Fprintf(w,string(resJson))
	}

}

// type ProductGetAllResponse struct {
// 	Status string `json:"status"`
// 	Data []models.Product `json:"data"`
// }

// func (p *Product) GetAll(w http.ResponseWriter, r *http.Request) {

// 	if r.Method == "GET" {
// 		r.ParseForm()                    
//     	category := r.Form.Get("category")

//      	res := ProductGetAllResponse{}

//     	mps, err := p.service.GetAll(context.TODO(),category)
//     	if err != nil {
//     		log.Println(err)
//     		res.Status = "error get product by id"
//     	} else {
//     		res.Status = "ok"
//     		res.Data = mps
//     	}

//     	resJson, err := json.Marshal(res)
// 		if err != nil {
// 			log.Println("Marshal json req")
// 		}
// 		fmt.Fprintf(w,string(resJson))
// 	}

// }











