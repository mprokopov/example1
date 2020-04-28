pipeline {
   agent any

   tools {
    go { 'go-1.14' }
   }

   stages {
      stage('Build') {
         steps {
           sh 'go build'
         }
      }
   }
}
