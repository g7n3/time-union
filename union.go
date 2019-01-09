package time_union

import (
	"errors"
	"sort"
)

type T struct {
	Start int64
	End int64
}

type TSlice []T

func (t TSlice) Len() int { return len(t) }

func (t TSlice) Swap(i, j int) { t[i], t[j] = t[j], t[i] }

func (t TSlice) Less(i, j int) bool { return t[i].Start < t[j].Start }

func (t TSlice) Union() (TSlice, error) {
	var s TSlice

	if len(t) > 0 {
		sort.Stable(t)
		s = append(s, t[0])

		for _, v := range t {
			if v.Start >= s[len(s)-1].Start && v.Start <= s[len(s)-1].End {
				// combine
				if v.End > s[len(s)-1].End {
					s[len(s)-1].End = v.End
				}
			} else if v.Start > s[len(s)-1].End {
				// split
				inner := T{Start:v.Start,End:v.End}
				s = append(s, inner)
			} else {
				return s, errors.New("start time is larger than end time")
			}
		}
	}

	return s, nil
}

