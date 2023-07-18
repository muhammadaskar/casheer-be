pipeline {
    agent any
    
    stages {
        stage('Set Environment Variables'){
            environment {
                DB_HOST = credentials('DB_HOST')
                DB_PORT = credentials('DB_PORT')
                DB_USER = credentials('DB_USER')
                DB_PASSWORD = credentials('DB_PASSWORD')
                DB_NAME = credentials('DB_NAME')
                SECRET_KEY = credentials('SECRET_KEY')
                SERVER_PORT = credentials('SERVER_PORT')
            }
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
                sh 'docker stop dev-casheer-be-container || true'
                sh 'docker rm dev-casheer-be-container || true'
                echo 'Container stopped.'
            }
        }
        
        stage('Docker Images') {
            steps {
                echo 'Building Docker images...'
                
                // Removing previous image
                sh 'docker rmi dev-casheer-be-image:latest || true'
                
                echo 'Building process...'
                sh 'docker build -t dev-casheer-be-image:latest .'
                echo 'Showing image results'
                sh 'docker images'
                
                // Copying and modifying .env file
                sh 'docker run --name temp-container -d dev-casheer-be-image:latest sleep 1d'
                sh 'docker cp /path/to/.env.example temp-container:/app/.env.example'
                sh 'docker exec temp-container cp /app/.env.example /app/.env'
                sh "docker exec temp-container sed -i \"s/DB_HOST=.*/DB_HOST=${DB_HOST}/\" /app/.env"
                sh "docker exec temp-container sed -i \"s/DB_PORT=.*/DB_PORT=${DB_PORT}/\" /app/.env"
                sh "docker exec temp-container sed -i \"s/DB_USER=.*/DB_USER=${DB_USER}/\" /app/.env"
                sh "docker exec temp-container sed -i \"s/DB_PASSWORD=.*/DB_PASSWORD=${DB_PASSWORD}/\" /app/.env"
                sh "docker exec temp-container sed -i \"s/DB_NAME=.*/DB_NAME=${DB_NAME}/\" /app/.env"
                sh "docker exec temp-container sed -i \"s/SECRET_KEY=.*/SECRET_KEY=${SECRET_KEY}/\" /app/.env"
                sh "docker exec temp-container sed -i \"s/SERVER_PORT=.*/SERVER_PORT=${SERVER_PORT}/\" /app/.env"
                sh 'docker cp temp-container:/app/.env .env'
                sh 'docker rm -f temp-container'
            }
        }
        
        stage('Deploy') {
            steps {
                echo 'Running the container...'
                
                sh 'docker run -d --name dev-casheer-be-container -p 3030:3030 --env-file .env dev-casheer-be-image:latest'
                echo 'Container is now running.'
                sh 'docker ps'
            }
        }
    }
}
