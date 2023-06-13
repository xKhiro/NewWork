package mysql

import (
	"NewWork/pkg/model"
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

func (db *GormStorage) CreateWorkspace(ctx context.Context, workspace *model.WorkspaceDTO) error {
	// TODO: Implement
	return nil
}

func (db *GormStorage) GetWorkspace(ctx context.Context, workspaceId int) (*model.WorkspaceDTO, error) {
	// TODO: Implement
	return nil, nil
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

func (db *GormStorage) CreateBooking(ctx context.Context, booking model.Booking) error {

	err := db.gormClient.Create(booking).Error
	if err != nil {
		log.Panicln("Database Error Create Booking: ", err)
		return err
	}

	return nil
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
	// TODO: Implement
	return nil
}
