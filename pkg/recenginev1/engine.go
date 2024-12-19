package recenginev1

import "github.com/google/uuid"

type Track struct {
	MBID        uuid.UUID
	Appearances int
	Metadata    *TrackMetadata
}

type TrackMetadata struct {
}

type Population struct {
	Tracks map[uuid.UUID]*Track
}

type Policy interface {
	Epoch(*Population) (*Population, error)
}

func ExecutePolicy(pop *Population, p Policy) (*Population, error) {
	resp, err := p.Epoch(pop)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
