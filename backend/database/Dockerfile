FROM node:18.5-slim

# set workdir, copy required files.
WORKDIR /usr/src/app

COPY data/users.json ./
COPY data/counties.json ./

COPY init/init-db.js ./
COPY init/package.json ./
COPY init/package-lock.json ./
COPY init/collections.json ./

# install dependencies.
RUN npm install

# run database initialization. 
CMD ["npm", "run", "init-color-my-country-db"]
