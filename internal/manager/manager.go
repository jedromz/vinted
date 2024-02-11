package manager

import (
	"context"
	"errors"
	"sync"
)

// Manager struct now includes a mutex for concurrent access management
type Manager struct {
	processes map[string]Process
	mu        sync.RWMutex // Protects processes
}

func New() *Manager {
	return &Manager{
		processes: make(map[string]Process),
	}
}

// AddProcess safely adds a process with concurrency control
func (m *Manager) AddProcess(name string, process Process) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.processes[name] = process
}

// StartProcess attempts to start a named process, now with improved error handling
func (m *Manager) StartProcess(ctx context.Context, name string) error {
	m.mu.RLock()
	process, ok := m.processes[name]
	m.mu.RUnlock()

	if !ok {
		return errors.New("process not found")
	}
	return process.Start(ctx)
}

// StopProcess attempts to stop a named process, safely and with proper error handling
func (m *Manager) StopProcess(ctx context.Context, name string) error {
	m.mu.RLock()
	process, ok := m.processes[name]
	m.mu.RUnlock()

	if !ok {
		return errors.New("process not found")
	}
	return process.Stop(ctx)
}

// StartAll starts all processes in the manager's control
func (m *Manager) StartAll(ctx context.Context) error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for name, process := range m.processes {
		if err := process.Start(ctx); err != nil {
			return errors.New("failed to start process " + name + ": " + err.Error())
		}
	}
	return nil
}

// StopAll stops all processes and optionally clears the list
func (m *Manager) StopAll(ctx context.Context, clear bool) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for name, process := range m.processes {
		if err := process.Stop(ctx); err != nil {
			return errors.New("failed to stop process " + name + ": " + err.Error())
		}
		if clear {
			delete(m.processes, name)
		}
	}
	return nil
}
