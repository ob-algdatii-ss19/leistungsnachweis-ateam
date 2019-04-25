pipeline {
    agent none
    stages {
        stage('Test') {
            agent {
                docker { image 'obraun/vss-jenkins' }
            }
            steps {
                sh 'go get github.com/t-yuki/gocover-cobertura' // install Code Coverage Tool
                sh 'go test ./... -v -coverpkg=./... -coverprofile=cover.out’ // save coverage info to file
                sh 'gocover-cobertura < cover.out > coverage.xml’ // transform coverage info to jenkins readable format
                sh 'go test ./... -bench=.'
                publishCoverage adapters: [coberturaAdapter('coverage.xml’)] // publish report on Jenkins
            }
        }
        stage('Lint') {
            agent {
                docker { image 'obraun/vss-jenkins' }
            }   
            steps {
                sh 'golangci-lint run --disable-all -E errcheck'
            }
        }
        stage('Build Docker Image') {
            agent {
                label 'master'
            }
            steps {
                sh "docker-build-and-push -b ${BRANCH_NAME}"
            }
        }
    }
    post {
        changed {
            script {
                if (currentBuild.currentResult == 'FAILURE') { // Other values: SUCCESS, UNSTABLE
                    // Send an email only if the build status has changed from green/unstable to red
                    emailext subject: '$DEFAULT_SUBJECT',
                        body: '$DEFAULT_CONTENT',
                        recipientProviders: [
                            [$class: 'DevelopersRecipientProvider']
                        ], 
                        replyTo: '$DEFAULT_REPLYTO'
                }
            }
        }
    }
}
