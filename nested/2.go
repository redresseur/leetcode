package nested

import (
	"encoding/json"
	"strings"
)

/*
2.Given a “flatten” dictionary object, whose keys are dot-separated.
For example, { ‘A’: 1, ‘B.A’: 2, ‘B.B’: 3, ‘CC.D.E’: 4, ‘CC.D.F’: 5}.
Implement a function in any language to transform it to a “nested” dictionary object.
In the above case, the nested version is like:
{
  ‘A’: 1,
  ‘B’: {
    ‘A’: 2,
    ‘B’: 3,
  },
  ‘CC’: {
    ‘D’: {
      ‘E’: 4,
      ‘F’: 5,
    }
  }
}
It’s guaranteed that no keys in dictionary are prefixes of other keys.
*/

type NestedDict struct {
	root map[string]interface{}
}

func NewNestedDict() *NestedDict {
	return &NestedDict{root: map[string]interface{}{}}
}

func path(k string) []string {
	return strings.Split(k, ".")
}

func (n NestedDict) insert(k string, v interface{}) {
	keys := path(k)
	leaf := n.root
	for i := 0; i < len(keys)-1; i++ {
		if c, isOK := leaf[keys[i]]; !isOK {
			leaf[keys[i]] = map[string]interface{}{}
			leaf = leaf[keys[i]].(map[string]interface{})
		} else {
			leaf = c.(map[string]interface{})
		}
	}

	leaf[keys[len(keys)-1]] = v
}

// 转化给定的数据
func (n NestedDict) Transform(data []byte) (err error) {
	if len(n.root) != 0 {
		n.root = map[string]interface{}{}
	}

	sets := map[string]interface{}{}
	if err = json.Unmarshal(data, &sets); err != nil {
		return
	}

	for k, v := range sets {
		n.insert(k, v)
	}

	return
}

func (n NestedDict) marshal(path string, value interface{}) map[string]interface{} {
	res := map[string]interface{}{}
	if leaf, isOK := value.(map[string]interface{}); isOK {
		for k, v := range leaf {
			if path != "" {
				k = path + "." + k
			}

			pairs := n.marshal(k, v)
			for k, v = range pairs {
				res[k] = v
			}
		}
	} else {
		res[path] = value
	}

	return res
}

// 转换为 { ‘A’: 1, ‘B.A’: 2, ‘B.B’: 3,
// ‘CC.D.E’: 4, ‘CC.D.F’: 5} 格式.
func (n NestedDict) Marshal() ([]byte, error) {
	res := n.marshal("", n.root)
	return json.Marshal(&res)
}
