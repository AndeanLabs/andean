#!/bin/bash

# setup-reviewer-local.sh - Setup completo para pruebas locales de Andean Chain

set -e

echo "🚀 Configurando Andean Chain para pruebas locales..."

# --- Verificación de Prerrequisitos ---
echo "🔍 Verificando prerrequisitos..."
command -v docker >/dev/null 2>&1 || { echo "❌ Docker no instalado. Instálalo desde https://docs.docker.com/get-docker/." >&2; exit 1; }
command -v git >/dev/null 2>&1 || { echo "❌ Git no instalado. Instálalo primero." >&2; exit 1; }

echo "✅ Prerrequisitos cumplidos."

# --- Clonar y Construir ---
REPO_URL="https://github.com/AndeanLabs/andean.git"
if [ ! -d "andean" ]; then
    echo "📥 Clonando repositorio..."
    git clone "$REPO_URL"
fi

cd andean

echo "🏗️  Construyendo imagen Docker (puede tardar varios minutos)..."
docker build -t andean-review . > /dev/null

# --- Preparar Contenedor ---
echo "🐳 Iniciando contenedor..."
CONTAINER_NAME="andean-review"

# Limpiar contenedor anterior
docker stop "$CONTAINER_NAME" >/dev/null 2>&1 || true
docker rm "$CONTAINER_NAME" >/dev/null 2>&1 || true

# Ejecutar con montaje de volumen para usar config.yml local
docker run -d --name "$CONTAINER_NAME" \
    -v $(pwd):/workspace \
    -p 1317:1317 \
    -p 26657:26657 \
    andean-review sleep infinity > /dev/null

# Verificar contenedor
sleep 5
if ! docker ps | grep -q "$CONTAINER_NAME"; then
    echo "❌ Contenedor no corriendo. Logs:"
    docker logs "$CONTAINER_NAME"
    exit 1
fi

echo "✅ Contenedor listo."

# --- Configurar Cadena ---
echo "⚙️  Configurando cadena..."

# Función para ejecutar en contenedor
exec_in_container() {
    docker exec "$CONTAINER_NAME" bash -c "$1"
}

exec_in_container "go install ./cmd/andeand"
exec_in_container "andeand init reviewer-demo --chain-id andean-demo-1 --home /workspace/.andean"
exec_in_container "andeand keys add reviewer --keyring-backend test --home /workspace/.andean"
exec_in_container "andeand genesis add-genesis-account reviewer 1000000000000aand --keyring-backend test --home /workspace/.andean"
exec_in_container "andeand genesis gentx reviewer 1000000000aand --chain-id andean-demo-1 --keyring-backend test --home /workspace/.andean"
exec_in_container "andeand genesis collect-gentxs --home /workspace/.andean"

# --- Iniciar Cadena ---
echo "🔥 Iniciando cadena..."
docker exec -d "$CONTAINER_NAME" bash -c "andeand start --home /workspace/.andean --minimum-gas-prices 0stake"

sleep 8

# Verificar
if exec_in_container "andeand status --node tcp://localhost:26657" >/dev/null 2>&1; then
    REVIEWER_ADDR=$(exec_in_container "andeand keys show reviewer -a --keyring-backend test --home /workspace/.andean")
    echo "✅ ¡Setup completo!"
    echo "Dirección: $REVIEWER_ADDR"
    echo "RPC: http://localhost:26657"
    echo "API: http://localhost:1317"
    echo "🧪 Prueba: docker exec -it $CONTAINER_NAME andeand status --node tcp://localhost:26657"
    echo "🛑 Detener: docker stop $CONTAINER_NAME && docker rm $CONTAINER_NAME"
else
    echo "❌ Error iniciando cadena."
    docker logs "$CONTAINER_NAME"
    exit 1
fi

# Mantener vivo
trap "docker stop $CONTAINER_NAME && docker rm $CONTAINER_NAME" EXIT
echo "Presiona Ctrl+C para salir."
while true; do sleep 1; done
