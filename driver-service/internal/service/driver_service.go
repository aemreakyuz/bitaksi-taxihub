package service

import (
	"errors"

	"github.com/aemreakyuz/bitaksi-taxihub/driver-service/internal/model"
	"github.com/aemreakyuz/bitaksi-taxihub/driver-service/internal/repository"
)

type DriverService struct {
	repo *repository.DriverRepository
}

func NewDriverService(repo *repository.DriverRepository) *DriverService {
	return &DriverService{
		repo: repo,
	}
}

func (s *DriverService) CreateDriver(driver *model.Driver) error {
	if driver.FirstName == "" {
		return errors.New("first name is required")
	}
	if driver.LastName == "" {
		return errors.New("last name is required")
	}
	if driver.Plate == "" {
		return errors.New("taxi plate is required")
	}
	if driver.TaxiType == "" {
		return errors.New("taxi type is required")
	}

	return s.repo.Create(driver)
}

func (s *DriverService) UpdateDriver(id string, driver *model.Driver) error {

	if id == "" {
		return errors.New("user id is required")
	}
	if driver.FirstName == "" {
		return errors.New("first name is required")
	}
	if driver.LastName == "" {
		return errors.New("last name is required")
	}
	if driver.Plate == "" {
		return errors.New("taxi plate is required")
	}

	return s.repo.Update(id, driver)
}

func (s *DriverService) GetAllDrivers(page, pageSize int) ([]*model.Driver, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}

	return s.repo.FindAll(page, pageSize)
}

func (s *DriverService) GetNearbyDrivers(lat, lon float64, taxiType string) ([]*model.Driver, error) {
	if lat < -90 || lat > 90 {
		return nil, errors.New("invalid latitude: must be between -90 and 90")
	}
	if lon < -180 || lon > 180 {
		return nil, errors.New("invalid longitude: must be between -180 and 180")
	}

	const radiusKm = 6.0

	return s.repo.FindNearby(lat, lon, taxiType, radiusKm)
}
