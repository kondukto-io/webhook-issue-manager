# Create Postgresql DB on Docker

### Compose

**Note**: This yaml inside the project already

```yml
version: "3.8"

services:
  database:
    container_name: postgres-database
    image: postgres
    restart: always
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=12345
      - POSTGRES_DB=issue
    ports:
      - 5432:5432
    volumes:
      - $HOME/docker/volumes/postgres:/var/lib/postgresql/data
```
