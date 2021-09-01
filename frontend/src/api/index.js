const socket = new WebSocket("ws://localhost:8080/ws");

let connect = cb => {
    console.log("Attempting Connection...");

    socket.onopen = () => {
        console.log("Successfully Connected");
    };

    socket.onmessage = msg => {
        console.log(msg);
        cb(msg);
    };

    socket.onclose = event => {
        console.log("Socket Closed Connection: ", event);
    };

    socket.onerror = error => {
        console.log("Socket Error: ", error);
    };
};

let sendMsg = (type, msg) => {
    console.log("sending msg: ", type, msg);
    socket.send(JSON.stringify({type: type, body: msg}));
};

export { connect, sendMsg };