import zeromq from "zeromq";
import { ZmqMiddlewareManager } from "./middleware/zmq-middleware-manager.js";
import { jsonMiddleware } from "./middleware/json-middleware.js";
import { zlibMiddleware } from "./middleware/zlib-middleware.js";

async function main() {
  const socket = new zeromq.Reply();
  await socket.bind("tcp://127.0.0.1:5000");

  const zmqm = new ZmqMiddlewareManager(socket);

  zmqm.use(zlibMiddleware());
  zmqm.use(jsonMiddleware());
  zmqm.use({
    async inbound(message) {
      console.log("Received", message);

      if (message.action === "ping") {
        await this.send({ action: "pong", echo: message.echo });
      }

      return message;
    },
  });

  console.log("Server started");
}

main();
