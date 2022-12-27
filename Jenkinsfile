pipeline {
    agent {label 'k8s-jenkins-slave'}
    stages {
        stage('环境变量') {
            steps {
                // 在script中给环境变量赋值
                script {
                    sh 'git log --oneline -n 1 > gitlog.file'
                    env.GIT_LOG = readFile("gitlog.file").trim()
                }
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
                    // 由于使用了切割字符串的语法，所以要使用bash执行
                    sh '''#!/bin/bash
                        docker build -t cicd_demo:${GIT_COMMIT:0:5} .
                    '''
                    echo '构建镜像成功，信息：' $GIT_LOG
                }
            }
        }

        stage('更新pod') {
            steps {
                // 步骤的超时时间
                timeout(time: 1, unit: 'MINUTES') {
                    echo '更新pod'
                }
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