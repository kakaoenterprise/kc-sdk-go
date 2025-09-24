package network

// AnyOf is a fallback for inline `anyOf` schemas that weren't emitted as concrete models.
// It preserves arbitrary JSON payloads (objects, arrays, scalars).
// NOTE: If the generator later emits a real AnyOf model, remove or rename this file to avoid conflicts.
type AnyOf map[string]interface{}
