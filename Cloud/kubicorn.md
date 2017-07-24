# Cloud Native Infrastructure with Kubicorn

### Abstract

Kubicorn is an open source Go project that is aimed an solving the Kubernetes infrastructure problem.
The day after the project was open-sourced it had already climbed to the number 1 fastest growing Go project on GitHub.
This talk will dicuss the reasoning behind the project and how the core of the tool is developed strictly for cloud native application.
We will clearly identify the problem space with existing infrastructure tooling and discuss how Kubicorn satisfies new patterns that can easily be vendored into control loops and operators.

Managing infrastructure via a cloud native application is a ground breaking idea and the project is a shining example of how we can begin to reason about infrastructure in this exciting new cloud native world we are living in. We explore the bootstrap problem of needing infrastructure in place to run the infrastructure management application and look at the importance of atomic infrastructure changes.

We will learn how the tool is a lovely implementation of the infrastructure reconciler pattern defined in my book "Cloud Native Infrastructure" and learn the dangers of managing infrastructure in other ways.
The lessons discussed are battle tested and have been proven over time. The audience will gain a rich understanding of what it means to run a cloud native application that manages underlying cloud native infrastructure through concrete examples from the Kubicorn project.