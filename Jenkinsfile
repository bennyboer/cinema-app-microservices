pipeline {
    agent none
    stages {
        stage('Build') {
            agent {
                docker { image 'obraun/vss-protoactor-jenkins' }
            }
            steps {
                sh '''
                    go version
                    go env
                    cd ci/install/protoc
                    chmod +x ./install.sh
                    . ./install.sh
                    cd ../../..
                    chmod +x ./build.sh
                    . ./build.sh
                '''
            }
        }
        stage('Test') {
            agent {
                docker { image 'obraun/vss-protoactor-jenkins' }
            }
            steps {
                sh 'echo run tests with code coverage...'
                sh 'go test ./... -cover'
            }
        }
        stage('Lint') {
            agent {
                docker { image 'obraun/vss-protoactor-jenkins' }
            }   
            steps {
                sh 'golangci-lint run --deadline 20m --enable-all --disable=goimports --disable=lll --disable=dupl --tests=false'
            }
        }
        stage('Build Docker Image') {
            agent any
            steps {
                sh "docker-build-and-push -b ${BRANCH_NAME} -s user-service -f user/service.dockerfile"
                sh "docker-build-and-push -b ${BRANCH_NAME} -s movie-service -f movie/service.dockerfile"
                sh "docker-build-and-push -b ${BRANCH_NAME} -s presentation-service -f presentation/service.dockerfile"
                sh "docker-build-and-push -b ${BRANCH_NAME} -s reservation-service -f reservation/service.dockerfile"
            }
        }
    }
}
