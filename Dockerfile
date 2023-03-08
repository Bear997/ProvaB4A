FROM mysql:latest
RUN chown -R mysql:root /var/lib/mysql/

ARG MYSQL_DATABASE
ARG MYSQL_ROOT_PASSWORD
ARG MYSQL_PASSWORD
ARG MYSQL_USER

ENV MYSQL_DATABASE=${MYSQL_DATABASE:-prova}
ENV MYSQL_USER=${MYSQL_ROOT_PASSWORD:-password}
ENV MYSQL_PASSWORD=${MYSQL_PASSWORD:-passwordMatteo}
ENV MYSQL_ROOT_PASSWORD=${MYSQL_USER:-matteo}

# ADD data.sql /etc/mysql/data.sql

# RUN sed -i 's/MYSQL_DATABASE/'$MYSQL_DATABASE'/g' /etc/mysql/data.sql
# RUN cp /etc/mysql/data.sql /docker-entrypoint-initdb.d

EXPOSE 3306
RUN echo "ciao"