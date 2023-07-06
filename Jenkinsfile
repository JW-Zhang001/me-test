pipeline {
  agent any
  stages {
  environment {
        COSMOS_SDK_PATH = "${WORKSPACE}/../cosmos-sdk-0.46.0"
  }
    stage('Copy Code to Cosmos SDK Directory') {
      steps {
          sh "mkdir -p ${COSMOS_SDK_PATH}"
      }
    }

    stage('Checkout Cosmos SDK') {
      steps {
          // 使用 checkout 步骤拉取第一个项目并检出指定的分支或提交记录
          checkout([
              $class: 'GitSCM',
              branches: [[name: 'main']],
              doGenerateSubmoduleConfigurations: false,
              extensions: [],
              submoduleCfg: [],
              userRemoteConfigs: [[
                url: 'git@github.com:stchain2022/cosmos-sdk-0.46.0.git',
                credentialsId: '2432580e-3beb-42f7-ab8a-9859617d43f1'
                ]],
              sh "mv . ${COSMOS_SDK_PATH}"
          ])
      }
    }
    
    
    // stage('Build me-chain') {
    //   steps {
    //     sh 'cd me-chain && make clean && make build'
    //   }
    // }

    // stage('Copy me-chaind to qa-home') {
    //   steps {
    //     sh 'cp ./me-chain/build/me-chaind /home/xingdao/qa-home/roles/me-chain/files'
    //   }
    // }

    // stage('Deploy') {
    //   steps {
    //       script {
    //         if (params.ENVIRONMENT == 'develop') {
    //           sh 'cd /home/xingdao/qa-home && ansible-playbook roles/me-chain/tests/test.yml -i roles/me-chain/tests/develop/inventory --tags "$ansible_tags"'
    //         } else if (params.ENVIRONMENT == 'alpha-test') {
    //           sh 'cd /home/xingdao/qa-home && ansible-playbook roles/me-chain/tests/test.yml -i roles/me-chain/tests/alpha/inventory --tags "$ansible_tags"'
    //         } else if (params.ENVIRONMENT == 'beta-test') {
    //           sh 'cd /home/xingdao/qa-home && ansible-playbook roles/me-chain/tests/test.yml -i roles/me-chain/tests/beta/inventory --tags "$ansible_tags"'
    //         } else {
    //           error('Invalid environment specified!')
    //         }
    //       }
    //   }
    //   post {
    //     failure {
    //       echo 'Deployment failed!'
    //     }
    //   }
    // }
  }
}
