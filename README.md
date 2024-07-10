# Bount.ing

Bount.ing is a gamified platform designed to incentivize and reward open source contributions. By integrating game mechanics, Bount.ing aims to make contributing to open source projects more engaging and rewarding.

## Technologies Used
Frontend: Vue 3 (Vite)
Backend: Golang
Containerization: Docker Compose

## Getting Started
These instructions will help you get a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites
Make sure you have the following installed on your system:

- `docker`
- `docker-compose`

### Download the repo

```
git clone git@github.com:Bount-ing/Bount.ing.git &&
cd Bount.ing
```

### Set up Env

Fill the files in `api/.env` and `front/.env`

### Run the docker

`sudo docker-compose build  && sudo docker-compose up -d && sudo docker-compose logs -f`
