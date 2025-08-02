#!/bin/bash

# Total de requisições
TOTAL=20

# Quantas em paralelo
PARALLEL=10

# URL do endpoint
URL="http://localhost:8080/product"

# Token JWT
TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJmMzE5NGE2Ni1iYzExLTQxYjktYjZjYi04NmExZjI2ZjhlMzQiLCJleHAiOjE3NTQxNDg1MjN9.RsrGCA6gE0F9DXQ40CPwWf0z_SDAk_yAHNq1D-9vk00"

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
