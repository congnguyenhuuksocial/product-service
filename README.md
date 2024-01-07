1. **Overall System Architecture**

   The system is split into several microservices, each following the Clean Architecture principles. Key microservices could be:

    `Product Service:` Manages product-related operations. (Tobe implemented)

    `Order Service:` Handles order processing and management. (Tobe implemented)

    `Inventory Service:` Manages stock and inventory levels. (Tobe implemented)

    `API Gateway:` Acts as the single entry point to the system, routing requests to appropriate services.

    Each microservice is designed with its own database to ensure service independence and scalability. They communicate asynchronously via Kafka for event-driven interactions and synchronously via gRPC or REST APIs for direct service-to-service communication.

2. **Microservice Internal Structure**
   
    Each microservice follows a similar internal structure:

- `API Layer (Handlers/Controllers):` Handles incoming requests and responses. It interacts only with the Application Layer.
- `Application Layer (Command/Query Handlers):` Implements CQRS pattern; it handles business logic for processing commands (writes) and queries (reads).
- `Domain Layer (Core):`
  - `Entities/Aggregates`: Business objects and logic.
  - `Events`: Changes or actions represented as domain events.
  - `Ports`: Defines interfaces for external communications and data persistence.
- `Infrastructure Layer:`
  - `Repositories:` Implements data access and storage mechanisms.
  - `Event Store:` Manages event persistence and retrieval.
  - `Message Bus (Kafka Producer/Consumer):` Handles publishing to and consuming from Kafka topics.
  - `External Services (e.g., Payment Gateway):` Interfaces to external systems.
3. **Event Sourcing and Kafka Integration**
   
    Domain events (e.g., ProductCreated, OrderPlaced) are produced and stored in the Event Store upon state changes.
   These events are also published to Kafka, enabling other microservices to react to these changes.
   Microservices subscribe to relevant Kafka topics to listen for domain events they are interested in.
4. **Saga Orchestrator for Distributed Transactions**

   Orchestrates long-running business processes spanning multiple microservices.
   Manages compensating transactions for rollback mechanisms in case of failures to ensure data consistency across services.
5. **Data Storage and Caching**
   
    Each service has its own relational or NoSQL database managed via ORMs like GORM.
   Redis is used for caching, especially for read-heavy operations or data that requires quick access.
6. **Containerization and Orchestration**
   
    Each microservice, along with its dependent services like databases and Kafka, is containerized using Docker.
    Docker Compose or Kubernetes can be used for local development and production deployment, respectively.
7. **Logging, Monitoring, and Tracing**

   Implement logging in each service for audit trails and debugging.
   Use tools like Prometheus and Grafana for monitoring and alerting.
   Implement distributed tracing with tools like Jaeger or Zipkin.
8. *API Gateway*

    An API Gateway acts as the single entry point to the system, routing requests to appropriate services.
## Diagram
   Here's a simplified diagram to visualize the architecture:
```
  [ API Gateway ]
        |
  [ Service Mesh ]
  /       |       \
[Product] [Order] [Inventory]
   |        |        |
[Kafka] [Kafka]  [Kafka]
   \        |       /
    [Kafka Cluster]
```
# Coding structure
```
/product-service
  /api
    /proto
    /handlers
  /cmd
    /server
      main.go
  /internal
    /core
      /domain
      /ports
      /services
    /infrastructure
      /repository
      /eventstore
    /application
      /commandhandlers
      /queryhandlers
    /adapter
      /grpc
      /kafka
      /elasticsearch
  /pkg
    /config
    utils.go
  Dockerfile
  go.mod
  .env
```
## Clean Architecture and CQRS
Clean architecture promotes separation of business logic from interfaces. CQRS stands for Command Query Responsibility Segregation, where you separate the write (command) and read (query) aspects of your application.

The `/core/domain` directory would contain your business entities and aggregates.
The `/core/services` would contain the application logic that orchestrates commands and queries.
The `/core/ports` would be interfaces that represent the expected operations for external tools like databases or client interfaces.

## Event Sourcing
The Event Sourcing pattern stores state changes as a series of events. These events are then used to rebuild current state or transfer state in between services. Each microservice would have:

An `/infrastructure/eventstore` directory where you would implement the logic to save and replay these events.

## How to deploy on AWS
### Step 1: Create an EKS Cluster
Use eksctl, a simple CLI tool for creating clusters on EKS:

```
eksctl create cluster --name my-cluster --region us-west-2
```
## Step 2: Set Up AWS App Mesh
Install the App Mesh Kubernetes components. First, add the EKS chart repo and update Helm:
```
helm repo add eks https://aws.github.io/eks-charts
helm repo update
```
Install the App Mesh CRDs and controller:
```
kubectl apply -k "github.com/aws/eks-charts/stable/appmesh-controller/crds?ref=master"
helm upgrade -i appmesh-controller eks/appmesh-controller \
    --namespace appmesh-system \
    --set region=us-west-2 \
    --set serviceAccount.create=false \
    --set serviceAccount.name=appmesh-controller
``` 
## Step 3: Deploy a Sample Application with App Mesh Integration
Create a Helm chart for your application, including App Mesh resources like VirtualNode and VirtualService.

### chart structure
```
/product-app-chart
  /charts
  /templates
    deployment.yaml
    service.yaml
    virtual-node.yaml
    virtual-service.yaml
  Chart.yaml
  values.yaml
```
Include App Mesh sidecar containers in deployment.yaml:
```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ include "product-app.fullname" . }}
  labels:
    {{- include "product-app.labels" . | nindent 4 }}
spec:
  replicas: {{ .Values.replicaCount }}
  selector:
    matchLabels:
      {{- include "product-app.selectorLabels" . | nindent 6 }}
  template:
    metadata:
      {{- with .Values.podAnnotations }}
      annotations:
        {{- toYaml . | nindent 8 }}
      {{- end }}
      labels:
        {{- include "product-app.selectorLabels" . | nindent 8 }}
    spec:
      containers:
        - name: {{ .Chart.Name }}
          image: "{{ .Values.image.repository }}:{{ .Values.image.tag | default .Chart.AppVersion }}"
          ports:
            - containerPort: {{ .Values.service.internalPort }}
          livenessProbe:
            httpGet:
              path: /healthz
              port: http
          readinessProbe:
            httpGet:
              path: /healthz
              port: http
          resources:
            {{- toYaml .Values.resources | nindent 12 }}
      
      # App Mesh Envoy sidecar container
      - name: envoy
        # Specify the Envoy image
        image: "840364872350.dkr.ecr.us-west-2.amazonaws.com/aws-appmesh-envoy:v1.15.1.0-prod"
        securityContext:
          runAsUser: 1337
        env:
          - name: APPMESH_VIRTUAL_NODE_NAME
            value: "mesh/{{ .Values.appMesh.meshName }}/virtualNode/{{ .Values.appMesh.virtualNodeName }}"
          - name: AWS_REGION
            value: "{{ .Values.appMesh.region }}"
          - name: ENVOY_LOG_LEVEL
            value: "{{ .Values.appMesh.envoyLogLevel }}"
        resources:
          {{- toYaml .Values.appMesh.envoyResources | nindent 10 }}
```
## Step 4: Implement Autoscaling
`/k8s/templates/hpa.yaml`
```
apiVersion: autoscaling/v2beta2
kind: HorizontalPodAutoscaler
metadata:
  name: {{ include "product-app-chart.fullname" . }}
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: {{ include "product-app-chart.fullname" . }}
  minReplicas: 1
  maxReplicas: 100
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 80
```
## Step 5: Deploy Your Application with Helm
```
helm install product ./product-app-chart
```

# Data flow
1. Actors:

- `User:` Interacts with the web application.
- `Admin:` Performs administrative tasks.
2. System Components:

`Web Application:` The main interface for users.
`API Gateway:` Routes requests to appropriate services.
`Service Layer:` Business logic (e.g., ProductService).
`Data Stores:` MySQL, Redis, Elasticsearch.
`Message Queue:` Kafka for asynchronous processing.

3. Data Flow:

- Users send requests to the Web Application.
- The Web Application communicates with the Service Layer via the API Gateway.
- The Service Layer performs business logic, interacts with Data Stores, and publishes messages to Kafka.
- Kafka Consumers process these messages and might interact back with Services or Data Stores.

**UML Data Flow Diagram**
Open by [PlantUML](https://plantuml.com/zh/)
```
@startuml
!define RECTANGLE class

actor User
actor Admin

RECTANGLE WebApplication {
  :User Interface:
  :API Calls:
}

RECTANGLE APIGateway {
  :Request Routing:
}

RECTANGLE Services {
  :Business Logic:
  :Data Processing:
}

database MySQL {
  :Relational Data:
}

database Redis {
  :Caching:
}

database Elasticsearch {
  :Search Index:
}

queue Kafka {
  :Message Queue:
}

User --> WebApplication : uses
Admin --> WebApplication : manages
WebApplication --> APIGateway : routes
APIGateway --> Services : requests
Services --> MySQL : CRUD Operations
Services --> Redis : Read/Write Cache
Services --> Elasticsearch : Index/Search
Services --> Kafka : Publish Messages
Kafka --> Services : Consumes Messages
@enduml
```
![dataflow.png](images%2Fdataflow.png)

# How to run locally

Step 1:  Install Docker and Docker Compose
```
brew install docker
```

Step 2: Clone the repository
```
git clone https://github.com/congnguyenhuuksocial/product-service.git
```
Step 3: Run the application
```
docker-compose up
```

Step 4: Test the application by gRPC client

Step 5: double check result in database


![Screenshot 2024-01-07 at 17.52.24.png](images%2FScreenshot%202024-01-07%20at%2017.52.24.png)
![Screenshot 2024-01-07 at 17.53.08.png](images%2FScreenshot%202024-01-07%20at%2017.53.08.png)
![Screenshot 2024-01-07 at 17.53.36.png](images%2FScreenshot%202024-01-07%20at%2017.53.36.png)

# libraries used
- [Golang](https://golang.org/)
- [gRPC](https://grpc.io/)
- [GORM](https://gorm.io/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Kafka](https://kafka.apache.org/)
- [Elasticsearch](https://www.elastic.co/)
- [Redis](https://redis.io/)
- [Validator] validation for golang v10
- [Gin] web framework for golang
- [Mysql] database
# Todo list
- implement rest api for gateway
- implement kafka consumer
- implement elasticsearch
- implement redis
- implement order service
- implement inventory service
