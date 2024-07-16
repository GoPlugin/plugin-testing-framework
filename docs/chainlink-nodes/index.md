---
layout: default
title: Plugin Nodes
nav_order: 4
has_children: false
---

# Plugin Nodes

Make sure to have your environment setup code in place, which will deploy a few Plugin nodes, and give you an `env` environment.

```go
// Get a list of all Plugin nodes deployed in your test environment
pluginNodes, err := client.ConnectPluginNodes(env)
```

From here, you can [interact with](https://pkg.go.dev/github.com/goplugin/plugin-testing-framework/client#Plugin) each Plugin node to manage keys, jobs, bridges, and transactions.

## Plugin Jobs

The most common interaction you'll have with Plugin nodes will likely be creating jobs, using the `pluginNode.CreateJob(JobSpec)` method. Plugin jobs are how the Plugin nodes know what actions they're expected to perform on chain, and how thy should perform them. A typical test consists of launching your resources, deploying contracts to the blockchain, and telling the Plugin node to interact with those contracts by creating a job. Read more about Plugin jobs and the specifics on using them [here](https://docs.chain.link/docs/jobs/).

There are plenty of built in [JobSpecs](https://pkg.go.dev/github.com/goplugin/plugin-testing-framework/client#JobSpec) like the [Keeper Job Spec](https://pkg.go.dev/github.com/goplugin/plugin-testing-framework/client#KeeperJobSpec) and the [OCR Job Spec](https://pkg.go.dev/github.com/goplugin/plugin-testing-framework/client#OCRTaskJobSpec) that you can use. But if for whatever reason, those don't do the job for you, you can create a raw TOML job with `CreateJobRaw(string)` like below.

```go
jobData, err := pluginNode.CreateJobRaw(`
schemaVersion = 1
otherField    = true
`)
```
