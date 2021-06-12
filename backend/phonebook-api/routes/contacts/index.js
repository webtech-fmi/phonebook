const express = require("express");
const contacts = express.Router();
const db = require("./queries");

contacts.use(express.json());

contacts.post("/", db.getContacts);
contacts.get("/:id", db.getContactById);
contacts.post("/", db.createContact);
contacts.delete("/:id", db.deleteContact);

module.exports = contacts;