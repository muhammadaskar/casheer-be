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
            }
        }
        
        stage('Deploy') {
            steps {
                echo 'Running the container...'
                
                sh 'docker run -d --name dev-casheer-be-container -p 3030:3030 dev-casheer-be-image:latest'
                echo 'Container is now running.'
                sh 'docker ps'
            }
        }
    }
}
