#!/bin/bash
SECONDS=0

cd $HOME/rides

msg () {
  echo -e "$1\n--------------------\n"
}

msg "Pulling from GitHub"
git pull

msg "Building Docker image"
sudo docker build --tag app .

msg "Stopping containers"
sudo docker compose down

msg "Starting containers"
sudo docker compose up -d

msg "Pruning stale Docker images"
sudo docker image prune -f


msg "Starting server"
nohup sudo ./main &>/dev/null &

duration=$SECONDS

echo
msg "Deploy finished in $(($duration % 60)) seconds."
msg "Press Enter to exit"
read
