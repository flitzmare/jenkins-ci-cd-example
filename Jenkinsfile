pipeline {
  environment {
    registry = "riksasuviana/hello-world-api"
    registryCredential = 'a7ca3b0f-ae79-4834-a148-cfc1c82e16aa'
    dockerImage = ''
  }
  agent any
  stages {
    stage('Cloning Git') {
      steps {
        git 'https://github.com/flitzmare/jenkins-ci-cd-example.git'
      }
    }
    stage('Test') {
        steps {
            sh 'go test ./...'
        }
    }
    stage('Building image') {
      steps{
        script {
          dockerImage = docker.build registry + ":$BUILD_NUMBER"
        }
      }
    }
    stage('Deploy Image') {
      steps{
        script {
          docker.withRegistry( '', registryCredential ) {
            dockerImage.push()
          }
        }
      }
    }
    stage('Remove Unused docker image') {
      steps{
        sh "docker rmi $registry:$BUILD_NUMBER"
      }
    }
  }
}