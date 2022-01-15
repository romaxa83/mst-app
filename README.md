# MST-APP

### microservice-test-app
The project is experimental and consists of several services written in golang.

The goal of the project is to study golang and related technologies.
and approaches to building a web application and services on golang, 
as well as the interaction between them

<div align="center"><a name="menu"></a>
  <h4>
    <a href="https://github.com/romaxa83/mst-app/tree/master/library-app#library-top">
      Library
    </a>
    <span> | </span>
    <a href="https://github.com/romaxa83/mst-app/tree/master/gin-app#todo-top">
      User/Todo
    </a>

  </h4>
</div>

### Local deploy project
🚀🚀 up services such as a database, file storage, etc. , for correct working
our webservices

```sh
$ make up
```

👀 show additional services

```sh
$ make show
```

in the folder <i>storage/data</i> there are files of different formats for importing data

#### 👨‍💻 Full list what has been used:
[Kafka](https://github.com/segmentio/kafka-go) as messages broker<br/>
[gRPC](https://github.com/grpc/grpc-go) Go implementation of gRPC<br/>
[PostgreSQL](https://github.com/jackc/pgx) as database<br/>
[Jaeger](https://www.jaegertracing.io/) open source, end-to-end distributed [tracing](https://opentracing.io/)<br/>
[Prometheus](https://prometheus.io/) monitoring and alerting<br/>
[Grafana](https://grafana.com/) for to compose observability dashboards with everything from Prometheus<br/>
[MongoDB](https://github.com/mongodb/mongo-go-driver) Web and API based SMTP testing<br/>
[Redis](https://github.com/go-redis/redis) Type-safe Redis client for Golang<br/>
[swag](https://github.com/swaggo/swag) Swagger for Go<br/>
[Echo](https://github.com/labstack/echo) web framework<br/>