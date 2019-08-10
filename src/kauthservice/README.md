# starter

This is your empty project. Ensure you have [GraalVM](https://www.graalvm.org) installed
on your path or use the provided `Dockerfile` to build your image.

**WARNING**: If you need to add more verticles to your application (so it can run them using
the standard `java -jar ... run your.other.Verticle`) you need to list it on:

[src/main/resources/META-INF/native-image/com.starter/starter/reflection.json](src/main/resources/META-INF/native-image/com.starter/starter/reflection.json)


## Build

`mvn package`

or

`docker build -t kauthservice .`

## Run

`./target/kauthservice`

or

`docker run --rm -it --net=host kauthservice`
