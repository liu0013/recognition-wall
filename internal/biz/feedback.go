package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type Feedback struct {
	ID        int64
	Name      string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Like      int32
}

type FeedbackRepo interface {
	CreateFeedback(ctx context.Context, feedback *Feedback) error
	UpdateFeedback(ctx context.Context, id int64, feedback *Feedback) error
	ListFeedback(ctx context.Context) ([]*Feedback, error)
	GetFeedback(ctx context.Context, id int64) (*Feedback, error)
	DeleteFeedback(ctx context.Context, id int64) error
	IncLikeCount(ctx context.Context, feedback *Feedback, delta int32) error
	GetLikeCount(ctx context.Context, feedback *Feedback) int32
}

type FeedbackUsecase struct {
	repo FeedbackRepo
	log  *log.Helper
}

func NewFeedbackUsecase(repo FeedbackRepo, logger log.Logger) *FeedbackUsecase {
	return &FeedbackUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *FeedbackUsecase) Create(ctx context.Context, feedback *Feedback) error {
	return uc.repo.CreateFeedback(ctx, feedback)
}

func (uc *FeedbackUsecase) Update(ctx context.Context, id int64, feedback *Feedback) error {
	return uc.repo.UpdateFeedback(ctx, id, feedback)
}

func (uc *FeedbackUsecase) Delete(ctx context.Context, id int64) error {
	return uc.repo.DeleteFeedback(ctx, id)
}

func (uc *FeedbackUsecase) ListAll(ctx context.Context) ([]*Feedback, error) {
	fs, err := uc.repo.ListFeedback(ctx)
	if err != nil {
		return nil, err
	}
	for _, f := range fs {
		f.Like = uc.repo.GetLikeCount(ctx, f)
	}
	return fs, err
}

func (uc *FeedbackUsecase) GetFeedback(ctx context.Context, id int64) (*Feedback, error) {
	f, err := uc.repo.GetFeedback(ctx, id)
	if err != nil {
		return nil, err
	}
	f.Like = uc.repo.GetLikeCount(ctx, f)
	return f, err
}

func (uc *FeedbackUsecase) LikeFeedback(ctx context.Context, id int64, unlike bool) error {
	f, err := uc.repo.GetFeedback(ctx, id)
	if err != nil {
		return err
	}
	if unlike {
		if uc.repo.GetLikeCount(ctx, f) > 0 {
			uc.repo.IncLikeCount(ctx, f, -1)
		}
		return nil
	}
	return uc.repo.IncLikeCount(ctx, f, 1)
}
