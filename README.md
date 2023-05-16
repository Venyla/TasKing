# Welcome to TasKing
This repository contains the source code of the TasKing application created as part of the challenge  task for the module Distributed Systems (DSy FS2023) at the Eastern Switzerland University of Applied Sciences.

TasKing is a simple application that allows users to complete tasks on the campus of Rapperswil-Jona. The person who has completed a task the most is nominated "TasKing", indicated by the presence of a crown above the task. As required by the assignment, the application consists of a frontend, a backend and a load balancer.

TODO: Insert Screenshot


## Infrastructure
The following diagram explains the infrastructure of our application. Details about the different components (frontend, backend and load balancer) can be found in the sections below.

![Infrastructure overview](./doc/TasKing.drawio.png)

### Frontend
**Technologies:** HTML, TS, SCSS

The frontend consists of a single `index.html` that communicates with the backend using a Single Page Application (SPA) architecture.
Since no framework like Vue, React, etc. was used, all logic is located in the `scripts.ts` file, which retrieves and updates each task using an ugly (but for this challenge task sufficient) polling mechanism. 
Both the `scripts.ts` and `default-theme.scss` are compiled/transpiled using npm, both locally and within the Dockerfile.


### Backend


### Load Balancer

## Usage
### Running Locally

```
# Running
docker compose up --build

# Tear Down
docker compose down --volumes
```

TODO: Insert localhost link/port to website

## API

**Return all task:**

GET: `/api/tasks`

**Return task by id:**

GET: `/api/tasks/{task-id}`

**Return rankings of a task:**

GET: `/api/tasks/{task-id}/rankings`

**Return all task histories:**

GET: `/api/history/{task-id}`

**Save task:**

POST: `/api/tasks`

**Save history entry:**

POST: `/api/history`

## Author
👤 Vina Zahnd, Vanessa Gyger, Lukas Messmer