package verify

import (
	"github.com/flinttack/cli/internal/diff"
	"github.com/flinttack/cli/pkg/lockfile"
)

func Compare(locked, current lockfile.Document) diff.Result {
	return diff.Compare(locked, current)
}
