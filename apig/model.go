package apig

// AssociationNone ...
const (
	AssociationNone      = 0
	AssociationBelongsTo = 1
	AssociationHasMany   = 2
	AssociationHasOne    = 3
)

// Model ...
type Model struct {
	Name   string
	Fields []*Field
}

// AllPreloadAssocs ...
func (m *Model) AllPreloadAssocs() []string {
	result := []string{}

	for _, field := range m.Fields {
		result = append(result, field.PreloadAssocs()...)
	}

	return result
}

// Models ...
type Models []*Model

// Len ...
func (m Models) Len() int {
	return len(m)
}

// Less ...
func (m Models) Less(i, j int) bool {
	return m[i].Name < m[j].Name
}

// Swap ...
func (m Models) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

// Field ...
type Field struct {
	Name        string
	JSONName    string
	Type        string
	Tag         string
	Association *Association
}

// PreloadAssocs ...
func (f *Field) PreloadAssocs() []string {
	if f.Association == nil || f.Association.Type == AssociationNone {
		return []string{}
	}

	result := []string{
		f.Name,
	}

	for _, field := range f.Association.Model.Fields {
		if field.Association == nil || field.Association.Type == AssociationNone {
			continue
		}

		result = append(result, f.Name+"."+field.Name)
	}

	return result
}

// IsAssociation ...
func (f *Field) IsAssociation() bool {
	return f.Association != nil && f.Association.Type != AssociationNone
}

// IsBelongsTo ...
func (f *Field) IsBelongsTo() bool {
	return f.Association != nil && f.Association.Type == AssociationBelongsTo
}

// Association ...
type Association struct {
	Type  int
	Model *Model
}
