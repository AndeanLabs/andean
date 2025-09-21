#!/bin/bash

# setup-reviewer.sh - Setup completo para revisión de Andean Chain

set -e

echo "🚀 Configurando Andean Chain para revisión técnica..."

# --- Verificación de Prerrequisitos ---
command -v docker >/dev/null 2>&1 || { echo "❌ Docker no instalado. Instálalo primero." >&2; exit 1; }
command -v git >/dev/null 2>&1 || { echo "❌ Git no instalado." >&2; exit 1; }

echo "✅ Prerrequisitos cumplidos."

# --- Clonar y Construir ---
if [ ! -d "andean" ]; then
    echo "📥 Clonando repositorio..."
    git clone https://github.com/AndeanLabs/andean.git
fi

cd andean

echo "🏗️  Construyendo imagen Docker (puede tardar varios minutos)..."
docker build -t andean-review . > /dev/null

# --- Preparar Contenedor ---
echo "🐳 Iniciando contenedor de setup..."
docker run -d --rm --name andean-review \
    -p 1317:1317 \
    -p 26657:26657 \
    andean-review sleep infinity > /dev/null

# Esperar a que el contenedor esté listo
sleep 5

# --- Secuencia de Inicio Dentro del Contenedor ---
echo "⚙️  Configurando la cadena de prueba (esto puede tardar un minuto)..."

# 1. Compilar
docker exec andean-review bash -c "go install ./cmd/andeand"

# 2. Inicializar
docker exec andean-review bash -c "andeand init reviewer-demo --chain-id andean-demo-1 --home /workspace/.andean"

# 3. Crear llave
docker exec andean-review bash -c "andeand keys add reviewer --keyring-backend test --home /workspace/.andean"

# 4. Añadir cuenta al genesis (Sintaxis corregida)
docker exec andean-review bash -c "andeand genesis add-genesis-account reviewer 1000000000000aand --keyring-backend test --home /workspace/.andean"

# 5. Crear Gentx
docker exec andean-review bash -c "andeand genesis gentx reviewer 1000000000aand --chain-id andean-demo-1 --keyring-backend test --home /workspace/.andean"

# 6. Recolectar Gentx (Sintaxis corregida)
docker exec andean-review bash -c "andeand genesis collect-gentxs --home /workspace/.andean"

# --- Iniciar la Cadena ---
echo "🔥 Iniciando la cadena en segundo plano..."
docker exec -d andean-review bash -c "andeand start --home /workspace/.andean --minimum-gas-prices 0stake"

# Esperar a que el primer bloque se produzca
sleep 8

# --- Verificación Final ---
REVIEWER_ADDR=$(docker exec andean-review bash -c "andeand keys show reviewer -a --keyring-backend test --home /workspace/.andean")

echo "✅ ¡Setup completo! La cadena está corriendo en segundo plano."

echo ""

echo "Dirección de la cuenta del revisor: $REVIEWER_ADDR"

echo ""

echo "🌐 Endpoints disponibles:"

echo "  - RPC      -> http://localhost:26657"

echo "  - API REST -> http://localhost:1317"

echo ""

echo "🧪 Comandos de ejemplo para probar:"

echo "  # Consultar el estado del nodo"

echo "  docker exec -it andean-review andeand status --node tcp://localhost:26657"

echo ""

echo "  # Consultar el balance de la cuenta del revisor"

echo "  docker exec -it andean-review andeand query bank balances $REVIEWER_ADDR --home /workspace/.andean --node tcp://localhost:26657"

echo ""

echo "🛑 Para detener y limpiar el contenedor:"

echo "  docker stop andean-review && docker rm andean-review"
