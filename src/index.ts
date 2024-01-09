import express from "express";
import { users } from "./routes/users";
import { config } from "./config";
import { db } from "./resources/db";
import { posts } from "./routes/posts";
import morgan from "morgan";

const app = express();
app.use(express.json());
app.use(morgan("dev"));

app.use("/users", users);
app.use("/posts", posts);

(async () => {
  await db.connect();
  app.listen(config.port, () => {
    console.log(`listening on port ${config.port}`);
  });
})();
