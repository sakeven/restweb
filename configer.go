package restweb

type Configer interface {
	ReadConfig(configFile string)
	Get(key string) interface{}
}

var Configerd Configer

func RegisterConfiger(uConfiger Configer) {
	Configerd = uConfiger
}
