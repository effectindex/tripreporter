FROM golang:alpine AS build
WORKDIR /tripreporter
COPY . .
RUN apk add --no-cache make npm && make

FROM alpine
WORKDIR /tripreporter
COPY --from=build /tripreporter/tripreporter /tripreporter/.env ./
COPY --from=build /tripreporter/ui/dist ./ui/dist
ENTRYPOINT ["/tripreporter/tripreporter", "-docker"]
