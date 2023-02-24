source ~/.bash_profile

cd /vagrant

docker compose -f docker-compose.yml pull
docker compose -f docker-compose.yml up -d