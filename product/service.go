package product

type Service interface {
	DeleteProductById(params *deleteProductRequest) (int64, error)
	GetBestSellingProducts() (*ProductTopResponse, error)
	GetProductById(param *getProductByIDRequest) (*Product, error)
	GetProducts(params *getProductsRequest) (*ProductsList, error)
	InsertProduct(params *getAddProductRequest) (*Product, error)
	UpdateProduct(params *getUpdateProductRequest) (*Product, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) DeleteProductById(params *deleteProductRequest) (int64, error) {
	return s.repo.DeleteProductById(params)
}

func (s *service) GetBestSellingProducts() (*ProductTopResponse, error) {
	products, err := s.repo.GetBestSellingProducts()
	if err != nil {
		panic(err)
	}

	totalSellings, err := s.repo.GetTotalSellings()
	if err != nil {
		panic(err)
	}

	return &ProductTopResponse{Data: products, TotalSellings: totalSellings}, nil
}

func (s *service) GetProductById(param *getProductByIDRequest) (*Product, error) {
	//Business Logic
	return s.repo.GetProductById(param.ProductID)
}

func (s *service) GetProducts(params *getProductsRequest) (*ProductsList, error) {
	products, err := s.repo.GetProducts(params)
	if err != nil {
		panic(err)
	}

	totalProducts, err := s.repo.GetTotalProducts()
	if err != nil {
		panic(err)
	}

	return &ProductsList{Data: products, TotalRecords: totalProducts}, nil
}

func (s *service) InsertProduct(params *getAddProductRequest) (*Product, error) {
	idProduct, err := s.repo.InsertProduct(params)
	if err != nil {
		panic(err)
	}
	return s.repo.GetProductById(int(idProduct))
}

func (s *service) UpdateProduct(params *getUpdateProductRequest) (*Product, error) {
	productId, err := s.repo.UpdateProduct(params)
	if err != nil {
		panic(err)
	}
	return s.repo.GetProductById(int(productId))
}
