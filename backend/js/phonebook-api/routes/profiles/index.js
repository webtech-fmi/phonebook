const express = require("express");
const profiles = express.Router();
const db = require("./queries");

profiles.use(express.json());

profiles.get("/:id", db.getProfileByProfileId);
profiles.get("/:user_id", db.getProfileByUserId);
profiles.post("/", db.createProfile);
profiles.put("/:id", db.editProfile);
profiles.delete("/:id", db.deleteProfile);


module.exports = profiles;
