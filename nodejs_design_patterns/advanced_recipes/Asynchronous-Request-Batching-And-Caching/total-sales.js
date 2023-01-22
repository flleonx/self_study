import { Level } from "level";
import sublevel from "subleveldown";

const db = new Level("example-db");
const salesDb = sublevel(db, "sales", { valueEncoding: "json" });

export async function totalSales(product) {
  const now = Date.now();
  let sum = 0;

  for await (const transacion of salesDb.createValueStream()) {
    if (!product || transacion.product === product) {
      sum += transacion.amount;
    }
  }

  console.log(`totalSales() took: ${Date.now() - now} ms`);
  return sum;
}
