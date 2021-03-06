// This file was generated by counterfeiter
package terminalhelperfakes

import (
	"io"
	"sync"

	"github.com/cloudfoundry/cli/cf/ssh/terminal"
	"github.com/docker/docker/pkg/term"
)

type FakeTerminalHelper struct {
	StdStreamsStub        func() (stdin io.ReadCloser, stdout io.Writer, stderr io.Writer)
	stdStreamsMutex       sync.RWMutex
	stdStreamsArgsForCall []struct{}
	stdStreamsReturns     struct {
		result1 io.ReadCloser
		result2 io.Writer
		result3 io.Writer
	}
	GetFdInfoStub        func(in interface{}) (fd uintptr, isTerminal bool)
	getFdInfoMutex       sync.RWMutex
	getFdInfoArgsForCall []struct {
		in interface{}
	}
	getFdInfoReturns struct {
		result1 uintptr
		result2 bool
	}
	SetRawTerminalStub        func(fd uintptr) (*term.State, error)
	setRawTerminalMutex       sync.RWMutex
	setRawTerminalArgsForCall []struct {
		fd uintptr
	}
	setRawTerminalReturns struct {
		result1 *term.State
		result2 error
	}
	RestoreTerminalStub        func(fd uintptr, state *term.State) error
	restoreTerminalMutex       sync.RWMutex
	restoreTerminalArgsForCall []struct {
		fd    uintptr
		state *term.State
	}
	restoreTerminalReturns struct {
		result1 error
	}
	IsTerminalStub        func(fd uintptr) bool
	isTerminalMutex       sync.RWMutex
	isTerminalArgsForCall []struct {
		fd uintptr
	}
	isTerminalReturns struct {
		result1 bool
	}
	GetWinsizeStub        func(fd uintptr) (*term.Winsize, error)
	getWinsizeMutex       sync.RWMutex
	getWinsizeArgsForCall []struct {
		fd uintptr
	}
	getWinsizeReturns struct {
		result1 *term.Winsize
		result2 error
	}
}

func (fake *FakeTerminalHelper) StdStreams() (stdin io.ReadCloser, stdout io.Writer, stderr io.Writer) {
	fake.stdStreamsMutex.Lock()
	fake.stdStreamsArgsForCall = append(fake.stdStreamsArgsForCall, struct{}{})
	fake.stdStreamsMutex.Unlock()
	if fake.StdStreamsStub != nil {
		return fake.StdStreamsStub()
	} else {
		return fake.stdStreamsReturns.result1, fake.stdStreamsReturns.result2, fake.stdStreamsReturns.result3
	}
}

func (fake *FakeTerminalHelper) StdStreamsCallCount() int {
	fake.stdStreamsMutex.RLock()
	defer fake.stdStreamsMutex.RUnlock()
	return len(fake.stdStreamsArgsForCall)
}

func (fake *FakeTerminalHelper) StdStreamsReturns(result1 io.ReadCloser, result2 io.Writer, result3 io.Writer) {
	fake.StdStreamsStub = nil
	fake.stdStreamsReturns = struct {
		result1 io.ReadCloser
		result2 io.Writer
		result3 io.Writer
	}{result1, result2, result3}
}

func (fake *FakeTerminalHelper) GetFdInfo(in interface{}) (fd uintptr, isTerminal bool) {
	fake.getFdInfoMutex.Lock()
	fake.getFdInfoArgsForCall = append(fake.getFdInfoArgsForCall, struct {
		in interface{}
	}{in})
	fake.getFdInfoMutex.Unlock()
	if fake.GetFdInfoStub != nil {
		return fake.GetFdInfoStub(in)
	} else {
		return fake.getFdInfoReturns.result1, fake.getFdInfoReturns.result2
	}
}

func (fake *FakeTerminalHelper) GetFdInfoCallCount() int {
	fake.getFdInfoMutex.RLock()
	defer fake.getFdInfoMutex.RUnlock()
	return len(fake.getFdInfoArgsForCall)
}

func (fake *FakeTerminalHelper) GetFdInfoArgsForCall(i int) interface{} {
	fake.getFdInfoMutex.RLock()
	defer fake.getFdInfoMutex.RUnlock()
	return fake.getFdInfoArgsForCall[i].in
}

func (fake *FakeTerminalHelper) GetFdInfoReturns(result1 uintptr, result2 bool) {
	fake.GetFdInfoStub = nil
	fake.getFdInfoReturns = struct {
		result1 uintptr
		result2 bool
	}{result1, result2}
}

func (fake *FakeTerminalHelper) SetRawTerminal(fd uintptr) (*term.State, error) {
	fake.setRawTerminalMutex.Lock()
	fake.setRawTerminalArgsForCall = append(fake.setRawTerminalArgsForCall, struct {
		fd uintptr
	}{fd})
	fake.setRawTerminalMutex.Unlock()
	if fake.SetRawTerminalStub != nil {
		return fake.SetRawTerminalStub(fd)
	} else {
		return fake.setRawTerminalReturns.result1, fake.setRawTerminalReturns.result2
	}
}

func (fake *FakeTerminalHelper) SetRawTerminalCallCount() int {
	fake.setRawTerminalMutex.RLock()
	defer fake.setRawTerminalMutex.RUnlock()
	return len(fake.setRawTerminalArgsForCall)
}

func (fake *FakeTerminalHelper) SetRawTerminalArgsForCall(i int) uintptr {
	fake.setRawTerminalMutex.RLock()
	defer fake.setRawTerminalMutex.RUnlock()
	return fake.setRawTerminalArgsForCall[i].fd
}

func (fake *FakeTerminalHelper) SetRawTerminalReturns(result1 *term.State, result2 error) {
	fake.SetRawTerminalStub = nil
	fake.setRawTerminalReturns = struct {
		result1 *term.State
		result2 error
	}{result1, result2}
}

func (fake *FakeTerminalHelper) RestoreTerminal(fd uintptr, state *term.State) error {
	fake.restoreTerminalMutex.Lock()
	fake.restoreTerminalArgsForCall = append(fake.restoreTerminalArgsForCall, struct {
		fd    uintptr
		state *term.State
	}{fd, state})
	fake.restoreTerminalMutex.Unlock()
	if fake.RestoreTerminalStub != nil {
		return fake.RestoreTerminalStub(fd, state)
	} else {
		return fake.restoreTerminalReturns.result1
	}
}

func (fake *FakeTerminalHelper) RestoreTerminalCallCount() int {
	fake.restoreTerminalMutex.RLock()
	defer fake.restoreTerminalMutex.RUnlock()
	return len(fake.restoreTerminalArgsForCall)
}

func (fake *FakeTerminalHelper) RestoreTerminalArgsForCall(i int) (uintptr, *term.State) {
	fake.restoreTerminalMutex.RLock()
	defer fake.restoreTerminalMutex.RUnlock()
	return fake.restoreTerminalArgsForCall[i].fd, fake.restoreTerminalArgsForCall[i].state
}

func (fake *FakeTerminalHelper) RestoreTerminalReturns(result1 error) {
	fake.RestoreTerminalStub = nil
	fake.restoreTerminalReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTerminalHelper) IsTerminal(fd uintptr) bool {
	fake.isTerminalMutex.Lock()
	fake.isTerminalArgsForCall = append(fake.isTerminalArgsForCall, struct {
		fd uintptr
	}{fd})
	fake.isTerminalMutex.Unlock()
	if fake.IsTerminalStub != nil {
		return fake.IsTerminalStub(fd)
	} else {
		return fake.isTerminalReturns.result1
	}
}

func (fake *FakeTerminalHelper) IsTerminalCallCount() int {
	fake.isTerminalMutex.RLock()
	defer fake.isTerminalMutex.RUnlock()
	return len(fake.isTerminalArgsForCall)
}

func (fake *FakeTerminalHelper) IsTerminalArgsForCall(i int) uintptr {
	fake.isTerminalMutex.RLock()
	defer fake.isTerminalMutex.RUnlock()
	return fake.isTerminalArgsForCall[i].fd
}

func (fake *FakeTerminalHelper) IsTerminalReturns(result1 bool) {
	fake.IsTerminalStub = nil
	fake.isTerminalReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeTerminalHelper) GetWinsize(fd uintptr) (*term.Winsize, error) {
	fake.getWinsizeMutex.Lock()
	fake.getWinsizeArgsForCall = append(fake.getWinsizeArgsForCall, struct {
		fd uintptr
	}{fd})
	fake.getWinsizeMutex.Unlock()
	if fake.GetWinsizeStub != nil {
		return fake.GetWinsizeStub(fd)
	} else {
		return fake.getWinsizeReturns.result1, fake.getWinsizeReturns.result2
	}
}

func (fake *FakeTerminalHelper) GetWinsizeCallCount() int {
	fake.getWinsizeMutex.RLock()
	defer fake.getWinsizeMutex.RUnlock()
	return len(fake.getWinsizeArgsForCall)
}

func (fake *FakeTerminalHelper) GetWinsizeArgsForCall(i int) uintptr {
	fake.getWinsizeMutex.RLock()
	defer fake.getWinsizeMutex.RUnlock()
	return fake.getWinsizeArgsForCall[i].fd
}

func (fake *FakeTerminalHelper) GetWinsizeReturns(result1 *term.Winsize, result2 error) {
	fake.GetWinsizeStub = nil
	fake.getWinsizeReturns = struct {
		result1 *term.Winsize
		result2 error
	}{result1, result2}
}

var _ sshTerminal.TerminalHelper = new(FakeTerminalHelper)
