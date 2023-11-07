FROM mongo:7.0

# copy seed data and the initialization script to
# the correct location to be loaded by mongo
COPY ./users.js /docker-entrypoint-initdb.d/
COPY ./counties.js /docker-entrypoint-initdb.d/
COPY ./init-db.js /docker-entrypoint-initdb.d/

# this is the standard mongo port
EXPOSE 27017
