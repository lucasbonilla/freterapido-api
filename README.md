

# freterapido-api
## comandos make
Os comando abaixo devem ser rodados na sequência correta para realizar o build e servir a aplicação em localhost:8080

    # realiza o build da aplicação
    make build
    
    # serve a aplicação em localhost:8080
    make env-serve
    
    # Realiza os testes automatizados
    make test
    
    # para os containers
    make env-stop
    
    # remove os containers
    make env-remove
    

## Endpoints
Para a solução proposta foram desenvolvidos dois endpoints:

 - http://localhost:8080/api/v1/quote
 - http://localhost:8080/api/v1/metrics

### Quote - POST
Recebe um json como o seguinte:

    {
        "recipient":{
            "address":{
                "zipcode":"75106785"
            }
        },
        "volumes":[
            {
                "category":7,
                "amount":1,
                "unitary_weight":5,
                "price":349,
                "sku":"abc-teste-123",
                "height":0.2,
                "width":0.2,
                "length":0.2
            },
            {
                "category":7,
                "amount":2,
                "unitary_weight":4,
                "price":556,
                "sku":"abc-teste-527",
                "height":0.4,
                "width":0.6,
                "length":0.15
            },
            {
                ...
            }
        ]
    }

A aplicação captura o json enviado e realiza uma requisição no endpoint https://sp.freterapido.com/api/v3/quote/simulate
O resultado é então manipulado e armazenado em uma base de dados PostgreSQL.
### Metrics - GET
O endpoint metrics recebe via querystring um ou dois parâmetros não obrigatórios:

 - last_quotes
	 - Faz referência a quantas informações serão buscadas da base de dados
 - page
	 - Paginação do endpoint
	 
A resposta à chamada ao endpoint resulta no seguinte json:

    {
        "message": "operation from route: metrics successful",
        "metrics": {
            "quotes_quantity": {
                "quotes_by_quantity": [
                    {
                        "carrier_name": "CORREIOS",
                        "id_carrier": 281,
                        "number_of_quotes": 22
                    }
                ]
            },
            "total_quotes_price": {
                "total_quotes_by_price": [
                    {
                        "carrier_name": "PRESSA FR (TESTE)",
                        "id_carrier": 346,
                        "total_price_quote": 12795.12
                    }
                ]
            },
            "total_quotes_average_price": {
                "total_quotes_by_average_price": [
                    {
                        "carrier_name": "PRESSA FR (TESTE)",
                        "id_carrier": 346,
                        "average_price_quote": 1599.39
                    }
                ]
            },
            "total_quotes_cheapest_price": {
                "total_quotes_for_cheapest_price": [
                    {
                        "carrier_name": "PRESSA FR (TESTE)",
                        "id_carrier": 346,
                        "price_quote_cheapest": 1599.39
                    }
                ]
            },
            "total_quote_most_expensive_price": {
                "total_quote_by_most_expensive_price": [
                    {
                        "carrier_name": "PRESSA FR (TESTE)",
                        "id_carrier": 346,
                        "price_quote_most_expensive": 1599.39
                    }
                ]
            }
        }
    }
Cada chave faz referência a cada uma das informações solicitadas, sendo elas:
- quotes_quantity:
	- Quantidade de resultados por transportadora.
- total_quotes_price:
	- Total de “preco_frete” por transportadora; (final_price na API).
- total_quotes_average_price:
	- Média de “preco_frete” por transportadora; (final_price na API).
- total_quotes_cheapest_price:
	- O frete mais barato geral.
- total_quote_most_expensive_price:
	- O frete mais caro geral.

A paginação foi adicionada para o caso de muitos resultados. Dessa forma, essa opção pode ser utilizada.
As consultas de frete mais barato geral e mais caro geral foi desenvolvida para buscar a cotação mais barata e mais cara das n primeiras cotações. A primeira cotação dessa chave, por teoria matemática, será a mais barata geral ou mais cara geral.