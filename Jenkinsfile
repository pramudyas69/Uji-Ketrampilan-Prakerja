pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build') {
            steps {
                sh 'go build -o myapp'
            }
        }

        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }

        stage('Docker Compose') {
            steps {
                script {
                    // Stop and remove existing containers defined in docker-compose.yml
                    sh 'docker-compose down'

                    // Build and start services defined in docker-compose.yml
                    sh 'docker-compose up -d'
                }
            }
        }
    }

    post {
        always {
            // Stop and remove containers after the pipeline finishes, regardless of the result
            sh 'docker-compose down'
        }
    }
}
