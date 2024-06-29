# Log Analysis using Kafka-ELK stack   
This final-year engineering project demonstrates how to analyze and visually represent logs using Kafka and Elasticsearch, Logstash, Kibana (ELK) stack.

### Prerequisites
Before you begin, you should have the following tools installed on your local machine:
- Docker (Running)
- GoLang
- Postman
- Git

### Demonstration
Follow these steps to install and run the log analysis service on your local machine:
1. Clone this repository to your local machine: `https://github.com/akhil3shh/Monitoring-as-a-Service.git`
2. Navigate to the directory: `cd Monitoring-as-a-Service\ELK-Kafka-Docker\kafka-elk`
3. Start/restart all the services defined in docker-compose.yml using: `docker-compose up`   
4. Leave this window running in the background and wait until all containers have been created. This step may take a while since you're pulling the base images for the first time.  
5. Once done, open another window in your terminal and navigate to the directory: `cd Monitoring-as-a-Service\log-gen`
6. To start generating logs run: `go run main.go`
7. Open Postman and use `http://localhost:8080/gen/start` with `POST` method to buffer logs

### Testing
1. Open `http://localhost:5601` in your browser and head towards Discover section under Kibana. Use `codespotify-index` to get a visual representation of incoming logs.
2. Feel free to play around with the interface to discover more insights and build more dashboards depending on your usecase.
3. Optionally, you may also modify Go files present in the `Monitoring-as-a-Service\log-gen` directory according to your needs.

### Clean-up
1. Once you're done exploring, head towards Postman and use `http://localhost:8080/gen/stop` with `POST` method to stop generating the logs.
2. Press `Ctrl+C` on both terminals to stop log generation and gracefully stop the containers.
3. If you wish, use `docker rmi <IMAGE_ID>` and `docker rm <CONTAINER_ID>` commands to delete the images and containers that were created earlier.
4. Additionally, you may use `docker images` and `docker ps -a` to list down the images and containers.

### Contributing
- If you'd like to contribute to this project, please open an issue or submit a pull request.

