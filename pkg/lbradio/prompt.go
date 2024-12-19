package lbradio

import (
	"bytes"
	"strconv"
)

type EntityType string

const (
	EntityTypeArtist     EntityType = "artist"
	EntityTypeTag        EntityType = "tag"
	EntityTypeCollection EntityType = "collection"
	EntityTypePlaylist   EntityType = "playlist"
	EntityTypeStats      EntityType = "stats"
	EntityTypeRecs       EntityType = "recs"
	EntityTypeCountry    EntityType = "country"
)

type OptionType string

const (
	OptionTypeEasy   OptionType = "easy"
	OptionTypeMedium OptionType = "medium"
	OptionTypeHard   OptionType = "hard"

	OptionTypeNosim OptionType = "nosim"

	OptionTypeAnd OptionType = "and"
	OptionTypeOr  OptionType = "or"

	OptionTypeRecsListened   OptionType = "listened"
	OptionTypeRecsUnlistened OptionType = "unlistened"

	OptionTypeStatsWeek       OptionType = "week"
	OptionTypeStatsMonth      OptionType = "month"
	OptionTypeStatsQuarter    OptionType = "quarter"
	OptionTypeStatsHalfYearly OptionType = "half_yearly"
	OptionTypeStatsYearly     OptionType = "yearly"
	OptionTypeStatsAllTime    OptionType = "all_time"
	OptionTypeStatsThisWeek   OptionType = "this_week"
	OptionTypeStatsThisMonth  OptionType = "this_month"
	OptionTypeStatsThisYear   OptionType = "this_year"
)

type PromptItems []PromptItem

type PromptItem struct {
	Entity EntityType `json:"entity"`
	Values []string   `json:"values"`
	Weight int        `json:"weight"`
	Option OptionType `json:"options"`
}

func (r PromptItems) MarshalText() ([]byte, error) {
	o := &bytes.Buffer{}
	for pidx, v := range r {
		o.WriteString((string)(v.Entity))
		o.WriteString(":(")
		for idx, vv := range v.Values {
			o.WriteString(vv)
			if len(v.Values) > 1 && idx != len(v.Values)-1 {
				o.WriteString(",")
			}
		}
		o.WriteString(")")
		if v.Weight > 0 {
			o.WriteString(":" + strconv.Itoa(v.Weight))
		}
		if v.Option != "" {
			o.WriteString(":" + (string)(v.Option))
		}
		if len(r) > 1 && pidx != len(r)-1 {
			o.WriteString(" ")
		}
	}
	return o.Bytes(), nil
}
