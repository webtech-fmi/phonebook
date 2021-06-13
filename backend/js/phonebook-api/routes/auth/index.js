const express = require("express");
const users = express.Router();
const db = require("./queries");
const axios = require("axios");

users.use(express.json());

users.post("/signup", async (req, res, next) => {
  const userPayload = {
    full_name: req.body.full_name,
    email: req.body.email,
    password: req.body.password,
  };

  const profilePayload = {
    full_name: req.body.full_name,
    email: req.body.email,
    phone: req.body.phone,
    consent: req.body.consent,
  };

  await axios
    .post(process.env.AUTHENTICATION_SERVICE + "/users/create", userPayload)
    .then((resp) => {
      const { id: userID } = resp.data;
      profilePayload.user_id = userID;
    })
    .catch((error) => {
      console.error(error);
    });

  await axios
    .post(process.env.PROFILE_SERVICE + "/profiles/create", profilePayload)
    .then((respp) => {
      res.status(200).json({ status: "success" });
    })
    .catch((error) => {
      console.error(error);
    });
});

users.get("/", db.getUsers);
users.get("/:id", db.getUserById);
users.post("/", db.createUser);
users.put("/:id", db.updateUser);
users.delete("/:id", db.deleteUser);

module.exports = users;
