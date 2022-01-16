package data

import (
	"context"
	"recognition-wall/internal/biz"
	"recognition-wall/internal/data/ent"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type feedbackRepo struct {
	data *Data
	log  *log.Helper
}

func NewFeedbackRepo(data *Data, logger log.Logger) biz.FeedbackRepo {
	return &feedbackRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *feedbackRepo) CreateFeedback(ctx context.Context, feedback *biz.Feedback) error {
	_, err := r.data.db.Feedback.Create().SetName(feedback.Name).SetContent(feedback.Content).Save(ctx)
	return err
}

func (r *feedbackRepo) UpdateFeedback(ctx context.Context, id int64, feedback *biz.Feedback) error {
	p, err := r.data.db.Feedback.Get(ctx, id)
	if err != nil {
		return err
	}
	_, err = p.Update().SetName(feedback.Name).SetContent(feedback.Content).SetUpdatedAt(time.Now()).Save(ctx)
	return err
}

func (r *feedbackRepo) ListFeedback(ctx context.Context) ([]*biz.Feedback, error) {
	ps, err := r.data.db.Feedback.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Feedback, 0)
	for _, p := range ps {
		rv = append(rv, &biz.Feedback{
			ID:        p.ID,
			Name:      p.Name,
			Content:   p.Content,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}
	return rv, nil
}

func (r *feedbackRepo) GetFeedback(ctx context.Context, id int64) (*biz.Feedback, error) {
	p, err := r.data.db.Feedback.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &biz.Feedback{
		ID:        p.ID,
		Name:      p.Name,
		Content:   p.Content,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}, nil
}

func (r *feedbackRepo) DeleteFeedback(ctx context.Context, id int64) error {
	return r.data.db.Feedback.DeleteOneID(id).Exec(ctx)
}

func (r *feedbackRepo) IncLikeCount(ctx context.Context, feedback *biz.Feedback, delta int32) error {
	f, err := r.data.db.Feedback.Get(ctx, feedback.ID)
	if err != nil {
		return err
	}
	like, err := r.data.db.Feedback.QueryLike(f).First(ctx)
	if err != nil && ent.IsNotFound(err) {
		r.data.db.Like.Create().SetCount(delta).SetFeedback(f).Save(ctx)
		return nil
	}
	r.data.db.Like.UpdateOne(like).AddCount(delta).Save(ctx)
	return nil
}

func (r *feedbackRepo) GetLikeCount(ctx context.Context, feedback *biz.Feedback) int32 {
	f, err := r.data.db.Feedback.Get(ctx, feedback.ID)
	if err != nil {
		return 0
	}
	like, err := r.data.db.Feedback.QueryLike(f).First(ctx)
	if err != nil && ent.IsNotFound(err) {
		return 0
	}
	return like.Count
}
