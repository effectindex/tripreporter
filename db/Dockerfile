FROM library/postgres
RUN mkdir -p /docker-entrypoint-initdb.d/
COPY ./init.sql /docker-entrypoint-initdb.d/init.sql
