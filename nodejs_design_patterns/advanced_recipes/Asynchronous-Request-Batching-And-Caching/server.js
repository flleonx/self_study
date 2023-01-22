import { createServer } from "http";
// import { totalSales } from "./total-sales.js";
// import { totalSales } from "./total-sales-batch.js";
import { totalSales } from "./total-sales-cache.js";

const PORT = 57000;

createServer(async (req, res) => {
  const url = new URL(req.url, "http://localhost");
  const product = url.searchParams.get("product");
  console.log(`Processing query: ${url.search}`);
  const sum = await totalSales(product);

  res.setHeader("Content-Type", "application/json");
  res.writeHead(200);
  res.end(
    JSON.stringify({
      product,
      sum,
    })
  );
}).listen(PORT, () => console.log("Server started"));
