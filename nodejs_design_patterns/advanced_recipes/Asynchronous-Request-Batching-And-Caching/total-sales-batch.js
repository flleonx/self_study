import { totalSales as totalSalesRaw } from "./total-sales.js";
const runningRequests = new Map();

export function totalSales(product) {
  if (runningRequests.has(product)) {
    console.log("Batching");
    return runningRequests.get(product);
  }

  const resultPromise = totalSalesRaw(product);
  runningRequests.set(product, resultPromise);
  resultPromise.finally(() => {
    runningRequests.delete(product);
  });

  return resultPromise;
}
