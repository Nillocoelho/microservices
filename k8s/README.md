# Kubernetes deployment for the mini e-commerce

## Imagens
Construa e publique (ou carregue em um cluster local como kind):

```bash
# na raiz do repo (GO-Projeto)
docker build -t microservices/order:latest -f microservices/Dockerfile --build-arg SERVICE=order --build-arg SERVICE_PORT=50051 .
docker build -t microservices/payment:latest -f microservices/Dockerfile --build-arg SERVICE=payment --build-arg SERVICE_PORT=50052 .
docker build -t microservices/shipping:latest -f microservices/Dockerfile --build-arg SERVICE=shipping --build-arg SERVICE_PORT=50053 .
# se usar kind:
kind load docker-image microservices/order:latest microservices/payment:latest microservices/shipping:latest
```

## Aplicar
```bash
kubectl apply -f microservices/k8s/stack.yaml
kubectl get pods -n microservices
```

## Testes rápidos (grpcurl)
Use `-plaintext` e faça port-forward do service:

```bash
kubectl port-forward svc/order 50051:50051 -n microservices &
grpcurl -plaintext -d '{
  "customer_id": 1,
  "order_items": [
    {"product_code": "PROD001", "unit_price": 99.99, "quantity": 2},
    {"product_code": "PROD002", "unit_price": 149.99, "quantity": 1}
  ]
}' localhost:50051 Order.Create
```

## Notas
- Probes são TCP porque os serviços não expõem `grpc_health`; pode adicionar depois.
- Postgres é StatefulSet com init SQL via ConfigMap; ajuste storage class conforme o cluster.
- As envs seguem o docker-compose original; personalize senhas/URLs via Secret/ConfigMap.
