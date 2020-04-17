#!groovy

@Library('build-golang-image@master') _

buildGolangDocker(
  branch: 'master',
  gitCredentialsId: 'Github',
  gitUrl: 'https://github.com/naijab/demo_microservice_post_service',
  dockerRegistry: 'naijabcom/demo-ms-post-service'
)