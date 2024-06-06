package hours

import (
	"fmt"
	"regexp"
	"strconv"
)

var (
	hoursRegexp = regexp.MustCompile(
		`^\s*(?P<hour>\d{1,2})(:(?P<minutes>\d{2}))?\s*(?<ampm>AM|PM)?\s*$`,
	)
)

// Hours represents time as hour:minute.
type Hours struct {
	Hour   int
	Minute int
}

func (h Hours) String() string {
	return fmt.Sprintf("%d:%02d", h.Hour, h.Minute)
}

// Parse parses a time in hour[:minute][AM|PM] format.
func Parse(when string) (*Hours, error) {
	matches := hoursRegexp.FindStringSubmatch(when)
	hourIndex := hoursRegexp.SubexpIndex("hour")
	minutesIndex := hoursRegexp.SubexpIndex("minutes")
	ampmIndex := hoursRegexp.SubexpIndex("ampm")
	var hour, minute int
	if len(matches) == 0 {
		return nil, fmt.Errorf("invalid time string: %q", when)
	}
	if len(matches)-1 <= hourIndex {
		return nil, fmt.Errorf("cannot find hour in time pattern %q", when)
	}
	hourStr := matches[hourIndex]
	n, err := strconv.ParseInt(hourStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("non-numeric hour %q: %w", hourStr, err)
	}
	hour = int(n)
	if minutesIndex != -1 {
		// if the minutes is set, use it instead of 0
		if len(matches)-1 >= minutesIndex {
			minuteStr := matches[minutesIndex]
			if minuteStr != "" {
				n, err := strconv.ParseInt(minuteStr, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("non-numeric minute %q: %w", minuteStr, err)
				}
				minute = int(n)
			}
		}
	}
	if hour < 0 {
		return nil, fmt.Errorf("hour cannot be < 0")
	}
	if minute < 0 || minute > 59 {
		return nil, fmt.Errorf("minute must be in range [0..59], got %d", minute)
	}
	if ampmIndex != -1 {
		switch matches[ampmIndex] {
		case "":
			// 24-hour format
			if hour > 23 {
				return nil, fmt.Errorf("hour in 24-hour format cannot be > 23, got %d", hour)
			}
		case "AM":
			// 12-hour format, before noon
			if hour > 12 {
				return nil, fmt.Errorf("hour in 12-hour format cannot be > 12, got %d", hour)
			}
			// convert to 24-hour format
			if hour == 12 {
				hour = 0
			}
		case "PM":
			// 12-hour format, after noon
			if hour > 12 {
				return nil, fmt.Errorf("hour in 12-hour format cannot be > 12, got %d", hour)
			}
			// convert to 24-hour format
			if hour != 12 {
				hour += 12
			}
		}
	}
	return &Hours{
		Hour:   hour,
		Minute: minute,
	}, nil
}
