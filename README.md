## Sobre projeto


Este projeto visa ilustrar o versionamento de API através de headers. Para tanto, oferece uma API simplificada, 
desenvolvida em Go, que inclui uma rota principal.

---

<details>
 <summary><code>GET</code><code><b>/data</b></code><code>(obtém uma mensagem contendo um autor ou não)</code></summary>

##### Headers

> | name          | type     | data type | description                                                      |
> |---------------|----------|-----------|------------------------------------------------------------------|
> | X-Api-Version | required | string    | Usar o [versionamento semântico](https://semver.org/lang/pt-BR/) |

##### Responses

> | http code | content-type       | response                                                          | version |
> |-----------|--------------------|-------------------------------------------------------------------|---------|
> | `200`     | `application/json` | `{ "message": "Hello from the data store!", "author": "Elias" } ` | v0.1.1  |
> | `200`     | `application/json` | `{ "message": "Hello from the data store!" } `                    | v0.1.0  |

##### Example cURL

> ```shell
> curl -H "X-Api-Version: v0.1.1" http://localhost:8080/data
> ```

</details>

---

Esta técnica de versionamento possibilita que os usuários incluam o número da versão como um cabeçalho na solicitação 
da API, separando assim a versão da API da estrutura da URL.

## Pré requisitos

* Instale o [Go](https://go.dev/dl/) com versão igual ou superior a 1.22
* Tenha disponível um editor de código. Algumas opções incluem [Neovim](https://neovim.io/), [Zed](https://zed.dev/)
ou [VS Code](https://code.visualstudio.com/), No desenvolvimento deste projeto, utilizei o 
[GoLand](https://www.jetbrains.com/pt-br/go/promo/).

## Como executar

**#1**: Clone o repositório no seu computador e acesse o diretório raiz da aplicação.

---

**#2**: Inicie a aplicação utilizando o comando abaixo.

```shell
go run .
```

---

## Referências

- Mais sobre versionamento de APIs: [What is API versioning?](https://www.postman.com/api-platform/api-versioning/)
- Como funciona o versionamento semântico: [Semantic Versioning 2.0.0](https://semver.org)
- Exemplo de como o GitHub versiona suas APIs: [API Versions](https://docs.github.com/en/rest/about-the-rest-api/api-versions?apiVersion=2022-11-28)
- Ótimo livro sobre design de APIs: [API Design Patterns](https://www.manning.com/books/api-design-patterns)
- Como a equipe da Apigee no Google Cloud aborda o tema de versionamento: [Versioning in API design](https://cloud.google.com/blog/products/api-management/api-design-which-version-of-versioning-is-right-for-you?m=1)