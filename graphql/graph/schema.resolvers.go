package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"log"

	"github.com/guilhermewolke/studio-sol-test/graphql/graph/internal"
	"github.com/guilhermewolke/studio-sol-test/graphql/graph/model"
)

// Verify is the resolver for the verify field.
func (r *queryResolver) Verify(ctx context.Context, input model.Request) (*model.Response, error) {
	log.Printf("graph.Verify - input passado: %#v", input)
	verifier := internal.NewVerifier(input)
	verifier.ReleaseTheKraken()
	return &verifier.Response, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.

type mutationResolver struct{ *Resolver }
