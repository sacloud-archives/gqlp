# sacloud/gqlp

`sacloud/gqlp` provides a proxy that enables GraphQL to be used with SAKURA Cloud API.

## Status

`sacloud/gqlp` is still in a very early development stage.


## Usage

```shell
$ go run github.com/sacloud/gqlp/cmd/gqlp
```

## Query Example

```gql
query foobar {
  servers(zone: "is1b") {
    id
    name
    tags
    description
    availability
    hostName
    interfaceDriver
    planID
    planName
    cpu
    memory
    commitment
    planGeneration
    instanceHostName
    instanceStatus
    disks{
      id
      name
      description
      tags
    }
  }
}
```

```
# result
{
  "data": {
    "servers": [
      {
        "id": "100000000001",
        "name": "example",
        "tags": [
          "stage=development"
        ],
        "description": "example server for GraphQL",
        "availability": "available",
        "hostName": "www.example.com",
        "interfaceDriver": "virtio",
        "planID": "4002",
        "planName": "プラン/2Core-4GB",
        "cpu": 2,
        "memory": 4,
        "commitment": "standard",
        "planGeneration": 100,
        "instanceHostName": "sac-is1b-sv999",
        "instanceStatus": "up",
        "disks": [
          {
            "id": "200000000001",
            "name": "example",
            "description": "",
            "tags": [
              "stage=development"
            ]
          }
        ]
      }
    ]
  }
}
```

## License

`sacloud/gqlp` Copyright (C) 2021 The gqlp Authors.

This project is published under [Apache 2.0 License](LICENSE).
