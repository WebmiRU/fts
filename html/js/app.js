
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
// let message = textarea.value;
let messages = document.querySelector(".messages");

// textarea.onkeydown = function (e) {
//     if (e.key === "Enter") {
//         e.preventDefault();
//     }
// }

textarea.onkeyup = function (e) {
    if (e.key === "Enter") {
        // socket.send();
        //  let message = textarea.value;
        // textarea.value



        messages.insertAdjacentHTML('beforeend', `
            <div class="messages_1" id="${999}">
                <div class="icon-avatar"></div>
                <div class="nickname">nickname</div>
                <div class="time">${(new Date).toLocaleTimeString()}</div>
                <div class="text">${textarea.value}</div>
                <div class="replies">
                    <div class="mini-avatar"></div>
                    <div class="reply">0 reply</div>
                    <div class="time_day">7 day ago</div>
                </div>
            </div>        
        `);
        textarea.value = null;
    }
}

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
