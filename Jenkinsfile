pipeline {
   agent any

   tools {
    go { 'go-1.14' }
   }

   stages {
      stage('Test') {
         steps {
           sh 'go test'
         }
      }
      stage('Build') {
         steps {
           sh 'go build -o example1'
         }
      }
      stage('Publish artifact') {
         steps {
           archiveArtifacts 'example1'
         }
      }
   }
}
