# Projeto Go Broker
### Visão Geral
O projeto Go Broker é uma implementação de um broker de mensagens em Go, que suporta múltiplos backends, incluindo Kafka e Redis. Este projeto foi desenvolvido para fornecer uma solução flexível e extensível para o consumo e produção de mensagens em diferentes sistemas de mensageria.

### Motivação
A motivação principal para a criação deste projeto é a necessidade de uma solução de broker de mensagens que possa ser facilmente configurada para trabalhar com diferentes backends de mensageria. Muitas aplicações modernas dependem de sistemas de mensageria para comunicação assíncrona entre serviços, e ter uma solução que suporte múltiplos backends permite maior flexibilidade e adaptabilidade às necessidades específicas de cada aplicação.

## Instalação

Para instalar o projeto Go Broker como uma biblioteca em seu projeto Go, siga os passos abaixo:

### Adicione o módulo ao seu projeto:

No diretório do seu projeto, execute o comando abaixo para adicionar o módulo Go Broker como uma dependência:

```go 
go get github.com/sandronister/go_broker
````

## Exemplo

``` go

package main

import (
    "log"
    "github.com/seu-usuario/go_broker/pkg/broker/factory"
    "github.com/seu-usuario/go_broker/pkg/broker/types"
)

func main() {
    // Configuração do broker
    config := types.BrokerConfig{
        Type: "kafka", // Pode ser "kafka" ou "redis"
        KafkaConfig: types.KafkaConfig{
            Brokers: []string{"localhost:9092"},
            Topic:   "example-topic",
        }
    }

    // Criação do broker usando a fábrica
    broker, err := factory.NewBroker(config)
    if err != nil {
        log.Fatalf("Erro ao criar broker: %v", err)
    }

    // Envio de uma mensagem
    err = broker.Produce("chave", []byte("mensagem"))
    if err != nil {
        log.Fatalf("Erro ao enviar mensagem: %v", err)
    }

    log.Println("Mensagem enviada com sucesso!")
}
```
