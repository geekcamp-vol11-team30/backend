{
  capacityProviderStrategy: [
    {
      base: 1,
      capacityProvider: 'FARGATE_SPOT',
      weight: 1,
    },
    {
      base: 0,
      capacityProvider: 'FARGATE',
      weight: 0,
    },
  ],
  deploymentConfiguration: {
    deploymentCircuitBreaker: {
      enable: true,
      rollback: true,
    },
    maximumPercent: 200,
    minimumHealthyPercent: 100,
  },
  deploymentController: {
    type: 'ECS',
  },
  desiredCount: 1,
  enableECSManagedTags: false,
  enableExecuteCommand: false,
  healthCheckGracePeriodSeconds: 0,
  launchType: '',
  loadBalancers: [
    {
      containerName: '{{ tfstate `module.base.aws_ecs_service.api.load_balancer[0].container_name` }}',
      containerPort: 8080,
      targetGroupArn: '{{ tfstate `module.base.aws_ecs_service.api.load_balancer[0].target_group_arn` }}',
    },
  ],
  networkConfiguration: {
    awsvpcConfiguration: {
      assignPublicIp: 'ENABLED',
      securityGroups: [
        '{{ tfstate `module.base.aws_ecs_service.api.network_configuration[0].security_groups[0]` }}',
      ],
      subnets: [
        '{{ tfstate `module.base.aws_ecs_service.api.network_configuration[0].subnets[0]` }}',
        '{{ tfstate `module.base.aws_ecs_service.api.network_configuration[0].subnets[1]` }}',
        '{{ tfstate `module.base.aws_ecs_service.api.network_configuration[0].subnets[2]` }}',
      ],
    },
  },
  platformFamily: 'Linux',
  // platformVersion: '1.4.0',
  platformVersion: '{{ tfstate `module.base.aws_ecs_service.api.platform_version` }}',
  propagateTags: 'NONE',
  schedulingStrategy: 'REPLICA',
  tags: [
    {
      key: 'Env',
      value: 'dev',
    },
    {
      key: 'Service',
      value: 'api',
    },
    {
      key: 'Name',
      value: 'magische-dev-api-ecs-service',
    },
  ],
}
