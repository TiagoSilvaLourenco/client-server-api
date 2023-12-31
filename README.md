## Para rodar o projeto

### Client:

Na pasta do projeto Client rode: go run main.go
Após o comando ele deve fazer a requisição
Receber a resposta
Salvar em um arquivo cotacao.txt a informação recebida

### Server:

Na pasta do projeto Server rode: go run main.go

Para testar o server apos rodar o comando acima abra outro terminal e use o comando: curl localhost:8080/cotacao

# client-server-api

Neste desafio vamos aplicar o que aprendemos sobre webserver http, contextos,
banco de dados e manipulação de arquivos com Go.

Você precisará nos entregar dois sistemas em Go:

- client.go
- server.go

Os requisitos para cumprir este desafio são:

[x] O client.go deverá realizar uma requisição HTTP no server.go solicitando a cotação do dólar.

[x] O server.go deverá consumir a API contendo o câmbio de Dólar e Real no endereço: https://economia.awesomeapi.com.br/json/last/USD-BRL e em seguida deverá retornar no formato JSON o resultado para o cliente.

[x] Usando o package "context", o server.go deverá registrar no banco de dados SQLite cada cotação recebida, sendo que o timeout máximo para chamar a API de cotação do dólar deverá ser de 200ms e o timeout máximo para conseguir persistir os dados no banco deverá ser de 10ms.

[x] O client.go precisará receber do server.go apenas o valor atual do câmbio (campo "bid" do JSON).

[x] Utilizando o package "context", o client.go terá um timeout máximo de 300ms para receber o resultado do server.go.

[x] Os 3 contextos deverão retornar erro nos logs caso o tempo de execução seja insuficiente.

[X] O client.go terá que salvar a cotação atual em um arquivo "cotacao.txt" no formato: Dólar: {valor}

[x] O endpoint necessário gerado pelo server.go para este desafio será: /cotacao e a porta a ser utilizada pelo servidor HTTP será a 8080.

[x] Ao finalizar, envie o link do repositório para correção.
