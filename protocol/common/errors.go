package common

import sdkerrors "cosmossdk.io/errors"

var (
	ContextDeadlineExceededError                = sdkerrors.New("ContextDeadlineExceeded Error", 300, "context deadline exceeded")
	StatusCodeError504                          = sdkerrors.New("Disallowed StatusCode Error", 504, "Disallowed status code error")
	StatusCodeError429                          = sdkerrors.New("Disallowed StatusCode Error", 429, "Disallowed status code error")
	StatusCodeErrorStrict                       = sdkerrors.New("Disallowed StatusCode Error", 800, "Disallowed status code error")
	APINotSupportedError                        = sdkerrors.New("APINotSupported Error", 900, "api not supported")
	ProviderFinalizationDataError               = sdkerrors.New("ProviderFinalizationData Error", 3365, "provider did not sign finalization data correctly")
	ProviderFinalizationDataAccountabilityError = sdkerrors.New("ProviderFinalizationDataAccountability Error", 3366, "provider returned invalid finalization data, with accountability")
)
