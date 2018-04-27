echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
docker push goboilerplates/micro-websocket
docker push goboilerplates/micro-websocket:0.0.0