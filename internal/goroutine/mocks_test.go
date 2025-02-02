// Code generated by go-mockgen 1.3.6; DO NOT EDIT.
//
// This file was generated by running `sg generate` (or `go-mockgen`) at the root of
// this repository. To add additional mocks to this or another package, add a new entry
// to the mockgen.yaml file in the root of this repository.

package goroutine

import (
	"context"
	"sync"
)

// MockBackgroundRoutine is a mock implementation of the BackgroundRoutine
// interface (from the package
// github.com/sourcegraph/sourcegraph/internal/goroutine) used for unit
// testing.
type MockBackgroundRoutine struct {
	// StartFunc is an instance of a mock function object controlling the
	// behavior of the method Start.
	StartFunc *BackgroundRoutineStartFunc
	// StopFunc is an instance of a mock function object controlling the
	// behavior of the method Stop.
	StopFunc *BackgroundRoutineStopFunc
}

// NewMockBackgroundRoutine creates a new mock of the BackgroundRoutine
// interface. All methods return zero values for all results, unless
// overwritten.
func NewMockBackgroundRoutine() *MockBackgroundRoutine {
	return &MockBackgroundRoutine{
		StartFunc: &BackgroundRoutineStartFunc{
			defaultHook: func() {
				return
			},
		},
		StopFunc: &BackgroundRoutineStopFunc{
			defaultHook: func() {
				return
			},
		},
	}
}

// NewStrictMockBackgroundRoutine creates a new mock of the
// BackgroundRoutine interface. All methods panic on invocation, unless
// overwritten.
func NewStrictMockBackgroundRoutine() *MockBackgroundRoutine {
	return &MockBackgroundRoutine{
		StartFunc: &BackgroundRoutineStartFunc{
			defaultHook: func() {
				panic("unexpected invocation of MockBackgroundRoutine.Start")
			},
		},
		StopFunc: &BackgroundRoutineStopFunc{
			defaultHook: func() {
				panic("unexpected invocation of MockBackgroundRoutine.Stop")
			},
		},
	}
}

// NewMockBackgroundRoutineFrom creates a new mock of the
// MockBackgroundRoutine interface. All methods delegate to the given
// implementation, unless overwritten.
func NewMockBackgroundRoutineFrom(i BackgroundRoutine) *MockBackgroundRoutine {
	return &MockBackgroundRoutine{
		StartFunc: &BackgroundRoutineStartFunc{
			defaultHook: i.Start,
		},
		StopFunc: &BackgroundRoutineStopFunc{
			defaultHook: i.Stop,
		},
	}
}

// BackgroundRoutineStartFunc describes the behavior when the Start method
// of the parent MockBackgroundRoutine instance is invoked.
type BackgroundRoutineStartFunc struct {
	defaultHook func()
	hooks       []func()
	history     []BackgroundRoutineStartFuncCall
	mutex       sync.Mutex
}

// Start delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockBackgroundRoutine) Start() {
	m.StartFunc.nextHook()()
	m.StartFunc.appendCall(BackgroundRoutineStartFuncCall{})
	return
}

// SetDefaultHook sets function that is called when the Start method of the
// parent MockBackgroundRoutine instance is invoked and the hook queue is
// empty.
func (f *BackgroundRoutineStartFunc) SetDefaultHook(hook func()) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Start method of the parent MockBackgroundRoutine instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *BackgroundRoutineStartFunc) PushHook(hook func()) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *BackgroundRoutineStartFunc) SetDefaultReturn() {
	f.SetDefaultHook(func() {
		return
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *BackgroundRoutineStartFunc) PushReturn() {
	f.PushHook(func() {
		return
	})
}

func (f *BackgroundRoutineStartFunc) nextHook() func() {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *BackgroundRoutineStartFunc) appendCall(r0 BackgroundRoutineStartFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of BackgroundRoutineStartFuncCall objects
// describing the invocations of this function.
func (f *BackgroundRoutineStartFunc) History() []BackgroundRoutineStartFuncCall {
	f.mutex.Lock()
	history := make([]BackgroundRoutineStartFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// BackgroundRoutineStartFuncCall is an object that describes an invocation
// of method Start on an instance of MockBackgroundRoutine.
type BackgroundRoutineStartFuncCall struct{}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c BackgroundRoutineStartFuncCall) Args() []interface{} {
	return []interface{}{}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c BackgroundRoutineStartFuncCall) Results() []interface{} {
	return []interface{}{}
}

// BackgroundRoutineStopFunc describes the behavior when the Stop method of
// the parent MockBackgroundRoutine instance is invoked.
type BackgroundRoutineStopFunc struct {
	defaultHook func()
	hooks       []func()
	history     []BackgroundRoutineStopFuncCall
	mutex       sync.Mutex
}

// Stop delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockBackgroundRoutine) Stop() {
	m.StopFunc.nextHook()()
	m.StopFunc.appendCall(BackgroundRoutineStopFuncCall{})
	return
}

// SetDefaultHook sets function that is called when the Stop method of the
// parent MockBackgroundRoutine instance is invoked and the hook queue is
// empty.
func (f *BackgroundRoutineStopFunc) SetDefaultHook(hook func()) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Stop method of the parent MockBackgroundRoutine instance invokes the hook
// at the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *BackgroundRoutineStopFunc) PushHook(hook func()) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *BackgroundRoutineStopFunc) SetDefaultReturn() {
	f.SetDefaultHook(func() {
		return
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *BackgroundRoutineStopFunc) PushReturn() {
	f.PushHook(func() {
		return
	})
}

func (f *BackgroundRoutineStopFunc) nextHook() func() {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *BackgroundRoutineStopFunc) appendCall(r0 BackgroundRoutineStopFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of BackgroundRoutineStopFuncCall objects
// describing the invocations of this function.
func (f *BackgroundRoutineStopFunc) History() []BackgroundRoutineStopFuncCall {
	f.mutex.Lock()
	history := make([]BackgroundRoutineStopFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// BackgroundRoutineStopFuncCall is an object that describes an invocation
// of method Stop on an instance of MockBackgroundRoutine.
type BackgroundRoutineStopFuncCall struct{}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c BackgroundRoutineStopFuncCall) Args() []interface{} {
	return []interface{}{}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c BackgroundRoutineStopFuncCall) Results() []interface{} {
	return []interface{}{}
}

// MockErrorHandler is a mock implementation of the ErrorHandler interface
// (from the package github.com/sourcegraph/sourcegraph/internal/goroutine)
// used for unit testing.
type MockErrorHandler struct {
	// HandleErrorFunc is an instance of a mock function object controlling
	// the behavior of the method HandleError.
	HandleErrorFunc *ErrorHandlerHandleErrorFunc
}

// NewMockErrorHandler creates a new mock of the ErrorHandler interface. All
// methods return zero values for all results, unless overwritten.
func NewMockErrorHandler() *MockErrorHandler {
	return &MockErrorHandler{
		HandleErrorFunc: &ErrorHandlerHandleErrorFunc{
			defaultHook: func(error) {
				return
			},
		},
	}
}

// NewStrictMockErrorHandler creates a new mock of the ErrorHandler
// interface. All methods panic on invocation, unless overwritten.
func NewStrictMockErrorHandler() *MockErrorHandler {
	return &MockErrorHandler{
		HandleErrorFunc: &ErrorHandlerHandleErrorFunc{
			defaultHook: func(error) {
				panic("unexpected invocation of MockErrorHandler.HandleError")
			},
		},
	}
}

// NewMockErrorHandlerFrom creates a new mock of the MockErrorHandler
// interface. All methods delegate to the given implementation, unless
// overwritten.
func NewMockErrorHandlerFrom(i ErrorHandler) *MockErrorHandler {
	return &MockErrorHandler{
		HandleErrorFunc: &ErrorHandlerHandleErrorFunc{
			defaultHook: i.HandleError,
		},
	}
}

// ErrorHandlerHandleErrorFunc describes the behavior when the HandleError
// method of the parent MockErrorHandler instance is invoked.
type ErrorHandlerHandleErrorFunc struct {
	defaultHook func(error)
	hooks       []func(error)
	history     []ErrorHandlerHandleErrorFuncCall
	mutex       sync.Mutex
}

// HandleError delegates to the next hook function in the queue and stores
// the parameter and result values of this invocation.
func (m *MockErrorHandler) HandleError(v0 error) {
	m.HandleErrorFunc.nextHook()(v0)
	m.HandleErrorFunc.appendCall(ErrorHandlerHandleErrorFuncCall{v0})
	return
}

// SetDefaultHook sets function that is called when the HandleError method
// of the parent MockErrorHandler instance is invoked and the hook queue is
// empty.
func (f *ErrorHandlerHandleErrorFunc) SetDefaultHook(hook func(error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// HandleError method of the parent MockErrorHandler instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *ErrorHandlerHandleErrorFunc) PushHook(hook func(error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *ErrorHandlerHandleErrorFunc) SetDefaultReturn() {
	f.SetDefaultHook(func(error) {
		return
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *ErrorHandlerHandleErrorFunc) PushReturn() {
	f.PushHook(func(error) {
		return
	})
}

func (f *ErrorHandlerHandleErrorFunc) nextHook() func(error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *ErrorHandlerHandleErrorFunc) appendCall(r0 ErrorHandlerHandleErrorFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of ErrorHandlerHandleErrorFuncCall objects
// describing the invocations of this function.
func (f *ErrorHandlerHandleErrorFunc) History() []ErrorHandlerHandleErrorFuncCall {
	f.mutex.Lock()
	history := make([]ErrorHandlerHandleErrorFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// ErrorHandlerHandleErrorFuncCall is an object that describes an invocation
// of method HandleError on an instance of MockErrorHandler.
type ErrorHandlerHandleErrorFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c ErrorHandlerHandleErrorFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c ErrorHandlerHandleErrorFuncCall) Results() []interface{} {
	return []interface{}{}
}

// MockFinalizer is a mock implementation of the Finalizer interface (from
// the package github.com/sourcegraph/sourcegraph/internal/goroutine) used
// for unit testing.
type MockFinalizer struct {
	// OnShutdownFunc is an instance of a mock function object controlling
	// the behavior of the method OnShutdown.
	OnShutdownFunc *FinalizerOnShutdownFunc
}

// NewMockFinalizer creates a new mock of the Finalizer interface. All
// methods return zero values for all results, unless overwritten.
func NewMockFinalizer() *MockFinalizer {
	return &MockFinalizer{
		OnShutdownFunc: &FinalizerOnShutdownFunc{
			defaultHook: func() {
				return
			},
		},
	}
}

// NewStrictMockFinalizer creates a new mock of the Finalizer interface. All
// methods panic on invocation, unless overwritten.
func NewStrictMockFinalizer() *MockFinalizer {
	return &MockFinalizer{
		OnShutdownFunc: &FinalizerOnShutdownFunc{
			defaultHook: func() {
				panic("unexpected invocation of MockFinalizer.OnShutdown")
			},
		},
	}
}

// NewMockFinalizerFrom creates a new mock of the MockFinalizer interface.
// All methods delegate to the given implementation, unless overwritten.
func NewMockFinalizerFrom(i Finalizer) *MockFinalizer {
	return &MockFinalizer{
		OnShutdownFunc: &FinalizerOnShutdownFunc{
			defaultHook: i.OnShutdown,
		},
	}
}

// FinalizerOnShutdownFunc describes the behavior when the OnShutdown method
// of the parent MockFinalizer instance is invoked.
type FinalizerOnShutdownFunc struct {
	defaultHook func()
	hooks       []func()
	history     []FinalizerOnShutdownFuncCall
	mutex       sync.Mutex
}

// OnShutdown delegates to the next hook function in the queue and stores
// the parameter and result values of this invocation.
func (m *MockFinalizer) OnShutdown() {
	m.OnShutdownFunc.nextHook()()
	m.OnShutdownFunc.appendCall(FinalizerOnShutdownFuncCall{})
	return
}

// SetDefaultHook sets function that is called when the OnShutdown method of
// the parent MockFinalizer instance is invoked and the hook queue is empty.
func (f *FinalizerOnShutdownFunc) SetDefaultHook(hook func()) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// OnShutdown method of the parent MockFinalizer instance invokes the hook
// at the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *FinalizerOnShutdownFunc) PushHook(hook func()) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *FinalizerOnShutdownFunc) SetDefaultReturn() {
	f.SetDefaultHook(func() {
		return
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *FinalizerOnShutdownFunc) PushReturn() {
	f.PushHook(func() {
		return
	})
}

func (f *FinalizerOnShutdownFunc) nextHook() func() {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *FinalizerOnShutdownFunc) appendCall(r0 FinalizerOnShutdownFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of FinalizerOnShutdownFuncCall objects
// describing the invocations of this function.
func (f *FinalizerOnShutdownFunc) History() []FinalizerOnShutdownFuncCall {
	f.mutex.Lock()
	history := make([]FinalizerOnShutdownFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// FinalizerOnShutdownFuncCall is an object that describes an invocation of
// method OnShutdown on an instance of MockFinalizer.
type FinalizerOnShutdownFuncCall struct{}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c FinalizerOnShutdownFuncCall) Args() []interface{} {
	return []interface{}{}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c FinalizerOnShutdownFuncCall) Results() []interface{} {
	return []interface{}{}
}

// MockHandler is a mock implementation of the Handler interface (from the
// package github.com/sourcegraph/sourcegraph/internal/goroutine) used for
// unit testing.
type MockHandler struct {
	// HandleFunc is an instance of a mock function object controlling the
	// behavior of the method Handle.
	HandleFunc *HandlerHandleFunc
}

// NewMockHandler creates a new mock of the Handler interface. All methods
// return zero values for all results, unless overwritten.
func NewMockHandler() *MockHandler {
	return &MockHandler{
		HandleFunc: &HandlerHandleFunc{
			defaultHook: func(context.Context) (r0 error) {
				return
			},
		},
	}
}

// NewStrictMockHandler creates a new mock of the Handler interface. All
// methods panic on invocation, unless overwritten.
func NewStrictMockHandler() *MockHandler {
	return &MockHandler{
		HandleFunc: &HandlerHandleFunc{
			defaultHook: func(context.Context) error {
				panic("unexpected invocation of MockHandler.Handle")
			},
		},
	}
}

// NewMockHandlerFrom creates a new mock of the MockHandler interface. All
// methods delegate to the given implementation, unless overwritten.
func NewMockHandlerFrom(i Handler) *MockHandler {
	return &MockHandler{
		HandleFunc: &HandlerHandleFunc{
			defaultHook: i.Handle,
		},
	}
}

// HandlerHandleFunc describes the behavior when the Handle method of the
// parent MockHandler instance is invoked.
type HandlerHandleFunc struct {
	defaultHook func(context.Context) error
	hooks       []func(context.Context) error
	history     []HandlerHandleFuncCall
	mutex       sync.Mutex
}

// Handle delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockHandler) Handle(v0 context.Context) error {
	r0 := m.HandleFunc.nextHook()(v0)
	m.HandleFunc.appendCall(HandlerHandleFuncCall{v0, r0})
	return r0
}

// SetDefaultHook sets function that is called when the Handle method of the
// parent MockHandler instance is invoked and the hook queue is empty.
func (f *HandlerHandleFunc) SetDefaultHook(hook func(context.Context) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Handle method of the parent MockHandler instance invokes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *HandlerHandleFunc) PushHook(hook func(context.Context) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *HandlerHandleFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(context.Context) error {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *HandlerHandleFunc) PushReturn(r0 error) {
	f.PushHook(func(context.Context) error {
		return r0
	})
}

func (f *HandlerHandleFunc) nextHook() func(context.Context) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *HandlerHandleFunc) appendCall(r0 HandlerHandleFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of HandlerHandleFuncCall objects describing
// the invocations of this function.
func (f *HandlerHandleFunc) History() []HandlerHandleFuncCall {
	f.mutex.Lock()
	history := make([]HandlerHandleFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// HandlerHandleFuncCall is an object that describes an invocation of method
// Handle on an instance of MockHandler.
type HandlerHandleFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c HandlerHandleFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c HandlerHandleFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}
