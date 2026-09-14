package service

import (
	"context"

	"github.com/jimmymuthoni/distributed_ride_sharing_application/services/trip-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type service struct {
	repo domain.TripRepository

}

func NewService(repo domain.TripRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateTrip(ctx context.Context, fare *domain.RideFareModel)(*domain.TripModel, error) {
	t := &domain.TripModel{
		ID:   primitive.NewObjectID(),
		UserID:  fare.UserID,
		Status: "pending",
		RideFare: fare,
	}

	return s.repo.CreateTrip(ctx, t)

}