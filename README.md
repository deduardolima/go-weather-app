# Go-Weather-App

## Descrição

Este é um sistema em Go que receba um CEP, identifica a cidade e retorna o clima atual (temperatura em graus celsius, fahrenheit e kelvin). Está disponivel para consulta em:

https://go-weather-app-rhaqnpdasa-uc.a.run.app/

## Funcionalidades

- Recebe um CEP válido de 8 dígitos.
- Identifica a cidade correspondente ao CEP.
- Retorna a temperatura atual em graus Celsius, Fahrenheit e Kelvin.
- Configurado para fácil implantação usando Docker e Docker Compose.


## Estrutura do Projeto


- go-weather-app/
  - cmd/
    - main.go
    - .env
  - internal/
    - location/
      - client.go
      - service.go
    - weather/ 
      - client.go 
      - handler_test.go
      - handler.go
  - docker-compose.yml
  - Dockerfile
  - go.mod
  - go.sum
  - README.md


## Pré-requisitos

- Docker
- Docker Compose

## Configuração
```
git clone https://github.com/deduardolima/go-weather-app.git
cd go-weather-app

```

No arquivo `.env` inclua suas credenciais para a API WeatherAPI. Se você ainda não tiver uma chave de API, crie uma conta para obter acesso em:
[WEATHER API](https://www.weatherapi.com/)

```
WEATHER_API_KEY=SUA-API-KEY-AQUI
PORT=8080

```

## Instalação e Execução com Docker
Construa e inicie os containers:
```
docker-compose up --build -d
```

isso irá construir a imagem do aplicativo e iniciar o serviço definido no docker-compose.yml

## Testes

Para executar os testes automatizados, use o seguinte comando:

```
go test ./...

```

## Utilização 

Para utilizar o serviço, basta fazer uma requisição HTTP com um CEP válido. Por exemplo:


```
http://localhost:8080/weather?cep=80010100

```

Para utilizar a versão pblicada no Google Cloud Run:

```
https://go-weather-app-rhaqnpdasa-uc.a.run.app/weather?cep=82940110

```

## Exemplo de Resposta

Em caso de sucesso

```
{
  "temp_C": 19,
  "temp_F": 66.2,
  "temp_K": 292.15
}

```

Em caso de falha, quando o CEP não é válido (com formato correto):

```
{
  "error": "invalid zipcode"
}

```

Em caso de falha, quando o CEP não é encontrado:

```
{
  "error": "can not find zipcode"
}

```


## Créditos

Este projeto foi criado por [Diego Eduardo](http://github.com/deduardolima)







