const express = require("express");
const profiles = require("./profiles");
const router = express.Router();
const users = require("./users");
const contacts = require("./contacts");

router.get("/", (request, response) => {
  response.json({ info: "Node.js, Express, and Postgres API" });
});

router.use("/users", users);
router.use("/contacts", contacts);

router.use("/profiles", profiles)
module.exports = router;
