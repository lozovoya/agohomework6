package app

import (
	bankgrpcv1 "agohomework6/pkg/bank/v1"
	"context"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) CreateTemplate(ctx context.Context, request *bankgrpcv1.MakeTemplate) (*bankgrpcv1.TemplateId, error) {
	return nil, nil
}

func (s *Server) GetAllTemplates(ctx context.Context, request *bankgrpcv1.All) (*bankgrpcv1.TemplatesList, error) {
	return nil, nil
}

func (s *Server) GetTemplateById(ctx context.Context, request *bankgrpcv1.TemplateId) (*bankgrpcv1.TemplatesList, error) {
	return nil, nil
}

func (s *Server) EditTemplate(ctx context.Context, request *bankgrpcv1.TemplateFixes) (*bankgrpcv1.Template, error) {
	return nil, nil
}

func (s *Server) RemoveTemplate(ctx context.Context, request *bankgrpcv1.TemplateId) (*bankgrpcv1.TemplatesList, error) {
	return nil, nil
}
