#!/bin/bash

# setup-reviewer-improved.sh - Setup completo y robusto para revisión de Andean Chain

set -e  # Salir en errores, pero con verificaciones manuales

echo "🚀 Configurando Andean Chain para revisión técnica..."

# --- Verificación de Prerrequisitos ---
echo "🔍 Verificando prerrequisitos..."
command -v docker >/dev/null 2>&1 || { echo "❌ Docker no instalado. Instálalo primero (https://docs.docker.com/get-docker/)." >&2; exit 1; }
command -v git >/dev/null 2>&1 || { echo "❌ Git no instalado. Instálalo primero." >&2; exit 1; }

# Verificar versión de Docker (mínimo 20.10)
DOCKER_VERSION=$(docker --version | awk '{print $3}' | sed 's/,//')
if [[ "$DOCKER_VERSION" < "20.10" ]]; then
    echo "⚠️  Docker versión $DOCKER_VERSION detectada. Recomendado: 20.10+. Puede causar problemas."
fi

echo "✅ Prerrequisitos cumplidos."

# --- Clonar y Construir ---
REPO_URL="https://github.com/AndeanLabs/andean.git"  # Ajusta si es privado
if [ ! -d "andean" ]; then
    echo "📥 Clonando repositorio desde $REPO_URL..."
    git clone "$REPO_URL"
else
    echo "📁 Directorio 'andean' ya existe. Saltando clonado."
fi

cd andean || { echo "❌ Error: No se pudo acceder a 'andean'. Verifica permisos."; exit 1; }

echo "🏗️  Construyendo imagen Docker (puede tardar varios minutos)..."
if ! docker build -t andean-review . ; then
    echo "❌ Error al construir la imagen. Revisa logs de Docker."
    exit 1
fi

echo "✅ Imagen construida exitosamente."

# --- Preparar Contenedor ---
echo "🐳 Iniciando contenedor de setup..."
CONTAINER_NAME="andean-review"

# Detener y remover contenedor si existe
docker stop "$CONTAINER_NAME" >/dev/null 2>&1 || true
docker rm "$CONTAINER_NAME" >/dev/null 2>&1 || true

# Ejecutar contenedor en background
if ! docker run -d --name "$CONTAINER_NAME" \
    -p 1317:1317 \
    -p 26657:26657 \
    andean-review sleep infinity; then
    echo "❌ Error al iniciar el contenedor. Revisa 'docker logs $CONTAINER_NAME'."
    exit 1
fi

# Verificar que el contenedor esté corriendo
sleep 3
if ! docker ps | grep -q "$CONTAINER_NAME"; then
    echo "❌ El contenedor no está corriendo. Logs:"
    docker logs "$CONTAINER_NAME" || true
    docker rm "$CONTAINER_NAME" >/dev/null 2>&1 || true
    exit 1
fi

echo "✅ Contenedor corriendo."

# --- Secuencia de Inicio Dentro del Contenedor ---
echo "⚙️  Configurando la cadena de prueba (esto puede tardar un minuto)..."

# Función para ejecutar comandos en el contenedor con verificación
exec_in_container() {
    local cmd="$1"
    echo "Ejecutando: $cmd"
    if ! docker exec "$CONTAINER_NAME" bash -c "$cmd"; then
        echo "❌ Error ejecutando '$cmd'. Logs del contenedor:"
        docker logs "$CONTAINER_NAME" || true
        cleanup
        exit 1
    fi
}

# 1. Verificar Go en el contenedor
exec_in_container "go version || { echo 'Go no disponible en contenedor'; exit 1; }"

# 2. Compilar
exec_in_container "go install ./cmd/andeand"

# 3. Inicializar
exec_in_container "andeand init reviewer-demo --chain-id andean-demo-1 --home /workspace/.andean"

# 4. Crear llave
exec_in_container "andeand keys add reviewer --keyring-backend test --home /workspace/.andean"

# 5. Añadir cuenta al genesis
exec_in_container "andeand genesis add-genesis-account reviewer 1000000000000aand --keyring-backend test --home /workspace/.andean"

# 6. Crear Gentx
exec_in_container "andeand genesis gentx reviewer 1000000000aand --chain-id andean-demo-1 --keyring-backend test --home /workspace/.andean"

# 7. Recolectar Gentx
exec_in_container "andeand genesis collect-gentxs --home /workspace/.andean"

# --- Iniciar la Cadena ---
echo "🔥 Iniciando la cadena en segundo plano..."
if ! docker exec -d "$CONTAINER_NAME" bash -c "andeand start --home /workspace/.andean --minimum-gas-prices 0stake"; then
    echo "❌ Error al iniciar la cadena."
    cleanup
    exit 1
fi

# Esperar a que el primer bloque se produzca
echo "⏳ Esperando inicialización de la cadena..."
sleep 10

# Verificar que la cadena esté corriendo
if ! docker exec "$CONTAINER_NAME" bash -c "andeand status --node tcp://localhost:26657" >/dev/null 2>&1; then
    echo "❌ La cadena no responde. Logs:"
    docker logs "$CONTAINER_NAME" || true
    cleanup
    exit 1
fi

# --- Verificación Final ---
REVIEWER_ADDR=$(docker exec "$CONTAINER_NAME" bash -c "andeand keys show reviewer -a --keyring-backend test --home /workspace/.andean" 2>/dev/null || echo "Error obteniendo dirección")

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
echo "  docker exec -it $CONTAINER_NAME andeand status --node tcp://localhost:26657"
echo ""
echo "  # Consultar el balance de la cuenta del revisor"
echo "  docker exec -it $CONTAINER_NAME andeand query bank balances $REVIEWER_ADDR --home /workspace/.andean --node tcp://localhost:26657"
echo ""
echo "🛑 Para detener y limpiar:"
echo "  docker stop $CONTAINER_NAME && docker rm $CONTAINER_NAME"

# Función de limpieza
cleanup() {
    echo "🧹 Limpiando..."
    docker stop "$CONTAINER_NAME" >/dev/null 2>&1 || true
    docker rm "$CONTAINER_NAME" >/dev/null 2>&1 || true
}

# Mantener el script corriendo para que el contenedor no se elimine automáticamente
echo "Presiona Ctrl+C para salir. El contenedor seguirá corriendo."
trap cleanup EXIT
while true; do sleep 1; done

