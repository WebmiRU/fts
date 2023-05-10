let socket = new WebSocket("ws://localhost:13900");

socket.onopen = function(e) {
    console.log("SOCKET OPEN");
    socket.send(msg1);
};

socket.onmessage = function(e) {
    console.log("ПОЛУЧЕНЫ ДАННЫЕ", e.data)
};

socket.onclose = function(e) {
    console.log("SOCKET CLOSED", e)
};

socket.onerror = function(error) {
    console.log("SOCKET ERROR", error)
};

let textarea = document.querySelector(".editor textarea");
let message = textarea.value;

let msg1 = JSON.stringify({ // Auth (login)
    type: "auth/login",
    payload: {
        token: "token2"
    }
});

let msg2 = JSON.stringify({ // Channel message
    type: "channel/message",
    payload: {
        channel_id: 3,
        message: message,
        text: message,
    },
});