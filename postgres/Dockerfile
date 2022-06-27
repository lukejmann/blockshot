FROM postgres:12.1

RUN apt-get update && apt-get install -y wget \
    && rm -rf /var/lib/apt/lists/*

RUN wget -nv https://github.com/golang-migrate/migrate/releases/download/v4.13.0/migrate.linux-amd64.tar.gz \ 
  && tar -xzf migrate.linux-amd64.tar.gz \
  && rm migrate.linux-amd64.tar.gz \
  && cp migrate.linux-amd64 /bin/migrate

COPY init.sh /docker-entrypoint-initdb.d
