package service

import (
	"errors"
	"pizzaria/internal/models"
)

func ValidateRatingReview(review models.Review) error {
	if review.Rating < 1 || review.Rating > 5 {
		return errors.New("A avaliação deve estar entre 1 e 5")
	}

	return nil
}
