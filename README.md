# Distributed System

[picture]

The purpose of this project was to build a scalable application using microservice architecture.

### What is microservice architecture?

Microservices architecture is a method of building large applications out of smaller, modular services or components. A given service can handle functionality like user authentication, alert generation, data logging, and a number of other functions. Each service handles its own database and typically runs a distinct process.

### What are some advantages of using microservice architecture?

Microservices have a wide variety of benefits, such as:

- **Flexibility in Development and Deployment**: Microservices are language agnostic, and can be written and deployed using different programming languages. This flexibility allows developers to choose the most suitable options for specific project requirements.

- **Rapid Scalability**: As the demand for a given service grows, it can be easily expanded by deploying additional instances of it. This elasticity ensures that the application can seamlessly accommodate increased user traffic and workloads, leading to fewer performance bottlenecks.

- **Reusability Across Projects**: Microservices foster reusability by enabling components to be extracted and shared among different projects. This modular approach reduces redundancy and accelerates development cycles, as developers can leverage existing microservices rather than reinventing the wheel for each new project.

- **Enhanced Fault Isolation**: Each microservice operates independently, encapsulating its own functionality and resources. This isolation prevents issues within one microservice from cascading across the entire system, enhancing overall system stability and resilience.

- **Suitability for Small Teams**: Each microservice can be developed, tested, and maintained by a relatively smaller set of engineers, minimizing communication overhead and allowing for a more agile and efficient development process.

- **Compatibility with Containerization**: Containers encapsulate individual microservices and their dependencies, creating a consistent and isolated environment for their execution. This compatibility facilitates efficient deployment, scaling, and management of microservices in diverse computing environments.

### How to run

1. Run `docker-compose up -d` in the root directory
2. `cd` into the `client` directory
3. Run `npm install`, followed by `npm run dev`

### Technologies

Built using:

- [React.js](https://react.dev/)
- [TypeScript](https://www.typescriptlang.org/)
- [Go](https://go.dev/)
- [PostgreSQL](https://www.postgresql.org/)
- [MongoDB](https://www.mongodb.com/)
- [gRPC](https://grpc.io/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
