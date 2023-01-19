# teststocks

You will need to create 2 secrets for the db user and password

echo "user" | docker secret create stocksmongouser -

echo "pass" | docker secret create stocksmongouserpassword -

To run the stack in swarm mode to use the secrets:
docker stack deploy --compose-file docker-compose.yaml stocksapp --with-registry-auth