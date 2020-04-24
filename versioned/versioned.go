// Package versioned provides the versioned interface and semantic versioning constants
package versioned

// Versioned describes a system that can be versioned and promoted between different 'gates'
type Versioned interface {
	// Ready verifies that the system can be operated on and versioned correctly.
	Ready() bool

	// Release increments the version.
	Release() error
	// Promote promotes the current version to a new stage - moving past a 'gate'.
	Promote(stage string) error
	// Version returns the highest version. Varying at what it looks at depending on the optional provided modifiers.
	Version(stage string, headOnly bool) (string, error)

	// Distribute distributes the current state of the system to remotes.
	Distribute() error
}

// Semantic version field - X.Y.Z
type Field int32

const (
	Major Field = 0 // X
	Minor Field = 1 // Y
	Patch Field = 2 // Z
)
