import express from "express";
import os from "os";

const app = express();
const PORT = 3000;

app.get("/", (_req, res) => {
  const helloMessage = `VERSION 4: Hello from the ${os.hostname()}`;
  console.log(helloMessage);
  res.json(helloMessage);
});

app.listen(PORT, () => {
  console.log(`Web server is listening at port ${PORT}`);
});
