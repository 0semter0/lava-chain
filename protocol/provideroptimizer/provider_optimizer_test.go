package provideroptimizer

import (
	"strconv"
	"testing"
	"time"

	"github.com/lavanet/lava/utils"
	"github.com/lavanet/lava/utils/rand"
	spectypes "github.com/lavanet/lava/x/spec/types"
	"github.com/stretchr/testify/require"
)

const (
	TEST_AVERAGE_BLOCK_TIME = 10 * time.Second
	TEST_BASE_WORLD_LATENCY = 150 * time.Millisecond
)

func setupProviderOptimizer(maxProvidersCount int) *ProviderOptimizer {
	averageBlockTIme := TEST_AVERAGE_BLOCK_TIME
	baseWorldLatency := TEST_BASE_WORLD_LATENCY
	return NewProviderOptimizer(STRATEGY_BALANCED, averageBlockTIme, baseWorldLatency, 1)
}

type providersGenerator struct {
	providersAddresses []string
}

func (pg *providersGenerator) setupProvidersForTest(count int) *providersGenerator {
	pg.providersAddresses = make([]string, count)
	for i := range pg.providersAddresses {
		pg.providersAddresses[i] = "lava@test_" + strconv.Itoa(i)
	}
	return pg
}

func TestProbabilitiesCalculations(t *testing.T) {
	value := cumulativeProbabilityFunctionForPoissonDist(1, 10)
	value2 := cumulativeProbabilityFunctionForPoissonDist(10, 10)
	require.Greater(t, value2, value)
}

func TestProviderOptimizerSetGet(t *testing.T) {
	providerOptimizer := setupProviderOptimizer(1)
	providersGen := (&providersGenerator{}).setupProvidersForTest(1)
	providerAddress := providersGen.providersAddresses[0]
	for i := 0; i < 100; i++ {
		providerData := ProviderData{SyncBlock: uint64(i)}
		address := providerAddress + strconv.Itoa(i)
		set := providerOptimizer.providersStorage.Set(address, providerData, 1)
		if set == false {
			utils.LavaFormatWarning("set in cache dropped", nil)
		}
	}
	time.Sleep(4 * time.Millisecond)
	for i := 0; i < 100; i++ {
		address := providerAddress + strconv.Itoa(i)
		providerData, found := providerOptimizer.getProviderData(address)
		require.Equal(t, uint64(i), providerData.SyncBlock, "failed getting entry %s", address)
		require.True(t, found)
	}
}

func TestProviderOptimizerBasic(t *testing.T) {
	providerOptimizer := setupProviderOptimizer(1)
	providersGen := (&providersGenerator{}).setupProvidersForTest(10)

	requestCU := uint64(10)
	requestBlock := int64(1000)
	pertrubationPercentage := 0.0

	returnedProviders := providerOptimizer.ChooseProvider(providersGen.providersAddresses, nil, requestCU, requestBlock, pertrubationPercentage)
	require.Equal(t, 1, len(returnedProviders))
	providerOptimizer.AppendProbeRelayData(providersGen.providersAddresses[1], TEST_BASE_WORLD_LATENCY*3, true)
	time.Sleep(4 * time.Millisecond)
	returnedProviders = providerOptimizer.ChooseProvider(providersGen.providersAddresses, nil, requestCU, requestBlock, pertrubationPercentage)
	require.Equal(t, 1, len(returnedProviders))
	require.NotEqual(t, returnedProviders[0], providersGen.providersAddresses[1]) // we shouldn't pick the wrong provider
	providerOptimizer.AppendProbeRelayData(providersGen.providersAddresses[0], TEST_BASE_WORLD_LATENCY/2, true)
	time.Sleep(4 * time.Millisecond)
	returnedProviders = providerOptimizer.ChooseProvider(providersGen.providersAddresses, nil, requestCU, requestBlock, pertrubationPercentage)
	require.Equal(t, 1, len(returnedProviders))
	require.Equal(t, providersGen.providersAddresses[0], returnedProviders[0]) // we should pick the best provider
}

func TestProviderOptimizerBasicRelayData(t *testing.T) {
	providerOptimizer := setupProviderOptimizer(1)
	providersGen := (&providersGenerator{}).setupProvidersForTest(10)

	requestCU := uint64(1)
	requestBlock := int64(1000)
	pertrubationPercentage := 0.0
	syncBlock := uint64(requestBlock)

	providerOptimizer.AppendRelayData(providersGen.providersAddresses[1], TEST_BASE_WORLD_LATENCY*4, false, requestCU, syncBlock)
	returnedProviders := providerOptimizer.ChooseProvider(providersGen.providersAddresses, nil, requestCU, requestBlock, pertrubationPercentage)
	require.Equal(t, 1, len(returnedProviders))
	require.NotEqual(t, returnedProviders[0], providersGen.providersAddresses[1]) // we shouldn't pick the wrong provider
	providerOptimizer.AppendRelayData(providersGen.providersAddresses[0], TEST_BASE_WORLD_LATENCY/4, false, requestCU, syncBlock)
	returnedProviders = providerOptimizer.ChooseProvider(providersGen.providersAddresses, nil, requestCU, requestBlock, pertrubationPercentage)
	require.Equal(t, 1, len(returnedProviders))
	require.Equal(t, providersGen.providersAddresses[0], returnedProviders[0]) // we should pick the best provider
}

func TestProviderOptimizerAvailability(t *testing.T) {
	providerOptimizer := setupProviderOptimizer(1)
	providersCount := 100
	providersGen := (&providersGenerator{}).setupProvidersForTest(providersCount)

	requestCU := uint64(10)
	requestBlock := int64(1000)
	pertrubationPercentage := 0.0
	skipIndex := rand.Intn(providersCount)
	for i := range providersGen.providersAddresses {
		// give all providers a worse availability score
		if i == skipIndex {
			// skip 0
			continue
		}
		providerOptimizer.AppendProbeRelayData(providersGen.providersAddresses[i], TEST_BASE_WORLD_LATENCY, false)
	}
	time.Sleep(4 * time.Millisecond)
	returnedProviders := providerOptimizer.ChooseProvider(providersGen.providersAddresses, nil, requestCU, requestBlock, pertrubationPercentage)
	require.Equal(t, 1, len(returnedProviders))
	require.Equal(t, providersGen.providersAddresses[skipIndex], returnedProviders[0])
	returnedProviders = providerOptimizer.ChooseProvider(providersGen.providersAddresses, map[string]struct{}{providersGen.providersAddresses[skipIndex]: {}}, requestCU, requestBlock, pertrubationPercentage)
	require.Equal(t, 1, len(returnedProviders))
	require.NotEqual(t, providersGen.providersAddresses[skipIndex], returnedProviders[0])
}

func TestProviderOptimizerAvailabilityRelayData(t *testing.T) {
	providerOptimizer := setupProviderOptimizer(1)
	providersCount := 100
	providersGen := (&providersGenerator{}).setupProvidersForTest(providersCount)
	requestCU := uint64(10)
	requestBlock := int64(1000)
	pertrubationPercentage := 0.0
	skipIndex := rand.Intn(providersCount)
	for i := range providersGen.providersAddresses {
		// give all providers a worse availability score
		if i == skipIndex {
			// skip one provider
			continue
		}
		providerOptimizer.AppendRelayFailure(providersGen.providersAddresses[i])
	}
	time.Sleep(4 * time.Millisecond)
	returnedProviders := providerOptimizer.ChooseProvider(providersGen.providersAddresses, nil, requestCU, requestBlock, pertrubationPercentage)
	require.Equal(t, 1, len(returnedProviders))
	require.Equal(t, providersGen.providersAddresses[skipIndex], returnedProviders[0])
	returnedProviders = providerOptimizer.ChooseProvider(providersGen.providersAddresses, map[string]struct{}{providersGen.providersAddresses[skipIndex]: {}}, requestCU, requestBlock, pertrubationPercentage)
	require.Equal(t, 1, len(returnedProviders))
	require.NotEqual(t, providersGen.providersAddresses[skipIndex], returnedProviders[0])
}

func TestProviderOptimizerAvailabilityBlockError(t *testing.T) {
	providerOptimizer := setupProviderOptimizer(1)
	providersCount := 10
	providersGen := (&providersGenerator{}).setupProvidersForTest(providersCount)

	requestCU := uint64(10)
	requestBlock := int64(1000)
	pertrubationPercentage := 0.0
	syncBlock := uint64(requestBlock)
	chosenIndex := rand.Intn(providersCount)
	for i := range providersGen.providersAddresses {
		time.Sleep(4 * time.Millisecond)
		// give all providers a worse availability score
		if i == chosenIndex {
			// give better syncBlock, worse latency by a little
			providerOptimizer.AppendRelayData(providersGen.providersAddresses[i], TEST_BASE_WORLD_LATENCY+10*time.Millisecond, false, requestCU, syncBlock)
			continue
		}
		providerOptimizer.AppendRelayData(providersGen.providersAddresses[i], TEST_BASE_WORLD_LATENCY, false, requestCU, syncBlock-1) // update that he doesn't have the latest requested block
	}
	time.Sleep(4 * time.Millisecond)
	returnedProviders := providerOptimizer.ChooseProvider(providersGen.providersAddresses, nil, requestCU, requestBlock, pertrubationPercentage)
	require.Equal(t, 1, len(returnedProviders))
	require.Equal(t, providersGen.providersAddresses[chosenIndex], returnedProviders[0])
	// now try to get a previous block, our chosenIndex should be inferior in latency and blockError chance should be the same
	returnedProviders = providerOptimizer.ChooseProvider(providersGen.providersAddresses, nil, requestCU, requestBlock-1, pertrubationPercentage)
	require.Equal(t, 1, len(returnedProviders))
	require.NotEqual(t, providersGen.providersAddresses[chosenIndex], returnedProviders[0])
}

func TestProviderOptimizerUpdatingLatency(t *testing.T) {
	providerOptimizer := setupProviderOptimizer(1)
	providersCount := 2
	providersGen := (&providersGenerator{}).setupProvidersForTest(providersCount)
	providerAddress := providersGen.providersAddresses[0]
	requestCU := uint64(10)
	requestBlock := int64(1000)
	syncBlock := uint64(requestBlock)
	// in this test we are repeatedly adding better results, and latency score should improve
	for i := 0; i < 10; i++ {
		providerData, _ := providerOptimizer.getProviderData(providerAddress)
		currentLatencyScore := providerOptimizer.calculateLatencyScore(providerData, requestCU, requestBlock)
		providerOptimizer.AppendProbeRelayData(providerAddress, TEST_BASE_WORLD_LATENCY, true)
		time.Sleep(4 * time.Millisecond)
		providerData, found := providerOptimizer.getProviderData(providerAddress)
		require.True(t, found)
		newLatencyScore := providerOptimizer.calculateLatencyScore(providerData, requestCU, requestBlock)
		require.Greater(t, currentLatencyScore, newLatencyScore, i)
	}
	providerAddress = providersGen.providersAddresses[1]
	for i := 0; i < 10; i++ {
		providerData, _ := providerOptimizer.getProviderData(providerAddress)
		currentLatencyScore := providerOptimizer.calculateLatencyScore(providerData, requestCU, requestBlock)
		providerOptimizer.AppendRelayData(providerAddress, TEST_BASE_WORLD_LATENCY, false, requestCU, syncBlock)
		time.Sleep(4 * time.Millisecond)
		providerData, found := providerOptimizer.getProviderData(providerAddress)
		require.True(t, found)
		newLatencyScore := providerOptimizer.calculateLatencyScore(providerData, requestCU, requestBlock)
		require.Greater(t, currentLatencyScore, newLatencyScore, i)
	}
}

func TestProviderOptimizerStrategiesProviderCount(t *testing.T) {
	providerOptimizer := setupProviderOptimizer(3)
	providersCount := 5
	providersGen := (&providersGenerator{}).setupProvidersForTest(providersCount)
	requestCU := uint64(10)
	requestBlock := int64(1000)
	syncBlock := uint64(requestBlock)
	pertrubationPercentage := 0.0
	// set a basic state for all of them
	for i := 0; i < 10; i++ {
		for _, address := range providersGen.providersAddresses {
			providerOptimizer.AppendRelayData(address, TEST_BASE_WORLD_LATENCY*2, false, requestCU, syncBlock)
		}
		time.Sleep(4 * time.Millisecond)
	}
	testProvidersCount := func(iterations int) float64 {
		exploration := 0.0
		for i := 0; i < iterations; i++ {
			returnedProviders := providerOptimizer.ChooseProvider(providersGen.providersAddresses, nil, requestCU, requestBlock, pertrubationPercentage)
			if len(returnedProviders) > 1 {
				exploration++
			}
		}
		return exploration
	}

	// with a cost strategy we expect only one provider, two with a chance of 1/100
	providerOptimizer.strategy = STRATEGY_COST
	providerOptimizer.wantedNumProvidersInConcurrency = 2
	iterations := 10000
	exploration := testProvidersCount(iterations)
	require.Less(t, exploration, float64(1.3)*float64(iterations*providersCount)*COST_EXPLORATION_CHANCE) // allow mistake buffer of 30% because of randomness

	// with a cost strategy we expect only one provider, two with a chance of 10/100
	providerOptimizer.strategy = STRATEGY_BALANCED
	exploration = testProvidersCount(iterations)
	require.Greater(t, exploration, float64(1.3)*float64(iterations*providersCount)/100.0)
	require.Less(t, exploration, float64(1.3)*float64(iterations*providersCount)*DEFAULT_EXPLORATION_CHANCE) // allow mistake buffer of 30% because of randomness

	providerOptimizer.strategy = STRATEGY_PRIVACY
	exploration = testProvidersCount(iterations)
	require.Equal(t, exploration, float64(0))
}

func TestProviderOptimizerSyncScore(t *testing.T) {
	providerOptimizer := setupProviderOptimizer(1)
	providersGen := (&providersGenerator{}).setupProvidersForTest(10)

	requestCU := uint64(10)
	requestBlock := spectypes.LATEST_BLOCK
	pertrubationPercentage := 0.0
	syncBlock := uint64(1000)

	chosenIndex := rand.Intn(len(providersGen.providersAddresses))
	sampleTime := time.Now()
	for j := 0; j < 3; j++ { // repeat several times because a sync score is only correct after all providers sent their first block otherwise its giving favor to the first one
		for i := range providersGen.providersAddresses {
			time.Sleep(4 * time.Millisecond)
			if i == chosenIndex {
				// give better syncBlock, latency is a tiny bit worse for the second check
				providerOptimizer.appendRelayData(providersGen.providersAddresses[i], TEST_BASE_WORLD_LATENCY*2+1*time.Microsecond, false, true, requestCU, syncBlock+5, sampleTime)
				continue
			}
			providerOptimizer.appendRelayData(providersGen.providersAddresses[i], TEST_BASE_WORLD_LATENCY*2, false, true, requestCU, syncBlock, sampleTime) // update that he doesn't have the latest requested block
		}
		sampleTime = sampleTime.Add(time.Millisecond * 5)
	}
	time.Sleep(4 * time.Millisecond)
	returnedProviders := providerOptimizer.ChooseProvider(providersGen.providersAddresses, nil, requestCU, requestBlock, pertrubationPercentage)
	require.Equal(t, 1, len(returnedProviders))
	require.Equal(t, providersGen.providersAddresses[chosenIndex], returnedProviders[0]) // we should pick the best sync score

	returnedProviders = providerOptimizer.ChooseProvider(providersGen.providersAddresses, nil, requestCU, int64(syncBlock), pertrubationPercentage)
	require.Equal(t, 1, len(returnedProviders))
	require.NotEqual(t, providersGen.providersAddresses[chosenIndex], returnedProviders[0]) // sync score doesn't matter now
}

func TestProviderOptimizerStrategiesScoring(t *testing.T) {
	providerOptimizer := setupProviderOptimizer(1)
	providersCount := 5
	providersGen := (&providersGenerator{}).setupProvidersForTest(providersCount)
	requestCU := uint64(10)
	requestBlock := spectypes.LATEST_BLOCK
	syncBlock := uint64(1000)
	pertrubationPercentage := 0.0
	// set a basic state for all of them
	for i := 0; i < 10; i++ {
		for _, address := range providersGen.providersAddresses {
			providerOptimizer.AppendRelayData(address, TEST_BASE_WORLD_LATENCY*2, false, requestCU, syncBlock)
		}
		time.Sleep(4 * time.Millisecond)
	}
	time.Sleep(4 * time.Millisecond)
	// provider 2 doesn't get a probe availability hit, this is the most meaningful factor
	for idx, address := range providersGen.providersAddresses {
		if idx != 2 {
			providerOptimizer.AppendProbeRelayData(address, TEST_BASE_WORLD_LATENCY*2, false)
			time.Sleep(4 * time.Millisecond)
		}
		providerOptimizer.AppendProbeRelayData(address, TEST_BASE_WORLD_LATENCY*2, true)
		time.Sleep(4 * time.Millisecond)
		providerOptimizer.AppendProbeRelayData(address, TEST_BASE_WORLD_LATENCY*2, false)
		time.Sleep(4 * time.Millisecond)
		providerOptimizer.AppendProbeRelayData(address, TEST_BASE_WORLD_LATENCY*2, true)
		time.Sleep(4 * time.Millisecond)
		providerOptimizer.AppendProbeRelayData(address, TEST_BASE_WORLD_LATENCY*2, true)
		time.Sleep(4 * time.Millisecond)
	}

	sampleTime := time.Now()
	improvedLatency := 270 * time.Millisecond
	normalLatency := TEST_BASE_WORLD_LATENCY * 2
	improvedBlock := syncBlock + 1
	// provider 0 gets a good latency
	providerOptimizer.appendRelayData(providersGen.providersAddresses[0], improvedLatency, false, true, requestCU, syncBlock, sampleTime)

	// providers 3,4 get a regular entry
	providerOptimizer.appendRelayData(providersGen.providersAddresses[3], normalLatency, false, true, requestCU, syncBlock, sampleTime)
	providerOptimizer.appendRelayData(providersGen.providersAddresses[4], normalLatency, false, true, requestCU, syncBlock, sampleTime)

	// provider 1 gets a good sync
	providerOptimizer.appendRelayData(providersGen.providersAddresses[1], normalLatency, false, true, requestCU, improvedBlock, sampleTime)

	sampleTime = sampleTime.Add(10 * time.Millisecond)
	// now repeat to modify all providers scores across sync calculation
	providerOptimizer.appendRelayData(providersGen.providersAddresses[0], improvedLatency, false, true, requestCU, syncBlock, sampleTime)
	providerOptimizer.appendRelayData(providersGen.providersAddresses[3], normalLatency, false, true, requestCU, syncBlock, sampleTime)
	providerOptimizer.appendRelayData(providersGen.providersAddresses[4], normalLatency, false, true, requestCU, syncBlock, sampleTime)
	providerOptimizer.appendRelayData(providersGen.providersAddresses[1], normalLatency, false, true, requestCU, improvedBlock, sampleTime)

	time.Sleep(4 * time.Millisecond)
	providerOptimizer.strategy = STRATEGY_BALANCED
	// a balanced strategy should pick provider 2 because of it's high availability
	returnedProviders := providerOptimizer.ChooseProvider(providersGen.providersAddresses, nil, requestCU, requestBlock, pertrubationPercentage)
	require.Equal(t, 1, len(returnedProviders))
	require.Equal(t, providersGen.providersAddresses[2], returnedProviders[0])

	providerOptimizer.strategy = STRATEGY_COST
	// with a cost strategy we expect the same as balanced
	returnedProviders = providerOptimizer.ChooseProvider(providersGen.providersAddresses, nil, requestCU, requestBlock, pertrubationPercentage)
	require.Equal(t, 1, len(returnedProviders))
	require.Equal(t, providersGen.providersAddresses[2], returnedProviders[0])

	providerOptimizer.strategy = STRATEGY_LATENCY
	// latency strategy should pick the best latency
	returnedProviders = providerOptimizer.ChooseProvider(providersGen.providersAddresses, map[string]struct{}{providersGen.providersAddresses[2]: {}}, requestCU, requestBlock, pertrubationPercentage)
	require.Equal(t, 1, len(returnedProviders))
	require.Equal(t, providersGen.providersAddresses[0], returnedProviders[0])

	providerOptimizer.strategy = STRATEGY_SYNC_FRESHNESS
	// freshness strategy should pick the most advanced provider
	returnedProviders = providerOptimizer.ChooseProvider(providersGen.providersAddresses, map[string]struct{}{providersGen.providersAddresses[2]: {}}, requestCU, requestBlock, pertrubationPercentage)
	require.Equal(t, 1, len(returnedProviders))
	require.Equal(t, providersGen.providersAddresses[1], returnedProviders[0])
}

func TestPerturbation(t *testing.T) {
	origValue1 := 1.0
	origValue2 := 0.5
	pertrubationPercentage := 0.03 // this is statistical and we don;t want this failing
	runs := 100000
	success := 0
	for i := 0; i < runs; i++ {
		res1 := pertrubWithNormalGaussian(origValue1, pertrubationPercentage)
		res2 := pertrubWithNormalGaussian(origValue2, pertrubationPercentage)
		if res1 > res2 {
			success++
		}
	}
	require.GreaterOrEqual(t, float64(success), float64(runs)*0.9)
}

func TestProviderOptimizerPerturbation(t *testing.T) {
	providerOptimizer := setupProviderOptimizer(1)
	providersCount := 100
	providersGen := (&providersGenerator{}).setupProvidersForTest(providersCount)
	requestCU := uint64(10)
	requestBlock := spectypes.LATEST_BLOCK
	syncBlock := uint64(1000)
	pertrubationPercentage := 0.03 // this is statistical and we don;t want this failing
	// set a basic state for all of them
	sampleTime := time.Now()
	for i := 0; i < 10; i++ {
		for idx, address := range providersGen.providersAddresses {
			if idx < len(providersGen.providersAddresses)/2 {
				// first half are good
				providerOptimizer.appendRelayData(address, TEST_BASE_WORLD_LATENCY, false, true, requestCU, syncBlock, sampleTime)
			} else {
				// second half are bad
				providerOptimizer.appendRelayData(address, TEST_BASE_WORLD_LATENCY*10, false, true, requestCU, syncBlock, sampleTime)
			}
		}
		sampleTime = sampleTime.Add(time.Millisecond * 5)
	}
	seed := time.Now().UnixNano() // constant seed.
	rand.SetSpecificSeed(seed)
	utils.LavaFormatDebug("rand seed", utils.Attribute{Key: "seed", Value: seed})
	same := 0
	pickFaults := 0
	chosenProvider := providerOptimizer.ChooseProvider(providersGen.providersAddresses, nil, requestCU, requestBlock, 0)[0]
	runs := 1000
	for i := 0; i < runs; i++ {
		returnedProviders := providerOptimizer.ChooseProvider(providersGen.providersAddresses, nil, requestCU, requestBlock, pertrubationPercentage)
		require.Equal(t, 1, len(returnedProviders))
		if chosenProvider == returnedProviders[0] {
			same++
		}
		for idx, address := range providersGen.providersAddresses {
			if address == returnedProviders[0] {
				if idx > len(providersGen.providersAddresses)/2 {
					pickFaults++
				}
			}
		}
	}
	require.Less(t, float64(pickFaults), float64(runs)*0.01)
	require.Less(t, same, runs/2)
}

func TestExcellence(t *testing.T) {
	floatVal := 0.25
	dec := turnFloatToDec(floatVal, 8)
	floatNew, err := dec.Float64()
	require.NoError(t, err)
	require.Equal(t, floatVal, floatNew)

	providerOptimizer := setupProviderOptimizer(1)
	providersCount := 5
	providersGen := (&providersGenerator{}).setupProvidersForTest(providersCount)
	requestCU := uint64(10)
	syncBlock := uint64(1000)
	// set a basic state for all of them
	sampleTime := time.Now()
	for i := 0; i < 10; i++ {
		for _, address := range providersGen.providersAddresses {
			providerOptimizer.appendRelayData(address, TEST_BASE_WORLD_LATENCY*2, false, true, requestCU, syncBlock, sampleTime)
		}
		time.Sleep(4 * time.Millisecond)
	}
	report := providerOptimizer.GetExcellenceQoSReportForProvider(providersGen.providersAddresses[0])
	require.NotNil(t, report)
	report2 := providerOptimizer.GetExcellenceQoSReportForProvider(providersGen.providersAddresses[1])
	require.NotNil(t, report2)
	require.Equal(t, report, report2)
}
