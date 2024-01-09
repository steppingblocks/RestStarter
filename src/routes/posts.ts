import express from "express";
import { db } from "../resources/db";

const posts = express.Router();

const queries = {
  GET_POSTS_FROM_USER: ``,
  INSERT_POST_FOR_USER: ``,
};

/**
 * GET endpoint /posts/:username retrieves posts for a given user
 *
 * returns a 404 if the user does not exist
 *
 * BONUS: add pagination
 */
posts.get("/:username", async (req, res) => {
  //... YOUR WORK HERE ...//
  res.json([]);
});

/**
 * POST endpoint /posts/:username creates a post for a given user from { "title": "Title", "content": "Content" } request JSON body
 *
 * returns a 404 if the user does not exist
 */
posts.post("/:username", async (req, res) => {
  //... YOUR WORK HERE ...//
  res.json({});
});

export { posts };
