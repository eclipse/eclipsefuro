package orderedmap

import (
	"encoding/json"
	"strconv"
	"strings"
)

func (m *OrderedMap) MarshalJSON() ([]byte, error) {
	response := []string{}

	for pair := m.Oldest(); pair != nil; pair = pair.Next() {
		v, _ := json.Marshal(pair.Value)
		f := strconv.Quote(pair.Key.(string)) + ":" + string(v)
		response = append(response, f)
	}
	r := "{" + strings.Join(response, ",") + "}"
	return []byte(r), nil
}
