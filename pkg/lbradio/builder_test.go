package lbradio_test

import (
	"testing"

	"anime.bike/gobrains/pkg/lbradio"
	"github.com/stretchr/testify/require"
)

func TestBuilderSimple(t *testing.T) {
	b := lbradio.PromptBuilder{}

	b.Add("recs", "tuxpaint")
	str := b.String()

	require.Equal(t, "recs:(tuxpaint)", str)
}

func TestBuilderWeight(t *testing.T) {
	b := lbradio.PromptBuilder{}

	// can use enum or not enum
	b.Add(lbradio.EntityTypeArtist, "Motion City Soundtrack")
	b.AddWithWeight("artist", 10, "Brand New")
	str := b.String()

	require.Equal(t, "artist:(Motion City Soundtrack) artist:(Brand New):10", str)
}

func TestBuilderOptions(t *testing.T) {
	b := lbradio.PromptBuilder{}

	b.AddWithOption("artist", lbradio.OptionTypeNosim, "Motion City Soundtrack")
	str := b.String()

	require.Equal(t, "artist:(Motion City Soundtrack):nosim", str)
}

func TestBuilderPromptItem(t *testing.T) {
	b := lbradio.PromptBuilder{}

	b.AddPromptItem(lbradio.PromptItem{
		Entity: "artist",
		Values: []string{"Motion City Soundtrack"},
		Weight: 10,
		Option: "nosim",
	})

	require.Equal(t, "artist:(Motion City Soundtrack):10:nosim", b.String())
}

func TestBuilderPopurri(t *testing.T) {
	b := lbradio.PromptBuilder{}

	b.AddPromptItem(lbradio.PromptItem{
		Entity: "artist",
		Values: []string{"Motion City Soundtrack"},
		Weight: 10,
		Option: lbradio.OptionTypeHard,
	})
	b.Add("artist", "Brand New")
	b.AddWithWeight("recs", 10, "tuxpaint")

	require.Equal(t, "artist:(Motion City Soundtrack):10:hard artist:(Brand New) recs:(tuxpaint):10", b.String())

}
