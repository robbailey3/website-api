package activities

import (
  "cloud.google.com/go/firestore"
  "context"
  "github.com/pkg/errors"
  "github.com/robbailey3/website-api/activities/auth"
)

type Service interface {
  GetActivities(ctx context.Context, request *GetActivitiesRequest) ([]*Activity, error)
  GetActivityById(ctx context.Context, id string) (*Activity, error)
  VerifyWebhook(req WebhookChallengeRequest) bool
}

type service struct {
  repo             Repository
  stravaApiService StravaApiService
}

func NewService(db *firestore.Client) (Service, error) {
  authService, err := auth.NewService()
  if err != nil {
    return nil, err
  }
  return &service{
    repo:             NewRepository(db),
    stravaApiService: NewStravaService(authService),
  }, nil
}

func (s *service) GetActivities(ctx context.Context, request *GetActivitiesRequest) ([]*Activity, error) {
  activities, err := s.repo.GetActivities(ctx, request.Limit, request.Skip)

  if err != nil {
    return nil, errors.Wrap(err, "failed to get activities from DB")
  }

  return activities, nil
}

func (s *service) GetActivityById(ctx context.Context, id string) (*Activity, error) {
  return s.repo.GetActivityById(ctx, id)
}

func (s *service) VerifyWebhook(req WebhookChallengeRequest) bool {
  return s.stravaApiService.WebhookIsValid(req)
}