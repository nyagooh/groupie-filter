package services

import (
	"fmt"
	"strings"
	"time"
)

// FilterCriteria defines the structure for the filtering parameters
type FilterCriteria struct {
	CreationDateFrom  int    `json:"creationDateFrom"`
	CreationDateTo    int    `json:"creationDateTo"`
	FirstAlbumFrom    string `json:"firstAlbumFrom"`
	FirstAlbumTo      string `json:"firstAlbumTo"`
	MemberCount       int    `json:"memberCount"`
	LocationSubstring string `json:"locationSubstring"`
}

// FilterArtists filters artists based on the given criteria
func FilterArtists(artists []Artist, criteria FilterCriteria) ([]Artist, error) {
	filteredArtists := []Artist{}

	// Parse date range for first album
	var firstAlbumFrom time.Time
	var firstAlbumTo time.Time
	var err error
	if criteria.FirstAlbumFrom != "" {
		firstAlbumFrom, err = time.Parse("2006-01-02", criteria.FirstAlbumFrom)
		if err != nil {
			return nil, fmt.Errorf("invalid FirstAlbumFrom date format: %w", err)
		}
	}
	if criteria.FirstAlbumTo != "" {
		firstAlbumTo, err = time.Parse("2006-01-02", criteria.FirstAlbumTo)
		if err != nil {
			return nil, fmt.Errorf("invalid FirstAlbumTo date format: %w", err)
		}
	}

	// Parse date range for creation date

	// Apply filters
	for _, artist := range artists {
		// Filter by creation date range
		if criteria.CreationDateFrom > 0 && artist.CreationDate < criteria.CreationDateFrom {
			continue
		}
		if criteria.CreationDateTo > 0 && artist.CreationDate > criteria.CreationDateTo {
			continue
		}

		// Filter by first album date range
		if criteria.FirstAlbumFrom != "" || criteria.FirstAlbumTo != "" {
			firstAlbumDate, err := time.Parse("02-01-2006", artist.FirstAlbum)
			if err != nil {
				continue // Skip invalid date formats
			}
			if !firstAlbumFrom.IsZero() && firstAlbumDate.Before(firstAlbumFrom) {
				continue
			}
			if !firstAlbumTo.IsZero() && firstAlbumDate.After(firstAlbumTo) {
				continue
			}
		}

		// Filter by member count
		if criteria.MemberCount > 0 && len(artist.Members) != criteria.MemberCount {
			continue
		}

		// Filter by location substring
		if criteria.LocationSubstring != "" {
			matched := false
			for _, location := range artist.Locations {
				if strings.Contains(strings.ToLower(location), strings.ToLower(criteria.LocationSubstring)) {
					matched = true
					break
				}
			}
			if !matched {
				continue
			}
		}

		// Add artist to filtered list
		filteredArtists = append(filteredArtists, artist)
	}

	// test if timestamp matches
	// fmt.Println("FirstAlbumFrom parsed:", firstAlbumFrom)
	// fmt.Println("FirstAlbumTo parsed:", firstAlbumTo)

	return filteredArtists, nil
}
