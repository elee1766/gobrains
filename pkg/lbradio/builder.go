package lbradio

type PromptBuilder PromptItems

func (o *PromptBuilder) Add(entity EntityType, values ...string) *PromptBuilder {
	*o = append(*o, PromptItem{
		Entity: entity,
		Values: values,
	})
	return o
}

func (o *PromptBuilder) AddPromptItem(item PromptItem) *PromptBuilder {
	*o = append(*o, item)
	return o
}

func (o *PromptBuilder) AddWithWeight(entity EntityType, count int, values ...string) *PromptBuilder {
	*o = append(*o, PromptItem{
		Entity: entity,
		Weight: count,
		Values: values,
	})
	return o
}

func (o *PromptBuilder) AddWithOption(entity EntityType, option OptionType, values ...string) *PromptBuilder {
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
