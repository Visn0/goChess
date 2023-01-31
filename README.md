# GoChess

## Description

The project consists of a chess game. With the intention of learning Domain Driven Design (DDD), a [Hexagonal architecture](<https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)>) has been used. So far, the game allows creating matches, connecting to them and playing the game, including calculation of valid moves, check detection and so on.

The project is developed using **Golang** and **[goFiber](https://gofiber.io/)** for the backend and **[Vue.js](https://vuejs.org/)** for frontend.

## Technologies summary

- **Frontend**: [Vue.js](https://vuejs.org/), HTML, CSS, [Bootstrap 5](https://getbootstrap.com/docs/5.0/getting-started/introduction/), WebSockets,
- **Backend**: [go1.18](https://go.dev/), Hexagonal Architecture, WebSockets.

<!-- HOW TO RUN -->

## How to run the project

### Prerequisites

<ul>
  <li><a href="https://nodejs.org/en/">Node.js</a> version 18.x.</li>
  <li><a href="https://go.dev/">Golang</a> version 1.18.</li>
</ul>

### Local deploy

In the root folder of the project:

```bash
# To install project dependencies (this may take a few minutes).
# Required only the first time.
$ make install_project
```

```bash
# To run the frontend
$ make up
```

In another shell:

```bash
# To run the backend
$ make up.backend
```

Once the server is up and running, open your browser and type `http://localhost:5173/` and you will see the chess.

<!-- AUTHORS -->

## Authors

<ul>
  <li>
    <p>
      <b>Antón Chernysh</b>:
        <ul>
          <li>
            LinkedIn: <a href="https://www.linkedin.com/in/anton-chernysh/">https://www.linkedin.com/in/anton-chernysh/</a>
          </li>
          <li>
            Email: <a href="mailto:anton_chernysh@outlook.es">anton_chernysh@outlook.es</a>
          </li>
        </ul>
    </p>
  </li>
    <li>
    <p>
      <b>Carlos Eduardo Arismendi Sánchez</b>:
        <ul>
          <li>
            LinkedIn: <a href="https://www.linkedin.com/in/carlos-arismendi/">https://www.linkedin.com/in/carlos-arismendi/</a>
          </li>
          <li>
            Email: <a href="mailto:carlos.arismendisanchez@gmail.com">carlos.arismendisanchez@gmail.com</a>
          </li>
        </ul>
    </p>
  </li>
</ul>
