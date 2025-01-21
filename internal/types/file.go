package types

type StyleResourceKind string

const (
	FileStyleResourceKind StyleResourceKind = "file"
	DirStyleResourceKind  StyleResourceKind = "dir"
)

// StyleResource is a descriptor for a
// file or directory for a Wakue Style.
type StyleResource struct {
	Kind                    StyleResourceKind // Denotes the kind of resource
	StyleRelPath            string            // This is the path relative to the style root
	TemplatedProjectRelPath string            // This is the path relative to the 'to-be-templated' project root
}
