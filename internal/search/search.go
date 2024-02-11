package search

import (
	client2 "vinted-bidder/internal/client"
)

type Tool struct {
	client2.VintedClient
}

func New() (*Tool, error) {
	v, err := client2.NewVintedClient()
	if err != nil {
		return nil, err
	}
	return &Tool{*v}, nil
}

func (s *Tool) Search(qp client2.IQueryParams) (*client2.ItemsResponse, error) {
	items, err := s.FindItems(qp)
	if err != nil {
		return &client2.ItemsResponse{}, err
	}
	return items, nil
}
