package lbradio

import (
	"bytes"
	"strconv"
)

type PromptBuilder PromptItems

func (o *PromptBuilder) Add(Entity string, values ...string) *PromptBuilder {
	*o = append(*o, PromptItem{
		Entity: Entity,
		Values: values,
	})
	return o
}

func (o *PromptBuilder) AddPromptItem(item PromptItem) *PromptBuilder {
	*o = append(*o, item)
	return o
}

func (o *PromptBuilder) AddWithWeight(entity string, count int, values ...string) *PromptBuilder {
	*o = append(*o, PromptItem{
		Entity: entity,
		Weight: count,
		Values: values,
	})
	return o
}

func (o *PromptBuilder) AddWithOption(entity string, option string, values ...string) *PromptBuilder {
	*o = append(*o, PromptItem{
		Entity: entity,
		Option: option,
		Values: values,
	})
	return o
}

func (o *PromptBuilder) String() string {
	val, _ := (*PromptItems)(o).MarshalText()
	return string(val)
}

type PromptItems []PromptItem

func (r PromptItems) MarshalText() ([]byte, error) {
	o := &bytes.Buffer{}
	for pidx, v := range r {
		o.WriteString(v.Entity)
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
			o.WriteString(":" + v.Option)
		}
		if len(r) > 1 && pidx != len(r)-1 {
			o.WriteString(" ")
		}
	}
	return o.Bytes(), nil
}

type PromptItem struct {
	Entity string   `json:"entity"`
	Values []string `json:"values"`
	Weight int      `json:"weight"`
	Option string   `json:"options"`
}
