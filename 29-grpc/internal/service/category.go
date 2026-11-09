package service

import (
	"context"
	"io"

	"github.com/hiaguedes/mba-go/grpc/internal/database"
	"github.com/hiaguedes/mba-go/grpc/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDb database.Category
}

func NewCategoryService(categoryDb database.Category) *CategoryService {
	return &CategoryService{
		CategoryDb: categoryDb,
	}
}

func (s *CategoryService) CreateCategory(ctx context.Context, req *pb.CategoryRequest) (*pb.CategoryResponse, error) {
	category, err := s.CategoryDb.Create(req.Name, req.Description)
	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return &pb.CategoryResponse{
		Category: categoryResponse,
	}, nil
}

func (s *CategoryService) ListCategories(ctx context.Context, in *pb.Empty) (*pb.CategoryListResponse, error) {
	categories, err := s.CategoryDb.List()
	if err != nil {
		return nil, err
	}

	var categoryResponses []*pb.Category
	for _, category := range categories {
		categoryResponses = append(categoryResponses, &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}

	return &pb.CategoryListResponse{
		Categories: categoryResponses,
	}, nil
}

func (s *CategoryService) GetCategory(ctx context.Context, req *pb.CategoryIdRequest) (*pb.CategoryResponse, error) {
	category, err := s.CategoryDb.FindByID(req.Id)
	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return &pb.CategoryResponse{
		Category: categoryResponse,
	}, nil
}

func (s *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	categories := &pb.CategoryListResponse{}
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		if err == io.EOF {
			return stream.SendAndClose(categories)
		}

		categoryResult, err := s.CategoryDb.Create(req.Name, req.Description)
		if err != nil {
			return err
		}

		categories.Categories = append(categories.Categories, &pb.Category{
			Id:          categoryResult.ID,
			Name:        categoryResult.Name,	
		    Description: categoryResult.Description,
		})
	}
}

func (s *CategoryService) CreateCategoryBidirectionalStream(stream pb.CategoryService_CreateCategoryBidirectionalStreamServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		if err == io.EOF {
			return nil
		}

		categoryResult, err := s.CategoryDb.Create(req.Name, req.Description)
		if err != nil {
			return err
		}

		categoryResponse := &pb.Category{
			Id:          categoryResult.ID,
			Name:        categoryResult.Name,
			Description: categoryResult.Description,
		}

		if err := stream.Send(categoryResponse); err != nil {
			return err
		}
}
}