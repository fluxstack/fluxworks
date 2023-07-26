package log

type Field struct {
	K string
	V interface{}
}

func F(k string, v interface{}) Field {
	return Field{K: k, V: v}
}

type Fields []Field

func (fs Fields) Field(k string, v interface{}) Fields {
	nf := fs
	nf = append(nf, F(k, v))
	return nf
}

func (fs Fields) ToMap(k string, v interface{}) map[string]interface{} {
	m := map[string]interface{}{}
	for _, f := range fs {
		m[f.K] = f.V
	}
	return m
}

//
//func Field(k string, v interface{}) Fields {
//	return Fields{
//		k: v,
//	}
//}
