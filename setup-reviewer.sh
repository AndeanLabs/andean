#!/bin/bash
# setup-reviewer.sh - Setup mejorado para Andean Chain con soporte multi-arquitectura
set -e

echo "🚀 Configurando Andean Chain para pruebas locales..."

# --- Verificación de Prerrequisitos ---
echo "🔍 Verificando prerrequisitos..."
command -v docker >/dev/null 2>&1 || { 
    echo "❌ Docker no instalado. Instálalo desde https://docs.docker.com/get-docker/" >&2
    exit 1 
}
command -v git >/dev/null 2>&1 || { 
    echo "❌ Git no instalado. Instálalo primero." >&2
    exit 1 
}

# Verificar que Docker esté corriendo
if ! docker info >/dev/null 2>&1; then
    echo "❌ Docker no está corriendo. Inicia Docker Desktop o el daemon de Docker." >&2
    exit 1
fi

echo "✅ Prerrequisitos cumplidos."

# --- Detectar Arquitectura ---
ARCH=$(uname -m)
PLATFORM=""
if [[ "$ARCH" == "arm64" ]] || [[ "$ARCH" == "aarch64" ]]; then
    echo "🔍 Detectada arquitectura ARM (M1/M2 Mac). Usando emulación x86_64."
    PLATFORM="--platform linux/amd64"
else
    echo "🔍 Detectada arquitectura x86_64."
fi

# --- Clonar o Usar Existente ---
REPO_URL="https://github.com/AndeanLabs/andean.git"
CONTAINER_NAME="andean-dev-container"
IMAGE_NAME="andean-dev"

if [ ! -d "andean" ]; then
    echo "📥 Clonando repositorio..."
    git clone "$REPO_URL" andean
else
    echo "📁 Usando directorio andean existente."
fi

cd andean

# --- Limpiar Estado Anterior ---
echo "🧹 Limpiando contenedores e imágenes anteriores..."
docker stop "$CONTAINER_NAME" >/dev/null 2>&1 || true
docker rm "$CONTAINER_NAME" >/dev/null 2>&1 || true
docker rmi "$IMAGE_NAME" >/dev/null 2>&1 || true

# --- Construir Imagen ---
echo "🏗️  Construyendo imagen Docker con soporte multi-arquitectura..."
echo "   (Esto puede tardar varios minutos la primera vez)"

if docker build $PLATFORM -t "$IMAGE_NAME" . 2>&1 | tee /tmp/docker_build.log; then
    echo "✅ Imagen construida exitosamente."
else
    echo "❌ Error construyendo imagen. Log completo:"
    cat /tmp/docker_build.log
    exit 1
fi

# Verificar que la imagen existe
if ! docker images | grep -q "$IMAGE_NAME"; then
    echo "❌ La imagen $IMAGE_NAME no se creó correctamente."
    exit 1
fi

# --- Iniciar Contenedor ---
echo "🐳 Iniciando contenedor..."

docker run -d \
    $PLATFORM \
    --name "$CONTAINER_NAME" \
    -v "$(pwd):/workspace" \
    -p 1317:1317 \
    -p 26656:26656 \
    -p 26657:26657 \
    "$IMAGE_NAME" \
    /bin/bash -c "while true; do sleep 30; done" >/dev/null

# Verificar contenedor con reintentos
sleep 3
RETRIES=5
for i in $(seq 1 $RETRIES); do
    if docker ps | grep -q "$CONTAINER_NAME"; then
        echo "✅ Contenedor iniciado correctamente."
        break
    fi
    
    if [ $i -eq $RETRIES ]; then
        echo "❌ Contenedor no pudo iniciarse después de $RETRIES intentos. Logs:"
        docker logs "$CONTAINER_NAME"
        exit 1
    fi
    
    echo "⏳ Esperando contenedor (intento $i/$RETRIES)..."
    sleep 2
done

# --- Función para Ejecutar en Contenedor ---
exec_in_container() {
    docker exec "$CONTAINER_NAME" bash -c "cd /workspace && $1"
}

# --- Configurar Blockchain ---
echo "⚙️  Configurando blockchain..."

# Instalar binario
echo "   📦 Instalando andeand..."
if ! exec_in_container "go install ./cmd/andeand"; then
    echo "❌ Error instalando andeand"
    docker logs "$CONTAINER_NAME"
    exit 1
fi

# Inicializar chain
echo "   🔧 Inicializando chain..."
exec_in_container "andeand init reviewer-test --chain-id andean-test-1 --home /workspace/.andean"

# Crear cuenta
echo "   🔑 Creando cuenta alice..."
exec_in_container "andeand keys add alice --keyring-backend test --home /workspace/.andean"

# Configurar genesis
echo "   ⚡ Configurando genesis..."
exec_in_container "andeand genesis add-genesis-account alice 1000000000000aand --keyring-backend test --home /workspace/.andean"
exec_in_container "andeand genesis gentx alice 1000000000aand --chain-id andean-test-1 --keyring-backend test --home /workspace/.andean"
exec_in_container "andeand genesis collect-gentxs --home /workspace/.andean"

# --- Iniciar Blockchain ---
echo "🔥 Iniciando blockchain..."
docker exec -d "$CONTAINER_NAME" bash -c "cd /workspace && andeand start --home /workspace/.andean --minimum-gas-prices 0stake" >/dev/null

# Esperar y verificar que inició
echo "⏳ Esperando que la blockchain inicie..."
sleep 10

# Verificar con reintentos
RETRIES=6
for i in $(seq 1 $RETRIES); do
    if exec_in_container "andeand status --node tcp://localhost:26657" >/dev/null 2>&1; then
        echo "✅ Blockchain iniciada correctamente."
        break
    fi
    
    if [ $i -eq $RETRIES ]; then
        echo "❌ Blockchain no pudo iniciar después de $((RETRIES * 10)) segundos."
        echo "Logs del contenedor:"
        docker logs "$CONTAINER_NAME" --tail 20
        exit 1
    fi
    
    echo "   Intento $i/$RETRIES - esperando 10 segundos más..."
    sleep 10
done

# --- Información Final ---
ALICE_ADDR=$(exec_in_container "andeand keys show alice -a --keyring-backend test --home /workspace/.andean")
LATEST_BLOCK=$(exec_in_container "andeand status --node tcp://localhost:26657 2>/dev/null | grep -o '\"latest_block_height\":\"[^\"]*\"' | cut -d'\"' -f4" || echo "N/A")

echo ""
echo "🎉 ¡Setup completado exitosamente!"
echo ""
echo "📊 Información de la blockchain:"
echo "   🔗 Chain ID: andean-test-1"
echo "   🏠 Dirección Alice: $ALICE_ADDR"
echo "   📦 Último bloque: #$LATEST_BLOCK"
echo ""
echo "🌐 Endpoints disponibles:"
echo "   📡 RPC: http://localhost:26657 o http://127.0.0.1:26657"
echo "   🌍 API: http://localhost:1317 o http://127.0.0.1:1317"
echo ""
echo "🧪 Comandos de prueba:"
echo "   # Verificar estado"
echo "   curl http://127.0.0.1:26657/status"
echo ""
echo "   # Conectar al contenedor"
echo "   docker exec -it $CONTAINER_NAME bash"
echo ""
echo "   # Ver balance de Alice"
echo "   docker exec $CONTAINER_NAME andeand query bank balances $ALICE_ADDR --home /workspace/.andean --node tcp://localhost:26657"
echo ""
echo "🛑 Para detener todo:"
echo "   docker stop $CONTAINER_NAME && docker rm $CONTAINER_NAME"
echo ""
echo "✨ Puedes cerrar esta terminal. La blockchain seguirá corriendo en background."
echo "   Continúa con el Paso 2 del README para usar el CLI nativo."
