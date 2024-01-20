package main

import (
	"fmt"
	"time"
	search2 "vinted-bidder/internal/search"
)

func main() {

	t := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-t.C:
			find()
		}
	}
}

func find() {
	searchTool, err := search2.New()
	if err != nil {
		fmt.Println(err)
		return
	}
	params := search2.QueryParams{
		SizeIDs:     nil,
		Catalog:     nil,
		MaterialIDs: nil,
		ColorIDs:    nil,
		BrandIDs:    nil,
		PriceFrom:   0,
		PriceTo:     0,
		StatusIDs:   nil,
		Order:       "",
		Currency:    "",
		SearchText:  "stone island",
	}
	items, err := searchTool.FindItems(&params)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, item := range items.Items {
		fmt.Println(item.Title + " " + item.Price + "PLN " + item.Photo.Url + " " + item.SizeTitle)
	}
}
