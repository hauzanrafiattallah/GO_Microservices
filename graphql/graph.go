package main

type Server struct {
	// accountClient *account.Client
	// catalogClient *catalog.Client
	// orderClient *order.Client
}

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

func (s *Server) Query() QueryResolver {
	return &queryResolver{server: s}
}

func (s *Server) Account() AccountResolver {
	return &accountResolver{server: s}
}