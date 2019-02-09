## squash deploy

deploy the squash agent or a demo microservice

### Synopsis

deploy the squash agent or a demo microservice

### Options

```
  -h, --help   help for deploy
```

### Options inherited from parent commands

```
      --container string           Container to debug
      --container-repo string      debug container repo to use (default "soloio")
      --container-version string   debug container version to use (default "v0.1.9")
      --crisock string             The path to the CRI socket (default "/var/run/dockershim.sock")
      --debug-server               start a debug server instead of an interactive session
      --debugger string            Debugger to use (default "dlv")
      --json                       output json format
      --lite                       run in lite mode (default) (default true)
      --localport int              port to use to connect to debugger (defaults to 1235)
      --machine                    machine mode input and output
      --namespace string           Namespace to debug
      --no-clean                   don't clean temporar pod when existing
      --no-guess-debugger          don't auto detect debugger to use
      --no-guess-pod               don't auto detect pod to use
      --pod string                 Pod to debug
      --timeout int                timeout in seconds to wait for debug pod to be ready (default 300)
```

### SEE ALSO

* [squash](squash.md)	 - debug microservices with squash
* [squash deploy agent](squash_deploy_agent.md)	 - deploy a squash agent
* [squash deploy demo](squash_deploy_demo.md)	 - deploy a demo microservice
