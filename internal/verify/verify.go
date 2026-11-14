package verify

import (
	"github.com/flinttack/flinttack/internal/diff"
	"github.com/flinttack/flinttack/pkg/lockfile"
)

func Compare(locked, current lockfile.Document) diff.Result {
	return diff.Compare(locked, current)
}
