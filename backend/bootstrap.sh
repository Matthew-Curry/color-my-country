### Script to startup the server ###

logMessage "Setting script vars"

# email for certbot
EMAIL="matt.curry56@gmail.com"

# directories to store source
CODE_DIR="/code" 
CERT_DIR="$CODE_DIR/cert"

# install docker
logMessage "Installing Docker"
apt-get install \
    ca-certificates \
    curl \
    gnupg \
    lsb-release

mkdir -p /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor -o /etc/apt/keyrings/docker.gpg

echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null

apt-get update
apt-get install docker-ce docker-ce-cli containerd.io docker-compose-plugin

# install git
apt-get install git

# to run docker without sudo
logMessage "Adding Ubuntu user to the docker group"
groupadd docker
usermod -aG docker $USER
newgrp docker 

# pull latest docker images
logMessage "Pulling latest docker images for the application"

# pull needed files to run app
logMessage "Pulling code from github"
cd $CODE_DIR
git clone git@github.com:Matthew-Curry/color-my-country.git

# setup certs
logMessage "Setting up the certs"
certbot certonly --webroot -w certs \ 
              --preferred-challenges http \
              -d reregion.com -d www.reregion.com \ 
              --non-interactive --agree-tos -m $EMAIL

# start the app
logMessage "Starting the app"
chmod 700 run_server.sh
./run_server.sh

logMessage "Startup script has run successfully."
