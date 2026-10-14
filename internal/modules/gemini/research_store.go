package gemini

import (
	"sync"
	"time"

	"gemini-web-to-api/internal/modules/gemini/dto"
)

// taskStatus enumerates possible research task states
type taskStatus string

const (
	taskStatusInProgress taskStatus = "in_progress"
	taskStatusCompleted  taskStatus = "completed"
	taskStatusFailed     taskStatus = "failed"
)

// researchTask holds a single background deep research task
type researchTask struct {
	ID        string
	Status    taskStatus
	Query     string
	CreatedAt int64
	Result    *dto.DeepResearchResponse // non-nil when completed
	Error     string                    // non-empty when failed
}

// taskStore is a concurrency-safe in-memory store for research tasks
type taskStore struct {
	mu    sync.RWMutex
	tasks map[string]*researchTask
}

func newTaskStore() *taskStore {
	return &taskStore{tasks: make(map[string]*researchTask)}
}

func (s *taskStore) set(t *researchTask) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tasks[t.ID] = t
}

func (s *taskStore) get(id string) (*researchTask, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	t, ok := s.tasks[id]
	return t, ok
}

// update atomically replaces the task with the given ID (must already exist)
func (s *taskStore) update(id string, fn func(*researchTask)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if t, ok := s.tasks[id]; ok {
		fn(t)
	}
}

// purgeOlderThan removes tasks older than the given duration (call periodically if needed)
func (s *taskStore) purgeOlderThan(d time.Duration) {
	cutoff := time.Now().Add(-d).Unix()
	s.mu.Lock()
	defer s.mu.Unlock()
	for id, t := range s.tasks {
		if t.CreatedAt < cutoff {
			delete(s.tasks, id)
		}
	}
}

// toDTO converts an internal researchTask to the public DTO format
func taskToDTO(t *researchTask) dto.InteractionResponse {
	resp := dto.InteractionResponse{
		ID:        t.ID,
		Status:    string(t.Status),
		Query:     t.Query,
		CreatedAt: t.CreatedAt,
		Error:     t.Error,
	}
	if t.Result != nil {
		resp.Outputs = []dto.InteractionOutput{
			{Text: t.Result.Summary},
		}
		resp.Sources = t.Result.Sources
		resp.Steps = t.Result.Steps
		resp.DurationMs = t.Result.DurationMs
	}
	return resp
}

