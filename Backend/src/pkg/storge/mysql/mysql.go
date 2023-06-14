package mysql

import (
	"NewWork/pkg/model"
	"context"
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

type GormStorage struct {
	gormClient *gorm.DB
}

const connectionURL = "root:root@tcp(localhost:3306)/workspace_booking?charset=utf8&parseTime=True&loc=Local"

func NewGormStorage() *GormStorage {

	gormClient, err := gorm.Open(mysql.Open(connectionURL), &gorm.Config{})
	log.Println()
	if err != nil {
		log.Panicln("cannot Connection to Database")
	}

	return &GormStorage{
		gormClient: gormClient,
	}
}

func (db *GormStorage) GetWorkspace(ctx context.Context, workspaceId int) (model.Workspace, error) {
	var workspace model.Workspace

	err := db.gormClient.Where("Id = ?", workspaceId).Find(&workspace).Error
	if err != nil {
		log.Panicln("Database Error find GetWorkspaces: ", err)
		return model.Workspace{}, err
	}

	return model.Workspace{}, nil
}

func (db *GormStorage) GetAllWorkspaces(ctx context.Context, filter *model.WorkspaceFilter) ([]model.Workspace, error) {
	var workspace []model.Workspace

	err := db.gormClient.Where(createQuery(filter)).Find(&workspace).Error
	if err != nil {
		log.Panicln("Database Error find GetAllWorkspaces: ", err)
		return nil, err
	}

	return workspace, nil
}

func createQuery(filter *model.WorkspaceFilter) clause.Expression {
	expr := clause.Where{}

	if filter.HasDockingStation {
		expr = clause.Where{Exprs: append(expr.Exprs, clause.Eq{Column: "docking_station_present", Value: filter.HasDockingStation})}
	}
	if filter.HasAdjustableDesk {
		expr = clause.Where{Exprs: append(expr.Exprs, clause.Eq{Column: "adjustable_desk_present", Value: filter.HasAdjustableDesk})}
	}
	if filter.HasTwoScreens {
		expr = clause.Where{Exprs: append(expr.Exprs, clause.Eq{Column: "has_two_screens", Value: filter.HasTwoScreens})}
	}
	if filter.RoomId != "" {
		expr = clause.Where{Exprs: append(expr.Exprs, clause.Eq{Column: "room_Id", Value: filter.RoomId})}
	}
	if filter.WorkspaceId != "" {
		expr = clause.Where{Exprs: append(expr.Exprs, clause.Eq{Column: "name", Value: "Arbeitsplatz 0" + filter.WorkspaceId})}
	}
	if len(expr.Exprs) == 0 {
		expr.Exprs = append(expr.Exprs, clause.Expr{SQL: "1 = 1"})
	}

	return expr
}

func (db *GormStorage) CreateBooking(ctx context.Context, booking model.Booking) model.Error {
	log.Println("CreateBooking")
	err := db.gormClient.AutoMigrate(&model.Booking{})

	var existingBooking model.Booking
	err = db.gormClient.Where("workspace_id = ? AND date = ?", booking.WorkspaceId, booking.Date).First(&existingBooking).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Panicln("Database Error checking for existing booking: ", err)
			return model.Error{
				Code:    500,
				Message: "Database Error",
			}
		}
	} else {
		log.Println("Booking already exists for the given WorkspaceId and Date.")
		return model.Error{
			Code:    409,
			Message: "Conflict, the workspace is already booked",
		}
	}

	err = db.gormClient.Create(&booking).Error
	if err != nil {
		log.Panicln("Database Error Create Booking: ", err)
		return model.Error{
			Code:    500,
			Message: "Database Error",
		}
	}

	return model.Error{}
}

func (db *GormStorage) GetBookingsWithWorkspaceId(ctx context.Context, workspaceId int) ([]model.Booking, error) {
	var bookings []model.Booking
	log.Println("GetBookingsWithWorkspaceId: ", workspaceId)

	err := db.gormClient.Where("workspace_Id = ?", workspaceId).Find(&bookings).Error
	if err != nil {
		log.Panicln("Database Error find GetAllWorkspaces: ", err)
		return nil, err
	}

	return bookings, nil
}

func (db *GormStorage) GetAllBookings(ctx context.Context, personId string) ([]model.Booking, error) {

	var bookings []model.Booking

	err := db.gormClient.Where("PersonId = ?", personId).Find(&bookings).Error
	if err != nil {
		log.Panicln("Database Error find GetAllWorkspaces: ", err)
		return nil, err
	}

	return bookings, nil
}

func (db *GormStorage) DeleteBooking(ctx context.Context, bookingId int) error {
	var booking model.Booking

	err := db.gormClient.Where("Id = ?", bookingId).First(&booking).Error // Find the booking
	if err != nil {
		log.Println("Database Error dont Find booking: ", err)
		return err
	}
	err = db.gormClient.Delete(&booking).Error
	if err != nil {
		log.Panicln("Database Error by delete: ", err)
		return err
	}

	return nil
}
