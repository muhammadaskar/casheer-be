pipeline {
    agent any
    
    stages {
        stage('Set Environment Variables'){
            steps {
                echo 'Environment variables set.'
            }
        }
        
        stage('Pull Repositories') {
            steps {
                echo 'Hello, World Koding!'
            }
        }
        
        stage('Stop Container') {
            steps {
                echo 'Stopping the running container...'
                sh 'docker stop casheer-be-dev-container || true'
                sh 'docker rm casheer-be-dev-container || true'
                echo 'Container stopped.'
            }
        }
        
        stage('Docker Images') {
            steps {
                echo 'Building Docker images...'
                
                // Removing previous image
                sh 'docker rmi casheer-be-dev-image:latest || true'
                
                echo 'Building process...'
                sh 'docker build -t casheer-be-dev-image:latest .'
                echo 'Showing image results'
                sh 'docker images'
            }
        }
        
        stage('Deploy') {
            steps {
                echo 'Running the container...'
                
                sh 'docker run -d --name casheer-be-dev-container -p 2020:$(SERVER_PORT) \
                        -env DB_HOST=$(DB_HOST) \
                        -env DB_PORT=$(DB_PORT) \
                        -env DB_USER=$(DB_USER) \
                        -env DB_PASSWORD=$(DB_PASSWORD) \
                        -env DB_NAME=$(DB_NAME) \
                        -env SECRET_KEY=$(SECRET_KEY) \
                        -env SERVER_PORT=$(SERVER_PORT) \
                    casheer-be-dev-image:latest'
                echo 'Container is now running.'
                sh 'docker ps'
            }
        }
    }
}
