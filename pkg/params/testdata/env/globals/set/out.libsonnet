local params = std.extVar("__ksonnet/params");
local globals = import "globals.libsonnet";
local envParams = params + {
  components+: {
    guestbook+: {
      name: "guestbook-dev",
      replicas: params.global.replicas,
      containerPort: 8080
    }
  }
};

{
  components: {
    [x]: envParams.components[x] + globals for x in std.objectFields(envParams.components)
  }
}