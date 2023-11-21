pipeline {
    agent any
    
    stages {
        stage('Pull Repositories') {
            steps {
                echo 'Pull Repositories!'
            }
        }
        
        stage('Stop Docker Container and Remove Container') {
            steps {
                echo 'Stopping the running container...'
                sh 'docker stop dev-casheer-be-container || true'
                sh 'docker rm dev-casheer-be-container || true'
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
                sh "sed -i 's/DB_HOST=.*/DB_HOST=${DB_HOST}/' .env"
                sh "sed -i 's/DB_PORT=.*/DB_PORT=${DB_PORT}/' .env"
                sh "sed -i 's/DB_USER=.*/DB_USER=${DB_USER}/' .env"
                sh "sed -i 's/DB_PASSWORD=.*/DB_PASSWORD=${DB_PASSWORD}/' .env"
                sh "sed -i 's/DB_NAME=.*/DB_NAME=${DB_NAME}/' .env"
                sh "sed -i 's/SECRET_KEY=.*/SECRET_KEY=${SECRET_KEY}/' .env"
                sh "sed -i 's/SERVER_PORT=.*/SERVER_PORT=${SERVER_PORT_DEV}/' .env"
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
        
        stage('Deploy with Container Run') {
            steps {
                echo 'Deploying with Docker Container Run...'
                sh 'docker run -d -p ${SERVER_PORT_DEV}:${SERVER_PORT_DEV} -v dev_asset_image:/app/asset/image/ --name dev-casheer-be-container dev-casheer-be-image:latest'
                echo 'docker ps'
            }
        }

    }
    post {
        success {
            // Script to be executed if the deployment is successful
            slackSend color: 'good', message: 'Deployment successful for *dev-casheer-be* :white_check_mark:. The application has been deployed successfully.'
        }
        failure {
            // Script to be executed if the deployment fails
            slackSend color: 'danger', message: 'Deployment failed for *dev-casheer-be* :x:. There was an issue during the deployment process.'
        }
    }
}