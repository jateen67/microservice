# Distributed System

[main gif]

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

### Services

This project is divided into 3 services + a frontend. They are all accessed through the Broker service, which acts as a centralized point of contact.

Below is a picture of the overall architecture of this application.

![architecture](./media/architecture.png)

**Broker Service**

This serves as a main point of entry into the microservice cluster. It is optional, but in this application, all requests to the various microservices from the client go through the Broker first, which then communicates with the respective services to return a response.

Since the Broker's main purpose is to communicate with the other services, sending a request to it directly will give a basic response back to the user, as a simple way of indicating that it is functioning well and ready to communicate.

![broker](./media/broker.gif)

**Logger Service**

The Logger's job is to log some kind of simulated activity, whatever that may be. When the user sends a request, it will insert some data into a Mongo database, indicating that the user has done some activity that's been successfully stored/logged, kind of like a traditional activity logger in any other application.

The Broker and Logger communicate with one another via gRPC.

The database containing the successful user activity logs can be accessed locally using a Mongo client like [MongoDBCompass](https://www.mongodb.com/products/compass) (Connection String: `mongodb://mongo:password@localhost:27017/logs_db?&ssl=false`)

![logger](./media/grpc.gif)

**Authentication Service**

The Authenticator simulates attempting to "sign in" a user given the proper credentials. When the user sends a request, a username and password is sent alongside it (admin@example.com/password123). It will take these credentials, which are hashed using the [bcrypt](https://en.wikipedia.org/wiki/Bcrypt) algorithm, and compare them to the credentials stored in a Postgres database to try to find a match. If there is a match, then a success message will be sent to the client.

The Broker and Authenticator communicate with one another via JSON.

The database containing the user credentials can be accessed locally using a database manager like [Beekeeper Studio](https://www.beekeeperstudio.io/) (Connection String: `host=localhost port=5432 user=postgres password=password dbname=users_db sslmode=disable timezone=UTC`)

![json authentication](./media/json-auth.gif)

**Listener Service**

The Listener is another way that the user can send a request to the Authenticator to sign in. It accomplishes the exact same things as the standard Authenticator, but through a different method.

When the user sends a request to the Authenticator via this alternative method, the Broker will not communicate with the it directly like normal, but instead with the Listener, which will then in turn communicate with the Authenticator.

This works by first pushing an event to a RabbitMQ server from the Broker via AMQP. RabbitMQ then takes that event and adds it to a queue. The Listener looks at that queue and constantly monitors it to see if there are any messages present that it should read. If so, it reads it, figures out what to do with it, and then calls the appropriate service to perform the action. In this case, it calls the Authenticator to attempt a signin.

All RabbitMQ activity can be viewed and monitored by accessing the [Management UI](https://www.rabbitmq.com/management.html) on `localhost:15672`

![rabbitmq authentication](./media/rabbitmq-auth.gif)

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
- [RabbitMQ](https://www.rabbitmq.com/)
- [Bootstrap](https://getbootstrap.com/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
