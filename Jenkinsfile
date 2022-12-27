pipeline {
    agent {label 'k8s-jenkins-slave'}
    stages {
        stage('环境变量') {
            steps {
                sh 'printenv'
            }
        }

        stage('拉取源码') {
            steps {
                // 直接从Jenkinsfile所在的scm拉取代码
                checkout scm
            }
        }

        stage('构建镜像') {
            steps {
                // 如果失败，重试两次
                retry (2) {
                    sh 'docker build -t cicd_demo:${GIT_COMMIT:0:5} .'
                    echo '构建镜像成功'
                }
            }
        }

        stage('更新pod') {
            steps {
                echo '更新pod'
            }
        }
    }

    // 后处理
    post {
        success {
            echo 'success'
        }
        failure {
            echo 'failure'
        }
        always {
            echo 'always'
        }
    }
}