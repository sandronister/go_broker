package kafka

type ConfigMap map[string]string

type Broker struct {
	server string
	port   int
}

func NewBroker(server string, port int) *Broker {
	return &Broker{server: server, port: port}
}
