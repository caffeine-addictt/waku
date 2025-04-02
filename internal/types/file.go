package types

// StyleResource is a descriptor for a
// file or directory for a Wakue Style.
type StyleResource struct {
	TemplateResourceRelPath string // This is the path relative to the template root
	TemplatedProjectRelPath string // This is the path relative to the 'to-be-templated' project root
}
