package provenance

import "github.com/flinttack/flinttack/pkg/lockfile"

func State() lockfile.EvidenceState {
	return lockfile.EvidenceUnknown
}
