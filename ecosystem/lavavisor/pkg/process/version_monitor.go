package processmanager

import (
	"context"
	"path/filepath"

	lvutil "github.com/lavanet/lava/ecosystem/lavavisor/pkg/util"
	"github.com/lavanet/lava/utils"
	protocoltypes "github.com/lavanet/lava/x/protocol/types"
)

type VersionMonitor struct {
	BinaryPath       string
	LavavisorPath    string
	updateTriggered  chan bool
	mismatchType     lvutil.MismatchType
	lastknownversion *protocoltypes.Version
	processes        []*ServiceProcess
	autoDownload     bool
}

func NewVersionMonitor(initVersion string, lavavisorPath string, processes []*ServiceProcess, autoDownload bool) *VersionMonitor {
	versionDir := filepath.Join(lavavisorPath, "upgrades", "v"+initVersion)
	binaryPath := filepath.Join(versionDir, "lava-protocol")

	return &VersionMonitor{
		BinaryPath:      binaryPath,
		LavavisorPath:   lavavisorPath,
		updateTriggered: make(chan bool),
		processes:       processes,
		autoDownload:    autoDownload,
	}
}

func (vm *VersionMonitor) MonitorVersionUpdates(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-vm.updateTriggered:
				// fetch new version from consensus
				utils.LavaFormatInfo("Update detected. Lavavisor starting the auto-upgrade...")

				// 1. check lavavisor directory first and attempt to fetch new binary from there
				var versionToUpgrade string
				if vm.mismatchType == 1 {
					versionToUpgrade = vm.lastknownversion.ProviderMin
				} else if vm.mismatchType == 2 {
					versionToUpgrade = vm.lastknownversion.ProviderTarget
				} else {
					utils.LavaFormatFatal("Unknown mismatch type detected in Version Monitor!", nil)
				}
				versionDir := filepath.Join(vm.LavavisorPath, "upgrades", "v"+versionToUpgrade)
				binaryPath := filepath.Join(versionDir, "lava-protocol")
				vm.BinaryPath = binaryPath // updating new binary path for validating new binary

				// fetcher
				_, err := FetchProtocolBinary(vm.LavavisorPath, vm.autoDownload, vm.lastknownversion)
				if err != nil {
					utils.LavaFormatFatal("Lavavisor was not able to fetch updated version!", nil, utils.Attribute{Key: "Version", Value: versionToUpgrade})
				}
				// linker
				CreateLink(binaryPath)

				for _, process := range vm.processes {
					utils.LavaFormatInfo("Restarting process", utils.Attribute{Key: "Process", Value: process.Name})
					vm.processes = StartProcess(vm.processes, process.Name)
				}

				utils.LavaFormatInfo("Lavavisor successfully updated protocol version!", utils.Attribute{Key: "Upgraded version:", Value: versionToUpgrade})
			}
		}
	}()
}

func (vm *VersionMonitor) ValidateProtocolVersion(incoming *protocoltypes.Version) error {
	binaryVersion, err := getBinaryVersion(vm.BinaryPath)
	if err != nil {
		return utils.LavaFormatError("failed to get binary version", err)
	}
	minVersionMismatch := (incoming.ConsumerMin != binaryVersion || incoming.ProviderMin != binaryVersion)
	targetVersionMismatch := (incoming.ConsumerTarget != binaryVersion || incoming.ProviderTarget != binaryVersion)

	// Take action only if both mismatches are detected
	if minVersionMismatch && targetVersionMismatch {
		select {
		case vm.updateTriggered <- true:
		default:
		}
		vm.mismatchType = lvutil.MinVersionMismatch
		vm.lastknownversion = incoming
		return utils.LavaFormatError("Version mismatch detected!", nil)
	}

	// version is ok.
	utils.LavaFormatInfo("Validated protocol version", utils.Attribute{Key: "current binary", Value: binaryVersion})

	return nil
}
