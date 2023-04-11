# SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
#
# SPDX-License-Identifier: OSL-3.0

FROM library/postgres
RUN mkdir -p /docker-entrypoint-initdb.d/
COPY ./init.sql /docker-entrypoint-initdb.d/init.sql
