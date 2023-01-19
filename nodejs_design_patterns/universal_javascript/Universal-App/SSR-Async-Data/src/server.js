import { resolve, dirname } from "path";
import { fileURLToPath } from "url";
import React from "react";
import reactServer from "react-dom/server.js";
import htm from "htm";
import fastify from "fastify";
import fastifyStatic from "@fastify/static";
import { StaticRouter } from "react-router-dom";
import { App } from "./frontend/App.js";

const __dirname = dirname(fileURLToPath(import.meta.url));
const html = htm.bind(React.createElement);

const template = ({ content }) => `<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8">
    <title>My library</title>
  </head>
  <body>
    <div id="root">${content}</div>
    <script type="text/javascript" src="/dist/main.js"></script>
  </body>
</html>`;

const server = fastify({ logger: true });

server.register(fastifyStatic, {
  root: resolve(__dirname, "..", "dist"),
  prefix: "/dist/",
});

server.get("*", async (req, reply) => {
  const location = req.url;
  const staticContext = {};

  const serverApp = html`
    <${StaticRouter} location=${location} context=${staticContext}>
      <${App}/>
    </${StaticRouter}>
  `;

  const content = reactServer.renderToString(serverApp);
  const responseHtml = template({ content });

  let code = 200;
  if (staticContext.statusCode) {
    code = staticContext.statusCode;
  }

  reply.code(code).type("text/html").send(responseHtml);
});

const port = Number.parseInt(process.env.PORT) || 3000;
const host = process.env.ADDRESS || '127.0.0.1';

server.listen({ port, host }, function(err) {
  if (err) {
    console.error(err);
    process.exit(1);
  }
})

