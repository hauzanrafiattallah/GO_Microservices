package main

import "context"

type queryResolver struct {
	server *Server
}

func (r *queryResolver) Accounts(ctx context.Context, pagination *PaginationInput, id *string) ([]*Account, error) {

	return nil, nil
}

func (r *queryResolver) Products(ctx context.Context, pagination *PaginationInput, query *string) ([]*Product, error) {

	return nil, nil
}
