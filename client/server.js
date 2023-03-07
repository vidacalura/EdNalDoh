require("dotenv").config();
const express = require("express");
const socket = require("socket.io");
const cors = require("cors");

let app = express();
let server = app.listen(process.env.PORT || 5000);

app.use(express.static(__dirname + "/public/CSS/"));
app.use(express.static(__dirname + "/public/JS/"));
app.use(express.static(__dirname + "/public/assets/"));
app.use(express.json());

/* Rotas */
app.get("/", (req, res) => {
    const codigoSala = Math.ceil(Math.random() * 10) * 1000 + Math.ceil(Math.random() * 1000);
    // Checar se sala já existe
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
        sairSala(socket.id);
    });

});


function entrarSala(id, codigoSala) {
    for (const sala of salas) {
        if (sala.codigoSala == codigoSala) {

            const salaCheia = sala.idJogador1 && sala.idJogador2;
            if (!salaCheia) {
                sala.idJogador2 = id;
                iniciarPartida(sala);
                return;
            }
            
            return;
        }
    }

    const sala = new Sala(id, null, codigoSala);
    salas.push(sala);
}

function sairSala(id) {
    for (let i = 0; i < salas.length; i++) {
        if (salas[i].idJogador1 == id || salas[i].idJogador2 == id) {
            io.sockets.emit("desconexaoDeAdversario", salas[i].codigoSala);
            salas.splice(i, 1);
            // Deletar sala na API
            return;
        }
    }
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
        io.sockets.emit("iniciarPartida", (res));
        io.sockets.emit("notificacaoJogoMudanca", null);
    })
    .catch((err) => {
        console.log(err);
    });

}