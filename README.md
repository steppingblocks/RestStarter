# REST Starter Repo

Steppingblocks starter repo for REST API coding challenge in golang.

### Setup Checklist

- [ ] **Golang verison**
      This project was created with Go 1.22. 
- [ ] **Install Dependencies**
      `go mod download`
- [ ] **Environment Variables**
      Rename `.sample.env` to `.env` in the root directory of your project, and we'll give you the database password to fill in
- [ ] **Run the dev script**
      `go run main.go`. This will require a rerun every code change as base go is not set up to do live refresh.
- [ ] **Checking your work**
      Use an API testing tool such as Postman or `curl` to test your work

### Starter Code

The application's entry point is in the `main.go` file. You can look here to find your way around. You'll see a basic gin server.

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

Complete the code in `/pkg/posts/controller.go` the `/posts` endpoint with a GET request to /posts/:username to retrieve a user's posts in reverse-chronological order and a POST request to /posts/:username to create a post for a given user. You do not need to worry about authorization.

Feel free to look at `/pkg/users/controller.go` if you need some example gin and sqlx code.

Bonus: add pagination to the GET request or describe how you would accomplish this
