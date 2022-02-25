k8s_yaml(['./deploy/postgresql/configmap.yml', './deploy/postgresql/persistent.yml', './deploy/app/config.yml'])


docker_build('db-postgres',
             context='.',
             dockerfile='./deploy/postgresql/postgres.Dockerfile',
)

k8s_yaml(['./deploy/postgresql/service.yml', './deploy/app/keycloak.yml'])

docker_build('my-app',
             context='.',
             dockerfile='./deploy/app/app.Dockerfile',
             only=['./'], 
             live_update=[
                sync('./app', '/src/'),
             ]
)

k8s_resource('db-migrate', 
   resource_deps=['postgres'],
   trigger_mode=TRIGGER_MODE_MANUAL,
   auto_init=False
)

k8s_yaml(['deploy/app/api.yml'])

local_resource('All pods',
   'kubectl get pods',
   resource_deps=['postgres', 'api']
)
