import express from "express";
import fetch from "node-fetch";
import os from "os";

const app = express();
const PORT = 3000;

app.get("/", (_req, res) => {
  const helloMessage = `Hello from the ${os.hostname()}`;
  console.log(helloMessage);
  res.json(helloMessage);
});

app.get("/nginx", async (_req, res) => {
  const url = "http://nginx";
  const response = await fetch(url);
  const body = await response.text();
  res.send(body);
});

app.listen(PORT, () => {
  console.log(`Web server is listening at port ${PORT}`);
});
