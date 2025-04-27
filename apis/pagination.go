package apis

import (
	"fmt"
	"golang-temp/models"
)

// Paginate returns a slice of items based on page and pageSize
func paginate(page, pageSize int, objectSet []models.Service) ([]models.Service, error) {
	if page <= 0 || pageSize <= 0 {
		return nil, fmt.Errorf("invalid pagination parameters")
	}

	// Calculate the start index based on the page and pageSize
	start := (page - 1) * pageSize
	if start > len(objectSet) {
		return nil, fmt.Errorf("page out of range")
	}

	// Calculate the end index
	end := start + pageSize
	if end > len(objectSet) {
		end = len(objectSet)
	}

	// Return a slice of items for the given page
	return objectSet[start:end], nil
}
