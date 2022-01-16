package service

import (
	"context"

	pb "recognition-wall/api/feedback/v1"
	v1 "recognition-wall/api/feedback/v1"
	"recognition-wall/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type FeedbackService struct {
	v1.UnimplementedFeedbackServiceServer
	feedbackUC *biz.FeedbackUsecase
	log        *log.Helper
}

func NewFeedbackService(feedbackUC *biz.FeedbackUsecase, logger log.Logger) *FeedbackService {
	return &FeedbackService{feedbackUC: feedbackUC, log: log.NewHelper(logger)}
}

func (s *FeedbackService) CreateFeedback(ctx context.Context, req *pb.CreateFeedbackRequest) (*pb.CreateFeedbackReply, error) {
	s.log.WithContext(ctx).Infof("create feedback content: %v", req.GetContent())
	if req.Content == "" {
		return nil, v1.ErrorFeedbackInvalid("content is empty: %s", req.GetContent())
	}
	err := s.feedbackUC.Create(ctx, &biz.Feedback{Name: req.Name, Content: req.Content})
	return &pb.CreateFeedbackReply{}, err
}
func (s *FeedbackService) UpdateFeedback(ctx context.Context, req *pb.UpdateFeedbackRequest) (*pb.UpdateFeedbackReply, error) {
	s.log.WithContext(ctx).Infof("update feedback content: %v", req.GetContent())
	err := s.feedbackUC.Update(ctx, req.Id, &biz.Feedback{
		Name:    req.Name,
		Content: req.Content,
	})
	return &pb.UpdateFeedbackReply{}, err
}
func (s *FeedbackService) DeleteFeedback(ctx context.Context, req *pb.DeleteFeedbackRequest) (*pb.DeleteFeedbackReply, error) {
	s.log.WithContext(ctx).Infof("delete feedback id: %v", req.GetId())
	err := s.feedbackUC.Delete(ctx, req.GetId())
	return &pb.DeleteFeedbackReply{}, err
}
func (s *FeedbackService) GetFeedback(ctx context.Context, req *pb.GetFeedbackRequest) (*pb.GetFeedbackReply, error) {
	s.log.WithContext(ctx).Infof("get feedback id: %v", req.GetId())
	f, err := s.feedbackUC.GetFeedback(ctx, req.GetId())
	if err != nil {
		return &pb.GetFeedbackReply{}, err
	}
	return &pb.GetFeedbackReply{Feedback: &pb.Feedback{Id: f.ID, Name: f.Name, Content: f.Content, Like: f.Like}}, err
}
func (s *FeedbackService) ListFeedback(ctx context.Context, req *pb.ListFeedbackRequest) (*pb.ListFeedbackReply, error) {
	fs, err := s.feedbackUC.ListAll(ctx)
	reply := &pb.ListFeedbackReply{}
	for _, p := range fs {
		reply.Results = append(reply.Results, &pb.Feedback{
			Id:      p.ID,
			Name:    p.Name,
			Content: p.Content,
			Like:    p.Like,
		})
	}
	return reply, err
}
func (s *FeedbackService) Like(ctx context.Context, req *pb.LikeRequest) (*pb.LikeReply, error) {
	s.log.WithContext(ctx).Infof("like feedback id: %v", req.GetId())
	err := s.feedbackUC.LikeFeedback(ctx, req.GetId(), req.GetUnlike())
	return &pb.LikeReply{}, err
}
