package phases

import (
	"ko3-gin/pkg/cluster/adm/workflow"
	"ko3-gin/pkg/host"
)

func NewMarkControlPlanePhase() workflow.Phase {
	return workflow.Phase{
		Name:   "bootstrap-token",
		Phases: nil,
		Run:    runBootstrapToken,
	}
}

func runMarkControlPlane(data workflow.RunData, host host.Host) error {
	return nil
}
