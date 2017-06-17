docker build -t "gotest" .
docker tag gotest docker-registry-default.cloud.expertsfactory.com/lab/gotest:latest
docker push docker-registry-default.cloud.expertsfactory.com/lab/gotest:latest
