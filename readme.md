Follow the instructions to run this project

- run this command to build docker:
    > docker compose build
- up the docker
    > docker compose up -d


See Logs:
- db logs:
    > docker logs go-todo-db
- app logs:
    > docker logs go-todo-app --follow