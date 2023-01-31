# goChess
Next competitor of chess.com

## Description
The project consists of a clone of the chess game. With the intention of learning DDD, a [Hexagonal architecture](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)) has been used. So far, the game allows creating matches, connecting to them and playing the game, including calculation of valid moves, check detection and others.

The project is developed using **Golang** and **[goFiber](https://gofiber.io/)** for the backend and **[Vue.js](https://vuejs.org/)** for frontend.

## Technologies summary
- **Frontend**: [Vue.js](https://vuejs.org/), HTML, CSS
- **Backend**: Golang1.18, Hexagonal architecture, REST Api.

<!-- HOW TO RUN -->
## How to run the project
First of all, you need to install <a href="https://golang.org/">Golang1.17</a> following the official documentation. Then, move to the root directory of this project and run:  
```bash
$ go get github.com/labstack/echo/v4
```
Then, if you are in a Windows machine:
```shell
C:\user> runserver.bat
```
If you are in a Linux machine:
```bash
foo@bar:~$ go run *
```
Once the server is up and running, open your browser and type ```localhost:8080``` and you will see the chess.

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

