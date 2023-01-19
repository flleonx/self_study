import { resolve, dirname } from "path";
import { fileURLToPath } from "url";
import React from "react";
import reactServer from "react-dom/server.js";
import htm from "htm";
import fastify from "fastify";
import fastifyStatic from "@fastify/static";
import { StaticRouter, matchPath } from "react-router-dom";
import { routes } from "./frontend/routes.js";
import { App } from "./frontend/App.js";

const __dirname = dirname(fileURLToPath(import.meta.url));
const html = htm.bind(React.createElement);

const template = ({ content, serverData }) => `<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8">
    <title>My library</title>
  </head>
  <body>
    <div id="root">${content}</div>
    ${
      serverData
        ? `<script type="text/javascript"> window.__STATIC_CONTEXT__=${JSON.stringify(
            serverData
          )}</script>`
        : ""
    }
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
  let component;
  let match;

  for (const route of routes) {
    component = route.component;
    match = matchPath(location, route);
    if (match) {
      break;
    }
  }

  let staticData;
  let staticError;
  let hasStaticContext = false;

  if (typeof component.preloadAsyncData === "function") {
    hasStaticContext = true;
    try {
      const data = await component.preloadAsyncData({ match });
      staticData = data;
    } catch (error) {
      staticError = err;
    }
  }

  const staticContext = {
    [location]: {
      data: staticData,
      err: staticError,
    },
  };

  const serverApp = html`
  <${StaticRouter} location=${location} context=${staticContext}>
    <${App} />
  </${StaticRouter}>
  `;

  const content = reactServer.renderToString(serverApp);
  const serverData = hasStaticContext ? staticContext : null;
  const responseHtml = template({ content, serverData });
  const code = staticContext.statusCode ? staticContext.statusCode : 200;

  reply.code(code).type("text/html").send(responseHtml);
});

const port = Number.parseInt(process.env.PORT) || 3000;
const host = process.env.ADDRESS || "127.0.0.1";

server.listen({ port, host }, function (err) {
  if (err) {
    console.error(err);
    process.exit(1);
  }
});
