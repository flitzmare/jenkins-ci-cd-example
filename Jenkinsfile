pipeline {
  environment {
    registry = "riksasuviana/hello-world-api"
    registryCredential = 'a7ca3b0f-ae79-4834-a148-cfc1c82e16aa'
    majorVersion = '0'
    minorVersion = '1'

//     dockerImage = ''
  }
  agent any
  stages {
    stage('Cloning Git') {
      steps {
        git branch: "${env.BRANCH_NAME}", credentialsId: 'cc63d301-5150-4b74-bf60-dd572bba744b', url: 'https://github.com/flitzmare/jenkins-ci-cd-example.git'
      }
    }
    stage('Test') {
        steps {
            sh "ls"
            sh 'go test ./...'
        }
    }
    stage('Building image') {
      steps{
        script {
                if (env.BRANCH_NAME == 'master') {
                    sh "docker image build -t $registry:$majorVersion.$minorVersion.$BUILD_NUMBER ."
                } else if (env.BRANCH_NAME == 'staging'){
                    sh "docker image build -t $registry:staging-$majorVersion.$minorVersion.$BUILD_NUMBER ."
                }
        }
//         script {
//           dockerImage = docker.build registry + ":$BUILD_NUMBER"
//         }
      }
    }
//     stage('Deploy Image') {
//       steps{
//         script {
//           docker.withRegistry( '', registryCredential ) {
//             dockerImage.push()
//           }
//         }
//       }
//     }
//     stage('Remove Unused docker image') {
//       steps{
//         sh "docker rmi $registry:$BUILD_NUMBER"
//       }
//     }
    stage('Running docker container') {
        steps{
            // sh "docker run --rm -itd --name run-hello-world $registry:$BUILD_NUMBER"
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