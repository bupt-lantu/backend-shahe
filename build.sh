bee generate docs
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bupt_tour .
docker build --rm -f "Dockerfile" -t wzekin/bupt_tour .
docker push wzekin/bupt_tour 

