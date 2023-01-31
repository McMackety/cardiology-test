package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"cardiology/graph/model"
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
)

// CreatePatient is the resolver for the createPatient field.
func (r *mutationResolver) CreatePatient(ctx context.Context, input model.PatientInput) (*model.Patient, error) {
	panic(fmt.Errorf("not implemented: CreatePatient - createPatient"))
}

// CreateScan is the resolver for the createScan field.
func (r *mutationResolver) CreateScan(ctx context.Context, input graphql.Upload) (*model.Scan, error) {
	panic(fmt.Errorf("not implemented: CreateScan - createScan"))
}

// Scans is the resolver for the scans field.
func (r *patientResolver) Scans(ctx context.Context, obj *model.Patient) ([]*model.Scan, error) {
	if obj.ID == "ptest" {
		return []*model.Scan{
			{
				ID:           "stest",
				TimeCreated:  1675129969,
				PresignedURL: "https://google.com",
			},
		}, nil
	}
	if obj.ID == "ptest1" {
		return []*model.Scan{
			{
				ID:           "stest",
				TimeCreated:  1675129969,
				PresignedURL: "https://google.com",
			},
		}, nil
	}
	if obj.ID == "ptest2" {
		return []*model.Scan{
			{
				ID:           "stest",
				TimeCreated:  1675129969,
				PresignedURL: "https://google.com",
			},
		}, nil
	}
	return []*model.Scan{}, nil
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	return &model.User{
		ID:    "utest",
		Name:  "Chase MacDonnell",
		Email: "macdonnell.chase@gmail.com",
	}, nil
}

// Patients is the resolver for the patients field.
func (r *queryResolver) Patients(ctx context.Context) ([]*model.Patient, error) {
	return []*model.Patient{
		{
			ID:   "ptest",
			Name: "John Doe",
		},
		{
			ID:   "ptest1",
			Name: "Jane Doe",
		},
		{
			ID:   "ptest2",
			Name: "James Doe",
		},
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Patient returns PatientResolver implementation.
func (r *Resolver) Patient() PatientResolver { return &patientResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type patientResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
