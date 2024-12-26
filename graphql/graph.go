package main

// type server struct {
// 	accountClient *account.Client
// 	catalogClient *catalog.Client
// 	orderClient *order.Client
// }

func NewGraphQLServer(accountUrl, catalogUrl, orderUrl string) (*Server, error) {
	accountClient, err := account.NewClient(accountUrl)

	if err != nil {
		return nil, err
	}

	cattalogClient, err := catalog.NewClient(catalogUrl)

	if err != nil {
		accountClient.Close()
		return nil, err
	}

	orderClient, err := order.NewClient(orderUrl)

	if err != nil {
		accountClient.Close()
		cattalogClient.Close()
		return nil, err
	}

	return &Server{
		accountClient,
		cattalogClient,
		orderClient,
	}, nil
}

func (s *Server) Mutation() MutationResolver {
	return &mutationResolver{server: s}
}
