pipeline {
    agent none
    stages {
        stage('Build') {
            agent {
                docker {
                    image 'obraun/vss-protoactor-jenkins'
                    args '-v $HOME/.cache/go-build:$HOME/gopath/pkg/mod'
                }
            }
            steps {
                sh '''
                    git config user.name "bennyboer-machine-user"
                    git config credential.helper store
                    echo https://51faa31d4b9f08c8e56d4fb23fc082a85e617df8:x-oauth-basic@github.com >> ~/.git-credentials
                '''
                sh '''
                    cat /etc/os-release
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
                docker {
                    image 'obraun/vss-protoactor-jenkins'
                    args '-v $HOME/.cache/go-build:$HOME/gopath/pkg/mod'
                }
            }
            steps {
                sh 'echo run tests with code coverage...'
                sh 'chmod +x ./test.sh'
                sh '. ./test.sh'
            }
        }
        stage('Lint') {
            agent {
                docker {
                    image 'obraun/vss-protoactor-jenkins'
                    args '-v $HOME/.cache/go-build:$HOME/gopath/pkg/mod'
                }
            }   
            steps {
                sh 'echo run tests with code coverage...'
                sh 'chmod +x ./lint.sh'
                sh '. ./lint.sh'
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
