#!/bin/bash

# Total de requisições
TOTAL=50

# Quantas em paralelo
PARALLEL=1

# URL do endpoint
URL="http://localhost:8080/product"

# Token JWT
TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJjZDkzMjQ3NC0zNDk0LTRiZWItOTJlMi1hZDE1ODYzYWE4ZDYiLCJleHAiOjE3NTM3NDYxNDJ9.1vsZBI6xctA2UnlddaZlosovGHuGD7DUQ12pGcMnxfY"

# Cabeçalhos
HEADER_CONTENT_TYPE="Content-Type: application/json"
HEADER_TOKEN="Token: $TOKEN"

# Função para enviar requisição com nome único
send_request() {
  i=$1
  NAME="lareira-$i"
  BODY=$(printf '{"name": "%s", "price": 5000, "quantity": 2}' "$NAME")

  curl -s  /dev/null -X POST "$URL" \
    -H "$HEADER_CONTENT_TYPE" \
    -H "$HEADER_TOKEN" \
    -d "$BODY"

  echo "Enviado produto: $NAME"
}

export -f send_request
export URL HEADER_CONTENT_TYPE HEADER_TOKEN

# Executar em paralelo
seq $TOTAL | xargs -n1 -P $PARALLEL -I{} bash -c 'send_request "$@"' _ {}
