package log

type Fields map[string]interface{}

func (f Fields) Field(k string, v interface{}) Fields {
	f[k] = v
	return f
}

func Field(k string, v interface{}) Fields {
	return Fields{
		k: v,
	}
}
