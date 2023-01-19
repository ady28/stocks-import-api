pipeline { 
    agent {
        node {
            label 'master'
        }
    }
    options {
        ansiColor('xterm')
    }
    environment { 
        DOCKERHUB_USER = 'ady28'
        IMAGE_NAME = 'stocks-import'
        IMAGE_VERSION = ''
    }
    stages {
        stage('Static code analysis') {
            steps {
                sh 'staticcheck ./...'
                sh 'gosec -exclude=G104 -fmt=yaml -out=security_code_analysis.yaml -stdout -verbose=text ./...'
            }
        }
        stage('Dockerfile check') {
            steps {
                sh "hadolint Dockerfile --info DL3008 -t warning -f json > dockerfile_lint.json"
            }
        }
        stage('Build') { 
            steps { 
                echo "Running build stage for ${env.IMAGE_NAME}"
                script {
                    IMAGE_VERSION = readFile('VERSION').trim()
                    image = docker.build("${env.DOCKERHUB_USER}/${env.IMAGE_NAME}:${IMAGE_VERSION}")
                    docker.withRegistry("https://registry.hub.docker.com", 'dockerhub-pub') {
                        image.push()
                    }
                }
            }
        }
        stage('Generate SBOM file') {
            steps {
                sh "docker sbom ${env.DOCKERHUB_USER}/${env.IMAGE_NAME}:${IMAGE_VERSION} --format syft-json -o sbom.json"
            }
        }
        stage('Security analysis') {
            steps {
                sh "trivy image ${env.DOCKERHUB_USER}/${env.IMAGE_NAME}:${IMAGE_VERSION} -f json -o trivycheck.json"
            }
        }
        stage('Tag Test Image'){
            steps {
                echo "Push image for testing"
                script {
                    docker.withRegistry("https://registry.hub.docker.com", 'dockerhub-pub') {
                        image.push("test")
                    }
                }
            }
        }
        stage('Deploy Test Infrastructure') {
            steps {
                build job: 'stocks-app-test-deploy'
            }
        }
        stage('Ask to go in qual') {
            steps {
                input('Do you want to proceed to the qual environment?')
            }
        }
        stage('Tag Qual Image'){
            steps {
                echo "Push image for qual"
                script {
                    docker.withRegistry("https://registry.hub.docker.com", 'dockerhub-pub') {
                        image.push("qual")
                    }
                }
            }
        }
        stage('Deploy Qual Infrastructure') {
            steps {
                build job: 'stocks-app-qual-deploy'
            }
        }
    }
    post {
        always {
            archiveArtifacts artifacts: 'sbom.json, dockerfile_lint.json,security_code_analysis.yaml,trivycheck.json', fingerprint: true
        }
    }
}