package songci

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Dependencies(
		newSongci,
	)
}
