FROM golang:1.21-alpine3.17 AS core


ARG arch=x86_64

RUN apk add --no-cache curl make git libc-dev bash gcc linux-headers eudev-dev python3
WORKDIR /go/src/github.com/quadrateorg
RUN git clone https://github.com/QubeLedger/core.git

WORKDIR /go/src/github.com/quadrateorg/core

ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.4.0/libwasmvm_muslc.aarch64.a /lib/libwasmvm_muslc.aarch64.a
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.4.0/libwasmvm_muslc.x86_64.a /lib/libwasmvm_muslc.x86_64.a
RUN sha256sum /lib/libwasmvm_muslc.aarch64.a | grep 2a72c7062e3c791792b3dab781c815c9a76083a7997ce6f9f2799aaf577f3c25
RUN sha256sum /lib/libwasmvm_muslc.x86_64.a | grep 8ea2e3b5fae83e671da2bb51115adc88591045953f509955ec38dc02ea5a7b94

# Copy the library you want to the final location that will be found by the linker flag `-lwasmvm_muslc`
RUN cp /lib/libwasmvm_muslc.${arch}.a /lib/libwasmvm_muslc.a

RUN BUILD_TAGS=muslc make install
RUN cp ./qubed /go/bin/

FROM alpine:3.17
COPY --from=core /go/bin/qubed /usr/local/bin
RUN apk add --no-cache bash
RUN apk add --no-cache gcc

EXPOSE 26657
EXPOSE 1317

ENTRYPOINT ["qubed"]