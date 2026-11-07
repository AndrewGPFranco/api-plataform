package projects

import (
	"sync"

	"github.com/google/uuid"
)

type scriptKey struct {
	projectID uuid.UUID
	path      string
}

type ScriptsEmExecucao struct {
	mu      sync.RWMutex
	scripts map[scriptKey]struct{}
}

func NewScriptsEmExecucao() *ScriptsEmExecucao {
	return &ScriptsEmExecucao{scripts: make(map[scriptKey]struct{})}
}

func (r *ScriptsEmExecucao) Rodando(projectID uuid.UUID, path string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	_, ok := r.scripts[scriptKey{projectID, path}]
	return ok
}

func (r *ScriptsEmExecucao) Marcar(projectID uuid.UUID, path string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	key := scriptKey{projectID, path}
	if _, ok := r.scripts[key]; ok {
		return false
	}
	r.scripts[key] = struct{}{}
	return true
}

func (r *ScriptsEmExecucao) Remover(projectID uuid.UUID, path string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.scripts, scriptKey{projectID, path})
}
