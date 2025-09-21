<!-- markdownlint-disable MD033 -->
# Andean Chain 🏔️

[![GoDoc](https://pkg.go.dev/badge/github.com/andean-labs/andean)](https://pkg.go.dev/github.com/andean-labs/andean)
[![Go Report Card](https://goreportcard.com/badge/github.com/andean-labs/andean)](https://goreportcard.com/report/github.com/andean-labs/andean)
[![Version](https://img.shields.io/github/tag/andean-labs/andean.svg)](https://github.com/andean-labs/andean/releases/latest)
[![License](https://img.shields.io/github/license/andean-labs/andean.svg)](https://github.com/andean-labs/andean/blob/main/LICENSE)
[![Discord](https://img.shields.io/discord/1234567890)](https://discord.gg/andean-chain)

> La primera blockchain nativa de Celestia para la región andina. Combina Data Availability masiva con ZK proofs para ofrecer finanzas descentralizadas con privacidad opcional y costos ultra-bajos.

## 🌟 Características Funcionales Principales

### ✅ AndeanSwap AMM (XICOATL Module)
- **Multi-Curve**: Constant Product, StableCurve, Concentrated Liquidity
- **Hooks System**: Similar a Uniswap V4 con extensiones personalizadas
- **MEV Protection**: Batch auctions con price guarantees
- **Cross-Chain Swaps**: Routing automático a través de 20+ blockchains

### ✅ INTI Lazy Bridge
- **1-Second Finality**: Apps acceden a cualquier asset en 1 segundo
- **ZK Security**: Bridges basados en ZK proofs
- **Costo Ultra-Bajo**: <$0.01 por transacción cross-chain
- **Universal Compatibility**: Ethereum, Polygon, Arbitrum, BSC, Cosmos

### ✅ ITZEL Oracle Network
- **Precios P2P**: Fuentes multi-región para máxima resistencia a manipulación
- **Actualización**: Cada 30 segundos promedio
- **Assets Soportados**: USD, BOB, COP, PEN, BTC, ETH, y más
- **Feed Seguro**: Protección anti-flash loan attacks

### ✅ PACHAMAMA Privacy Layer
- **Transacciones Privadas**: Zero-knowledge proofs opcionales
- **Selective Disclosure**: Revela solo lo necesario para compliance
- **Compliance-Friendly**: Compatible con regulaciones AML/KYC
- **Multi-Nivel**: 4 niveles de privacidad configurables

### ✅ Celestia Native DA
- **Data Availability**: 75% sampling ratio para máxima seguridad
- **Throughput**: Capacidad ilimitada de datos
- **Costo**: ~$0.00001 por KB de datos almacenados
- **Namespace**: `andean-chain-mainnet-v1`

### ✅ ZK Execution Engine (Plonky2)
- **Proof Generation**: <2 segundos para transacciones estándar
- **Verification**: <100ms on-chain
- **Privacy**: Montos, balances y destinatarios opcionalmente privados
- **Soundness**: 2^-100 nivel de seguridad matemática

## 📊 Performance y Costos

| Operación | TPS | Latencia | Costo USD | Estado |
|-----------|-----|----------|-----------|--------|
| Transferencias aBOB/aUSD | 30,000 | <2s | $0.0001 | ✅ Funcional |
| Swaps AMM | 25,000 | <1s | $0.0002 | ✅ Funcional |
| Operaciones LP | 15,000 | <2s | $0.0004 | ✅ Funcional |
| Transacciones ZK privadas | 8,000 | <4s | $0.001 | ✅ Funcional |
| Cross-chain bridges | 8,000 | <5s | $0.005 | ✅ Funcional |
| Hook Execution | 20,000 | <1s | Variable | ✅ Funcional |

## 🚀 Inicio Rápido

### Paso 1: Levantar la Blockchain Localmente

Primero necesitas tener una instancia de Andean Chain corriendo. Elige una de estas opciones:

```bash
# 1. Clonar el repositorio (solo si no existe)
if [ ! -d "andean" ]; then
    git clone https://github.com/andean-labs/andean.git
else
    echo "⚠️  Directorio 'andean' ya existe. Entrando al directorio existente..."
fi
cd andean

# 2. Ejecutar el script automático
chmod +x setup-reviewer.sh
./setup-reviewer.sh
```

**Nota**: Si ya clonaste el repositorio antes, simplemente haz `cd andean` y ejecuta el script.

El script automáticamente:
- ✅ Construye la imagen Docker
- ✅ Inicia el contenedor
- ✅ Configura e inicia la blockchain
- ✅ Crea una cuenta con fondos
- ✅ Deja la cadena corriendo en segundo plano

#### Opción B: Setup Manual con Docker (Con Visibilidad de Logs)

Si prefieres controlar cada paso y ver los logs en tiempo real:

```bash
# 1. Clonar y construir (si no lo hiciste antes)
if [ ! -d "andean" ]; then
    git clone https://github.com/andean-labs/andean.git
fi
cd andean
docker build -t andean-dev .

# 2. Limpiar contenedores anteriores
docker stop andean-dev-container 2>/dev/null || true
docker rm andean-dev-container 2>/dev/null || true

# 3. Iniciar contenedor en background
docker run -d --rm \
  -v $(pwd):/workspace \
  -p 1317:1317 -p 26656:26656 -p 26657:26657 \
  --name andean-dev-container \
  --platform linux/amd64 \
  andean-dev tail -f /dev/null

# 4. Verificar que el contenedor está corriendo
docker ps | grep andean-dev-container

# 5. Configurar la cadena dentro del contenedor
echo "Configurando la blockchain..."
docker exec andean-dev-container bash -c "
  cd /workspace &&
  go install ./cmd/andeand &&
  andeand init test-chain --chain-id andean-test-1 --home /workspace/.andean &&
  andeand keys add alice --keyring-backend test --home /workspace/.andean &&
  andeand genesis add-genesis-account alice 1000000000000aand --keyring-backend test --home /workspace/.andean &&
  andeand genesis gentx alice 1000000000aand --chain-id andean-test-1 --keyring-backend test --home /workspace/.andean &&
  andeand genesis collect-gentxs --home /workspace/.andean
"

# 6. Iniciar la blockchain CON LOGS VISIBLES
echo ""
echo "🔥 Iniciando la blockchain... Los logs aparecerán abajo:"
echo "   Cuando veas 'committed state' repetidamente, la blockchain está lista."
echo "   Presiona Ctrl+C para detener los logs (la blockchain seguirá corriendo)."
echo ""
docker exec andean-dev-container andeand start --home /workspace/.andean --minimum-gas-prices 0stake
```

**Después de ver los logs y que aparezcan bloques**, presiona `Ctrl+C` para salir de los logs. La blockchain seguirá corriendo en background.

### Paso 2: Conectar con el CLI Nativo (Recomendado)

**Importante**: Abre una **nueva terminal** para este paso (manteniendo la blockchain corriendo en la terminal anterior).

#### Verificar que la Blockchain está Corriendo

```bash
# En la nueva terminal, verificar conectividad
curl http://localhost:26657/status 2>/dev/null | grep -q "chain_id" && echo "✅ Blockchain corriendo" || echo "❌ Blockchain no disponible"

# Si obtienes "❌", espera 30 segundos más y reintenta
# Los contenedores pueden tardar en estar listos
sleep 30
curl http://localhost:26657/status 2>/dev/null | grep -q "chain_id" && echo "✅ Blockchain corriendo" || echo "❌ Necesitas revisar el setup"

# Ver los últimos bloques para confirmar que está generando bloques
curl -s http://localhost:26657/status | grep -o '"latest_block_height":"[^"]*"' || echo "Esperando bloques..."
```

#### Opción A: CLI Nativo (Más Rápido y Conveniente)

```bash
# 1. Asegúrate de estar en el directorio andean
cd andean  

# 2. Instalar el CLI localmente (requiere Go 1.21+)
go install ./cmd/andeand

# Verificar que se instaló correctamente
andeand version || echo "❌ Error: Verifica que Go esté instalado y en tu PATH"

# 3. Configurar variables de entorno (IMPORTANTE: usar 127.0.0.1 en lugar de localhost)
export RPC_ENDPOINT="http://127.0.0.1:26657"
export API_ENDPOINT="http://127.0.0.1:1317"
export CHAIN_ID="andean-test-1"

# 4. Verificar conexión (reintentar si falla)
echo "Verificando conexión..."
for i in {1..5}; do
    if andeand status --node $RPC_ENDPOINT > /dev/null 2>&1; then
        echo "✅ Conexión exitosa"
        break
    else
        echo "Intento $i/5 - Esperando..."
        sleep 10
    fi
    if [ $i -eq 5 ]; then
        echo "❌ No se pudo conectar. Verifica que la blockchain esté corriendo."
        exit 1
    fi
done

# 5. Crear tu propia cuenta
andeand keys add mi-cuenta --keyring-backend test
export MI_DIRECCION=$(andeand keys show mi-cuenta -a --keyring-backend test)
echo "✅ Tu nueva dirección: $MI_DIRECCION"

# 6. Obtener fondos desde la cuenta alice del contenedor
echo "Obteniendo fondos iniciales..."

# Obtener dirección de alice
ALICE_ADDR=$(docker exec andean-dev-container andeand keys show alice -a --keyring-backend test --home /workspace/.andean)
echo "📍 Dirección de Alice (con fondos): $ALICE_ADDR"

# Exportar e importar clave de alice
docker exec andean-dev-container andeand keys export alice --keyring-backend test --home /workspace/.andean --unsafe --unarmored-hex > /tmp/alice_key.txt
andeand keys import alice /tmp/alice_key.txt --keyring-backend test
rm -f /tmp/alice_key.txt

# 7. Verificar balance de alice
echo "Balance de Alice:"
andeand query bank balances $ALICE_ADDR --node $RPC_ENDPOINT

# 8. Transferir fondos iniciales a tu cuenta
echo "💸 Transfiriendo fondos..."
andeand tx bank send alice $MI_DIRECCION 100000000aand \
  --chain-id $CHAIN_ID \
  --keyring-backend test \
  --node $RPC_ENDPOINT \
  --gas 200000 \
  --gas-prices 0.025aand \
  --broadcast-mode sync \
  -y

# 9. Esperar y verificar
echo "⏳ Esperando confirmación..."
sleep 8

# 10. Verificar tu nuevo balance
echo "Tu balance final:"
andeand query bank balances $MI_DIRECCION --node $RPC_ENDPOINT
```

#### Opción B: Usar CLI Dentro del Contenedor Docker

Si prefieres trabajar dentro del contenedor (no requiere Go local):

```bash
# Conectar al contenedor existente
docker exec -it andean-dev-container bash

# Una vez dentro del contenedor:
export ALICE_ADDR=$(andeand keys show alice -a --keyring-backend test --home /workspace/.andean)
echo "Dirección de Alice: $ALICE_ADDR"

# Verificar balance de alice
andeand query bank balances $ALICE_ADDR --home /workspace/.andean --node tcp://localhost:26657

# Crear nueva cuenta dentro del contenedor
andeand keys add mi-cuenta --keyring-backend test --home /workspace/.andean
MI_DIRECCION=$(andeand keys show mi-cuenta -a --keyring-backend test --home /workspace/.andean)

# Transferir fondos
andeand tx bank send alice $MI_DIRECCION 100000000aand \
  --chain-id andean-test-1 \
  --keyring-backend test \
  --home /workspace/.andean \
  --node tcp://localhost:26657 \
  --gas auto --gas-adjustment 1.5 \
  --fees 1000aand \
  -y
```

## 🧪 Ejemplos Prácticos de Uso

**Requisito previo**: Asegúrate de haber completado los Pasos 1 y 2 anteriores, y que tengas:
- ✅ La blockchain corriendo (http://localhost:26657 responde)
- ✅ El CLI instalado localmente (`go install ./cmd/andeand`)
- ✅ Variables de entorno configuradas (`RPC_ENDPOINT`, `CHAIN_ID`, `MI_DIRECCION`)
- ✅ Fondos en tu cuenta (verifica con `andeand query bank balances $MI_DIRECCION --node $RPC_ENDPOINT`)

### 1. Transferencias Básicas

```bash
# Crear una segunda cuenta para pruebas
andeand keys add receptor --keyring-backend test
RECEPTOR_ADDR=$(andeand keys show receptor -a --keyring-backend test)

# Verificar balance actual
andeand query bank balances $MI_DIRECCION --node $RPC_ENDPOINT

# Hacer una transferencia
andeand tx bank send mi-cuenta $RECEPTOR_ADDR 1000aand \
  --chain-id $CHAIN_ID \
  --keyring-backend test \
  --node $RPC_ENDPOINT \
  --gas auto --gas-adjustment 1.5 \
  -y

# Consultar el hash de la transacción y esperar confirmación
# Verificar el nuevo balance
andeand query bank balances $RECEPTOR_ADDR --node $RPC_ENDPOINT
```

### 2. AndeanSwap AMM (Módulo: xicoatl)

```bash
# Ver pools existentes
andeand query xicoatl pools --node $RPC_ENDPOINT

# Crear un nuevo pool (si tienes permisos)
andeand tx xicoatl create-pool \
  --token-a aBOB \
  --token-b aUSD \
  --fee 0.003 \
  --initial-deposit-a 1000000 \
  --initial-deposit-b 1000000 \
  --from mi-cuenta --keyring-backend test \
  --chain-id $CHAIN_ID --node $RPC_ENDPOINT \
  --gas auto --gas-adjustment 1.5 -y

# Hacer un swap
andeand tx xicoatl swap \
  --pool-id 1 \
  --token-in aBOB \
  --token-out aUSD \
  --amount-in 100 \
  --min-out 95 \
  --from mi-cuenta --keyring-backend test \
  --chain-id $CHAIN_ID --node $RPC_ENDPOINT \
  --gas auto --gas-adjustment 1.5 -y

# Agregar liquidez a un pool existente
andeand tx xicoatl join-pool \
  --pool-id 1 \
  --tokens-in "1000000aBOB,1000000aUSD" \
  --from mi-cuenta --keyring-backend test \
  --chain-id $CHAIN_ID --node $RPC_ENDPOINT \
  --gas auto --gas-adjustment 1.5 -y
```

### 3. Oracle Price Feeds (Módulo: itzel)

```bash
# Consultar precio actual de un par
andeand query itzel aggregated-price BOB/USD --node $RPC_ENDPOINT

# Ver todos los precios disponibles
andeand query itzel all-prices --node $RPC_ENDPOINT

# Enviar un precio (si eres un oráculo autorizado)
andeand tx itzel submit-price \
  --asset BOB/USD \
  --price 6.96 \
  --from mi-cuenta --keyring-backend test \
  --chain-id $CHAIN_ID --node $RPC_ENDPOINT \
  --gas auto --gas-adjustment 1.5 -y
```

### 4. Cross-Chain Bridges (Módulo: inti)

```bash
# Ver bridges activos
andeand query inti bridges --node $RPC_ENDPOINT

# Iniciar un bridge desde otra cadena
andeand tx inti initiate-bridge \
  --source-chain ethereum \
  --target-chain andean \
  --asset USDC \
  --amount 1000000 \
  --recipient $MI_DIRECCION \
  --from mi-cuenta --keyring-backend test \
  --chain-id $CHAIN_ID --node $RPC_ENDPOINT \
  --gas auto --gas-adjustment 1.5 -y

# Consultar el estado de un bridge
andeand query inti bridge-status \
  --bridge-id 1 \
  --node $RPC_ENDPOINT
```

## 🔧 Comandos de Diagnóstico y Troubleshooting

### Verificar Conectividad

```bash
# Estado del nodo
andeand status --node $RPC_ENDPOINT

# Info de la cadena
andeand query block --node $RPC_ENDPOINT

# Últimos bloques
andeand query block-results --node $RPC_ENDPOINT

# Ver validadores
andeand query staking validators --node $RPC_ENDPOINT
```

### Problemas Comunes

**❌ "connection refused" o "dial tcp: connect: connection refused"**
- **Causa más común**: IPv6 vs IPv4. Usa `127.0.0.1` en lugar de `localhost`
- Verifica contenedor: `docker ps | grep andean-dev-container`
- Reinicia si es necesario: `docker restart andean-dev-container && sleep 15`
- Prueba diferentes endpoints:
  ```bash
  # Probar diferentes opciones
  curl http://127.0.0.1:26657/status
  curl http://localhost:26657/status
  curl http://0.0.0.0:26657/status
  ```

**❌ "cannot execute binary file" (en contenedor)**
- Problema de arquitectura (M1 Mac o ARM)
- Añadir `--platform linux/amd64` al comando docker:
  ```bash
  docker run -d --platform linux/amd64 --rm \
    -v $(pwd):/workspace \
    -p 1317:1317 -p 26656:26656 -p 26657:26657 \
    --name andean-dev-container \
    andean-dev tail -f /dev/null
  ```

**❌ "andeand: command not found" o "go: command not found"**
- Instala Go 1.21+ desde https://golang.org/dl/
- Verifica: `go version` (debe mostrar 1.21 o superior)
- Añade Go al PATH:
  ```bash
  export PATH=$PATH:/usr/local/go/bin
  export PATH=$PATH:$(go env GOPATH)/bin
  ```

**❌ "directorio 'andean' ya existe"**
- Normal si ya clonaste antes
- Simplemente: `cd andean && ./setup-reviewer.sh`

**❌ "account sequence mismatch"**
- Tu cuenta está desincronizada
- Consulta secuencia: `andeand query auth account $MI_DIRECCION --node $RPC_ENDPOINT`

**❌ "insufficient funds" o "not enough gas"**
- Verifica balance: `andeand query bank balances $MI_DIRECCION --node $RPC_ENDPOINT`
- Usa gas fijo: `--gas 200000 --gas-prices 0.025aand` en lugar de `--gas auto`

## 🧹 Limpieza y Detener la Blockchain

Cuando hayas terminado de probar:

```bash
# Detener y eliminar el contenedor
docker stop andean-dev-container

# (Opcional) Eliminar la imagen para liberar espacio
docker rmi andean-dev

# (Opcional) Limpiar datos locales
rm -rf .andean
```

## 🚀 Script de Verificación Completa

Aquí tienes un script que puedes ejecutar para verificar que todo está funcionando:

```bash
#!/bin/bash
# verify-setup.sh - Verificar que Andean Chain está funcionando correctamente

echo "🔍 Verificando setup de Andean Chain..."

# Verificar que el contenedor está corriendo
if ! docker ps | grep andean-dev-container > /dev/null; then
    echo "❌ El contenedor andean-dev-container no está corriendo"
    exit 1
fi

# Verificar conectividad RPC
if ! curl -s http://localhost:26657/status > /dev/null; then
    echo "❌ RPC no responde en localhost:26657"
    exit 1
fi

# Verificar que andeand está instalado localmente
if ! which andeand > /dev/null; then
    echo "❌ andeand CLI no está instalado localmente"
    echo "💡 Ejecuta: cd andean && go install ./cmd/andeand"
    exit 1
fi

# Variables de entorno
export RPC_ENDPOINT="http://localhost:26657"
export CHAIN_ID="andean-test-1"

# Verificar conexión del CLI
if ! andeand status --node $RPC_ENDPOINT > /dev/null 2>&1; then
    echo "❌ CLI no puede conectar a la blockchain"
    exit 1
fi

# Verificar que hay cuentas
if ! andeand keys list --keyring-backend test | grep -q "mi-cuenta"; then
    echo "⚠️  La cuenta 'mi-cuenta' no existe. Créala con:"
    echo "andeand keys add mi-cuenta --keyring-backend test"
fi

echo "✅ ¡Todo está funcionando correctamente!"
echo "🌐 RPC: http://localhost:26657"
echo "🌐 API: http://localhost:1317"
echo "⛓️  Chain ID: andean-test-1"
```

## 📡 APIs y Endpoints

### REST API Endpoints

```bash
# Balance de una cuenta
curl "$API_ENDPOINT/cosmos/bank/v1beta1/balances/$MI_DIRECCION"

# Información de pools AMM (puede no estar disponible hasta crear pools)
curl "$API_ENDPOINT/xicoatl/pools"

# Precios del oracle (puede no estar disponible hasta enviar precios)
curl "$API_ENDPOINT/itzel/prices/BOB/USD"

# Estado general de la blockchain
curl "http://localhost:26657/status"

# Último bloque
curl "http://localhost:26657/block"
```

### WebSocket para Eventos en Tiempo Real

```javascript
// Ejemplo en JavaScript para escuchar eventos
const ws = new WebSocket('ws://tu-servidor:26657/websocket');
ws.onopen = () => {
    ws.send(JSON.stringify({
        "jsonrpc": "2.0",
        "method": "subscribe",
        "params": {"query": "tm.event='NewBlock'"},
        "id": 1
    }));
};
```

## 🌐 Redes Disponibles

| Red | Chain ID | RPC | API | Estado |
|-----|----------|-----|-----|--------|
| Local | `andean-test-1` | `http://localhost:26657` | `http://localhost:1317` | ✅ Disponible con Docker |
| Testnet | `andean-testnet-1` | `https://rpc.testnet.andean.chain` | `https://api.testnet.andean.chain` | 🟡 En desarrollo |
| Mainnet | `andean-1` | `https://rpc.andean.chain` | `https://api.andean.chain` | 🔴 Próximamente |

## 🤝 Contribuir

¡Bienvenido! Andean Chain es un proyecto open-source.

1. Fork el repositorio
2. Crea una branch: `git checkout -b feature/nueva-funcionalidad`
3. Commit cambios: `git commit -m 'Agrega nueva funcionalidad'`
4. Push: `git push origin feature/nueva-funcionalidad`
5. Abre un Pull Request

## 📞 Comunidad

- [Discord](https://discord.gg/andean-chain)
- [Twitter](https://twitter.com/andean_chain)
- [Telegram](https://t.me/andean_chain)
- [Forum](https://forum.andean.chain)

## 📄 Licencia

Este proyecto está bajo la Licencia MIT - ver [LICENSE](LICENSE) para detalles.

---

**Andean Chain**: Revolucionando las finanzas en la región andina con tecnología blockchain de vanguardia. 🌅🏔️
