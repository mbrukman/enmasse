{
  generate(image_name)::
  {
    "apiVersion": "v1",
    "kind": "ReplicationController",
    "metadata": {
      "labels": {
        "name": "subserv"
      },
      "name": "subserv"
    },
    "spec": {
      "replicas": 1,
      "selector": {
        "name": "subserv"
      },
      "template": {
        "metadata": {
          "labels": {
            "name": "subserv"
          }
        },
        "spec": {
          "containers": [
            {
              "image": image_name,
              "name": "master",
              "ports": [
                {
                  "containerPort": 5672,
                  "protocol": "TCP"
                }
              ]
            }
          ]
        }
      }
    }
  }
}
