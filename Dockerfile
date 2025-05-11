# --- Build frontend ---
FROM node:22.14 AS frontend-build

WORKDIR /app/frontend
COPY frontend/package*.json ./
RUN npm install
COPY frontend/ ./
RUN npm run build

# --- Build backend ---
FROM golang:1.24 AS backend-build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Copy built frontend into backend's embed/build directory if needed
RUN rm -rf frontend/build && \
    mkdir -p frontend/build && \
    cp -r frontend/build /app/frontend/
RUN CGO_ENABLED=0 GOOS=linux go build -o josie ./cmd/api

# --- Final image ---
FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=backend-build /app/josie .
COPY --from=backend-build /app/frontend/build ./frontend/build

EXPOSE 7000

CMD ["./josie"]