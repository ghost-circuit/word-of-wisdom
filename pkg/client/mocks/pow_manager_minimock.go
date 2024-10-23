// Code generated by http://github.com/gojuno/minimock (v3.4.1). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/alisher-baizhumanov/word-of-wisdom/pkg/client.PowManager -o pow_manager_minimock.go -n PowManagerMock -p mocks

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// PowManagerMock implements mm_client.PowManager
type PowManagerMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcSolveCustomDifficulty          func(challenge []byte, difficulty uint8) (ba1 []byte, err error)
	funcSolveCustomDifficultyOrigin    string
	inspectFuncSolveCustomDifficulty   func(challenge []byte, difficulty uint8)
	afterSolveCustomDifficultyCounter  uint64
	beforeSolveCustomDifficultyCounter uint64
	SolveCustomDifficultyMock          mPowManagerMockSolveCustomDifficulty
}

// NewPowManagerMock returns a mock for mm_client.PowManager
func NewPowManagerMock(t minimock.Tester) *PowManagerMock {
	m := &PowManagerMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.SolveCustomDifficultyMock = mPowManagerMockSolveCustomDifficulty{mock: m}
	m.SolveCustomDifficultyMock.callArgs = []*PowManagerMockSolveCustomDifficultyParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mPowManagerMockSolveCustomDifficulty struct {
	optional           bool
	mock               *PowManagerMock
	defaultExpectation *PowManagerMockSolveCustomDifficultyExpectation
	expectations       []*PowManagerMockSolveCustomDifficultyExpectation

	callArgs []*PowManagerMockSolveCustomDifficultyParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// PowManagerMockSolveCustomDifficultyExpectation specifies expectation struct of the PowManager.SolveCustomDifficulty
type PowManagerMockSolveCustomDifficultyExpectation struct {
	mock               *PowManagerMock
	params             *PowManagerMockSolveCustomDifficultyParams
	paramPtrs          *PowManagerMockSolveCustomDifficultyParamPtrs
	expectationOrigins PowManagerMockSolveCustomDifficultyExpectationOrigins
	results            *PowManagerMockSolveCustomDifficultyResults
	returnOrigin       string
	Counter            uint64
}

// PowManagerMockSolveCustomDifficultyParams contains parameters of the PowManager.SolveCustomDifficulty
type PowManagerMockSolveCustomDifficultyParams struct {
	challenge  []byte
	difficulty uint8
}

// PowManagerMockSolveCustomDifficultyParamPtrs contains pointers to parameters of the PowManager.SolveCustomDifficulty
type PowManagerMockSolveCustomDifficultyParamPtrs struct {
	challenge  *[]byte
	difficulty *uint8
}

// PowManagerMockSolveCustomDifficultyResults contains results of the PowManager.SolveCustomDifficulty
type PowManagerMockSolveCustomDifficultyResults struct {
	ba1 []byte
	err error
}

// PowManagerMockSolveCustomDifficultyOrigins contains origins of expectations of the PowManager.SolveCustomDifficulty
type PowManagerMockSolveCustomDifficultyExpectationOrigins struct {
	origin           string
	originChallenge  string
	originDifficulty string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmSolveCustomDifficulty *mPowManagerMockSolveCustomDifficulty) Optional() *mPowManagerMockSolveCustomDifficulty {
	mmSolveCustomDifficulty.optional = true
	return mmSolveCustomDifficulty
}

// Expect sets up expected params for PowManager.SolveCustomDifficulty
func (mmSolveCustomDifficulty *mPowManagerMockSolveCustomDifficulty) Expect(challenge []byte, difficulty uint8) *mPowManagerMockSolveCustomDifficulty {
	if mmSolveCustomDifficulty.mock.funcSolveCustomDifficulty != nil {
		mmSolveCustomDifficulty.mock.t.Fatalf("PowManagerMock.SolveCustomDifficulty mock is already set by Set")
	}

	if mmSolveCustomDifficulty.defaultExpectation == nil {
		mmSolveCustomDifficulty.defaultExpectation = &PowManagerMockSolveCustomDifficultyExpectation{}
	}

	if mmSolveCustomDifficulty.defaultExpectation.paramPtrs != nil {
		mmSolveCustomDifficulty.mock.t.Fatalf("PowManagerMock.SolveCustomDifficulty mock is already set by ExpectParams functions")
	}

	mmSolveCustomDifficulty.defaultExpectation.params = &PowManagerMockSolveCustomDifficultyParams{challenge, difficulty}
	mmSolveCustomDifficulty.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmSolveCustomDifficulty.expectations {
		if minimock.Equal(e.params, mmSolveCustomDifficulty.defaultExpectation.params) {
			mmSolveCustomDifficulty.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSolveCustomDifficulty.defaultExpectation.params)
		}
	}

	return mmSolveCustomDifficulty
}

// ExpectChallengeParam1 sets up expected param challenge for PowManager.SolveCustomDifficulty
func (mmSolveCustomDifficulty *mPowManagerMockSolveCustomDifficulty) ExpectChallengeParam1(challenge []byte) *mPowManagerMockSolveCustomDifficulty {
	if mmSolveCustomDifficulty.mock.funcSolveCustomDifficulty != nil {
		mmSolveCustomDifficulty.mock.t.Fatalf("PowManagerMock.SolveCustomDifficulty mock is already set by Set")
	}

	if mmSolveCustomDifficulty.defaultExpectation == nil {
		mmSolveCustomDifficulty.defaultExpectation = &PowManagerMockSolveCustomDifficultyExpectation{}
	}

	if mmSolveCustomDifficulty.defaultExpectation.params != nil {
		mmSolveCustomDifficulty.mock.t.Fatalf("PowManagerMock.SolveCustomDifficulty mock is already set by Expect")
	}

	if mmSolveCustomDifficulty.defaultExpectation.paramPtrs == nil {
		mmSolveCustomDifficulty.defaultExpectation.paramPtrs = &PowManagerMockSolveCustomDifficultyParamPtrs{}
	}
	mmSolveCustomDifficulty.defaultExpectation.paramPtrs.challenge = &challenge
	mmSolveCustomDifficulty.defaultExpectation.expectationOrigins.originChallenge = minimock.CallerInfo(1)

	return mmSolveCustomDifficulty
}

// ExpectDifficultyParam2 sets up expected param difficulty for PowManager.SolveCustomDifficulty
func (mmSolveCustomDifficulty *mPowManagerMockSolveCustomDifficulty) ExpectDifficultyParam2(difficulty uint8) *mPowManagerMockSolveCustomDifficulty {
	if mmSolveCustomDifficulty.mock.funcSolveCustomDifficulty != nil {
		mmSolveCustomDifficulty.mock.t.Fatalf("PowManagerMock.SolveCustomDifficulty mock is already set by Set")
	}

	if mmSolveCustomDifficulty.defaultExpectation == nil {
		mmSolveCustomDifficulty.defaultExpectation = &PowManagerMockSolveCustomDifficultyExpectation{}
	}

	if mmSolveCustomDifficulty.defaultExpectation.params != nil {
		mmSolveCustomDifficulty.mock.t.Fatalf("PowManagerMock.SolveCustomDifficulty mock is already set by Expect")
	}

	if mmSolveCustomDifficulty.defaultExpectation.paramPtrs == nil {
		mmSolveCustomDifficulty.defaultExpectation.paramPtrs = &PowManagerMockSolveCustomDifficultyParamPtrs{}
	}
	mmSolveCustomDifficulty.defaultExpectation.paramPtrs.difficulty = &difficulty
	mmSolveCustomDifficulty.defaultExpectation.expectationOrigins.originDifficulty = minimock.CallerInfo(1)

	return mmSolveCustomDifficulty
}

// Inspect accepts an inspector function that has same arguments as the PowManager.SolveCustomDifficulty
func (mmSolveCustomDifficulty *mPowManagerMockSolveCustomDifficulty) Inspect(f func(challenge []byte, difficulty uint8)) *mPowManagerMockSolveCustomDifficulty {
	if mmSolveCustomDifficulty.mock.inspectFuncSolveCustomDifficulty != nil {
		mmSolveCustomDifficulty.mock.t.Fatalf("Inspect function is already set for PowManagerMock.SolveCustomDifficulty")
	}

	mmSolveCustomDifficulty.mock.inspectFuncSolveCustomDifficulty = f

	return mmSolveCustomDifficulty
}

// Return sets up results that will be returned by PowManager.SolveCustomDifficulty
func (mmSolveCustomDifficulty *mPowManagerMockSolveCustomDifficulty) Return(ba1 []byte, err error) *PowManagerMock {
	if mmSolveCustomDifficulty.mock.funcSolveCustomDifficulty != nil {
		mmSolveCustomDifficulty.mock.t.Fatalf("PowManagerMock.SolveCustomDifficulty mock is already set by Set")
	}

	if mmSolveCustomDifficulty.defaultExpectation == nil {
		mmSolveCustomDifficulty.defaultExpectation = &PowManagerMockSolveCustomDifficultyExpectation{mock: mmSolveCustomDifficulty.mock}
	}
	mmSolveCustomDifficulty.defaultExpectation.results = &PowManagerMockSolveCustomDifficultyResults{ba1, err}
	mmSolveCustomDifficulty.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmSolveCustomDifficulty.mock
}

// Set uses given function f to mock the PowManager.SolveCustomDifficulty method
func (mmSolveCustomDifficulty *mPowManagerMockSolveCustomDifficulty) Set(f func(challenge []byte, difficulty uint8) (ba1 []byte, err error)) *PowManagerMock {
	if mmSolveCustomDifficulty.defaultExpectation != nil {
		mmSolveCustomDifficulty.mock.t.Fatalf("Default expectation is already set for the PowManager.SolveCustomDifficulty method")
	}

	if len(mmSolveCustomDifficulty.expectations) > 0 {
		mmSolveCustomDifficulty.mock.t.Fatalf("Some expectations are already set for the PowManager.SolveCustomDifficulty method")
	}

	mmSolveCustomDifficulty.mock.funcSolveCustomDifficulty = f
	mmSolveCustomDifficulty.mock.funcSolveCustomDifficultyOrigin = minimock.CallerInfo(1)
	return mmSolveCustomDifficulty.mock
}

// When sets expectation for the PowManager.SolveCustomDifficulty which will trigger the result defined by the following
// Then helper
func (mmSolveCustomDifficulty *mPowManagerMockSolveCustomDifficulty) When(challenge []byte, difficulty uint8) *PowManagerMockSolveCustomDifficultyExpectation {
	if mmSolveCustomDifficulty.mock.funcSolveCustomDifficulty != nil {
		mmSolveCustomDifficulty.mock.t.Fatalf("PowManagerMock.SolveCustomDifficulty mock is already set by Set")
	}

	expectation := &PowManagerMockSolveCustomDifficultyExpectation{
		mock:               mmSolveCustomDifficulty.mock,
		params:             &PowManagerMockSolveCustomDifficultyParams{challenge, difficulty},
		expectationOrigins: PowManagerMockSolveCustomDifficultyExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmSolveCustomDifficulty.expectations = append(mmSolveCustomDifficulty.expectations, expectation)
	return expectation
}

// Then sets up PowManager.SolveCustomDifficulty return parameters for the expectation previously defined by the When method
func (e *PowManagerMockSolveCustomDifficultyExpectation) Then(ba1 []byte, err error) *PowManagerMock {
	e.results = &PowManagerMockSolveCustomDifficultyResults{ba1, err}
	return e.mock
}

// Times sets number of times PowManager.SolveCustomDifficulty should be invoked
func (mmSolveCustomDifficulty *mPowManagerMockSolveCustomDifficulty) Times(n uint64) *mPowManagerMockSolveCustomDifficulty {
	if n == 0 {
		mmSolveCustomDifficulty.mock.t.Fatalf("Times of PowManagerMock.SolveCustomDifficulty mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmSolveCustomDifficulty.expectedInvocations, n)
	mmSolveCustomDifficulty.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmSolveCustomDifficulty
}

func (mmSolveCustomDifficulty *mPowManagerMockSolveCustomDifficulty) invocationsDone() bool {
	if len(mmSolveCustomDifficulty.expectations) == 0 && mmSolveCustomDifficulty.defaultExpectation == nil && mmSolveCustomDifficulty.mock.funcSolveCustomDifficulty == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmSolveCustomDifficulty.mock.afterSolveCustomDifficultyCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmSolveCustomDifficulty.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// SolveCustomDifficulty implements mm_client.PowManager
func (mmSolveCustomDifficulty *PowManagerMock) SolveCustomDifficulty(challenge []byte, difficulty uint8) (ba1 []byte, err error) {
	mm_atomic.AddUint64(&mmSolveCustomDifficulty.beforeSolveCustomDifficultyCounter, 1)
	defer mm_atomic.AddUint64(&mmSolveCustomDifficulty.afterSolveCustomDifficultyCounter, 1)

	mmSolveCustomDifficulty.t.Helper()

	if mmSolveCustomDifficulty.inspectFuncSolveCustomDifficulty != nil {
		mmSolveCustomDifficulty.inspectFuncSolveCustomDifficulty(challenge, difficulty)
	}

	mm_params := PowManagerMockSolveCustomDifficultyParams{challenge, difficulty}

	// Record call args
	mmSolveCustomDifficulty.SolveCustomDifficultyMock.mutex.Lock()
	mmSolveCustomDifficulty.SolveCustomDifficultyMock.callArgs = append(mmSolveCustomDifficulty.SolveCustomDifficultyMock.callArgs, &mm_params)
	mmSolveCustomDifficulty.SolveCustomDifficultyMock.mutex.Unlock()

	for _, e := range mmSolveCustomDifficulty.SolveCustomDifficultyMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.ba1, e.results.err
		}
	}

	if mmSolveCustomDifficulty.SolveCustomDifficultyMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSolveCustomDifficulty.SolveCustomDifficultyMock.defaultExpectation.Counter, 1)
		mm_want := mmSolveCustomDifficulty.SolveCustomDifficultyMock.defaultExpectation.params
		mm_want_ptrs := mmSolveCustomDifficulty.SolveCustomDifficultyMock.defaultExpectation.paramPtrs

		mm_got := PowManagerMockSolveCustomDifficultyParams{challenge, difficulty}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.challenge != nil && !minimock.Equal(*mm_want_ptrs.challenge, mm_got.challenge) {
				mmSolveCustomDifficulty.t.Errorf("PowManagerMock.SolveCustomDifficulty got unexpected parameter challenge, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmSolveCustomDifficulty.SolveCustomDifficultyMock.defaultExpectation.expectationOrigins.originChallenge, *mm_want_ptrs.challenge, mm_got.challenge, minimock.Diff(*mm_want_ptrs.challenge, mm_got.challenge))
			}

			if mm_want_ptrs.difficulty != nil && !minimock.Equal(*mm_want_ptrs.difficulty, mm_got.difficulty) {
				mmSolveCustomDifficulty.t.Errorf("PowManagerMock.SolveCustomDifficulty got unexpected parameter difficulty, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmSolveCustomDifficulty.SolveCustomDifficultyMock.defaultExpectation.expectationOrigins.originDifficulty, *mm_want_ptrs.difficulty, mm_got.difficulty, minimock.Diff(*mm_want_ptrs.difficulty, mm_got.difficulty))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSolveCustomDifficulty.t.Errorf("PowManagerMock.SolveCustomDifficulty got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmSolveCustomDifficulty.SolveCustomDifficultyMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmSolveCustomDifficulty.SolveCustomDifficultyMock.defaultExpectation.results
		if mm_results == nil {
			mmSolveCustomDifficulty.t.Fatal("No results are set for the PowManagerMock.SolveCustomDifficulty")
		}
		return (*mm_results).ba1, (*mm_results).err
	}
	if mmSolveCustomDifficulty.funcSolveCustomDifficulty != nil {
		return mmSolveCustomDifficulty.funcSolveCustomDifficulty(challenge, difficulty)
	}
	mmSolveCustomDifficulty.t.Fatalf("Unexpected call to PowManagerMock.SolveCustomDifficulty. %v %v", challenge, difficulty)
	return
}

// SolveCustomDifficultyAfterCounter returns a count of finished PowManagerMock.SolveCustomDifficulty invocations
func (mmSolveCustomDifficulty *PowManagerMock) SolveCustomDifficultyAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSolveCustomDifficulty.afterSolveCustomDifficultyCounter)
}

// SolveCustomDifficultyBeforeCounter returns a count of PowManagerMock.SolveCustomDifficulty invocations
func (mmSolveCustomDifficulty *PowManagerMock) SolveCustomDifficultyBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSolveCustomDifficulty.beforeSolveCustomDifficultyCounter)
}

// Calls returns a list of arguments used in each call to PowManagerMock.SolveCustomDifficulty.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSolveCustomDifficulty *mPowManagerMockSolveCustomDifficulty) Calls() []*PowManagerMockSolveCustomDifficultyParams {
	mmSolveCustomDifficulty.mutex.RLock()

	argCopy := make([]*PowManagerMockSolveCustomDifficultyParams, len(mmSolveCustomDifficulty.callArgs))
	copy(argCopy, mmSolveCustomDifficulty.callArgs)

	mmSolveCustomDifficulty.mutex.RUnlock()

	return argCopy
}

// MinimockSolveCustomDifficultyDone returns true if the count of the SolveCustomDifficulty invocations corresponds
// the number of defined expectations
func (m *PowManagerMock) MinimockSolveCustomDifficultyDone() bool {
	if m.SolveCustomDifficultyMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.SolveCustomDifficultyMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.SolveCustomDifficultyMock.invocationsDone()
}

// MinimockSolveCustomDifficultyInspect logs each unmet expectation
func (m *PowManagerMock) MinimockSolveCustomDifficultyInspect() {
	for _, e := range m.SolveCustomDifficultyMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to PowManagerMock.SolveCustomDifficulty at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterSolveCustomDifficultyCounter := mm_atomic.LoadUint64(&m.afterSolveCustomDifficultyCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.SolveCustomDifficultyMock.defaultExpectation != nil && afterSolveCustomDifficultyCounter < 1 {
		if m.SolveCustomDifficultyMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to PowManagerMock.SolveCustomDifficulty at\n%s", m.SolveCustomDifficultyMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to PowManagerMock.SolveCustomDifficulty at\n%s with params: %#v", m.SolveCustomDifficultyMock.defaultExpectation.expectationOrigins.origin, *m.SolveCustomDifficultyMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSolveCustomDifficulty != nil && afterSolveCustomDifficultyCounter < 1 {
		m.t.Errorf("Expected call to PowManagerMock.SolveCustomDifficulty at\n%s", m.funcSolveCustomDifficultyOrigin)
	}

	if !m.SolveCustomDifficultyMock.invocationsDone() && afterSolveCustomDifficultyCounter > 0 {
		m.t.Errorf("Expected %d calls to PowManagerMock.SolveCustomDifficulty at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.SolveCustomDifficultyMock.expectedInvocations), m.SolveCustomDifficultyMock.expectedInvocationsOrigin, afterSolveCustomDifficultyCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *PowManagerMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockSolveCustomDifficultyInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *PowManagerMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *PowManagerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockSolveCustomDifficultyDone()
}