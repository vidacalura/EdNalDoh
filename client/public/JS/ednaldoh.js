// Este arquivo lida com a lógica do jogo e cartas
// (eventos enviados e recebidos do servidor)
let socket = io();
const API = "http://localhost:4000/api/";

const codigoSala = window.location.href.split("/")[window.location.href.split("/").length - 1];
socket.emit("entrarSala", codigoSala);

const codigoSalaTextbox = document.getElementById("input-convite");
codigoSalaTextbox.value = window.location.href;


socket.on("iniciarPartida", (dadosPartida) => {
    if (codigoSala == dadosPartida.codigoSala) {
        const telaConvite = document.getElementById("convite-container");
        telaConvite.style.display = "none";
    }
});

socket.on("notificacaoJogoMudanca", atualizarJogo);

socket.on("desconexaoDeAdversario", (codigoSalaDesconexao) => {
    if (codigoSala == codigoSalaDesconexao) {
        alert("Oh não! Seu oponente saiu da partida :(");
        window.location.href = "/";
    }
});

function atualizarJogo() {
    fetch(API + "jogo/usuario/" + socket.id)
    .then((res) => { return res.json(); })
    .then((res) => {
        if (res.error) {
            alert("Ops! Ocorreu um erro tentando se conectar ao servidor! Voltando à página principal...");
            window.location.href = "/";
        }

        // Mostra sua mão
        const maoContainer = document.getElementById("sua-mao-div");
        for (const carta of res.voce.mao) {
            const cartaDiv = document.createElement("div");
            cartaDiv.classList.add("slot-carta-mao");
            cartaDiv.style.backgroundImage = `url(${carta.imgPath})`;

            maoContainer.appendChild(cartaDiv);
        }

        // Mostra mão do adversário
        const maoAdversarioContainer = document.getElementById("mao-oponente-div");
        for (let i = 0; i < res.maoAdversario; i++) {
            const cartaDiv = document.createElement("div");
            cartaDiv.classList.add("slot-carta-mao");
            cartaDiv.style.backgroundImage = `url('/carta-virada.webp')`;

            maoAdversarioContainer.appendChild(cartaDiv);
        }

        // Atualiza seu campo
        // Atualiza campo do adversário

        // Atualiza sua vida
        // Atualiza vida do adversário
    });
}