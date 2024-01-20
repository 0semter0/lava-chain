package types

const (
	MAX_PROJECT_NAME_LEN = 50
	MAX_KEYS_AMOUNT      = 10
)

// set policy enum
type SetPolicyEnum int

const (
	SET_ADMIN_POLICY        SetPolicyEnum = 1
	SET_SUBSCRIPTION_POLICY SetPolicyEnum = 2
)

const (
	AddProjectKeyEventName         = "add_key_to_project_event"
	DelProjectKeyEventName         = "del_key_from_project_event"
	SetAdminPolicyEventName        = "set_admin_policy_event"
	SetSubscriptionPolicyEventName = "set_subscription_policy_event"
	ProjectResetFailEventName      = "project_reset_failed"
)
