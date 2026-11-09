package product

type ProductUseCase struct {
	Repository ProductRepositoryInterface
}

func NewProductUseCase(repository ProductRepositoryInterface) *ProductUseCase {
	return &ProductUseCase{Repository: repository}
}

// GetProductByID retrieves a product by its ID using the repository.
// this product was not supposed to be returned, we should return a DTO instead, but for the sake of simplicity, we will return the entity directly.
func (uc *ProductUseCase) GetProductByID(id int) (*Product, error) {
	return uc.Repository.GetProductByID(id)
}