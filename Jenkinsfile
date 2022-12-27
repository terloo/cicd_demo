pipeline {
    agent {label 'k8s-jenkins-slave'}
    stages {
        stage('拉取源码') {
            steps {
                git credentialsId: 'git_pk', url: 'git@github.com:terloo/cicd_demo.git'
            }
        }

        stage('构建镜像') {
            steps {
                sh 'docker build -t cicd_demo:${GIT_COMMIT} .'
                echo '构建镜像成功'
            }
        }

        stage('更新pod') {
            steps {
                echo '更新pod'
            }
        }
    }
}