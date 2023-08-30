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

        stage('Docker Build') {
            steps {
                sh 'docker build -t myapp .'
            }
        }

        stage('Docker Deploy') {
            steps {
                sh 'docker run -d -p 8080:8080 myapp'
            }
        }
    }
}
