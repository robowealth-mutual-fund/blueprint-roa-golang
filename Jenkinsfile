#!/usr/bin/env groovy
@Library('jenkins-shared-library@3.3.2') _

properties([
    gitLabConnection('robowealth@gitlab')
])

stdPipeline([
    projectName: 'github.com/robowealth-mutual-fund/blueprint-roa-golang',
    pipelineType: 'golang',
    gcsBucket: 'robowealth-manifest-artifact-store/roa',
    utility: [
            gitlabPlugin: true,
            notify: [
                notifyStatus: ["FAILURE"],
                discordPlugin: [
                    webhookURL: "https://discordapp.com/api/webhooks/710772184032870472/QPyWPs2foamRYcuT2VpW6b0QkRoBmz_fF-Jc_UduaS6uNzclF2rCutXuE7MTnr5egaMy"
                ]
            ]
    ],
    jenkinsConfig: [
            agent : 'roa-nonprod'
    ],
    registry: [
            url: 'asia.gcr.io',
            group: 'cloud-management-robowealth',
            subgroup: 'roa',
            project: 'github.com/robowealth-mutual-fund/blueprint-roa-golang',
    ],
    branchToEnvironmentMapping: [
            'production/': ['uat', 'prod'],
            'release/': ['sit'],
            'develop': ['develop'],
            'feature/': ['develop'],
            '*': ['review'],
    ],
    branchEnableManualDeploy: ['feature/'],
    // branchToCreateArtifactOnly: ['production/'],
    environmentToDeploy: ['sit', 'uat', 'develop'],
    kubernetes: [
            helmConfig: [
                repositoryName: 'private',
                repositoryUrl: 'gs://robowealth-helm-chart-repository',
                chart: 'common-services:1.4.0'
            ],
            environmentToClusterMapping: [
                    'uat': [
                        'name': 'roa-uat-cluster',
                        'region': 'asia-southeast1',
                        'namespace': 'uat'
                    ],
                    'sit': [
                        'name': 'roa-nonprod-cluster',
                        'region': 'asia-southeast1',
                        'namespace': 'sit'
                    ],
                    'develop': [
                        'name': 'roa-nonprod-cluster',
                        'region': 'asia-southeast1',
                        'namespace': 'develop'
                    ],
            ],
    ],
]) { context ->
    stageBuild(context)
    stageBuildAndPushToRegistry(context)
    stageBuildHelmManifest(context)
    stageDeployHelm(context)
    stageArchiveArtifactGCS(context)
}
