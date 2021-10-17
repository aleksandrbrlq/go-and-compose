### Live-reloading command line utility https://github.com/cosmtrek/air for Go applications in development in docker.

Inspired by https://firehydrant.io/blog/develop-a-go-app-with-docker-compose/

We should run our Air init from inside of our app container (defined in our compose yaml file). To do this, let's run the follow command:

```docker-compose run --rm app air init```

Then run:

```docker-compose run```
