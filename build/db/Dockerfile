FROM library/mysql:8.0.25

ENV MYSQL_DATABASE app

COPY custom.cnf /etc/mysql/conf.d/

COPY sql /docker-entrypoint-initdb.d
