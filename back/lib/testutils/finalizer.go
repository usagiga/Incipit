package testutils

// Finalizer is finalizer of some testing values
// Once you received it, you must execute it before ending test.
type Finalizer func()
