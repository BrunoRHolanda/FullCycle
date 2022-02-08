package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/BrunoRHolanda/FullCycle/graph/generated"
	"github.com/BrunoRHolanda/FullCycle/graph/model"
)

func (r *categoryResolver) Cources(ctx context.Context, obj *model.Category) ([]*model.Cource, error) {
	var cources []*model.Cource

	for _, v := range r.Resolver.Cources {
		if v.Category.ID == obj.ID {
			cources = append(cources, v)
		}
	}

	return cources, nil
}

func (r *courceResolver) Chapters(ctx context.Context, obj *model.Cource) ([]*model.Chapter, error) {
	var chapters []*model.Chapter

	for _, v := range r.Resolver.Chapters {
		if v.Cource.ID == obj.ID {
			chapters = append(chapters, v)
		}
	}

	return chapters, nil
}

func (r *mutationResolver) CreateCategory(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	category := &model.Category{
		ID:          fmt.Sprintf("T%d", rand.Int()),
		Name:        input.Name,
		Description: &input.Description,
	}

	r.Categories = append(r.Categories, category)

	return category, nil
}

func (r *mutationResolver) CreateCource(ctx context.Context, input model.NewCource) (*model.Cource, error) {
	var category *model.Category

	for _, v := range r.Categories {
		if v.ID == input.CategoryID {
			category = v
		}
	}

	cource := &model.Cource{
		ID:          fmt.Sprintf("T%d", rand.Int()),
		Name:        input.Name,
		Description: &input.Description,
		Category:    category,
	}

	r.Cources = append(r.Cources, cource)

	return cource, nil
}

func (r *mutationResolver) CreateChapter(ctx context.Context, input model.NewChapter) (*model.Chapter, error) {
	var cource *model.Cource

	for _, v := range r.Cources {
		if v.ID == input.CourceID {
			cource = v
		}
	}

	chapter := &model.Chapter{
		ID:     fmt.Sprintf("T%d", rand.Int()),
		Name:   input.Name,
		Cource: cource,
	}

	r.Chapters = append(r.Chapters, chapter)

	return chapter, nil
}

func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	return r.Resolver.Categories, nil
}

func (r *queryResolver) Cources(ctx context.Context) ([]*model.Cource, error) {
	return r.Resolver.Cources, nil
}

func (r *queryResolver) Chapters(ctx context.Context) ([]*model.Chapter, error) {
	return r.Resolver.Chapters, nil
}

// Category returns generated.CategoryResolver implementation.
func (r *Resolver) Category() generated.CategoryResolver { return &categoryResolver{r} }

// Cource returns generated.CourceResolver implementation.
func (r *Resolver) Cource() generated.CourceResolver { return &courceResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type categoryResolver struct{ *Resolver }
type courceResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
