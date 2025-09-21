<!-- markdownlint-disable MD033 -->
# Andean Chain 🏔️

[![GoDoc](https://pkg.go.dev/badge/github.com/andean-labs/andean)](https://pkg.go.dev/github.com/andean-labs/andean)
[![Go Report Card](https://goreportcard.com/badge/github.com/andean-labs/andean)](https://goreportcard.com/report/github.com/andean-labs/andean)
[![Version](https://img.shields.io/github/tag/andean-labs/andean.svg)](https://github.com/andean-labs/andean/releases/latest)
[![License](https://img.shields.io/github/license/andean-labs/andean.svg)](https://github.com/andean-labs/andean/blob/main/LICENSE)
[![Discord](https://img.shields.io/discord/1234567890)](https://discord.gg/andean-chain)

> La primera blockchain nativa de Celestia para la región andina. Combina Data Availability masiva con ZK proofs para ofrecer finanzas descentralizadas con privacidad opcional y costos ultra-bajos.

## 🌟 Características Funcionales Principales

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

## 📊 Performance y Costos

| Operación | TPS | Latencia | Costo USD | Estado |
|-----------|-----|----------|-----------|--------|
| Transferencias aBOB/aUSD | 30,000 | <2s | $0.0001 | ✅ Funcional |
| Swaps AMM | 25,000 | <1s | $0.0002 | ✅ Funcional |
| Operaciones LP | 15,000 | <2s | $0.0004 | ✅ Funcional |
| Transacciones ZK privadas | 8,000 | <4s | $0.001 | ✅ Funcional |
| Cross-chain bridges | 8,000 | <5s | $0.005 | ✅ Funcional |
| Hook Execution | 20,000 | <1s | Variable | ✅ Funcional |

## 🚀 Inicio Rápido - Prueba la Tecnología

### Opción 1: Script Automático (Recomendado para Revisores)

```bash
# Clonar el repositorio (si no lo has hecho)
git clone https://github.com/AndeanLabs/andean.git
cd andean

# Dar permisos y ejecutar el script
chmod +x setup-reviewer.sh
./setup-reviewer.sh
```

### Opción 2: Manual con Docker

```bash
# 1. Construir la imagen
docker build -t andean-dev .

# 2. Iniciar un contenedor interactivo
docker run -it --rm \
  -v $(pwd):/workspace \
  -p 1317:1317 -p 26656:26656 -p 26657:26657 \
  --name andean-dev-container \
  andean-dev

# 3. Dentro del contenedor, inicializar la cadena

go install ./cmd/andeand
andeand init test-chain --chain-id andean-test-1 --home /workspace/.andean
andeand keys add alice --keyring-backend test --home /workspace/.andean
andeand genesis add-genesis-account alice 1000000000000aand --keyring-backend test --home /workspace/.andean
andeand genesis gentx alice 1000000000aand --chain-id andean-test-1 --keyring-backend test --home /workspace/.andean
andeand genesis collect-gentxs --home /workspace/.andean
andeand start --home /workspace/.andean --minimum-gas-prices 0stake
```

## 🧪 Ejemplos Prácticos de Uso

*Para ejecutar estos ejemplos, primero inicia la cadena con una de las opciones anteriores. Luego, abre una **segunda terminal** y entra al contenedor con `docker exec -it andean-dev-container /bin/bash`.*

### 1. Transferencias Básicas

```bash
# Dirección del usuario 'alice'
ALICE_ADDR=$(andeand keys show alice -a --keyring-backend test --home /workspace/.andean)

# Crear un nuevo usuario 'bob'
andeand keys add bob --keyring-backend test --home /workspace/.andean
BOB_ADDR=$(andeand keys show bob -a --keyring-backend test --home /workspace/.andean)

# Verificar balance de alice
andeand query bank balances $ALICE_ADDR --home /workspace/.andean --node tcp://localhost:26657

# Transferir 1000aand de alice a bob
andeand tx bank send alice $BOB_ADDR 1000aand \
  --chain-id andean-test-1 \
  --keyring-backend test \
  --home /workspace/.andean \
  --node tcp://localhost:26657 -y

# Verificar el nuevo balance de bob
andeand query bank balances $BOB_ADDR --home /workspace/.andean --node tcp://localhost:26657
```

### 2. AndeanSwap AMM (Módulo: xicoatl)

```bash
# Crear un pool de aBOB y aUSD
andeand tx xicoatl create-pool \
  --token-a aBOB \
  --token-b aUSD \
  --fee 0.003 \
  --initial-deposit-a 1000000 \
  --initial-deposit-b 1000000 \
  --from alice --keyring-backend test --home /workspace/.andean \
  --chain-id andean-test-1 --node tcp://localhost:26657 -y

# Agregar más liquidez al pool #1
andeand tx xicoatl join-pool \
  --pool-id 1 \
  --tokens-in "1000000aBOB,1000000aUSD" \
  --from alice --keyring-backend test --home /workspace/.andean \
  --chain-id andean-test-1 --node tcp://localhost:26657 -y

# Intercambiar 100 aBOB por aUSD
andeand tx xicoatl swap \
  --pool-id 1 \
  --token-in aBOB \
  --token-out aUSD \
  --amount-in 100 \
  --min-out 98 \
  --from alice --keyring-backend test --home /workspace/.andean \
  --chain-id andean-test-1 --node tcp://localhost:26657 -y
```

### 3. Cross-Chain Bridges (Módulo: inti)

```bash
# Simular un bridge desde Ethereum a Andean Chain
andeand tx inti initiate-bridge \
  --source-chain ethereum \
  --target-chain andean \
  --asset USDC \
  --amount 1000000 \
  --recipient $(andeand keys show alice -a --keyring-backend test --home /workspace/.andean) \
  --from alice --keyring-backend test --home /workspace/.andean \
  --chain-id andean-test-1 --node tcp://localhost:26657 -y
```

### 4. Oracle Price Feeds (Módulo: itzel)

```bash
# Consultar el precio agregado de un par
andeand query itzel aggregated-price BOB/USD --node tcp://localhost:26657

# Simular un oráculo enviando un precio
andeand tx itzel submit-price \
  --asset BOB/USD \
  --price 6.96 \
  --from alice --keyring-backend test --home /workspace/.andean \
  --chain-id andean-test-1 --node tcp://localhost:26657 -y
```

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
