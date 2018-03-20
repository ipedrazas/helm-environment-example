# Managing environments with Helm


Helm allows to define an application chart with dependencies. This means that we can create a chart that contains only dependencies and identifies an environment.

In this repo we will have 3 applications:

* A microservice written in Go.
* A microservice written in Python.
* A web frontend using vue.js.

Each application contains a chart that allows you to deploy that application in a kubernetes cluster. We wil define a set of environments as charts: __DEV__,  __UAT__, __PROD__

Each environment contains a set of configurations that allow you to deploy these applications in certain state.