package types

const (
	ProviderStakeEventName       = "stake_new_provider"
	ProviderStakeUpdateEventName = "stake_update_provider"
	ProviderUnstakeEventName     = "provider_unstake_commit"

	RelayPaymentEventName       = "relay_payment"
	ProviderJailedEventName     = "provider_jailed"
	ProviderReportedEventName   = "provider_reported"
	LatestBlocksReportEventName = "provider_latest_block_report"
	RejectedCuEventName         = "rejected_cu"
	UnstakeProposalEventName    = "unstake_gov_proposal"
)

// unstake description strings
const (
	UnstakeDescriptionClientUnstake     = "Client unstaked entry"
	UnstakeDescriptionProviderUnstake   = "Provider unstaked entry"
	UnstakeDescriptionInsufficientFunds = "client stake is below the minimum stake required"
)

const (
	FlagMoniker                  = "provider-moniker"
	FlagCommission               = "delegate-commission"
	FlagDelegationLimit          = "delegate-limit"
	FlagOperator                 = "operator"
	MAX_LEN_MONIKER              = 50
	MAX_ENDPOINTS_AMOUNT_PER_GEO = 5 // max number of endpoints per geolocation for provider stake entry
)

// unresponsiveness consts
const (
	// Consider changing back on mainnet when providers QoS benchmarks are better // EPOCHS_NUM_TO_CHECK_CU_FOR_UNRESPONSIVE_PROVIDER uint64 = 4 // number of epochs to sum CU that the provider serviced
	EPOCHS_NUM_TO_CHECK_CU_FOR_UNRESPONSIVE_PROVIDER uint64 = 8 // number of epochs to sum CU that the provider serviced
	EPOCHS_NUM_TO_CHECK_FOR_COMPLAINERS              uint64 = 2 // number of epochs to sum CU of complainers against the provider
)

type ClientUsedCU struct {
	TotalUsed uint64
	Providers map[string]uint64
}

type ClientProviderOverusedCUPercent struct {
	TotalOverusedPercent    float64
	OverusedPercentProvider float64
}
