pipeline {
  environment {
    registry = "riksasuviana/hello-world-api"
    registryCredential = 'a7ca3b0f-ae79-4834-a148-cfc1c82e16aa'
//     dockerImage = ''
  }
  agent any
  stages {
    stage('Cloning Git') {
        when {
          branch 'master'
        }
      steps {
                  sh "echo Branch: ${env.BRANCH_NAME}"
                                    sh "BRANCH MASTER"
        git branch: 'master', credentialsId: 'cc63d301-5150-4b74-bf60-dd572bba744b', url: 'https://github.com/flitzmare/jenkins-ci-cd-example.git'
      }
        when {
          branch 'staging'
        }
      steps {
                  sh "echo Branch: ${env.BRANCH_NAME}"
                  sh "BRANCH STAGING"
        git branch: 'staging', credentialsId: 'cc63d301-5150-4b74-bf60-dd572bba744b', url: 'https://github.com/flitzmare/jenkins-ci-cd-example.git'
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
        sh "docker image build -t $registry:$BUILD_NUMBER ."
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
            withEnv(["HELLO_WORLD_API_VERSION=$BUILD_NUMBER"]) {
                sh "docker-compose -f /Users/riksasuviana/docker-compose.yml up -d"
            }
        }
    }
  }
}