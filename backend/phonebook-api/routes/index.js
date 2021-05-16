const express = require("express");
const router = express.Router();
const users = require("./users");

router.get("/", (request, response) => {
  response.json({ info: "Node.js, Express, and Postgres API" });
});

router.use("/users", users);

module.exports = router;
