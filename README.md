# sample-grpc
## Steps to Run Locally

1. **Build the Go binary:**
   ```sh
   GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o server main.go
   ```

2. **Run the server:**
   ```
   go run main.go
   docker build -t sample-grpc-server .
   docker run -p 5001:5001 sample-grpc-server
   ```

## Steps to Build and Run with Docker

1. **Build the Docker image:**
   ```sh
   docker build -t sample-grpc-server .
   ```

2. **Run the Docker container:**
   ```sh
   docker build -t sample-grpc-server .
   docker run -p 5001:5001 sample-grpc-server
   ```

## Steps to Deploy on Minikube

1. **Start Minikube:**
   ```sh
   minikube start
   ```

2. **Enable metrics-server (for autoscaling):**
   ```sh
   minikube addons enable metrics-server
   ```

3. **Load Docker image into Minikube (if using local image):**
   ```sh
   minikube image load sample-grpc-server
   ```

4. **Apply Kubernetes manifests:**
   ```sh
   kubectl apply -f manifests/deployment.yaml
   kubectl apply -f manifests/service.yaml
   kubectl apply -f manifests/hpa.yaml
   ```

5. **Access the service:**
   ```sh
   minikube service grpc-server-service
   ```

## Testing the gRPC Server

- 
use postman
'{"user": {"id": "1", "name": "Satyam", "email": "satyam@example.com"}}' localhost:5001 samplegrpc.UserService/CreateUser
  ```