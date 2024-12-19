package metaclient

import (
	"context"
	"encoding/json"
	"fmt"

	"anime.bike/gobrains/pkg/labshttp"
	"anime.bike/gobrains/pkg/metahttputil"
	"github.com/google/uuid"
)

type Client struct {
	labs *labshttp.Client
}

func NewClient(
	labs *labshttp.Client,
) *Client {
	return &Client{
		labs: labs,
	}
}

type TrackBulkTags struct {
	Tags []TrackPopularityTag
}

type TrackPopularityTag struct {
	Percent       float32   `json:"percent"`
	RecordingMbid uuid.UUID `json:"recording_mbid"`
	Source        string    `json:"source"`
	Tag           string    `json:"tag"`
}

func (c *Client) GetBulkTagLookup(
	ctx context.Context,
	mbid uuid.UUID,
) (*TrackBulkTags, error) {
	resp, err := c.labs.GetBulkTagLookupJson(ctx, &labshttp.GetBulkTagLookupJsonParams{
		RecordingMbid: mbid,
	})
	if err != nil {
		return nil, err
	}
	if err := metahttputil.CheckHttpError(resp); err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		// TODO: create some error handling for the html errors...
		return nil, fmt.Errorf("unexpected error %s: %d", resp.Status, resp.StatusCode)
	}
	o := &TrackBulkTags{}
	if err := json.NewDecoder(resp.Body).Decode(o); err != nil {
		return nil, err
	}
	return o, nil
}
