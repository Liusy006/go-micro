package main

import "strconv"

type Product struct{
	Id int `json:"id"`
	Name string `json:"name"`
}

func newProduct(id int, name string)*Product{
	return &Product{
		Id:   id,
		Name: name,
	}
}

func newProductList(n int)[]*Product{
	ret := make([]*Product, 0)
	for i := 0; i < n; i++{
		ret = append(ret, newProduct(100+i, "product"+strconv.Itoa(100+i)))
	}
	return ret
}
