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
                sh 'docker stop prod-casheer-be-container || true'
                sh 'docker rm prod-casheer-be-container || true'
                echo 'Container stopped.'
            }
        }

        stage('Copy .env.example to .env') {
            steps {
                // Use cat to read the content of .env.example and tee to write it to .env
                sh 'cat .env.example | tee .env'
            }
        }

        stage('Set Environment Variables') {
            steps {
                // Use Jenkins environment variables to replace values in the .env file
                sh "sed -i 's#ALLOW_ORIGIN_PROD=.*#ALLOW_ORIGIN_PROD=${ALLOW_ORIGIN_PROD}#' .env"
                sh "sed -i 's#ALLOW_ORIGIN_DEV=.*#ALLOW_ORIGIN_DEV=${ALLOW_ORIGIN_DEV}#' .env"
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
                sh 'docker rmi prod-casheer-be-image:latest || true'
                
                echo 'Building process...'
                sh 'docker build -t prod-casheer-be-image:latest .'
                echo 'Showing image results'
                sh 'docker images'
            }
        }
        
        stage('Deploy') {
            steps {
                echo 'Running the container...'
                
                sh 'docker run -d --name prod-casheer-be-container -p ${SERVER_PORT}:${SERVER_PORT} --env-file .env prod-casheer-be-image:latest'
                echo 'Container is now running.'
                sh 'docker ps'
            }
        }
    }
}
