# syntax=docker/dockerfile:1

# Usar versión de Go específica
ARG GO_VERSION="1.24.1"
ARG ALPINE_VERSION="3.20"

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION}

# Metadatos
LABEL maintainer="Andean Chain Development"
LABEL description="Development environment for Andean Chain"

# Instalar dependencias del sistema
RUN apk add --no-cache \
    git \
    build-base \
    curl \
    bash \
    jq \
    wget \
    linux-headers \
    binutils-gold

# Instalar Ignite CLI como root primero (método manual)
ENV IGNITE_VERSION="v29.4.0"
RUN ARCH=$(uname -m | sed 's/x86_64/amd64/; s/aarch64/arm64/') && \
    wget "https://github.com/ignite/cli/releases/download/${IGNITE_VERSION}/ignite_29.4.0_linux_${ARCH}.tar.gz" -O ignite.tar.gz && \
    tar -xzf ignite.tar.gz && \
    mv ignite /usr/local/bin/ && \
    chmod +x /usr/local/bin/ignite && \
    rm ignite.tar.gz

# Crear usuario no-root para desarrollo
RUN adduser -D -s /bin/bash developer

# Cambiar a usuario developer
USER developer
WORKDIR /home/developer

# Configurar variables de entorno de Go
ENV GOPROXY=https://proxy.golang.org,direct
ENV GOSUMDB=sum.golang.org
ENV GO111MODULE=on
ENV PATH="/usr/local/bin:${PATH}"

# Crear directorio de trabajo del proyecto
WORKDIR /workspace

# Copiar go.mod y go.sum primero para aprovechar el caché de capas de Docker
COPY --chown=developer:developer go.mod go.sum ./

# Descargar dependencias
RUN go mod download && go mod verify

# Copiar el resto del código fuente
COPY --chown=developer:developer . .

# Exponer puertos necesarios para desarrollo
EXPOSE 1317 26656 26657 26658 4000 8545

# Script de entrada para desarrollo interactivo
ENTRYPOINT ["/bin/bash"]