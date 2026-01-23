# Mini E-commerce gRPC (Go + Hexagonal + Postgres)
Três microsserviços — **Order → Payment → Shipping** — em Go, falando via gRPC e persistindo em Postgres. O `docker-compose` sobe tudo de uma vez, com um DB por serviço.

## Arquitetura (hexagonal, resumão)
- `core/domain`: modelos puros (`Order`, `Payment`, `Shipping`) e regras (ex.: `TotalPrice`, cálculo de entrega).
- `core/api`: orquestra o caso de uso (ex.: `PlaceOrder`) sem saber de transporte ou banco.
- `ports`: interfaces que o core espera (DB, payment client, shipping client, gRPC server).
- `adapters/db`: GORM + Postgres implementando os ports de persistência.
- `adapters/grpc`: servers que expõem os casos de uso.
- `adapters/clients`: gRPC clients para Payment e Shipping, injetados no core. Tudo troca pela interface, deixando o core desacoplado.

## Fluxo CreateOrder (feliz/triste)
1) Order recebe `Create` com itens e cliente.  
2) Valida estoque por item e soma total das quantidades (limite 50).  
3) Persiste como `Pending`.  
4) Cobra no Payment; se falhar, marca `Canceled`.  
5) Se pago, marca `Paid`.  
6) Solicita frete no Shipping; se falhar, marca `PaymentSuccessfulShippingFailed`.  
7) Se ok, marca `Shipped` e devolve `order_id`, `shipping_id`, `delivery_days`.

## Regras de negócio
- Soma das quantidades > 50 → `InvalidArgument` (order é cancelado).
- `total_price` > 1000 no Payment → `InvalidArgument`.
- Produto inexistente no estoque → `NotFound`.
- Prazo de entrega: mínimo 1 dia + 1 dia a cada 5 itens (`floor(total/5)`).
- O total é calculado no servidor a partir dos itens; o cliente não envia `total_price`.

## Resiliência
- Cliente Payment tem timeout de 2s.  
- Retries só em `Unavailable` ou `ResourceExhausted`, com backoff linear de 1s, até 5 tentativas.  
- Credenciais inseguras (sem TLS) porque é ambiente de laboratório.

## Como rodar (compose)
```bash
cd microservices
docker-compose down -v   # limpa volumes e DBs
docker-compose up --build
```
Portas gRPC: Order `localhost:50051`, Payment `localhost:50052`, Shipping `localhost:50053`.  
Env vars já vêm do compose (`DATABASE_URL` por serviço, `APPLICATION_PORT`, URLs internas `payment:50052`, `shipping:50053`).

## Testes rápidos com grpcurl
Use `-plaintext` (sem TLS). O service name é o mesmo do proto (`Order`, `Payment`, `Shipping`).

1) Pedido OK  
```bash
grpcurl -plaintext -d '{
  "customer_id": 1,
  "order_items": [
    {"product_code": "PROD001", "unit_price": 99.99, "quantity": 2},
    {"product_code": "PROD002", "unit_price": 149.99, "quantity": 1}
  ]
}' localhost:50051 Order.Create
```

2) Produto inexistente (espera `NotFound`)  
```bash
grpcurl -plaintext -d '{
  "customer_id": 1,
  "order_items": [{"product_code": "PROD999", "unit_price": 10, "quantity": 1}]
}' localhost:50051 Order.Create
```

3) Total > 1000 (gatilha validação do Payment: `InvalidArgument`)  
```bash
grpcurl -plaintext -d '{
  "customer_id": 1,
  "order_items": [{"product_code": "PROD001", "unit_price": 600, "quantity": 2}]
}' localhost:50051 Order.Create
```

## Trade-offs e anotações
- Ambiente de lab sem TLS; credenciais inseguras nos dials gRPC.  
- O preço total é calculado server-side; o campo no request é ignorado como fonte de verdade.  
- Postgres roda em um container e cria `order_db`, `payment_db`, `shipping_db` via `init.sql` (idempotente).  
- Pensado para facilitar demo local; para prod, ajuste TLS, observabilidade e secrets. 
