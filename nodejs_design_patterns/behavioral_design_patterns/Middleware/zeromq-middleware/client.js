import zeromq from "zeromq";
import { ZmqMiddlewareManager } from "./middleware/zmq-middleware-manager.js";
import { jsonMiddleware } from "./middleware/json-middleware.js";
import { zlibMiddleware } from "./middleware/zlib-middleware.js";

async function main() {
  const socket = new zeromq.Request();
  await socket.connect("tcp://127.0.0.1:5000");

  const zmqm = new ZmqMiddlewareManager(socket);

  zmqm.use(zlibMiddleware());
  zmqm.use(jsonMiddleware());
  zmqm.use({
    inbound(message) {
      console.log("Echoed back", message);
      return message;
    },
  });

  setInterval(() => {
    zmqm
      .send({ action: "ping", echo: Date.now() })
      .catch((err) => console.error(err));
  }, 2000);

  console.log("Client connected");
}

main();
