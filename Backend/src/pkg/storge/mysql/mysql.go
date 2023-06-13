package mysql

import (
	"NewWork/pkg/model"
	"context"
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
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

func (db *GormStorage) CreateWorkspace(ctx context.Context, workspace *model.WorkspaceDTO) error {
	// TODO: Implement
	return nil
}

func (db *GormStorage) GetWorkspace(ctx context.Context, workspaceId int) (model.Workspace, error) {
	var workspace model.Workspace

	err := db.gormClient.Where("WorkspaceId = ?", workspaceId).Find(&workspace).Error
	if err != nil {
		log.Panicln("Database Error find GetWorkspaces: ", err)
		return model.Workspace{}, err
	}

	return model.Workspace{}, nil
}

func (db *GormStorage) GetAllWorkspaces(ctx context.Context, filter *model.WorkspaceFilter) ([]model.Workspace, error) {
	var workspace []model.Workspace

	err := db.gormClient.Find(&workspace).Error
	if err != nil {
		log.Panicln("Database Error find GetAllWorkspaces: ", err)
		return nil, err
	}

	return workspace, nil
}

func (db *GormStorage) DeleteWorkspace(ctx context.Context, workspaceId int) error {
	// TODO: Implement
	return nil
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

func (db *GormStorage) GetBooking(ctx context.Context, bookingId int) (*model.BookingDTO, error) {
	// TODO: Implement
	return nil, nil
}

func (db *GormStorage) GetAllBookings(ctx context.Context, personId string) ([]model.Booking, error) {

	var bookings []model.Booking

	err := db.gormClient.Where("PersonId = ?", personId).Find(&bookings).Error
	if err != nil {
		log.Panicln("Database Error find GetAllWorkspaces: ", err)
		return nil, err
	}
	log.Println("tttt: ", bookings)

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

func parseTime(date string) time.Time {

	time, err := time.Parse("2006-01-02", date)

	if err != nil {
		log.Panicln("Parse error: ", err)
	}
	return time

}
