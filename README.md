

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
    
## Desenvolvimento
Para o desenvolvimento da solução foi escolhido a arquitetura hexagonal com  ports e adapters. Dessa forma, cada pacote do sistema fica encapsulado permitindo uma maior desacoplagem entre os recursos do sistema.
Essa independência entre os pacotes permite que o sistema tenha refatorações ou mudanças de tecnologias sem afetar os demais módulos.
Por exemplo, se no futuro a base de dados passe a utilizar outra solução ao invés de PostgreSQL, basta alterar o pacote postgres para outra tecnologia e realizar ajustes básicos para adequar o projeto como um todo para essa nova solução.
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

## Testes
Os testes do sistema, escolheu-se os testes unitários na aplicação. Contudo, para termos de comprovação, somente um dos arquivos foi feita a cobertura dos testes, Os demais testes dos demais pacotes que não foram desenvolvidos seguem a mesma lógica de aplicação.
A arquitetura hexagonal favorece os testes do sistema pois dessa forma é possível realizar o mock dos dados e gerar uma chamada personalizada a um método de alguma biblioteca.
Um exemplo que gosto de exemplificar são as chamadas à base de dados em que não é necessário testar se o método db.Query realmente está armazenado dados na tabela ou buscando os mesmos da persistência.
Espera-se que as bibliotecas, que são amplamente utilizadas, estejam consistentes e muito bem testadas. Dessa forma não é necessário utilizar uma chamada real para a base de dados e nem subir um container específico somente para os testes.
Um exemplo mais claro, utilizado aqui nessa solução, é o método Do da biclioteca [http](https://pkg.go.dev/net/http) em que é executada uma requição http para a URL informada no método [http.NewRequestWithContext](https://pkg.go.dev/net/http#NewRequestWithContext) [pacote internal/adapters/handler/routes/quote/quote.go]. Não é necessário de fato enviar a requisição para um endpoint verdadeiro para realizar os testes. Bascara criar o mock do método em um MockedAdapter, como abaixo:

	mockedHttp = &httpH.MockedAdapter{
		DoFn: func(req *http.Request) (*http.Response, error) {
			return &http.Response{}, nil
		},
		CloseFn: func() error {
			return nil
		},
	}
Quando a aplicação, em modo de teste, realizar a chamada ao método Do, na verdade ela irá invocar o método DoFn que está parametrizado com um retorno personalizado.
Da mesma forma, para testar o erro da chamada, basta informar que no parâmetro error, deve ser retornado um erro. Como se o método Do do pacote http estivesse informando um erro real quando na verdade é um erro simulado, como exemplificado abaixo:

	mockedHttp = &httpH.MockedAdapter{
		DoFn: func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("Ocorreu um erro")
		},
	}
Por fim, para o pacote quote [internal/adapters/handler/routes/quote/quote.go] foram realizados os testes com cobertura de 100%. Isso pode ser conferido executando o comando make test na raiz do projeto.