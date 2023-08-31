package usecase

import (
	"github.com/nunenuh/iquote-fiber/internal/quote/domain"
	"github.com/nunenuh/iquote-fiber/internal/utils/exception"
)

type QuoteUseCase struct {
	repo domain.IQuoteRepository
}

func NewQuoteUsecase(r domain.IQuoteRepository) *QuoteUseCase {
	return &QuoteUseCase{
		repo: r,
	}
}

func (ucase *QuoteUseCase) GetAll(limit int, offset int) ([]*domain.Quote, error) {
	u, err := ucase.repo.GetAll(limit, offset)
	if err != nil {
		return nil, exception.NewRepositoryError(err.Error())
	}

	return u, nil
}

func (ucase *QuoteUseCase) GetByAuthorName(name string, limit int, offset int) ([]*domain.Quote, error) {
	u, err := ucase.repo.GetByAuthorName(name, limit, offset)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ucase *QuoteUseCase) GetByAuthorID(ID int, limit int, offset int) ([]*domain.Quote, error) {
	u, err := ucase.repo.GetByAuthorID(ID, limit, offset)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ucase *QuoteUseCase) GetByCategoryName(name string, limit int, offset int) ([]*domain.Quote, error) {
	u, err := ucase.repo.GetByCategoryName(name, limit, offset)
	if err != nil {
		return nil, exception.NewRepositoryError(err.Error())
	}

	return u, nil
}

func (ucase *QuoteUseCase) GetByCategoryID(ID int, limit int, offset int) ([]*domain.Quote, error) {
	u, err := ucase.repo.GetByCategoryID(ID, limit, offset)
	if err != nil {
		return nil, exception.NewRepositoryError(err.Error())
	}

	return u, nil
}

// func (ucase *QuoteUseCase) GetByTags(tags string) ([]*domain.Quote, error) {
// 	u, err := ucase.repo.GetByTags(tags)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return u, nil
// }
// func (ucase *QuoteUseCase) Search(keyword string) ([]*domain.Quote, error) {
// 	u, err := ucase.repo.Search(keyword)
// 	if err != nil {
// 		return nil, err
// 	}

//		return u, nil
//	}

func (ucase *QuoteUseCase) Like(quoteID int, userID int) (*domain.Quote, error) {
	u, err := ucase.repo.Like(quoteID, userID)
	if err != nil {
		return nil, exception.NewRepositoryError(err.Error())
	}

	return u, nil
}

func (ucase *QuoteUseCase) Unlike(quoteID int, userID int) (*domain.Quote, error) {
	u, err := ucase.repo.Unlike(quoteID, userID)
	if err != nil {
		return nil, exception.NewRepositoryError(err.Error())
	}

	return u, nil
}

func (ucase *QuoteUseCase) GetByID(ID int) (*domain.Quote, error) {
	u, err := ucase.repo.GetByID(ID)
	if err != nil {
		return nil, exception.NewRepositoryError(err.Error())
	}

	return u, nil
}

func (ucase *QuoteUseCase) Create(quote *domain.Quote) (*domain.Quote, error) {
	u, err := ucase.repo.Create(quote)
	if err != nil {
		return nil, exception.NewRepositoryError(err.Error())
	}
	return u, nil
}

func (ucase *QuoteUseCase) Update(ID int, quote *domain.Quote) (*domain.Quote, error) {
	u, err := ucase.repo.Update(ID, quote)
	if err != nil {
		return nil, exception.NewRepositoryError(err.Error())
	}

	return u, nil
}

func (ucase *QuoteUseCase) Delete(ID int) error {
	err := ucase.repo.Delete(ID)
	if err != nil {
		return exception.NewRepositoryError(err.Error())
	}

	return nil
}
