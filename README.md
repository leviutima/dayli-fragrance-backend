# Dayli Fragrance — Backend

API REST do e-commerce de perfumes Dayli Fragrance. Construída com Go, Clean Architecture e princípios SOLID.

---

## Instalação

Pré-requisitos: Go 1.26+ e Docker.

```bash
# Subir o banco de dados
docker compose up -d

# Copiar variáveis de ambiente
cp .env.example .env

# Rodar a aplicação
go run ./cmd/api/main.go
```

A API sobe em `http://localhost:8080`.

---

## Arquitetura — Clean Architecture

O projeto segue Clean Architecture, organizando o código em camadas com regras estritas de dependência. A regra central é simples: **dependências apontam sempre para dentro**. Camadas externas conhecem as internas, nunca o contrário.

```
presentation  →  application  →  domain
infrastructure →  application  →  domain
```

`domain` não importa de ninguém. É o núcleo do sistema.

### Estrutura de pastas

```
cmd/
└── api/
    └── main.go                  → ponto de entrada, composição de dependências

internal/
├── domain/
│   ├── entity/                  → structs que representam o negócio
│   └── repository/              → interfaces de acesso a dados (contratos)
│
├── application/
│   └── usecase/                 → casos de uso (regras de negócio)
│
├── infrastructure/
│   ├── database/                → implementações dos repositórios (SQL, Postgres)
│   └── http/                    → configuração do servidor HTTP
│
└── presentation/
    └── handler/                 → handlers HTTP (recebem requisição, chamam use case, devolvem resposta)

config/                          → leitura de variáveis de ambiente
```

### Por que essa separação?

Em projetos sem arquitetura, qualquer arquivo acessa o banco diretamente, mistura regra de negócio com SQL e acopla HTTP com persistência. Trocar o banco ou o framework HTTP exige reescrever metade do sistema.

Com Clean Architecture:

- Trocar Postgres por outro banco → cria nova implementação em `infrastructure/database`, o resto não muda
- Trocar Chi por outro roteador → muda só `infrastructure/http` e `presentation/handler`
- Regras de negócio ficam em `application/usecase`, sem dependência de banco ou HTTP

---

## Princípios SOLID na prática

### S — Single Responsibility
Cada arquivo tem uma única responsabilidade. O handler só trata HTTP. O use case só executa a regra de negócio. O repositório só fala com o banco.

### O — Open/Closed
Para adicionar um novo endpoint, você cria um novo handler e um novo use case — sem modificar o que já existe.

### D — Dependency Inversion
Use cases dependem de **interfaces** de repositório, nunca de implementações concretas. O Postgres é um detalhe de infraestrutura invisível para a camada de aplicação.

```
// use case depende da interface (domínio)
type GetProductsUseCase struct {
    repo repository.ProductRepository   // interface
}

// implementação concreta fica na infraestrutura
type ProductRepositoryPostgres struct {
    connection *pgx.Conn               // detalhe de implementação
}
```

A injeção de dependências acontece no `main.go` — único lugar do sistema que conhece todas as camadas.

---

## Fluxo de uma requisição

```
GET /api/products
    → Chi Router
        → ProductHandler.GetAll
            → GetProductsUseCase.Execute
                → ProductRepositoryPostgres.FindAll
                    → PostgreSQL
                        → []Product (JSON)
```

Cada camada recebe o resultado da anterior e devolve para a próxima. Nenhuma camada pula outra.

---

## Banco de dados

PostgreSQL 16 rodando em Docker. O schema é versionado em arquivos `.sql` numerados dentro de `internal/infrastructure/database/migrations/`.

Para executar uma migration:

```bash
Get-Content migrations/001_create_tables.sql | docker exec -i dayli_postgres psql -U dayli -d dayli_db
```

O banco não usa ORM. Queries são escritas em SQL puro com o driver `pgx`. Isso é intencional — SQL explícito é mais previsível, mais performático e mais fácil de otimizar do que queries geradas por ORM.

---

## Entidades de Domínio

### `Fragrance`
Representa o perfume como **composição olfativa** — família, notas, intensidade, estações e ocasiões. É a entidade mais central do sistema: tudo que envolve recomendação parte dela.

- `FragranceFamily` — floral, woody, citrus, oriental, fresh, aquatic, gourmand
- `FragranceNotes` — notas de topo, coração e base (piramide olfativa)
- `Intensity` — light, moderate, strong, intense
- `Season`, `Occasion` — usados pelo algoritmo de recomendação

### `Product`
Representa o perfume como **item comercial** — preço, estoque, SKU, volume, slug. Um `Fragrance` pode ter múltiplos `Products` (ex: mesma fragrância em 30ml, 50ml e 100ml).

### `Customer`
Cliente final com acesso ao e-commerce. Possui endereços, wishlist, histórico de visualizações e preferências olfativas — todas alimentam o recomendador.

### `Order`
Pedido de compra com itens, status e total.

- `OrderStatus` — pending → confirmed → shipped → delivered | cancelled

### `User`
Funcionário ou admin com acesso ao backoffice e ERP.

- `UserRole` — admin, manager, staff

### `Recommendation`
Resposta do algoritmo de recomendação.

- `score` — valor entre 0 e 1 indicando confiança da recomendação
- `reason` — viewing_history, fragrance_match, similar_customers, trending, seasonal

---

## Recomendador de Perfumes

O recomendador é **híbrido** — combina dois algoritmos:

1. **Fragrance Vectors** — converte notas olfativas em vetores numéricos e calcula similaridade cosseno com o histórico de visualizações do usuário
2. **Collaborative Filtering** — usuários com comportamento similar tendem a gostar dos mesmos perfumes

O recomendador recebe eventos de comportamento do frontend (`product_view`, `product_wishlist`, `product_cart`, `recommendation_click`, `search`) com pesos diferentes, e usa o `sessionId` do cookie `dayli_session` para rastrear usuários anônimos antes do login.

---

## Autenticação

A autenticação usa cookies `httpOnly` — o browser os envia automaticamente em toda requisição, e JavaScript não consegue lê-los (imune a XSS).

- `customer_token` — cookie do cliente final (e-commerce)
- `user_token` — cookie do funcionário (backoffice/ERP), com campo `role`
- `dayli_session` — cookie de sessão anônima para rastreamento de comportamento

O frontend (Next.js) atua como proxy — o browser nunca fala diretamente com o Go. Toda requisição passa pelo Next.js, que valida autenticação e repassa o cookie ao Go. O Go valida novamente no lado dele.

---

## Endpoints

Todos os endpoints têm prefixo `/api`.

```
GET  /api/health
GET  /api/products
GET  /api/product/:slug
GET  /api/fragrances
GET  /api/fragrances/:id
GET  /api/customers
GET  /api/customers/:id
GET  /api/customers/me
GET  /api/orders
GET  /api/orders/:id
GET  /api/users
GET  /api/users/:id
GET  /api/recommendations
POST /api/events
```

---

## Stack

| Tecnologia | Uso |
|---|---|
| Go 1.26 | Linguagem |
| Chi | Roteador HTTP |
| pgx v5 | Driver PostgreSQL |
| godotenv | Leitura de `.env` |
| PostgreSQL 16 | Banco de dados |
| Docker | Infraestrutura local |

---

## Variáveis de Ambiente

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=dayli
DB_PASSWORD=dayli123
DB_NAME=dayli_db

SERVER_PORT=8080
```

---

## Convenções

- Arquivos em `snake_case` (`product_repository.go`, `get_products.go`)
- Tipos, structs e funções exportadas em `PascalCase`
- Campos públicos de struct em `PascalCase` — campos privados em `camelCase`
- Funções construtoras seguem o padrão `NewXxx` — Go não tem `constructor`
- Tratamento de erro explícito em todo retorno — Go não usa `try/catch`
- Variáveis de erro sempre chamadas `err` por convenção da comunidade
- `context.Background()` em todo I/O — permite cancelamento futuro
- `defer` para fechar recursos (conexões, rows) — equivalente ao `finally`
