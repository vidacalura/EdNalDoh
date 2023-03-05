require("dotenv").config();
const express = require("express");
const socket = require("socket.io");
const cors = require("cors");

let app = express();
let server = app.listen(process.env.PORT || 5000);

app.use(express.static(__dirname + "/public/CSS/"));
app.use(express.static(__dirname + "/public/JS/"));
app.use(express.static(__dirname + "/public/assets/"));
app.use(cors({ origin: "*" }));
app.use(express.json());

/* Rotas */
app.get("/", (req, res) => {
    const codigoSala = Math.ceil(Math.random() * 10) * 1000 + Math.ceil(Math.random() * 100);
    res.redirect(307, `/jogo/${codigoSala}`);
});

app.get("/jogo/:codigoSala", (req, res) => {
    res.sendFile("public/index.html", { root: __dirname });
});

app.get("/help", (req, res) => {
    res.sendFile("./public/regras.html", { root: __dirname });
});

/* Jogo */
class Sala {
    constructor(idJogador1, idJogador2, codigoSala) {
        this.idJogador1 = idJogador1;
        this.idJogador2 = idJogador2;
        this.codigoSala = codigoSala;
    }
}

let salas = [];

let io = socket(server);
io.on("connection", (socket) => {

    socket.on("entrarSala", (codigoSala) => {
        entrarSala(socket.id, codigoSala);
    });

    socket.on("disconnect", () => {
        // Remover da sala
    });

});


function entrarSala(id, codigoSala) {
    for (const sala of salas) {
        if (sala.codigoSala == codigoSala) {
            // Checar se sala já está cheia
                // Redirecionar para /
            sala.idJogador2 = id;
            iniciarPartida(sala);
            return;
        }
    }

    const sala = new Sala(id, null, codigoSala);
    salas.push(sala);
}

async function iniciarPartida(sala) {

    await fetch(process.env.API + "jogo", {
        method: "POST",
        headers: {
            "Content-type": "Application/JSON"
        },
        body: JSON.stringify({
            idJogador1: sala.idJogador1,
            idJogador2: sala.idJogador2,
            codigoSala: sala.codigoSala
        })
    })
    .then((res) => { return res.json(); })
    .then((res) => {
        // Enviar aos sockets da sala para remover tela de convite
        console.log(res);
    })
    .catch((err) => {
        console.log(err);
    });

}