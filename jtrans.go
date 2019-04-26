package jtrans

// ValueType  the type of transform value
type ValueType string

const (
	Mapping  ValueType = "mapping"
	Constant ValueType = "constant"
)

// Map contains a transformation of one field value
type Map struct {
	Type ValueType
	From []string // depending on type a constant or mapping by key
	To   []string
}

// Transformation contains all information for a json-json transformation
type Transformation []Map

// Transform applies the transformation
func Transform(t Transformation, jv map[string]interface{}) map[string]interface{} {
	results := make(map[string]interface{})
	var fromIface, toIface interface{}
	var ok bool
	var current map[string]interface{}

OUTER:
	for _, m := range t {
		switch m.Type {
		case Constant:
			if len(m.From) == 0 {
				continue OUTER
			}
			fromIface = m.From[0]
		case Mapping:
			fromIface = jv
			for _, k := range m.From {
				current, ok = fromIface.(map[string]interface{})
				if !ok {
					continue OUTER
				}

				fromIface, ok = current[k]
				if !ok {
					continue OUTER
				}
			}
		default:
			continue OUTER
		}

		toIface = results
		for i, k := range m.To {
			current, ok = toIface.(map[string]interface{})
			if !ok {
				continue OUTER
			}
			toIface, ok = current[k]
			if !ok {
				if i == len(m.To)-1 {
					current[m.To[len(m.To)-1]] = fromIface
				} else {
					toIface = make(map[string]interface{})
				}
			}
		}
	}
	return results
}
