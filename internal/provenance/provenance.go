package provenance

import "github.com/flinttack/cli/pkg/lockfile"

func State() lockfile.EvidenceState {
	return lockfile.EvidenceUnknown
}
