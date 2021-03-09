package app

import (
	bankgrpcv1 "agohomework6/pkg/bank/v1"
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	time3 "time"
)

var (
	ErrCreate = errors.New("template creating error")
	ErrRead   = errors.New("template reading error")
)

type Server struct {
	pool *pgxpool.Pool
	ctx  context.Context
}

func NewServer(pool *pgxpool.Pool, ctx context.Context) *Server {
	return &Server{pool: pool, ctx: ctx}
}

func (s *Server) CreateTemplate(ctx context.Context, request *bankgrpcv1.MakeTemplate) (*bankgrpcv1.TemplateId, error) {
	log.Printf("Creating template name - %s, phone - %s", request.Name, request.Phone)

	var id int64 = 0
	err := s.pool.QueryRow(ctx,
		"INSERT INTO templates (name, phone) VALUES ($1, $2) RETURNING id", request.Name, request.Phone).Scan(&id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if id != 0 {
		log.Printf("Template was created, id - %d", id)
		return &bankgrpcv1.TemplateId{Id: id}, nil
	}

	return nil, ErrCreate
}

func (s *Server) GetAllTemplates(ctx context.Context, request *bankgrpcv1.All) (*bankgrpcv1.TemplatesList, error) {
	log.Println("All templates request")

	rows, err := s.pool.Query(ctx,
		"SELECT * FROM templates LIMIT 100")
	if err != nil {
		log.Println(err)
		return nil, ErrRead
	}
	defer rows.Close()

	var templates bankgrpcv1.TemplatesList

	for rows.Next() {
		var created, edited time3.Time
		var template bankgrpcv1.Template
		err = rows.Scan(&template.Id, &template.Name, &template.Phone, &created, &edited)
		if err != nil {
			log.Println(err)
			return nil, ErrRead
		}
		template.Created, err = ptypes.TimestampProto(created)
		if err != nil {
			log.Println(err)
			return nil, ErrRead
		}
		template.Edited, err = ptypes.TimestampProto(edited)
		if err != nil {
			log.Println(err)
			return nil, ErrRead
		}
		templates.Items = append(templates.Items, &template)
	}

	return &templates, nil
}

func (s *Server) GetTemplateById(ctx context.Context, request *bankgrpcv1.TemplateId) (*bankgrpcv1.TemplatesList, error) {
	log.Printf("template %d requested", request.Id)

	var templates bankgrpcv1.TemplatesList
	var template bankgrpcv1.Template
	var created, edited time3.Time

	err := s.pool.QueryRow(ctx,
		"SELECT id, name, phone, created, edited FROM templates WHERE id = $1",
		request.Id).Scan(&template.Id, &template.Name, &template.Phone, &created, &edited)
	if err != nil {
		log.Println(err)
		return nil, ErrRead
	}

	template.Created, err = ptypes.TimestampProto(created)
	if err != nil {
		log.Println(err)
		return nil, ErrRead
	}
	template.Edited, err = ptypes.TimestampProto(edited)
	if err != nil {
		log.Println(err)
		return nil, ErrRead
	}
	templates.Items = append(templates.Items, &template)

	return &templates, nil
}

func (s *Server) EditTemplate(ctx context.Context, request *bankgrpcv1.TemplateFixes) (*bankgrpcv1.Template, error) {
	return nil, nil
}

func (s *Server) RemoveTemplate(ctx context.Context, request *bankgrpcv1.TemplateId) (*bankgrpcv1.TemplatesList, error) {
	return nil, nil
}
