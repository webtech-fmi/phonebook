const express = require("express");
const users = express.Router();
const db = require("./queries");

users.use(express.json());

users.get("/", db.getUsers);
users.get("/:id", db.getUserById);
users.post("/", db.createUser);
users.put("/:id", db.updateUser);
users.delete("/:id", db.deleteUser);

module.exports = users;