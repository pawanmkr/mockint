package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"

	"github.com/pawanmkr/mockint/database"
	"github.com/pawanmkr/mockint/graph/model"
)

// ScheduleInterview is the resolver for the scheduleInterview field.
func (r *mutationResolver) ScheduleInterview(ctx context.Context, input model.InterviewInput) (*model.Interview, error) {
	return db.ScheduleInterview(input), nil
}

// UpdateInterview is the resolver for the updateInterview field.
func (r *mutationResolver) UpdateInterview(ctx context.Context, id string, input model.InterviewInput) (*model.Interview, error) {
	return db.UpdateMeeting(id, input), nil
}

// CancelInterview is the resolver for the cancelInterview field.
func (r *mutationResolver) CancelInterview(ctx context.Context, id string) (*model.DeleteResponse, error) {
	return db.CancelMeeting(id), nil
}

// BookInterview is the resolver for the bookInterview field.
func (r *mutationResolver) BookInterview(ctx context.Context, input model.BookInterview) (*model.Interview, error) {
	return db.BookMeeting(input), nil
}

// Interview is the resolver for the interview field.
func (r *queryResolver) Interview(ctx context.Context, id string) (*model.Interview, error) {
	return db.GetMeetingById(&id), nil
}

// AllInterviews is the resolver for the allInterviews field.
func (r *queryResolver) AllInterviews(ctx context.Context) ([]*model.Interview, error) {
	return db.GetAllMeetings(), nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
var db = database.Connect()
