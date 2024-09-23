// Code generated by http://github.com/gojuno/minimock (v3.4.0). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/asp3cto/auth/internal/repository.UserRepository -o user_repository_minimock.go -n UserRepositoryMock -p mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/asp3cto/auth/internal/model"
	"github.com/gojuno/minimock/v3"
)

// UserRepositoryMock implements mm_repository.UserRepository
type UserRepositoryMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcCreate          func(ctx context.Context, info *model.UserInfo) (i1 int64, err error)
	funcCreateOrigin    string
	inspectFuncCreate   func(ctx context.Context, info *model.UserInfo)
	afterCreateCounter  uint64
	beforeCreateCounter uint64
	CreateMock          mUserRepositoryMockCreate

	funcGet          func(ctx context.Context, id int64) (up1 *model.User, err error)
	funcGetOrigin    string
	inspectFuncGet   func(ctx context.Context, id int64)
	afterGetCounter  uint64
	beforeGetCounter uint64
	GetMock          mUserRepositoryMockGet
}

// NewUserRepositoryMock returns a mock for mm_repository.UserRepository
func NewUserRepositoryMock(t minimock.Tester) *UserRepositoryMock {
	m := &UserRepositoryMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateMock = mUserRepositoryMockCreate{mock: m}
	m.CreateMock.callArgs = []*UserRepositoryMockCreateParams{}

	m.GetMock = mUserRepositoryMockGet{mock: m}
	m.GetMock.callArgs = []*UserRepositoryMockGetParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mUserRepositoryMockCreate struct {
	optional           bool
	mock               *UserRepositoryMock
	defaultExpectation *UserRepositoryMockCreateExpectation
	expectations       []*UserRepositoryMockCreateExpectation

	callArgs []*UserRepositoryMockCreateParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// UserRepositoryMockCreateExpectation specifies expectation struct of the UserRepository.Create
type UserRepositoryMockCreateExpectation struct {
	mock               *UserRepositoryMock
	params             *UserRepositoryMockCreateParams
	paramPtrs          *UserRepositoryMockCreateParamPtrs
	expectationOrigins UserRepositoryMockCreateExpectationOrigins
	results            *UserRepositoryMockCreateResults
	returnOrigin       string
	Counter            uint64
}

// UserRepositoryMockCreateParams contains parameters of the UserRepository.Create
type UserRepositoryMockCreateParams struct {
	ctx  context.Context
	info *model.UserInfo
}

// UserRepositoryMockCreateParamPtrs contains pointers to parameters of the UserRepository.Create
type UserRepositoryMockCreateParamPtrs struct {
	ctx  *context.Context
	info **model.UserInfo
}

// UserRepositoryMockCreateResults contains results of the UserRepository.Create
type UserRepositoryMockCreateResults struct {
	i1  int64
	err error
}

// UserRepositoryMockCreateOrigins contains origins of expectations of the UserRepository.Create
type UserRepositoryMockCreateExpectationOrigins struct {
	origin     string
	originCtx  string
	originInfo string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmCreate *mUserRepositoryMockCreate) Optional() *mUserRepositoryMockCreate {
	mmCreate.optional = true
	return mmCreate
}

// Expect sets up expected params for UserRepository.Create
func (mmCreate *mUserRepositoryMockCreate) Expect(ctx context.Context, info *model.UserInfo) *mUserRepositoryMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("UserRepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &UserRepositoryMockCreateExpectation{}
	}

	if mmCreate.defaultExpectation.paramPtrs != nil {
		mmCreate.mock.t.Fatalf("UserRepositoryMock.Create mock is already set by ExpectParams functions")
	}

	mmCreate.defaultExpectation.params = &UserRepositoryMockCreateParams{ctx, info}
	mmCreate.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmCreate.expectations {
		if minimock.Equal(e.params, mmCreate.defaultExpectation.params) {
			mmCreate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreate.defaultExpectation.params)
		}
	}

	return mmCreate
}

// ExpectCtxParam1 sets up expected param ctx for UserRepository.Create
func (mmCreate *mUserRepositoryMockCreate) ExpectCtxParam1(ctx context.Context) *mUserRepositoryMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("UserRepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &UserRepositoryMockCreateExpectation{}
	}

	if mmCreate.defaultExpectation.params != nil {
		mmCreate.mock.t.Fatalf("UserRepositoryMock.Create mock is already set by Expect")
	}

	if mmCreate.defaultExpectation.paramPtrs == nil {
		mmCreate.defaultExpectation.paramPtrs = &UserRepositoryMockCreateParamPtrs{}
	}
	mmCreate.defaultExpectation.paramPtrs.ctx = &ctx
	mmCreate.defaultExpectation.expectationOrigins.originCtx = minimock.CallerInfo(1)

	return mmCreate
}

// ExpectInfoParam2 sets up expected param info for UserRepository.Create
func (mmCreate *mUserRepositoryMockCreate) ExpectInfoParam2(info *model.UserInfo) *mUserRepositoryMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("UserRepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &UserRepositoryMockCreateExpectation{}
	}

	if mmCreate.defaultExpectation.params != nil {
		mmCreate.mock.t.Fatalf("UserRepositoryMock.Create mock is already set by Expect")
	}

	if mmCreate.defaultExpectation.paramPtrs == nil {
		mmCreate.defaultExpectation.paramPtrs = &UserRepositoryMockCreateParamPtrs{}
	}
	mmCreate.defaultExpectation.paramPtrs.info = &info
	mmCreate.defaultExpectation.expectationOrigins.originInfo = minimock.CallerInfo(1)

	return mmCreate
}

// Inspect accepts an inspector function that has same arguments as the UserRepository.Create
func (mmCreate *mUserRepositoryMockCreate) Inspect(f func(ctx context.Context, info *model.UserInfo)) *mUserRepositoryMockCreate {
	if mmCreate.mock.inspectFuncCreate != nil {
		mmCreate.mock.t.Fatalf("Inspect function is already set for UserRepositoryMock.Create")
	}

	mmCreate.mock.inspectFuncCreate = f

	return mmCreate
}

// Return sets up results that will be returned by UserRepository.Create
func (mmCreate *mUserRepositoryMockCreate) Return(i1 int64, err error) *UserRepositoryMock {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("UserRepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &UserRepositoryMockCreateExpectation{mock: mmCreate.mock}
	}
	mmCreate.defaultExpectation.results = &UserRepositoryMockCreateResults{i1, err}
	mmCreate.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmCreate.mock
}

// Set uses given function f to mock the UserRepository.Create method
func (mmCreate *mUserRepositoryMockCreate) Set(f func(ctx context.Context, info *model.UserInfo) (i1 int64, err error)) *UserRepositoryMock {
	if mmCreate.defaultExpectation != nil {
		mmCreate.mock.t.Fatalf("Default expectation is already set for the UserRepository.Create method")
	}

	if len(mmCreate.expectations) > 0 {
		mmCreate.mock.t.Fatalf("Some expectations are already set for the UserRepository.Create method")
	}

	mmCreate.mock.funcCreate = f
	mmCreate.mock.funcCreateOrigin = minimock.CallerInfo(1)
	return mmCreate.mock
}

// When sets expectation for the UserRepository.Create which will trigger the result defined by the following
// Then helper
func (mmCreate *mUserRepositoryMockCreate) When(ctx context.Context, info *model.UserInfo) *UserRepositoryMockCreateExpectation {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("UserRepositoryMock.Create mock is already set by Set")
	}

	expectation := &UserRepositoryMockCreateExpectation{
		mock:               mmCreate.mock,
		params:             &UserRepositoryMockCreateParams{ctx, info},
		expectationOrigins: UserRepositoryMockCreateExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmCreate.expectations = append(mmCreate.expectations, expectation)
	return expectation
}

// Then sets up UserRepository.Create return parameters for the expectation previously defined by the When method
func (e *UserRepositoryMockCreateExpectation) Then(i1 int64, err error) *UserRepositoryMock {
	e.results = &UserRepositoryMockCreateResults{i1, err}
	return e.mock
}

// Times sets number of times UserRepository.Create should be invoked
func (mmCreate *mUserRepositoryMockCreate) Times(n uint64) *mUserRepositoryMockCreate {
	if n == 0 {
		mmCreate.mock.t.Fatalf("Times of UserRepositoryMock.Create mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmCreate.expectedInvocations, n)
	mmCreate.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmCreate
}

func (mmCreate *mUserRepositoryMockCreate) invocationsDone() bool {
	if len(mmCreate.expectations) == 0 && mmCreate.defaultExpectation == nil && mmCreate.mock.funcCreate == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmCreate.mock.afterCreateCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmCreate.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Create implements mm_repository.UserRepository
func (mmCreate *UserRepositoryMock) Create(ctx context.Context, info *model.UserInfo) (i1 int64, err error) {
	mm_atomic.AddUint64(&mmCreate.beforeCreateCounter, 1)
	defer mm_atomic.AddUint64(&mmCreate.afterCreateCounter, 1)

	mmCreate.t.Helper()

	if mmCreate.inspectFuncCreate != nil {
		mmCreate.inspectFuncCreate(ctx, info)
	}

	mm_params := UserRepositoryMockCreateParams{ctx, info}

	// Record call args
	mmCreate.CreateMock.mutex.Lock()
	mmCreate.CreateMock.callArgs = append(mmCreate.CreateMock.callArgs, &mm_params)
	mmCreate.CreateMock.mutex.Unlock()

	for _, e := range mmCreate.CreateMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.i1, e.results.err
		}
	}

	if mmCreate.CreateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreate.CreateMock.defaultExpectation.Counter, 1)
		mm_want := mmCreate.CreateMock.defaultExpectation.params
		mm_want_ptrs := mmCreate.CreateMock.defaultExpectation.paramPtrs

		mm_got := UserRepositoryMockCreateParams{ctx, info}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmCreate.t.Errorf("UserRepositoryMock.Create got unexpected parameter ctx, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmCreate.CreateMock.defaultExpectation.expectationOrigins.originCtx, *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.info != nil && !minimock.Equal(*mm_want_ptrs.info, mm_got.info) {
				mmCreate.t.Errorf("UserRepositoryMock.Create got unexpected parameter info, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmCreate.CreateMock.defaultExpectation.expectationOrigins.originInfo, *mm_want_ptrs.info, mm_got.info, minimock.Diff(*mm_want_ptrs.info, mm_got.info))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreate.t.Errorf("UserRepositoryMock.Create got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmCreate.CreateMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreate.CreateMock.defaultExpectation.results
		if mm_results == nil {
			mmCreate.t.Fatal("No results are set for the UserRepositoryMock.Create")
		}
		return (*mm_results).i1, (*mm_results).err
	}
	if mmCreate.funcCreate != nil {
		return mmCreate.funcCreate(ctx, info)
	}
	mmCreate.t.Fatalf("Unexpected call to UserRepositoryMock.Create. %v %v", ctx, info)
	return
}

// CreateAfterCounter returns a count of finished UserRepositoryMock.Create invocations
func (mmCreate *UserRepositoryMock) CreateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.afterCreateCounter)
}

// CreateBeforeCounter returns a count of UserRepositoryMock.Create invocations
func (mmCreate *UserRepositoryMock) CreateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.beforeCreateCounter)
}

// Calls returns a list of arguments used in each call to UserRepositoryMock.Create.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreate *mUserRepositoryMockCreate) Calls() []*UserRepositoryMockCreateParams {
	mmCreate.mutex.RLock()

	argCopy := make([]*UserRepositoryMockCreateParams, len(mmCreate.callArgs))
	copy(argCopy, mmCreate.callArgs)

	mmCreate.mutex.RUnlock()

	return argCopy
}

// MinimockCreateDone returns true if the count of the Create invocations corresponds
// the number of defined expectations
func (m *UserRepositoryMock) MinimockCreateDone() bool {
	if m.CreateMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.CreateMock.invocationsDone()
}

// MinimockCreateInspect logs each unmet expectation
func (m *UserRepositoryMock) MinimockCreateInspect() {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserRepositoryMock.Create at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterCreateCounter := mm_atomic.LoadUint64(&m.afterCreateCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && afterCreateCounter < 1 {
		if m.CreateMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to UserRepositoryMock.Create at\n%s", m.CreateMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to UserRepositoryMock.Create at\n%s with params: %#v", m.CreateMock.defaultExpectation.expectationOrigins.origin, *m.CreateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && afterCreateCounter < 1 {
		m.t.Errorf("Expected call to UserRepositoryMock.Create at\n%s", m.funcCreateOrigin)
	}

	if !m.CreateMock.invocationsDone() && afterCreateCounter > 0 {
		m.t.Errorf("Expected %d calls to UserRepositoryMock.Create at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.CreateMock.expectedInvocations), m.CreateMock.expectedInvocationsOrigin, afterCreateCounter)
	}
}

type mUserRepositoryMockGet struct {
	optional           bool
	mock               *UserRepositoryMock
	defaultExpectation *UserRepositoryMockGetExpectation
	expectations       []*UserRepositoryMockGetExpectation

	callArgs []*UserRepositoryMockGetParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// UserRepositoryMockGetExpectation specifies expectation struct of the UserRepository.Get
type UserRepositoryMockGetExpectation struct {
	mock               *UserRepositoryMock
	params             *UserRepositoryMockGetParams
	paramPtrs          *UserRepositoryMockGetParamPtrs
	expectationOrigins UserRepositoryMockGetExpectationOrigins
	results            *UserRepositoryMockGetResults
	returnOrigin       string
	Counter            uint64
}

// UserRepositoryMockGetParams contains parameters of the UserRepository.Get
type UserRepositoryMockGetParams struct {
	ctx context.Context
	id  int64
}

// UserRepositoryMockGetParamPtrs contains pointers to parameters of the UserRepository.Get
type UserRepositoryMockGetParamPtrs struct {
	ctx *context.Context
	id  *int64
}

// UserRepositoryMockGetResults contains results of the UserRepository.Get
type UserRepositoryMockGetResults struct {
	up1 *model.User
	err error
}

// UserRepositoryMockGetOrigins contains origins of expectations of the UserRepository.Get
type UserRepositoryMockGetExpectationOrigins struct {
	origin    string
	originCtx string
	originId  string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmGet *mUserRepositoryMockGet) Optional() *mUserRepositoryMockGet {
	mmGet.optional = true
	return mmGet
}

// Expect sets up expected params for UserRepository.Get
func (mmGet *mUserRepositoryMockGet) Expect(ctx context.Context, id int64) *mUserRepositoryMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("UserRepositoryMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &UserRepositoryMockGetExpectation{}
	}

	if mmGet.defaultExpectation.paramPtrs != nil {
		mmGet.mock.t.Fatalf("UserRepositoryMock.Get mock is already set by ExpectParams functions")
	}

	mmGet.defaultExpectation.params = &UserRepositoryMockGetParams{ctx, id}
	mmGet.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmGet.expectations {
		if minimock.Equal(e.params, mmGet.defaultExpectation.params) {
			mmGet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGet.defaultExpectation.params)
		}
	}

	return mmGet
}

// ExpectCtxParam1 sets up expected param ctx for UserRepository.Get
func (mmGet *mUserRepositoryMockGet) ExpectCtxParam1(ctx context.Context) *mUserRepositoryMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("UserRepositoryMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &UserRepositoryMockGetExpectation{}
	}

	if mmGet.defaultExpectation.params != nil {
		mmGet.mock.t.Fatalf("UserRepositoryMock.Get mock is already set by Expect")
	}

	if mmGet.defaultExpectation.paramPtrs == nil {
		mmGet.defaultExpectation.paramPtrs = &UserRepositoryMockGetParamPtrs{}
	}
	mmGet.defaultExpectation.paramPtrs.ctx = &ctx
	mmGet.defaultExpectation.expectationOrigins.originCtx = minimock.CallerInfo(1)

	return mmGet
}

// ExpectIdParam2 sets up expected param id for UserRepository.Get
func (mmGet *mUserRepositoryMockGet) ExpectIdParam2(id int64) *mUserRepositoryMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("UserRepositoryMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &UserRepositoryMockGetExpectation{}
	}

	if mmGet.defaultExpectation.params != nil {
		mmGet.mock.t.Fatalf("UserRepositoryMock.Get mock is already set by Expect")
	}

	if mmGet.defaultExpectation.paramPtrs == nil {
		mmGet.defaultExpectation.paramPtrs = &UserRepositoryMockGetParamPtrs{}
	}
	mmGet.defaultExpectation.paramPtrs.id = &id
	mmGet.defaultExpectation.expectationOrigins.originId = minimock.CallerInfo(1)

	return mmGet
}

// Inspect accepts an inspector function that has same arguments as the UserRepository.Get
func (mmGet *mUserRepositoryMockGet) Inspect(f func(ctx context.Context, id int64)) *mUserRepositoryMockGet {
	if mmGet.mock.inspectFuncGet != nil {
		mmGet.mock.t.Fatalf("Inspect function is already set for UserRepositoryMock.Get")
	}

	mmGet.mock.inspectFuncGet = f

	return mmGet
}

// Return sets up results that will be returned by UserRepository.Get
func (mmGet *mUserRepositoryMockGet) Return(up1 *model.User, err error) *UserRepositoryMock {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("UserRepositoryMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &UserRepositoryMockGetExpectation{mock: mmGet.mock}
	}
	mmGet.defaultExpectation.results = &UserRepositoryMockGetResults{up1, err}
	mmGet.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmGet.mock
}

// Set uses given function f to mock the UserRepository.Get method
func (mmGet *mUserRepositoryMockGet) Set(f func(ctx context.Context, id int64) (up1 *model.User, err error)) *UserRepositoryMock {
	if mmGet.defaultExpectation != nil {
		mmGet.mock.t.Fatalf("Default expectation is already set for the UserRepository.Get method")
	}

	if len(mmGet.expectations) > 0 {
		mmGet.mock.t.Fatalf("Some expectations are already set for the UserRepository.Get method")
	}

	mmGet.mock.funcGet = f
	mmGet.mock.funcGetOrigin = minimock.CallerInfo(1)
	return mmGet.mock
}

// When sets expectation for the UserRepository.Get which will trigger the result defined by the following
// Then helper
func (mmGet *mUserRepositoryMockGet) When(ctx context.Context, id int64) *UserRepositoryMockGetExpectation {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("UserRepositoryMock.Get mock is already set by Set")
	}

	expectation := &UserRepositoryMockGetExpectation{
		mock:               mmGet.mock,
		params:             &UserRepositoryMockGetParams{ctx, id},
		expectationOrigins: UserRepositoryMockGetExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmGet.expectations = append(mmGet.expectations, expectation)
	return expectation
}

// Then sets up UserRepository.Get return parameters for the expectation previously defined by the When method
func (e *UserRepositoryMockGetExpectation) Then(up1 *model.User, err error) *UserRepositoryMock {
	e.results = &UserRepositoryMockGetResults{up1, err}
	return e.mock
}

// Times sets number of times UserRepository.Get should be invoked
func (mmGet *mUserRepositoryMockGet) Times(n uint64) *mUserRepositoryMockGet {
	if n == 0 {
		mmGet.mock.t.Fatalf("Times of UserRepositoryMock.Get mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmGet.expectedInvocations, n)
	mmGet.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmGet
}

func (mmGet *mUserRepositoryMockGet) invocationsDone() bool {
	if len(mmGet.expectations) == 0 && mmGet.defaultExpectation == nil && mmGet.mock.funcGet == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmGet.mock.afterGetCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmGet.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Get implements mm_repository.UserRepository
func (mmGet *UserRepositoryMock) Get(ctx context.Context, id int64) (up1 *model.User, err error) {
	mm_atomic.AddUint64(&mmGet.beforeGetCounter, 1)
	defer mm_atomic.AddUint64(&mmGet.afterGetCounter, 1)

	mmGet.t.Helper()

	if mmGet.inspectFuncGet != nil {
		mmGet.inspectFuncGet(ctx, id)
	}

	mm_params := UserRepositoryMockGetParams{ctx, id}

	// Record call args
	mmGet.GetMock.mutex.Lock()
	mmGet.GetMock.callArgs = append(mmGet.GetMock.callArgs, &mm_params)
	mmGet.GetMock.mutex.Unlock()

	for _, e := range mmGet.GetMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.up1, e.results.err
		}
	}

	if mmGet.GetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGet.GetMock.defaultExpectation.Counter, 1)
		mm_want := mmGet.GetMock.defaultExpectation.params
		mm_want_ptrs := mmGet.GetMock.defaultExpectation.paramPtrs

		mm_got := UserRepositoryMockGetParams{ctx, id}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmGet.t.Errorf("UserRepositoryMock.Get got unexpected parameter ctx, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmGet.GetMock.defaultExpectation.expectationOrigins.originCtx, *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.id != nil && !minimock.Equal(*mm_want_ptrs.id, mm_got.id) {
				mmGet.t.Errorf("UserRepositoryMock.Get got unexpected parameter id, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmGet.GetMock.defaultExpectation.expectationOrigins.originId, *mm_want_ptrs.id, mm_got.id, minimock.Diff(*mm_want_ptrs.id, mm_got.id))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGet.t.Errorf("UserRepositoryMock.Get got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmGet.GetMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGet.GetMock.defaultExpectation.results
		if mm_results == nil {
			mmGet.t.Fatal("No results are set for the UserRepositoryMock.Get")
		}
		return (*mm_results).up1, (*mm_results).err
	}
	if mmGet.funcGet != nil {
		return mmGet.funcGet(ctx, id)
	}
	mmGet.t.Fatalf("Unexpected call to UserRepositoryMock.Get. %v %v", ctx, id)
	return
}

// GetAfterCounter returns a count of finished UserRepositoryMock.Get invocations
func (mmGet *UserRepositoryMock) GetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.afterGetCounter)
}

// GetBeforeCounter returns a count of UserRepositoryMock.Get invocations
func (mmGet *UserRepositoryMock) GetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.beforeGetCounter)
}

// Calls returns a list of arguments used in each call to UserRepositoryMock.Get.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGet *mUserRepositoryMockGet) Calls() []*UserRepositoryMockGetParams {
	mmGet.mutex.RLock()

	argCopy := make([]*UserRepositoryMockGetParams, len(mmGet.callArgs))
	copy(argCopy, mmGet.callArgs)

	mmGet.mutex.RUnlock()

	return argCopy
}

// MinimockGetDone returns true if the count of the Get invocations corresponds
// the number of defined expectations
func (m *UserRepositoryMock) MinimockGetDone() bool {
	if m.GetMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.GetMock.invocationsDone()
}

// MinimockGetInspect logs each unmet expectation
func (m *UserRepositoryMock) MinimockGetInspect() {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserRepositoryMock.Get at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterGetCounter := mm_atomic.LoadUint64(&m.afterGetCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && afterGetCounter < 1 {
		if m.GetMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to UserRepositoryMock.Get at\n%s", m.GetMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to UserRepositoryMock.Get at\n%s with params: %#v", m.GetMock.defaultExpectation.expectationOrigins.origin, *m.GetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && afterGetCounter < 1 {
		m.t.Errorf("Expected call to UserRepositoryMock.Get at\n%s", m.funcGetOrigin)
	}

	if !m.GetMock.invocationsDone() && afterGetCounter > 0 {
		m.t.Errorf("Expected %d calls to UserRepositoryMock.Get at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.GetMock.expectedInvocations), m.GetMock.expectedInvocationsOrigin, afterGetCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *UserRepositoryMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockCreateInspect()

			m.MinimockGetInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *UserRepositoryMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *UserRepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateDone() &&
		m.MinimockGetDone()
}
