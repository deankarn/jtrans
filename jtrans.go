package jtrans

// Transformation contains all information for a json-json transformation
type Transformation []*Map

// Transform applies the transformation
func (t Transformation) Transform(jv map[string]interface{}) map[string]interface{} {
	results := make(map[string]interface{})
	var fromIface, toIface interface{}
	var ok bool
	var current map[string]interface{}

OUTER:
	for _, m := range t {
		if len(m.From) == 0 || len(m.To) == 0 {
			continue OUTER
		}
		switch m.Type {
		case Constant:
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
