package main

type mutationResolver struct {
	server *Server
}

// func (r *mutationResolver) createAccount(ctx context.Context, input AccountInput) (*Account, error) {
// 	return r.server.accountClient.CreateAccount(ctx, input)
// }

// func (r *mutationResolver) createProduct(ctx context.Context, input ProductInput) (*Product, error) {
// 	return r.server.catalogClient.CreateProduct(ctx, input)
// }

// func (r *mutationResolver) createOrder(ctx context.Context, input OrderInput) (*Order, error) {
// 	return r.server.orderClient.CreateOrder(ctx, input)
// }
