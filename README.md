# REST Starter Repo

Steppingblocks starter repo for REST API coding challenge.

### Setup Checklist

- [ ] **Node.js Version**
      This project was created with Node v20.2.0. If you use nvm, you can just run `nvm use`, and install this version if you don't have it installed
- [ ] **Install Dependencies**
      `npm install`
- [ ] **Environment Variables**
      Rename `.sample.env` to `.env` in the root directory of your project and we'll give you the database password to fill in
- [ ] **Run the dev script**
      `npm run dev`
- [ ] **Checking your work**
      Use an API testing tool such as Postman or `curl` to test your work

### Starter Code

The application's entry point is in the `src/index.ts` file. You can look here to find your way around. You'll see a basic express server.

We've got a couple of endpoints to do some basic CRUD actions on the `Users` table to get you started.

### Database

You'll be using a Postgres database that we've set up on [Supabase](https://supabase.com/). We've got two tables so far:

```
Users:
    id (uuid)
    username (varchar)
    first_name (varchar)
    last_name (varchar)
    created_at (timestamp)

Posts:
    id (uuid)
    title (varchar)
    content (text)
    user_id (foreign key, uuid)
    created_at (timestamp)
```

The `Users` have a one-to-many relationship with `Posts`.

### Assignment

Complete the code in `/src/routes/posts.ts` the `/posts` endpoint with a GET request to /posts/:username to retrieve a user's posts in reverse-chronological order and a POST request to /posts/:username to create a post for a given user. You do not need to worry about authorization.

Feel free to look at `/src/routes/users.ts` if you need some example express code.

Bonus: add pagination to the GET request or describe how you would accomplish this

NOTE: `allowJS` is turned on, feel free to rename `posts.ts` to `posts.js` and work in pure JS if you are more comfortable that way.
