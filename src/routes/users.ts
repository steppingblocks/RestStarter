import express from "express";
import { db } from "../resources/db";

const users = express.Router();

const queries = {
  ALL_USERS: `SELECT id, created_at, first_name, last_name, username FROM "Users"`,
  ONE_USER: `SELECT id, created_at, first_name, last_name, username FROM "Users" WHERE username=$1`,
  INSERT_USER: `INSERT INTO "Users" (first_name, last_name, username) VALUES ($1, $2, $3) returning id, first_name, last_name, created_at`,
};

const ERR = { error: "something went wrong" };

/**
 * GET endpoint /users returns all users
 */
users.get("/", async (_req, res) => {
  try {
    const result = await db.query(queries.ALL_USERS);
    res.json(result.rows);
  } catch (_e) {
    res.status(500);
    res.json(ERR);
  }
});

/**
 * GET endpoint /users/:username returns individual user or 404
 */
users.get("/:username", async (req, res) => {
  try {
    const { username } = req.params;
    const result = await db.query(queries.ONE_USER, [username]);

    if (result.rows.length < 1) {
      res.status(404);
      res.json({ error: "not found" });
    }

    console.log(result.rows);
    res.json(result.rows[0]);
  } catch (_e) {
    res.status(500);
    res.json(ERR);
  }
});

/**
 * POST endpoint /users creates a user from { first_name: "First", last_name: "Last"} request JSON body
 * returns JSON with the inserted values
 */
users.post("/", async (req, res) => {
  try {
    const json = req.body;
    const { first_name, last_name, username } = json;

    if (!first_name || !last_name || !username) {
      throw "missing values";
    }

    const result = await db.query(queries.INSERT_USER, [
      first_name,
      last_name,
      username,
    ]);

    if (result.rows.length !== 1) {
      throw "there was an error";
    }

    res.json(result.rows[0]);
  } catch (_e) {
    res.status(500);
    res.json(ERR);
  }
});

export { users };
