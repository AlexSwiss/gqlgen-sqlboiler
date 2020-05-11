package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/AlexSwiss/gqlgen-sqlboiler/graph/generated"
	"github.com/AlexSwiss/gqlgen-sqlboiler/graph/model"
)

func (r *mutationResolver) CreateAuthor(ctx context.Context, input model.AuthorCreateInput) (*model.AuthorPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateAuthors(ctx context.Context, input model.AuthorsCreateInput) (*model.AuthorsPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateAuthor(ctx context.Context, id string, input model.AuthorUpdateInput) (*model.AuthorPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateAuthors(ctx context.Context, filter *model.AuthorFilter, input model.AuthorUpdateInput) (*model.AuthorsUpdatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteAuthor(ctx context.Context, id string) (*model.AuthorDeletePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteAuthors(ctx context.Context, filter *model.AuthorFilter) (*model.AuthorsDeletePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateBook(ctx context.Context, input model.BookCreateInput) (*model.BookPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateBooks(ctx context.Context, input model.BooksCreateInput) (*model.BooksPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateBook(ctx context.Context, id string, input model.BookUpdateInput) (*model.BookPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateBooks(ctx context.Context, filter *model.BookFilter, input model.BookUpdateInput) (*model.BooksUpdatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteBook(ctx context.Context, id string) (*model.BookDeletePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteBooks(ctx context.Context, filter *model.BookFilter) (*model.BooksDeletePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Author(ctx context.Context, id string) (*model.Author, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Authors(ctx context.Context, filter *model.AuthorFilter) ([]*model.Author, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Book(ctx context.Context, id string) (*model.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Books(ctx context.Context, filter *model.BookFilter) ([]*model.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
