package chainlib

import (
	"errors"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/lavanet/lava/protocol/chainlib/extensionslib"
	"github.com/lavanet/lava/utils"
	pairingtypes "github.com/lavanet/lava/x/pairing/types"
	spectypes "github.com/lavanet/lava/x/spec/types"
)

type BaseChainParser struct {
	taggedApis      map[spectypes.FUNCTION_TAG]TaggedContainer
	spec            spectypes.Spec
	rwLock          sync.RWMutex
	serverApis      map[ApiKey]ApiContainer
	apiCollections  map[CollectionKey]*spectypes.ApiCollection
	headers         map[ApiKey]*spectypes.Header
	verifications   map[VerificationKey][]VerificationContainer
	allowedAddons   map[string]struct{}
	extensionParser extensionslib.ExtensionParser
	active          bool
}

func (bcp *BaseChainParser) Activate() {
	bcp.active = true
}

func (bcp *BaseChainParser) Active() bool {
	return bcp.active
}

func (bcp *BaseChainParser) UpdateBlockTime(newBlockTime time.Duration) {
	bcp.rwLock.Lock()
	defer bcp.rwLock.Unlock()
	utils.LavaFormatInfo("chainParser updated block time", utils.Attribute{Key: "newTime", Value: newBlockTime}, utils.Attribute{Key: "oldTime", Value: time.Duration(bcp.spec.AverageBlockTime) * time.Millisecond}, utils.Attribute{Key: "specID", Value: bcp.spec})
	bcp.spec.AverageBlockTime = newBlockTime.Milliseconds()
}

func (bcp *BaseChainParser) HandleHeaders(metadata []pairingtypes.Metadata, apiCollection *spectypes.ApiCollection, headersDirection spectypes.Header_HeaderType) (filteredHeaders []pairingtypes.Metadata, overwriteRequestedBlock string, ignoredMetadata []pairingtypes.Metadata) {
	bcp.rwLock.RLock()
	defer bcp.rwLock.RUnlock()
	if len(metadata) == 0 {
		return []pairingtypes.Metadata{}, "", []pairingtypes.Metadata{}
	}
	retMetadata := []pairingtypes.Metadata{}
	for _, header := range metadata {
		headerName := strings.ToLower(header.Name)
		apiKey := ApiKey{Name: headerName, ConnectionType: apiCollection.CollectionData.Type}
		headerDirective, ok := bcp.headers[apiKey]
		if !ok {
			// this header is not handled
			continue
		}
		if headerDirective.Kind == headersDirection || headerDirective.Kind == spectypes.Header_pass_both {
			retMetadata = append(retMetadata, header)
			if headerDirective.FunctionTag == spectypes.FUNCTION_TAG_SET_LATEST_IN_METADATA {
				// this header sets the latest requested block
				overwriteRequestedBlock = header.Value
			}
		} else if headerDirective.Kind == spectypes.Header_pass_ignore {
			ignoredMetadata = append(ignoredMetadata, header)
		}
	}
	return retMetadata, overwriteRequestedBlock, ignoredMetadata
}

func (bcp *BaseChainParser) isAddon(addon string) bool {
	_, ok := bcp.allowedAddons[addon]
	return ok
}

func (bcp *BaseChainParser) isExtension(extension string) bool {
	return bcp.extensionParser.AllowedExtension(extension)
}

func (bcp *BaseChainParser) SetConfiguredExtensions(extensions map[string]struct{}) error {
	bcp.rwLock.Lock()
	defer bcp.rwLock.Unlock()
	for extension := range extensions {
		if !bcp.isExtension(extension) {
			return utils.LavaFormatError("configured extensions contain an extension that is not allowed in the spec", nil, utils.Attribute{Key: "spec", Value: bcp.spec.Index}, utils.Attribute{Key: "configured extensions", Value: extensions}, utils.Attribute{Key: "allowed extensions", Value: bcp.extensionParser.AllowedExtensions})
		}
	}
	// reset the current one in case we configured it previously
	configuredExtensions := make(map[extensionslib.ExtensionKey]*spectypes.Extension)
	for collectionKey, apiCollection := range bcp.apiCollections {
		for _, extension := range apiCollection.Extensions {
			if extension.Name == "" {
				// skip empty extensions
				continue
			}
			if _, ok := extensions[extension.Name]; ok {
				extensionKey := extensionslib.ExtensionKey{
					Extension:      extension.Name,
					ConnectionType: collectionKey.ConnectionType,
					InternalPath:   collectionKey.InternalPath,
					Addon:          collectionKey.Addon,
				}
				configuredExtensions[extensionKey] = extension
			}
		}
	}
	bcp.extensionParser.SetConfiguredExtensions(configuredExtensions)
	return nil
}

// this function errors if it meets a value that is neither a n addon or an extension
func (bcp *BaseChainParser) SeparateAddonsExtensions(supported []string) (addons, extensions []string, err error) {
	checked := map[string]struct{}{}
	for _, supportedToCheck := range supported {
		// ignore repeated occurrences
		if _, ok := checked[supportedToCheck]; ok {
			continue
		}
		checked[supportedToCheck] = struct{}{}

		if bcp.isAddon(supportedToCheck) {
			addons = append(addons, supportedToCheck)
		} else {
			if supportedToCheck == "" {
				continue
			}
			if !bcp.isExtension(supportedToCheck) {
				// neither is an error
				return nil, nil, utils.LavaFormatError("invalid supported to check, is neither an addon or an extension", nil, utils.Attribute{Key: "spec", Value: bcp.spec.Index}, utils.Attribute{Key: "supported", Value: supportedToCheck})
			}
			extensions = append(extensions, supportedToCheck)
		}
	}
	return addons, extensions, nil
}

// gets all verifications for an endpoint supporting multiple addons and extensions
func (bcp *BaseChainParser) GetVerifications(supported []string) (retVerifications []VerificationContainer, err error) {
	// addons will contains extensions and addons,
	// extensions must exist in all verifications, addons must be split because they are separated
	addons, extensions, err := bcp.SeparateAddonsExtensions(supported)
	if err != nil {
		return nil, err
	}
	if len(extensions) == 0 {
		extensions = []string{""}
	}
	addons = append(addons, "") // always add the empty addon
	for _, addon := range addons {
		for _, extension := range extensions {
			verificationKey := VerificationKey{
				Extension: extension,
				Addon:     addon,
			}
			verifications, ok := bcp.verifications[verificationKey]
			if ok {
				retVerifications = append(retVerifications, verifications...)
			}
		}
	}
	return
}

func (bcp *BaseChainParser) Construct(spec spectypes.Spec, taggedApis map[spectypes.FUNCTION_TAG]TaggedContainer, serverApis map[ApiKey]ApiContainer, apiCollections map[CollectionKey]*spectypes.ApiCollection, headers map[ApiKey]*spectypes.Header, verifications map[VerificationKey][]VerificationContainer) {
	bcp.spec = spec
	bcp.serverApis = serverApis
	bcp.taggedApis = taggedApis
	bcp.headers = headers
	bcp.apiCollections = apiCollections
	bcp.verifications = verifications
	allowedAddons := map[string]struct{}{}
	allowedExtensions := map[string]struct{}{}
	for _, apoCollection := range apiCollections {
		for _, extension := range apoCollection.Extensions {
			allowedExtensions[extension.Name] = struct{}{}
		}
		allowedAddons[apoCollection.CollectionData.AddOn] = struct{}{}
	}
	bcp.allowedAddons = allowedAddons
	bcp.extensionParser = extensionslib.ExtensionParser{AllowedExtensions: allowedExtensions}
}

func (bcp *BaseChainParser) GetParsingByTag(tag spectypes.FUNCTION_TAG) (parsing *spectypes.ParseDirective, collectionData *spectypes.CollectionData, existed bool) {
	bcp.rwLock.RLock()
	defer bcp.rwLock.RUnlock()

	val, ok := bcp.taggedApis[tag]
	if !ok {
		return nil, nil, false
	}
	return val.Parsing, &val.ApiCollection.CollectionData, ok
}

func (bcp *BaseChainParser) ExtensionParsing(addon string, parsedMessageArg *parsedMessage, latestBlock uint64) {
	bcp.rwLock.RLock()
	defer bcp.rwLock.RUnlock()
	bcp.extensionParser.ExtensionParsing(addon, parsedMessageArg, latestBlock)
}

// getSupportedApi fetches service api from spec by name
func (apip *BaseChainParser) getSupportedApi(name, connectionType string) (*ApiContainer, error) {
	// Guard that the GrpcChainParser instance exists
	if apip == nil {
		return nil, errors.New("ChainParser not defined")
	}

	// Acquire read lock
	apip.rwLock.RLock()
	defer apip.rwLock.RUnlock()

	// Fetch server api by name
	apiCont, ok := apip.serverApis[ApiKey{
		Name:           name,
		ConnectionType: connectionType,
	}]

	// Return an error if spec does not exist
	if !ok {
		return nil, utils.LavaFormatError("api not supported", nil, utils.Attribute{Key: "name", Value: name}, utils.Attribute{Key: "connectionType", Value: connectionType})
	}

	// Return an error if api is disabled
	if !apiCont.api.Enabled {
		return nil, utils.LavaFormatError("api is disabled", nil, utils.Attribute{Key: "name", Value: name}, utils.Attribute{Key: "connectionType", Value: connectionType})
	}

	return &apiCont, nil
}

// getSupportedApi fetches service api from spec by name
func (apip *BaseChainParser) getApiCollection(connectionType, internalPath, addon string) (*spectypes.ApiCollection, error) {
	// Guard that the GrpcChainParser instance exists
	if apip == nil {
		return nil, errors.New("ChainParser not defined")
	}

	// Acquire read lock
	apip.rwLock.RLock()
	defer apip.rwLock.RUnlock()

	// Fetch server api by name
	api, ok := apip.apiCollections[CollectionKey{
		ConnectionType: connectionType,
		InternalPath:   internalPath,
		Addon:          addon,
	}]

	// Return an error if spec does not exist
	if !ok {
		return nil, utils.LavaFormatError("api not supported", nil, utils.Attribute{Key: "connectionType", Value: connectionType})
	}

	// Return an error if api is disabled
	if !api.Enabled {
		return nil, utils.LavaFormatError("api is disabled", nil, utils.Attribute{Key: "connectionType", Value: connectionType})
	}

	return api, nil
}

func getServiceApis(spec spectypes.Spec, rpcInterface string) (retServerApis map[ApiKey]ApiContainer, retTaggedApis map[spectypes.FUNCTION_TAG]TaggedContainer, retApiCollections map[CollectionKey]*spectypes.ApiCollection, retHeaders map[ApiKey]*spectypes.Header, retVerifications map[VerificationKey][]VerificationContainer) {
	serverApis := map[ApiKey]ApiContainer{}
	taggedApis := map[spectypes.FUNCTION_TAG]TaggedContainer{}
	headers := map[ApiKey]*spectypes.Header{}
	apiCollections := map[CollectionKey]*spectypes.ApiCollection{}
	verifications := map[VerificationKey][]VerificationContainer{}
	if spec.Enabled {
		for _, apiCollection := range spec.ApiCollections {
			if !apiCollection.Enabled {
				continue
			}
			if apiCollection.CollectionData.ApiInterface != rpcInterface {
				continue
			}
			collectionKey := CollectionKey{
				ConnectionType: apiCollection.CollectionData.Type,
				InternalPath:   apiCollection.CollectionData.InternalPath,
				Addon:          apiCollection.CollectionData.AddOn,
			}
			for _, parsing := range apiCollection.ParseDirectives {
				taggedApis[parsing.FunctionTag] = TaggedContainer{
					Parsing:       parsing,
					ApiCollection: apiCollection,
				}
			}

			for _, api := range apiCollection.Apis {
				if !api.Enabled {
					continue
				}
				//
				// TODO: find a better spot for this (more optimized, precompile regex, etc)
				if rpcInterface == spectypes.APIInterfaceRest {
					re := regexp.MustCompile(`{[^}]+}`)
					processedName := string(re.ReplaceAll([]byte(api.Name), []byte("replace-me-with-regex")))
					processedName = regexp.QuoteMeta(processedName)
					processedName = strings.ReplaceAll(processedName, "replace-me-with-regex", `[^\/\s]+`)
					serverApis[ApiKey{
						Name:           processedName,
						ConnectionType: collectionKey.ConnectionType,
					}] = ApiContainer{
						api:           api,
						collectionKey: collectionKey,
					}
				} else {
					serverApis[ApiKey{
						Name:           api.Name,
						ConnectionType: collectionKey.ConnectionType,
					}] = ApiContainer{
						api:           api,
						collectionKey: collectionKey,
					}
				}
			}
			for _, header := range apiCollection.Headers {
				headers[ApiKey{
					Name:           header.Name,
					ConnectionType: collectionKey.ConnectionType,
				}] = header
			}
			for _, verification := range apiCollection.Verifications {
				for _, parseValue := range verification.Values {
					verificationKey := VerificationKey{
						Extension: parseValue.Extension,
						Addon:     apiCollection.CollectionData.AddOn,
					}

					verCont := VerificationContainer{
						ConnectionType:  apiCollection.CollectionData.Type,
						Name:            verification.Name,
						ParseDirective:  *verification.ParseDirective,
						Value:           parseValue.ExpectedValue,
						LatestDistance:  parseValue.LatestDistance,
						VerificationKey: verificationKey,
					}

					if extensionVerifications, ok := verifications[verificationKey]; !ok {
						verifications[verificationKey] = []VerificationContainer{verCont}
					} else {
						verifications[verificationKey] = append(extensionVerifications, verCont)
					}
				}
			}
			apiCollections[collectionKey] = apiCollection
		}
	}
	return serverApis, taggedApis, apiCollections, headers, verifications
}

// matchSpecApiByName returns service api which match given name
func matchSpecApiByName(name, connectionType string, serverApis map[ApiKey]ApiContainer) (*ApiContainer, bool) {
	// TODO: make it faster and better by not doing a regex instead using a better algorithm
	foundNameOnDifferentConnectionType := ""
	for apiName, api := range serverApis {
		re, err := regexp.Compile("^" + apiName.Name + "$")
		if err != nil {
			utils.LavaFormatError("regex Compile api", err, utils.Attribute{Key: "apiName", Value: apiName})
			continue
		}
		if re.MatchString(name) {
			if apiName.ConnectionType == connectionType {
				return &api, true
			} else {
				foundNameOnDifferentConnectionType = apiName.ConnectionType
			}
		}
	}
	if foundNameOnDifferentConnectionType != "" { // its hard to notice when we have an API on only one connection type.
		utils.LavaFormatWarning("API was found on a different connection type", nil,
			utils.Attribute{Key: "connection_type_found", Value: foundNameOnDifferentConnectionType},
			utils.Attribute{Key: "connection_type_requested", Value: connectionType},
		)
	}
	return nil, false
}
