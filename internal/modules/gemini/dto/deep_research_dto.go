package dto

import models "gemini-web-to-api/internal/commons/models"

// DeepResearchRequest represents a deep research request
type DeepResearchRequest struct {
	// Query is the research topic or question
	Query string `json:"query"`
	// Model to use (optional, defaults to best available)
	Model string `json:"model,omitempty"`
	// Language for the response (optional, e.g. "en", "vi")
	Language string `json:"language,omitempty"`
	// MaxSources maximum number of web sources to consult (optional)
	MaxSources int `json:"max_sources,omitempty"`
	// Images optional base64 images to use as research context
	Images []models.Attachment `json:"images,omitempty"`
}

// DeepResearchResponse represents a deep research response
type DeepResearchResponse struct {
	// ID of this research session
	ID string `json:"id"`
	// Status: "in_progress", "completed", "failed"
	Status string `json:"status"`
	// Query that was researched
	Query string `json:"query"`
	// Summary is the synthesized research report
	Summary string `json:"summary,omitempty"`
	// Sources used in research
	Sources []ResearchSource `json:"sources,omitempty"`
	// Steps of research performed
	Steps []ResearchStep `json:"steps,omitempty"`
	// Model used
	Model string `json:"model,omitempty"`
	// Error message if status is "failed"
	Error string `json:"error,omitempty"`
	// CreatedAt timestamp (Unix seconds)
	CreatedAt int64 `json:"created_at"`
	// CompletedAt timestamp (Unix seconds), 0 if not completed
	CompletedAt int64 `json:"completed_at,omitempty"`
	// DurationMs total research duration in milliseconds
	DurationMs int64 `json:"duration_ms,omitempty"`
}

// ResearchSource represents a web source consulted during research
type ResearchSource struct {
	// Title of the source page
	Title string `json:"title,omitempty"`
	// URL of the source
	URL string `json:"url,omitempty"`
	// Snippet relevant excerpt from the source
	Snippet string `json:"snippet,omitempty"`
	// Domain of the source
	Domain string `json:"domain,omitempty"`
}

// ResearchStep represents a step in the deep research process
type ResearchStep struct {
	// StepNumber ordinal step number
	StepNumber int `json:"step_number"`
	// Type of step: "search", "read", "synthesize", "outline", "write"
	Type string `json:"type"`
	// Description of what this step did
	Description string `json:"description"`
	// Query used in this step (for search steps)
	Query string `json:"query,omitempty"`
	// Result summary of this step
	Result string `json:"result,omitempty"`
}

// DeepResearchStreamEvent represents a server-sent event during streaming deep research
type DeepResearchStreamEvent struct {
	// Event type: "progress", "source", "step", "result", "error", "done"
	Event string `json:"event"`
	// Message human-readable progress message
	Message string `json:"message,omitempty"`
	// Step details (for "step" events)
	Step *ResearchStep `json:"step,omitempty"`
	// Source details (for "source" events)
	Source *ResearchSource `json:"source,omitempty"`
	// Result is the final research result (for "result" events)
	Result *DeepResearchResponse `json:"result,omitempty"`
	// Progress percentage 0-100
	Progress int `json:"progress,omitempty"`
	// Error message (for "error" events)
	Error string `json:"error,omitempty"`
}

// ---------------------------------------------------------------------------
// Async Interaction API (create + poll pattern)
// ---------------------------------------------------------------------------

// InteractionCreateRequest is the body for POST /deepresearch/create
type InteractionCreateRequest struct {
	// Input is the research query / topic
	Input string `json:"input"`
	// Agent selects the research mode (e.g. "deep-research-pro-preview-12-2025")
	Agent string `json:"agent,omitempty"`
	// Background if true the task runs asynchronously (always true for this endpoint)
	Background bool `json:"background,omitempty"`
	// Stream if true the server returns an SSE stream of InteractionResponse events
	Stream bool `json:"stream,omitempty"`
	// Language for the output report
	Language string `json:"language,omitempty"`
	// MaxSources max number of web sources to consult
	MaxSources int `json:"max_sources,omitempty"`
	// Images optional base64 images to use as research context
	Images []models.Attachment `json:"images,omitempty"`
}

// InteractionOutput holds a single output item in the interaction result
type InteractionOutput struct {
	// Text is the research report text
	Text string `json:"text"`
}

// InteractionResponse is the response for both create and get interaction endpoints
type InteractionResponse struct {
	// ID unique task identifier
	ID string `json:"id"`
	// Status: "in_progress", "completed", "failed"
	Status string `json:"status"`
	// Query that was researched
	Query string `json:"query"`
	// Outputs list of output items (non-empty when completed)
	Outputs []InteractionOutput `json:"outputs,omitempty"`
	// Sources consulted during research
	Sources []ResearchSource `json:"sources,omitempty"`
	// Steps research steps performed
	Steps []ResearchStep `json:"steps,omitempty"`
	// Error message if status is "failed"
	Error string `json:"error,omitempty"`
	// CreatedAt Unix timestamp
	CreatedAt int64 `json:"created_at"`
	// CompletedAt Unix timestamp (0 if not yet completed)
	CompletedAt int64 `json:"completed_at,omitempty"`
	// DurationMs total research duration in milliseconds
	DurationMs int64 `json:"duration_ms,omitempty"`
}
