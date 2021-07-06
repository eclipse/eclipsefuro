package orderedmap

import (
	"container/list"
	"gopkg.in/yaml.v3"
)

func (m *OrderedMap) UnmarshalYAML(value *yaml.Node) error {
	var d yaml.Node
	value.Decode(&d)
	m.pairs = make(map[interface{}]*Pair)
	m.list = list.New()
	// every even n is the key
	for i, n := range d.Content {
		if i%2 == 1 {
			m.Set(d.Content[i-1].Value, n)
		}
	}
	return nil
}

func (m *OrderedMap) MarshalYAML() (interface{}, error) {

	value := yaml.Node{
		Kind: yaml.MappingNode,
	}
	for pair := m.Oldest(); pair != nil; pair = pair.Next() {
		field := &yaml.Node{}
		key := &yaml.Node{}

		field.Encode(pair.Value)
		key.Encode(pair.Key.(string))
		value.Content = append(value.Content, key, field)

	}
	return &value, nil
}
