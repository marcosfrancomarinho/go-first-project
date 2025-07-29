#!/bin/bash

# Total de requisições
TOTAL=100

# Quantas em paralelo
PARALLEL=10

# URL do endpoint
URL="https://tool-backend-na56.onrender.com/product"

# Token JWT
TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIwMjlhOWY0My05MTQzLTRkNjQtYjA0MC04Y2I4ODVhZWVlNzYiLCJleHAiOjE3NTM3NTQxMDB9.raeLM4kO27Vk8-HtBUMmMXEB7vn7AjyZ-woBkm6qCm0"

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
