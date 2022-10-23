import { Database } from "./database.js";

export const dbInstance = new Database("my-app-db", {
  url: "localhost:5432",
  username: "user",
  password: "password",
});
