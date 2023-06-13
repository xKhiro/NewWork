package storge

import (
	"NewWork/pkg/model"
	"context"
)

type Storage interface {
	CreateWorkspace(ctx context.Context, workspace *model.WorkspaceDTO) error
	GetWorkspace(ctx context.Context, workspaceId int) (*model.WorkspaceDTO, error)
	GetAllWorkspaces(ctx context.Context, filter *model.WorkspaceFilter) ([]*model.WorkspaceDTO, error)
	DeleteWorkspace(ctx context.Context, workspaceId int) error

	CreateBooking(ctx context.Context, booking *model.BookingDTO) error
	GetBooking(ctx context.Context, bookingId int) (*model.BookingDTO, error)
	GetAllBookings(ctx context.Context, personId string) ([]*model.BookingDTO, error)
	DeleteBooking(ctx context.Context, bookingId int) error
}
