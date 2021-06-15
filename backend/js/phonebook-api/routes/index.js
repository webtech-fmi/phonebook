const express = require("express");
const profiles = require("./profiles");
const router = express.Router();
const auth = require("./auth");
const contacts = require("./contacts");

router.get("/", (request, response) => {
  response.json({ info: "Phonebook API!" });
});

router.use("/auth", auth);
router.use("/contacts", contacts);

router.use("/profiles", profiles)
module.exports = router;
