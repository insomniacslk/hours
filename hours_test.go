package hours

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		hour  *Hours
		err   error
		input string
		name  string
	}{
		//// POSITIVE TESTING
		// midnight
		{
			&Hours{Hour: 0, Minute: 0},
			nil,
			"0:00",
			"midnight (0:00)",
		},
		{
			&Hours{Hour: 0, Minute: 0},
			nil,
			"0:00",
			"midnight (00:00)",
		},
		{
			&Hours{Hour: 0, Minute: 0},
			nil,
			"0",
			"midnight (0)",
		},
		{
			&Hours{Hour: 0, Minute: 0},
			nil,
			"0AM",
			"midnight (0AM)",
		},
		{
			&Hours{Hour: 0, Minute: 0},
			nil,
			"0 AM",
			"midnight (0 AM)",
		},
		{
			&Hours{Hour: 0, Minute: 0},
			nil,
			"0    AM",
			"midnight (0 AM)",
		},
		{
			&Hours{Hour: 0, Minute: 0},
			nil,
			"12 AM",
			"midnight (12 AM)",
		},
		// noon
		{
			&Hours{Hour: 12, Minute: 0},
			nil,
			"12:00",
			"noon (12:00)",
		},
		{
			&Hours{Hour: 12, Minute: 0},
			nil,
			"12",
			"noon (12)",
		},
		{
			&Hours{Hour: 12, Minute: 0},
			nil,
			"12PM",
			"noon (12PM)",
		},
		{
			&Hours{Hour: 12, Minute: 0},
			nil,
			"12 PM",
			"noon (12 PM)",
		},
		{
			&Hours{Hour: 12, Minute: 0},
			nil,
			"12  PM",
			"noon (12  PM)",
		},
		// morning
		{
			&Hours{Hour: 6, Minute: 37},
			nil,
			"6:37",
			"morning (6:37)",
		},
		{
			&Hours{Hour: 6, Minute: 37},
			nil,
			"6:37AM",
			"morning (6:37AM)",
		},
		{
			&Hours{Hour: 6, Minute: 37},
			nil,
			"6:37 AM",
			"morning (6:37 AM)",
		},
		{
			&Hours{Hour: 6, Minute: 37},
			nil,
			"6:37  AM",
			"morning (6:37  AM)",
		},
		// afternoon
		{
			&Hours{Hour: 13, Minute: 37},
			nil,
			"13:37",
			"afternoon (13:37)",
		},
		{
			&Hours{Hour: 13, Minute: 37},
			nil,
			"1:37PM",
			"afternoon (1:37PM)",
		},
		{
			&Hours{Hour: 13, Minute: 37},
			nil,
			"1:37 PM",
			"afternoon (1:37 PM)",
		},
		{
			&Hours{Hour: 13, Minute: 37},
			nil,
			"1:37  PM",
			"afternoon (1:37  PM)",
		},

		//// NEGATIVE TESTING
		{
			nil,
			errors.New("invalid time string: \"\""),
			"",
			"empty string",
		},
		{
			nil,
			errors.New("invalid time string: \"   \""),
			"   ",
			"spaces",
		},
		{
			nil,
			errors.New("invalid time string: \":37 PM\""),
			":37 PM",
			"missing hour (:37 PM)",
		},
		{
			nil,
			errors.New("hour in 12-hour format cannot be > 12, got 13"),
			"13:37 PM",
			"wrong afternoon (13:37 PM)",
		},
		{
			nil,
			errors.New("invalid time string: \"-13:37\""),
			"-13:37",
			"negative (-13:37)",
		},
	}

	for _, testCase := range testCases {
		h, err := Parse(testCase.input)
		require.Equal(t, testCase.err, err, testCase.name)
		if testCase.hour == nil {
			assert.Nil(t, testCase.hour, testCase.name)
		} else {
			assert.Equal(t, testCase.hour.Hour, h.Hour, testCase.name)
			assert.Equal(t, testCase.hour.Minute, h.Minute, testCase.name)
		}
	}
}
