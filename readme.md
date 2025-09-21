# Andean Chain 🏔️ - Blockchain Experimental Local

[![GoDoc](https://pkg.go.dev/badge/github.com/AndeanLabs/andean)](https://pkg.go.dev/github.com/AndeanLabs/andean)
[![Go Report Card](https://goreportcard.com/badge/github.com/AndeanLabs/andean)](https://goreportcard.com/report/github.com/AndeanLabs/andean)
[![Version](https://img.shields.io/github/tag/AndeanLabs/andean.svg)](https://github.com/AndeanLabs/andean/releases/latest)
[![License](https://img.shields.io/github/license/AndeanLabs/andean.svg)](https://github.com/AndeanLabs/andean/blob/main/LICENSE)

> Blockchain experimental para desarrollo y pruebas locales. Combina disponibilidad de datos (Data Availability) con pruebas ZK para finanzas descentralizadas con privacidad opcional, Lazy Bridge y costos ultra-bajos.

⚠️ **VERSIÓN DE DESARROLLO**: Esta implementación es solo para testing local. No se recomienda para uso en producción.

## 📋 Tabla de Contenidos

- [🌟 Módulos Experimentales](#-módulos-experimentales)
- [📊 Rendimiento Local](#-rendimiento-local)
- [🚀 Instalación y Configuración](#-instalación-y-configuración)
- [🎮 Guía de Interacción](#-guía-de-interacción)
- [🧪 Ejemplos Avanzados](#-ejemplos-avanzados)
- [📡 APIs y Monitoreo](#-apis-y-monitoreo)
- [🔧 Troubleshooting](#-troubleshooting)

---

## 🌟 Módulos Experimentales

### ✅ AndeanSwap AMM (XICOATL Module)
- **Multi-Curve**: Constant Product, StableCurve, Liquidez Concentrada
- **Hooks System**: Similar a Uniswap V4 con extensiones
- **MEV Protection**: Subastas por lotes con garantías de precio
- **Estado**: 🚧 En desarrollo para pruebas locales

### ✅ INTI Lazy Bridge
- **1-Second Finality**: Simulación de acceso rápido a assets
- **ZK Security**: Uso de ZK proofs en entorno de testing
- **Costo Simulado**: <$0.01 por transacción (modo local)
- **Estado**: 🚧 Prototipo experimental

### ✅ ITZEL Oracle Network
- **Precios P2P**: Fuentes simuladas multi-región
- **Actualización**: Cada 30s en la red local
- **Assets Mock**: USD, BOB, PEN, COP, BTC, ETH
- **Estado**: 🚧 Datos simulados para testing

### ✅ PACHAMAMA Privacy Layer
- **Privacidad ZK**: Transacciones privadas opcionales
- **Divulgación Selectiva**: En desarrollo
- **Multi-nivel**: 4 niveles de privacidad configurables
- **Estado**: 🚧 Experimental


## 🌟 Módulos Futuros


### ✅ CHASQUI P2P Network
- **Pagos Directos**: Envío de remesas y transferencias P2P sin intermediarios  
- **Privacidad Opcional**: Soporte para rutas privadas en múltiples saltos (multi-hop)  
- **Costo Ultra-Bajo**: Simulación de transacciones por ~$0.0001 en entorno local  
- **Enfoque Regional**: Diseñado para pagos en regiones andinas  
- **Estado**: 🚧 En desarrollo inicial

### ✅ TLAHUIZCAL MEV Shield
- **Batch Auctions**: Protección contra MEV mediante subastas por bloques  
- **Hooks Avanzados**: Integración con módulos AMM para ejecución justa  
- **MEV Redistribution**: Captura y redistribución de MEV a proveedores de liquidez  
- **Prevención de Ataques**: Evita *sandwich attacks* y manipulaciones de precio  
- **Estado**: 🚧 Prototipo en pruebas

### ✅ VIRACOCHA Settlement Layer
- **Validación ZK**: Verificación de Zero-Knowledge proofs on-chain  
- **Cross-chain Settlement**: Resolución de transacciones privadas entre cadenas (simulado)  
- **Gestión de Proofs**: Registro, verificación y tracking de pruebas criptográficas  
- **Estado**: 🚧 Implementación básica

### ✅ ANDES Staking Module
- **Token Staking**: Delegación de ANDES para asegurar la red  
- **Rewards Dinámicos**: APR ajustable entre 15% y 60% según condiciones de red  
- **Slashing**: Penalización por inactividad o mala conducta del validador  
- **Gobernanza**: Integración con sistema de votación on-chain  
- **Estado**: 🚧 Versión inicial funcional

### ✅ Governance Module
- **Votación On-Chain**: Decisiones sobre upgrades, parámetros y tesorería  
- **Privacy Voting**: Votaciones con privacidad opcional para participantes  
- **Integración Total**: Compatible con módulos económicos, sociales y técnicos  
- **Transparencia y Seguridad**: Registro auditado de todas las propuestas  
- **Estado**: 🚧 Activo en entorno de pruebas


---

## 📊 Rendimiento Local

| Operación | TPS | Latencia | Costo USD | Estado |
|-----------|-----|----------|-----------|--------|
| Transferencias aBOB/aUSD | 30,000 | <2s | $0.0001 | 🚧 No Implementado |
| Swaps AMM | 25,000 | <1s | $0.0002 |  ✅ Funcional |
| Operaciones LP | 15,000 | <2s | $0.0004 |  ✅ Funcional |
| Transacciones ZK privadas | 8,000 | <4s | $0.001 | 🚧 Experimental |
| Cross-chain bridges | 8,000 | <5s | $0.005 | 🚧 Prototipo |
| Hook Execution | 20,000 | <1s | Variable | 🚧 Experimental |

📌 *El rendimiento depende de tu hardware local.

---

## 🚀 Instalación y Configuración

### 🧹 Limpieza Previa (IMPORTANTE)

**Antes de empezar, ejecuta estos comandos de limpieza para evitar conflictos:**

#### Para Docker:
```bash
# Detener y eliminar contenedores relacionados
docker stop andean-container andean-node andean-chain 2>/dev/null || true
docker rm andean-container andean-node andean-chain 2>/dev/null || true

# Eliminar imágenes previas
docker rmi andean-chain andean-node andean/chain 2>/dev/null || true
docker image prune -f

# Limpiar volúmenes Docker (CUIDADO: elimina datos persistentes)
docker volume ls | grep andean | awk '{print $2}' | xargs docker volume rm 2>/dev/null || true

# Verificar limpieza
echo "🔍 Verificando limpieza..."
docker ps -a | grep -i andean || echo "✅ No hay contenedores Andean"
docker images | grep -i andean || echo "✅ No hay imágenes Andean"
```

#### Para instalación local:
```bash
# Detener procesos andeand en ejecución
pkill -f andeand || true

# Limpiar directorio de configuración (CUIDADO: elimina wallets y datos)
rm -rf ~/.andean || true
rm -rf ./.andean || true

# Limpiar binarios previos
rm -f $(which andeand) 2>/dev/null || true
rm -f $(go env GOPATH)/bin/andeand 2>/dev/null || true

# Limpiar caché de Go
go clean -modcache
go clean -cache

# Verificar limpieza
echo "🔍 Verificando limpieza..."
pgrep -f andeand || echo "✅ No hay procesos andeand ejecutándose"
ls ~/.andean 2>/dev/null || echo "✅ Directorio ~/.andean eliminado"
which andeand || echo "✅ Binario andeand eliminado"
```

#### Limpieza de puertos (si están ocupados):
```bash
# Verificar qué procesos usan los puertos necesarios
echo "🔍 Verificando puertos..."
lsof -i :1317 || echo "Puerto 1317 libre"
lsof -i :26656 || echo "Puerto 26656 libre" 
lsof -i :26657 || echo "Puerto 26657 libre"

# Si hay procesos ocupando los puertos, detenerlos:
# sudo kill -9 $(lsof -ti:1317) 2>/dev/null || true
# sudo kill -9 $(lsof -ti:26656) 2>/dev/null || true  
# sudo kill -9 $(lsof -ti:26657) 2>/dev/null || true
```

---

### 🐳 Opción A: Con Docker (Recomendado)

#### Requisitos
- Docker Desktop instalado
- 8GB RAM disponible
- Puertos 1317, 26656, 26657 libres

#### Instalación paso a paso

```bash
# 1. Clonar el repositorio
git clone https://github.com/AndeanLabs/andean.git
cd andean

# 2. Verificar limpieza (ya debería estar hecho arriba)
docker ps | grep andean || echo "✅ Listo para continuar"

# 3. Construir la imagen
docker build -t andean-chain . --no-cache

# 4. Crear y ejecutar contenedor
docker run -d --name andean-container \
    -v "$(pwd):/workspace" \
    -p 1317:1317 -p 26656:26656 -p 26657:26657 \
    andean-chain \
    tail -f /dev/null

# 5. Inicializar la blockchain con configuración corregida
docker exec andean-container bash -c "
  cd /workspace
  go install ./cmd/andeand
  andeand init test-chain --chain-id andean-test-1 --home /workspace/.andean
  andeand keys add alice --keyring-backend test --home /workspace/.andean
  
  # IMPORTANTE: Configuración de denominaciones correcta
  andeand genesis add-genesis-account alice 1000000000000stake --keyring-backend test --home /workspace/.andean
  andeand genesis gentx alice 1000000000stake --chain-id andean-test-1 --keyring-backend test --home /workspace/.andean
  andeand genesis collect-gentxs --home /workspace/.andean
  
  # Agregar tokens aand para transacciones normales
  andeand genesis add-genesis-account alice 1000000000000aand --keyring-backend test --home /workspace/.andean
"

# 6. Verificar que el contenedor esté ejecutándose
docker ps | grep andean-container || echo "⚠️ El contenedor no está ejecutándose"

# 7. Iniciar el nodo
docker exec -d andean-container bash -c "
  cd /workspace && 
  andeand start --home /workspace/.andean --minimum-gas-prices 0.025aand
"
```

### 💻 Opción B: Sin Docker

#### Requisitos
- Go 1.21+ instalado
- make (opcional)  
- git

#### Instalación paso a paso

```bash
# 1. Clonar el repositorio (si no lo has hecho)
git clone https://github.com/AndeanLabs/andean.git
cd andean

# 2. Verificar limpieza (ya debería estar hecho arriba)
pgrep -f andeand || echo "✅ Listo para continuar"

# 3. Instalar el binario
go install ./cmd/andeand

# 4. Verificar instalación
andeand version || echo "⚠️  Verifica que $GOPATH/bin esté en tu PATH"

# 5. Inicializar blockchain con denominaciones correctas
andeand init test-chain --chain-id andean-test-1
andeand keys add alice --keyring-backend test

# IMPORTANTE: Configuración de denominaciones correcta
andeand genesis add-genesis-account alice 1000000000000stake --keyring-backend test
andeand genesis gentx alice 1000000000stake --chain-id andean-test-1 --keyring-backend test
andeand genesis collect-gentxs

# Agregar tokens aand para transacciones normales
andeand genesis add-genesis-account alice 1000000000000aand --keyring-backend test

# 6. Iniciar nodo
andeand start --minimum-gas-prices 0.025aand
```

---

## 🎮 Guía de Interacción

### 🔑 Gestión de Cuentas

```bash
# Crear nueva cuenta
andeand keys add mi-cuenta --keyring-backend test

# Listar cuentas
andeand keys list --keyring-backend test

# Ver balance
andeand query bank balances $(andeand keys show mi-cuenta -a --keyring-backend test)

# Exportar address para usar en variables
export MI_ADDR=$(andeand keys show mi-cuenta -a --keyring-backend test)
echo "Mi dirección: $MI_ADDR"
```

### 💸 Transacciones Básicas

```bash
# Crear segunda cuenta para pruebas
andeand keys add bob --keyring-backend test
export BOB_ADDR=$(andeand keys show bob -a --keyring-backend test)

# Transferir tokens
andeand tx bank send $MI_ADDR $BOB_ADDR 1000000aand \
  --keyring-backend test --chain-id andean-test-1 \
  --gas 200000 --gas-prices 0.025aand -y

# Verificar transacción (usar el hash devuelto)
andeand query tx [HASH_DE_TRANSACCION]
```

### 📊 Consultas de Estado

```bash
# Estado general de la blockchain
andeand status

# Información del último bloque
andeand query block

# Ver validadores
andeand query staking validators

# Ver supply total
andeand query bank total
```

---

## 🧪 Ejemplos Avanzados

### 🔄 AndeanSwap (Xicoatl Module)

```bash
# Crear un pool de liquidez
andeand tx xicoatl create-pool \
  --token-a aand --token-b utest --fee 0.003 \
  --initial-deposit-a 1000000 --initial-deposit-b 1000000 \
  --from mi-cuenta --keyring-backend test \
  --chain-id andean-test-1 --gas 300000 --gas-prices 0.025aand -y

# Hacer un swap
andeand tx xicoatl swap \
  --pool-id 1 --token-in aand --amount-in 100000 \
  --token-out-min-amount 95000 \
  --from mi-cuenta --keyring-backend test \
  --chain-id andean-test-1 --gas 250000 --gas-prices 0.025aand -y

# Consultar pools disponibles
andeand query xicoatl pools

# Consultar precio de un pool
andeand query xicoatl pool-price --pool-id 1
```

### 🌐 Oracle Network (Itzel)

```bash
# Enviar precio al oracle
andeand tx itzel submit-price \
  --asset BTC/USD --price 45000.50 --confidence 0.95 \
  --from mi-cuenta --keyring-backend test \
  --chain-id andean-test-1 --gas 200000 --gas-prices 0.025aand -y

# Consultar precios disponibles
andeand query itzel prices

# Consultar precio específico
andeand query itzel price --asset BTC/USD

# Enviar precio de moneda local
andeand tx itzel submit-price \
  --asset BOB/USD --price 0.145 --confidence 0.98 \
  --from mi-cuenta --keyring-backend test \
  --chain-id andean-test-1 --gas 200000 --gas-prices 0.025aand -y
```

### 🌉 Bridge Operations (Inti)

```bash
# Iniciar operación de bridge
andeand tx inti initiate-bridge \
  --source-chain ethereum --target-chain andean \
  --asset USDT --amount 1000000 --recipient $MI_ADDR \
  --from mi-cuenta --keyring-backend test \
  --chain-id andean-test-1 --gas 300000 --gas-prices 0.025aand -y

# Consultar operaciones de bridge
andeand query inti bridge-operations

# Verificar estado de una operación
andeand query inti bridge-status --operation-id 1

# Completar bridge (simulated)
andeand tx inti complete-bridge \
  --operation-id 1 --proof "mock-zk-proof" \
  --from mi-cuenta --keyring-backend test \
  --chain-id andean-test-1 --gas 250000 --gas-prices 0.025aand -y
```

### 🔒 Transacciones Privadas (Pachamama)

```bash
# Crear transacción privada (nivel básico)
andeand tx pachamama private-transfer \
  --recipient $BOB_ADDR --amount 500000 --privacy-level 1 \
  --from mi-cuenta --keyring-backend test \
  --chain-id andean-test-1 --gas 400000 --gas-prices 0.025aand -y

# Consultar balance privado
andeand query pachamama private-balance --address $MI_ADDR

# Generar prueba ZK (simulado)
andeand tx pachamama generate-proof \
  --transaction-hash [HASH] --proof-type transfer \
  --from mi-cuenta --keyring-backend test \
  --chain-id andean-test-1 --gas 500000 --gas-prices 0.025aand -y
```

---


### Monitoreo con WebSocket

```javascript
// Conectar a WebSocket para eventos
const ws = new WebSocket('ws://127.0.0.1:26657/websocket');

ws.onopen = function() {
    // Suscribirse a nuevos bloques
    ws.send(JSON.stringify({
        "jsonrpc": "2.0",
        "method": "subscribe",
        "id": 1,
        "params": {
            "query": "tm.event='NewBlock'"
        }
    }));
};

ws.onmessage = function(event) {
    const data = JSON.parse(event.data);
    console.log('Nuevo bloque:', data);
};
```

### Herramientas de Desarrollo

```bash
# Ejecutar en modo debug
andeand start --log_level debug --minimum-gas-prices 0.025aand

# Ver logs del contenedor Docker
docker logs andean-container -f

# Resetear blockchain (CUIDADO: borra todos los datos)
# Docker:
docker exec andean-container andeand unsafe-reset-all --home /workspace/.andean
# Local:
andeand unsafe-reset-all
```

---

## 🔧 Troubleshooting

### Problemas Comunes

### ❌ Errores de instalación previa
```bash
# Si aparece "container name already exists"
docker rm -f andean-container

# Si aparece "image already exists" 
docker rmi andean-chain --force

# Si aparece "port already in use"
sudo kill -9 $(lsof -ti:26657) 2>/dev/null || true

# Si aparece "permission denied"
sudo chown -R $USER:$USER .
sudo usermod -aG docker $USER
newgrp docker
```

#### ❌ Datos corruptos o inconsistentes
```bash
# Resetear blockchain completamente (Docker)
docker exec andean-container andeand unsafe-reset-all --home /workspace/.andean

# Resetear blockchain completamente (Local)
andeand unsafe-reset-all

# Si persisten problemas, ejecutar limpieza completa:
./cleanup.sh  # Usar el script de arriba
```

#### ❌ `andeand: command not found`
```bash
# Verificar instalación de Go
go version

# Reinstalar binario
cd andean
go install ./cmd/andeand

# Verificar PATH
export PATH=$PATH:$(go env GOPATH)/bin
echo $PATH
```

#### ❌ Puertos ocupados
```bash
# Verificar qué está usando los puertos
lsof -i :26657
lsof -i :1317

# Cambiar puertos en Docker
docker run -d --name andean-container \
    -v "$(pwd):/workspace" \
    -p 1318:1317 -p 26658:26656 -p 26659:26657 \
    andean-chain tail -f /dev/null
```

#### ❌ Errores de permisos
```bash
# Docker: asegurar permisos correctos
sudo chown -R $USER:$USER .andean/

# Linux: agregar usuario al grupo docker
sudo usermod -aG docker $USER
newgrp docker
```

#### ❌ Blockchain no responde
```bash
# Docker: reiniciar contenedor
docker restart andean-container

# Local: reiniciar con logs
andeand start --log_level debug --minimum-gas-prices 0.025aand

# Verificar conectividad
curl http://127.0.0.1:26657/status
```

#### ❌ Transacciones fallan
```bash
# Verificar balance
andeand query bank balances $MI_ADDR

# Verificar secuencia de cuenta
andeand query account $MI_ADDR

# Incrementar gas
andeand tx bank send $MI_ADDR $BOB_ADDR 1000aand \
  --gas 300000 --gas-prices 0.050aand \
  --keyring-backend test --chain-id andean-test-1 -y
```

### Logs y Debug

```bash
# Logs completos (Docker)
docker logs andean-container --tail 100 -f

# Logs completos (Local)
andeand start --log_level debug 2>&1 | tee andean.log

# Ver solo errores
andeand start 2>&1 | grep ERROR

# Verificar estado de módulos
andeand query xicoatl params
andeand query itzel params
andeand query inti params
```

---

## 🧪 Scripts de Automatización

### Script de Limpieza Completa

```bash
#!/bin/bash
# cleanup.sh - Limpieza completa del sistema
set -e

echo "🧹 Iniciando limpieza completa de Andean Chain..."

# Función para confirmar acciones destructivas
confirm() {
    read -p "⚠️  $1 (y/N): " -n 1 -r
    echo
    [[ $REPLY =~ ^[Yy]$ ]]
}

# Docker cleanup
echo "🐳 Limpiando Docker..."
docker stop andean-container andean-node andean-chain 2>/dev/null || true
docker rm andean-container andean-node andean-chain 2>/dev/null || true
docker rmi andean-chain andean-node andean/chain 2>/dev/null || true

if confirm "¿Eliminar volúmenes Docker (esto borra datos persistentes)?"; then
    docker volume ls | grep andean | awk '{print $2}' | xargs docker volume rm 2>/dev/null || true
fi

# Local cleanup
echo "💻 Limpiando instalación local..."
pkill -f andeand || true

if confirm "¿Eliminar directorios de configuración (esto borra wallets y datos)?"; then
    rm -rf ~/.andean || true
    rm -rf ./.andean || true
fi

# Binary cleanup
echo "🗑️  Limpiando binarios..."
rm -f $(which andeand) 2>/dev/null || true
rm -f $(go env GOPATH)/bin/andeand 2>/dev/null || true

# Go cache cleanup
if confirm "¿Limpiar caché de Go?"; then
    go clean -modcache
    go clean -cache
fi

# Port cleanup check
echo "🔍 Verificando puertos..."
for port in 1317 26656 26657; do
    if lsof -i :$port >/dev/null 2>&1; then
        echo "⚠️  Puerto $port está ocupado:"
        lsof -i :$port
        if confirm "¿Terminar procesos en puerto $port?"; then
            sudo kill -9 $(lsof -ti:$port) 2>/dev/null || true
        fi
    else
        echo "✅ Puerto $port está libre"
    fi
done

# Verification
echo "🔍 Verificando limpieza completa..."
docker ps -a | grep -i andean && echo "⚠️  Aún hay contenedores Andean" || echo "✅ No hay contenedores Andean"
docker images | grep -i andean && echo "⚠️  Aún hay imágenes Andean" || echo "✅ No hay imágenes Andean"
pgrep -f andeand && echo "⚠️  Procesos andeand aún ejecutándose" || echo "✅ No hay procesos andeand"
ls ~/.andean 2>/dev/null && echo "⚠️  Directorio ~/.andean existe" || echo "✅ Directorio ~/.andean eliminado"
which andeand && echo "⚠️  Binario andeand encontrado en PATH" || echo "✅ Binario andeand eliminado"

echo "✅ Limpieza completa terminada!"
```

### Setup Completo (Docker)

```bash
#!/bin/bash
# setup.sh - Configuración automática
set -e

echo "🏔️ Configurando Andean Chain..."

# Limpieza automática
echo "🧹 Ejecutando limpieza previa..."
docker stop andean-container 2>/dev/null || true
docker rm andean-container 2>/dev/null || true
docker rmi andean-chain 2>/dev/null || true

# Verificar puertos
for port in 1317 26656 26657; do
    if lsof -i :$port >/dev/null 2>&1; then
        echo "⚠️  Puerto $port ocupado. Terminando procesos..."
        sudo kill -9 $(lsof -ti:$port) 2>/dev/null || true
        sleep 2
    fi
done

# Build and run
docker stop andean-container 2>/dev/null || true
docker rm andean-container 2>/dev/null || true

# Build and run
docker build -t andean-chain . --no-cache
docker run -d --name andean-container \
    -v "$(pwd):/workspace" \
    -p 1317:1317 -p 26656:26656 -p 26657:26657 \
    andean-chain tail -f /dev/null

# Initialize
docker exec andean-container bash -c "
  cd /workspace
  go install ./cmd/andeand
  andeand init test-chain --chain-id andean-test-1 --home /workspace/.andean
  andeand keys add alice --keyring-backend test --home /workspace/.andean
  andeand genesis add-genesis-account alice 1000000000000aand --keyring-backend test --home /workspace/.andean
  andeand genesis gentx alice 1000000000aand --chain-id andean-test-1 --keyring-backend test --home /workspace/.andean
  andeand genesis collect-gentxs --home /workspace/.andean
"

# Start node
docker exec -d andean-container bash -c "
  cd /workspace && andeand start --home /workspace/.andean --minimum-gas-prices 0.025aand
"

echo "✅ Andean Chain configurado y ejecutándose!"
echo "RPC: http://127.0.0.1:26657"
echo "REST: http://127.0.0.1:1317"
```

### Script de Testing

```bash
#!/bin/bash
# test-interactions.sh - Pruebas automatizadas
set -e

# Wait for node to be ready
echo "⏳ Esperando que el nodo esté listo..."
sleep 10

# Create test account
echo "🔑 Creando cuenta de prueba..."
docker exec andean-container andeand keys add testuser --keyring-backend test --home /workspace/.andean

# Get addresses
ALICE_ADDR=$(docker exec andean-container andeand keys show alice -a --keyring-backend test --home /workspace/.andean)
TEST_ADDR=$(docker exec andean-container andeand keys show testuser -a --keyring-backend test --home /workspace/.andean)

echo "Alice: $ALICE_ADDR"
echo "Test User: $TEST_ADDR"

# Send tokens
echo "💸 Enviando tokens de prueba..."
docker exec andean-container andeand tx bank send $ALICE_ADDR $TEST_ADDR 1000000aand \
  --keyring-backend test --chain-id andean-test-1 --home /workspace/.andean \
  --gas 200000 --gas-prices 0.025aand -y

sleep 5

# Check balance
echo "💰 Verificando balance..."
docker exec andean-container andeand query bank balances $TEST_ADDR

echo "✅ Testing completado!"
```

---

## 🤝 Contribución

### Cómo Contribuir

1. **Fork** el repositorio
2. Crea una rama para tu feature: `git checkout -b feature/nueva-funcionalidad`
3. Commit tus cambios: `git commit -m "Agrega nueva funcionalidad"`
4. Push a la rama: `git push origin feature/nueva-funcionalidad`
5. Abre un **Pull Request**

### Desarrollo Local

```bash
# Ejecutar tests
go test ./...

# Ejecutar linting
golangci-lint run

# Generar mocks
make generate

# Build local
make build

# Install local
make install
```

### Reportar Issues

Antes de reportar un issue, por favor:

1. Verifica que no exista un issue similar
2. Incluye información del sistema (OS, Docker version, Go version)
3. Proporciona logs relevantes
4. Describe los pasos para reproducir el problema

---

## 🌐 Comunidad y Soporte (No existe aún)

- 💬 [Discord](https://discord.gg/andean-chain) - Chat en tiempo real
- 🐦 [Twitter](https://twitter.com/andean_chain) - Actualizaciones y noticias
- 📱 [Telegram](https://t.me/andean_chain) - Comunidad hispanohablante
- 🗣️ [Forum](https://forum.andean-chain.org) - Discusiones técnicas
- 📧 [Email](mailto:hello@andean-chain.org) - Contacto directo


## 📄 Licencia

Este proyecto está licenciado bajo la **MIT License** - ver el archivo [LICENSE](LICENSE) para más detalles.



---

<div align="center">

**Andean Chain** 🏔️⚡

*Blockchain experimental de próxima generación para América Latina*

Made with ❤️ by the Andean Labs team

</div>
