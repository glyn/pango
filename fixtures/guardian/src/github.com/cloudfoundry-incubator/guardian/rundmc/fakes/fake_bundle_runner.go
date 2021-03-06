// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/garden"
	"github.com/cloudfoundry-incubator/guardian/rundmc"
)

type FakeBundleRunner struct {
	StartStub        func(bundlePath, id string, io garden.ProcessIO) (garden.Process, error)
	startMutex       sync.RWMutex
	startArgsForCall []struct {
		bundlePath string
		id         string
		io         garden.ProcessIO
	}
	startReturns struct {
		result1 garden.Process
		result2 error
	}
	ExecStub        func(id string, spec garden.ProcessSpec, io garden.ProcessIO) (garden.Process, error)
	execMutex       sync.RWMutex
	execArgsForCall []struct {
		id   string
		spec garden.ProcessSpec
		io   garden.ProcessIO
	}
	execReturns struct {
		result1 garden.Process
		result2 error
	}
	KillStub        func(bundlePath string) error
	killMutex       sync.RWMutex
	killArgsForCall []struct {
		bundlePath string
	}
	killReturns struct {
		result1 error
	}
}

func (fake *FakeBundleRunner) Start(bundlePath string, id string, io garden.ProcessIO) (garden.Process, error) {
	fake.startMutex.Lock()
	fake.startArgsForCall = append(fake.startArgsForCall, struct {
		bundlePath string
		id         string
		io         garden.ProcessIO
	}{bundlePath, id, io})
	fake.startMutex.Unlock()
	if fake.StartStub != nil {
		return fake.StartStub(bundlePath, id, io)
	} else {
		return fake.startReturns.result1, fake.startReturns.result2
	}
}

func (fake *FakeBundleRunner) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeBundleRunner) StartArgsForCall(i int) (string, string, garden.ProcessIO) {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return fake.startArgsForCall[i].bundlePath, fake.startArgsForCall[i].id, fake.startArgsForCall[i].io
}

func (fake *FakeBundleRunner) StartReturns(result1 garden.Process, result2 error) {
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 garden.Process
		result2 error
	}{result1, result2}
}

func (fake *FakeBundleRunner) Exec(id string, spec garden.ProcessSpec, io garden.ProcessIO) (garden.Process, error) {
	fake.execMutex.Lock()
	fake.execArgsForCall = append(fake.execArgsForCall, struct {
		id   string
		spec garden.ProcessSpec
		io   garden.ProcessIO
	}{id, spec, io})
	fake.execMutex.Unlock()
	if fake.ExecStub != nil {
		return fake.ExecStub(id, spec, io)
	} else {
		return fake.execReturns.result1, fake.execReturns.result2
	}
}

func (fake *FakeBundleRunner) ExecCallCount() int {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	return len(fake.execArgsForCall)
}

func (fake *FakeBundleRunner) ExecArgsForCall(i int) (string, garden.ProcessSpec, garden.ProcessIO) {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	return fake.execArgsForCall[i].id, fake.execArgsForCall[i].spec, fake.execArgsForCall[i].io
}

func (fake *FakeBundleRunner) ExecReturns(result1 garden.Process, result2 error) {
	fake.ExecStub = nil
	fake.execReturns = struct {
		result1 garden.Process
		result2 error
	}{result1, result2}
}

func (fake *FakeBundleRunner) Kill(bundlePath string) error {
	fake.killMutex.Lock()
	fake.killArgsForCall = append(fake.killArgsForCall, struct {
		bundlePath string
	}{bundlePath})
	fake.killMutex.Unlock()
	if fake.KillStub != nil {
		return fake.KillStub(bundlePath)
	} else {
		return fake.killReturns.result1
	}
}

func (fake *FakeBundleRunner) KillCallCount() int {
	fake.killMutex.RLock()
	defer fake.killMutex.RUnlock()
	return len(fake.killArgsForCall)
}

func (fake *FakeBundleRunner) KillArgsForCall(i int) string {
	fake.killMutex.RLock()
	defer fake.killMutex.RUnlock()
	return fake.killArgsForCall[i].bundlePath
}

func (fake *FakeBundleRunner) KillReturns(result1 error) {
	fake.KillStub = nil
	fake.killReturns = struct {
		result1 error
	}{result1}
}

var _ rundmc.BundleRunner = new(FakeBundleRunner)
