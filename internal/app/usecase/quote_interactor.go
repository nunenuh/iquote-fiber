package usecase

import (
	exception "github.com/nunenuh/iquote-fiber/internal/app/exeption"
	"github.com/nunenuh/iquote-fiber/internal/domain/entity"
	"github.com/nunenuh/iquote-fiber/internal/domain/repository"
)

type QuoteUseCase struct {
	repo repository.IQuoteRepository
}

func NewQuoteUsecase(r repository.IQuoteRepository) *QuoteUseCase {
	return &QuoteUseCase{
		repo: r,
	}
}

func (ucase *QuoteUseCase) GetAll(limit int, offset int) ([]*entity.Quote, error) {
	u, err := ucase.repo.GetAll(limit, offset)
	if err != nil {
		return nil, exception.NewRepositoryError(err.Error())
	}

	return u, nil
}

// func (ucase *QuoteUseCase) GetByAuthor(name string) ([]*entity.Quote, error) {
// 	u, err := ucase.repo.GetByAuthor(name)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return u, nil
// }
// func (ucase *QuoteUseCase) GetByCategory(category string) ([]*entity.Quote, error) {
// 	u, err := ucase.repo.GetByCategory(category)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return u, nil

// }
// func (ucase *QuoteUseCase) GetByTags(tags string) ([]*entity.Quote, error) {
// 	u, err := ucase.repo.GetByTags(tags)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return u, nil
// }
// func (ucase *QuoteUseCase) Search(keyword string) ([]*entity.Quote, error) {
// 	u, err := ucase.repo.Search(keyword)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return u, nil
// }
// func (ucase *QuoteUseCase) Like(quoteID int, userID int) (*entity.Quote, error) {
// 	u, err := ucase.repo.Like(quoteID, userID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return u, nil
// }
// func (ucase *QuoteUseCase) Unlike(quoteID int, userID int) (*entity.Author, error) {
// 	u, err := ucase.repo.Unlike(quoteID, userID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return u, nil
// }

func (ucase *QuoteUseCase) GetByID(ID int) (*entity.Quote, error) {
	u, err := ucase.repo.GetByID(ID)
	if err != nil {
		return nil, exception.NewRepositoryError(err.Error())
	}

	return u, nil
}

func (ucase *QuoteUseCase) Create(quote *entity.Quote) (*entity.Quote, error) {
	u, err := ucase.repo.Create(quote)
	if err != nil {
		return nil, exception.NewRepositoryError(err.Error())
	}
	return u, nil
}

func (ucase *QuoteUseCase) Update(ID int, quote *entity.Quote) (*entity.Quote, error) {
	u, err := ucase.repo.Update(ID, quote)
	if err != nil {
		return nil, err
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
