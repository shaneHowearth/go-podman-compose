## Goal
The goal of this project is to have an application that reads all versions of
docker-compose.yml files and send commands to [podman](https://podman.io) such that podman will build containers and associated resources as [docker](https://www.docker.com/) will when instructed by [docker-compose](https://docs.docker.com/compose/).

Initial work based on [podman-compose](https://github.com/containers/podman-compose).
