pipeline {
    agent {label 'k8s-jenkins-slave'}
    stages {
        stage('拉取源码') {
            git credentialsId: 'git_pk', url: 'git@github.com:terloo/cicd_demo.git'
        }

        stage('构建镜像') {
            sh 'docker build -t cicd_demo:${GIT_COMMIT} .'
            echo '构建镜像成功'
        }

        stage('更新pod') {
            echo '更新pod'
        }
    }
}