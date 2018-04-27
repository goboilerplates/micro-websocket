CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main.out .
docker build -t goboilerplates/micro-websocket .
docker tag  goboilerplates/micro-websocket goboilerplates/micro-websocket:0.0.0