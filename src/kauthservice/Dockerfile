# GraalVM docker image used for AoT compilation
FROM oracle/graalvm-ce:19.1.1 AS build-aot
# Add maven wrapper
ADD mvnw app/
ADD .mvn app/.mvn/
# Add pom
ADD pom.xml app/
# Add sources
ADD src app/src/
# Set working dir
WORKDIR /app
# Build (java side)
RUN ./mvnw package
# Create new image from debian-slim (20Mb)
FROM debian:stable-slim
# Copy generated native executable from build-aot
COPY --from=build-aot /app/target/kauthervice /kauthervice
COPY --from=build-aot /opt/graalvm-ce-19.1.1/jre/lib/amd64/libsunec.so /libsunec.so
# Set the entrypoint
EXPOSE 50051, 8080
ENTRYPOINT [ "./kauthervice" ]
