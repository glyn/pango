// This file was generated by counterfeiter
package fake_process_tracker

import (
	"os/exec"
	"sync"

	"github.com/cloudfoundry-incubator/garden"
	"github.com/cloudfoundry-incubator/guardian/rundmc/process_tracker"
)

type FakeProcessTracker struct {
	RunStub        func(processID string, cmd *exec.Cmd, io garden.ProcessIO, tty *garden.TTYSpec, signaller process_tracker.Signaller) (garden.Process, error)
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		processID string
		cmd       *exec.Cmd
		io        garden.ProcessIO
		tty       *garden.TTYSpec
		signaller process_tracker.Signaller
	}
	runReturns struct {
		result1 garden.Process
		result2 error
	}
	AttachStub        func(processID string, io garden.ProcessIO) (garden.Process, error)
	attachMutex       sync.RWMutex
	attachArgsForCall []struct {
		processID string
		io        garden.ProcessIO
	}
	attachReturns struct {
		result1 garden.Process
		result2 error
	}
	RestoreStub        func(processID string, signaller process_tracker.Signaller)
	restoreMutex       sync.RWMutex
	restoreArgsForCall []struct {
		processID string
		signaller process_tracker.Signaller
	}
	ActiveProcessesStub        func() []garden.Process
	activeProcessesMutex       sync.RWMutex
	activeProcessesArgsForCall []struct{}
	activeProcessesReturns     struct {
		result1 []garden.Process
	}
}

func (fake *FakeProcessTracker) Run(processID string, cmd *exec.Cmd, io garden.ProcessIO, tty *garden.TTYSpec, signaller process_tracker.Signaller) (garden.Process, error) {
	fake.runMutex.Lock()
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		processID string
		cmd       *exec.Cmd
		io        garden.ProcessIO
		tty       *garden.TTYSpec
		signaller process_tracker.Signaller
	}{processID, cmd, io, tty, signaller})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(processID, cmd, io, tty, signaller)
	} else {
		return fake.runReturns.result1, fake.runReturns.result2
	}
}

func (fake *FakeProcessTracker) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeProcessTracker) RunArgsForCall(i int) (string, *exec.Cmd, garden.ProcessIO, *garden.TTYSpec, process_tracker.Signaller) {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return fake.runArgsForCall[i].processID, fake.runArgsForCall[i].cmd, fake.runArgsForCall[i].io, fake.runArgsForCall[i].tty, fake.runArgsForCall[i].signaller
}

func (fake *FakeProcessTracker) RunReturns(result1 garden.Process, result2 error) {
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 garden.Process
		result2 error
	}{result1, result2}
}

func (fake *FakeProcessTracker) Attach(processID string, io garden.ProcessIO) (garden.Process, error) {
	fake.attachMutex.Lock()
	fake.attachArgsForCall = append(fake.attachArgsForCall, struct {
		processID string
		io        garden.ProcessIO
	}{processID, io})
	fake.attachMutex.Unlock()
	if fake.AttachStub != nil {
		return fake.AttachStub(processID, io)
	} else {
		return fake.attachReturns.result1, fake.attachReturns.result2
	}
}

func (fake *FakeProcessTracker) AttachCallCount() int {
	fake.attachMutex.RLock()
	defer fake.attachMutex.RUnlock()
	return len(fake.attachArgsForCall)
}

func (fake *FakeProcessTracker) AttachArgsForCall(i int) (string, garden.ProcessIO) {
	fake.attachMutex.RLock()
	defer fake.attachMutex.RUnlock()
	return fake.attachArgsForCall[i].processID, fake.attachArgsForCall[i].io
}

func (fake *FakeProcessTracker) AttachReturns(result1 garden.Process, result2 error) {
	fake.AttachStub = nil
	fake.attachReturns = struct {
		result1 garden.Process
		result2 error
	}{result1, result2}
}

func (fake *FakeProcessTracker) Restore(processID string, signaller process_tracker.Signaller) {
	fake.restoreMutex.Lock()
	fake.restoreArgsForCall = append(fake.restoreArgsForCall, struct {
		processID string
		signaller process_tracker.Signaller
	}{processID, signaller})
	fake.restoreMutex.Unlock()
	if fake.RestoreStub != nil {
		fake.RestoreStub(processID, signaller)
	}
}

func (fake *FakeProcessTracker) RestoreCallCount() int {
	fake.restoreMutex.RLock()
	defer fake.restoreMutex.RUnlock()
	return len(fake.restoreArgsForCall)
}

func (fake *FakeProcessTracker) RestoreArgsForCall(i int) (string, process_tracker.Signaller) {
	fake.restoreMutex.RLock()
	defer fake.restoreMutex.RUnlock()
	return fake.restoreArgsForCall[i].processID, fake.restoreArgsForCall[i].signaller
}

func (fake *FakeProcessTracker) ActiveProcesses() []garden.Process {
	fake.activeProcessesMutex.Lock()
	fake.activeProcessesArgsForCall = append(fake.activeProcessesArgsForCall, struct{}{})
	fake.activeProcessesMutex.Unlock()
	if fake.ActiveProcessesStub != nil {
		return fake.ActiveProcessesStub()
	} else {
		return fake.activeProcessesReturns.result1
	}
}

func (fake *FakeProcessTracker) ActiveProcessesCallCount() int {
	fake.activeProcessesMutex.RLock()
	defer fake.activeProcessesMutex.RUnlock()
	return len(fake.activeProcessesArgsForCall)
}

func (fake *FakeProcessTracker) ActiveProcessesReturns(result1 []garden.Process) {
	fake.ActiveProcessesStub = nil
	fake.activeProcessesReturns = struct {
		result1 []garden.Process
	}{result1}
}

var _ process_tracker.ProcessTracker = new(FakeProcessTracker)
