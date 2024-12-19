package metahttputil

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseError(t *testing.T) {

	errText := `<!doctype html>
<html lang=en>
<title>400 Bad Request</title>
<h1>Bad Request</h1>
<p>1 validation error for BulkTagLookupInput<br>recording_mbid<br>  field required (type=value_error.missing)</p>
`

	xmlErr, err := ParseHtmlError(strings.NewReader(errText))
	require.NoError(t, err)

	require.Equal(t, "400 Bad Request", xmlErr.Title)
	require.Equal(t, "1 validation error for BulkTagLookupInput recording_mbid   field required (type=value_error.missing)", xmlErr.Description)

}
