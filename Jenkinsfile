def output

pipeline {
    
   agent any

   stages {
      stage('Hello') {
         steps {
             script {
                 output = sh label: '', returnStdout: true, script: 'echo "Hello, World"'
             }
         }
      }
      stage('Step two') {
         steps {
            echo output
         }
      }
   }
}
