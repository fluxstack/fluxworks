package sagas

// SagaData interface
type SagaData interface {
	SagaDataName() string
}

type mapSagaData struct {
	name   string
	extras map[string]interface{}
}

func NewMapSagaData(name string, extras map[string]interface{}) SagaData {
	return &mapSagaData{name: name, extras: extras}

}

func (sd *mapSagaData) SagaDataName() string {
	return sd.name
}

func (sd *mapSagaData) Set(k string, v interface{}) {
	sd.extras[k] = v
}

func (sd *mapSagaData) Get(k string) (interface{}, bool) {
	v, ok := sd.extras[k]
	return v, ok
}
