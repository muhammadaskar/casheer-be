pipeline {
    agent any
    
    stages {
        stage('Pull Repositories') {
            steps {
                echo 'Pull Repositories!'
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

        stage('Set Environment Variables') {
            steps {
                // Use Jenkins environment variables to replace values in the .env file
                sh "sed -i 's/DB_HOST=.*/DB_HOST=${DB_HOST}/' .env"
                sh "sed -i 's/DB_PORT=.*/DB_PORT=${DB_PORT}/' .env"
                sh "sed -i 's/DB_USER=.*/DB_USER=${DB_USER}/' .env"
                sh "sed -i 's/DB_PASSWORD=.*/DB_PASSWORD=${DB_PASSWORD}/' .env"
                sh "sed -i 's/DB_NAME=.*/DB_NAME=${DB_NAME}/' .env"
                sh "sed -i 's/SECRET_KEY=.*/SECRET_KEY=${SECRET_KEY}/' .env"
                sh "sed -i 's/SERVER_PORT=.*/SERVER_PORT=${SERVER_PORT}/' .env"
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
                
                sh 'docker run -d --name casheer-be-dev-container -p 2020:$(SERVER_PORT) --env-file .env casheer-be-dev-image:latest'
                echo 'Container is now running.'
                sh 'docker ps'
            }
        }
    }
}
