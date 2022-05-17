FROM golang:1.17-alpine AS base

WORKDIR /app

COPY ./go.*  ./
RUN go mod download

COPY ./ ./

FROM base AS go-builder

RUN CGO_ENABLED=0 go build \
  -ldflags "-w -s" \
  -installsuffix 'static' \
  -o /wwdatabase main.go

FROM scratch AS final-image

COPY --from=go-builder /wwdatabase /wwdatabase

ENTRYPOINT [ "/wwdatabase" ]
