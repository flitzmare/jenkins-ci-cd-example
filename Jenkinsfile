pipeline {
  environment {
    registry = "riksasuviana/hello-world-api"
    registryCredential = $REGISTRY_CREDENTIAL
    majorVersion = '0'
    minorVersion = '1'
    dockerImage = ''
  }
  agent any
  stages {
    stage('Cloning Git') {
      steps {
        git branch: "${env.BRANCH_NAME}", credentialsId: $GIT_CREDENTIAL, url: 'https://github.com/flitzmare/jenkins-ci-cd-example.git'
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
                if (env.BRANCH_NAME == 'master') {
                    dockerImage = docker.build registry + ":$majorVersion.$minorVersion.$BUILD_NUMBER"+" -e VERSION="+$majorVersion.$minorVersion.$BUILD_NUMBER
                } else if (env.BRANCH_NAME == 'staging'){
                    dockerImage = docker.build registry + ":staging-$majorVersion.$minorVersion.$BUILD_NUMBER"+" -e VERSION=staging-"+$majorVersion.$minorVersion.$BUILD_NUMBER
                }
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
    stage('Running docker container') {
        steps{
            script {
                if (env.BRANCH_NAME == 'master') {
                    withEnv(["HELLO_WORLD_API_VERSION=$majorVersion.$minorVersion.$BUILD_NUMBER"]) {
                        sh "echo VERSION: $HELLO_WORLD_API_VERSION"
                        sh "docker-compose -f $DOCKER_COMPOSE_LOCATION up -d"
                    }
                } else if (env.BRANCH_NAME == 'staging'){
                    withEnv(["STAGING_HELLO_WORLD_API_VERSION=staging-$majorVersion.$minorVersion.$BUILD_NUMBER"]) {
                        sh "echo STAGING-VERSION: $STAGING_HELLO_WORLD_API_VERSION"
                        sh "docker-compose -f $DOCKER_COMPOSE_STAGING_LOCATION up -d"
                    }
                }
            }
        }
    }
  }
}