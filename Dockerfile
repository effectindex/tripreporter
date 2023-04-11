# SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
#
# SPDX-License-Identifier: OSL-3.0

FROM golang:alpine AS build
WORKDIR /tripreporter
COPY . .
RUN apk add --no-cache make npm && make

FROM alpine
WORKDIR /tripreporter
COPY --from=build /tripreporter/tripreporter /tripreporter/.env ./
COPY --from=build /tripreporter/ui/dist ./ui/dist
ENTRYPOINT ["/tripreporter/tripreporter", "-docker"]
