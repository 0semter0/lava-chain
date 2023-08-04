package lvchaintracker

import "time"

const (
	DefualtAssumedBlockMemory      = 20
	DefaultBlockCheckpointDistance = 100
)

type ChainTrackerConfig struct {
	ForkCallback             func(block int64)              // a function to be called when a fork is detected
	NewLatestCallback        func(block int64, hash string) // a function to be called when a new block is detected
	ServerAddress            string                         // if not empty will open up a grpc server for that address
	BlocksToSave             uint64
	AverageBlockTime         time.Duration // how often to query latest block
	ServerBlockMemory        uint64
	blocksCheckpointDistance uint64 // this causes the chainTracker to trigger it's checkpoint every X blocks
}

func (cnf *ChainTrackerConfig) validate() error {
	if cnf.BlocksToSave == 0 {
		return InvalidConfigErrorBlocksToSave
	}
	if cnf.AverageBlockTime == 0 {
		return InvalidConfigBlockTime
	}

	if cnf.ServerBlockMemory == 0 {
		cnf.ServerBlockMemory = DefualtAssumedBlockMemory
	}
	if cnf.blocksCheckpointDistance == 0 {
		cnf.blocksCheckpointDistance = DefaultBlockCheckpointDistance
	}
	// TODO: validate address is in the right format if not empty
	return nil
}
