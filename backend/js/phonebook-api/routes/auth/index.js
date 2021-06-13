const express = require("express");
const users = express.Router();
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

  try {
    const response = await axios.post(
      process.env.AUTHENTICATION_SERVICE + "/users/create",
      userPayload
    );

    profilePayload.user_id = response.data.id;
    await axios.post(
      process.env.PROFILE_SERVICE + "/profiles/create",
      profilePayload
    );
    res.status(200).json({ status: "success" });
  } catch (error) {
    res.status(500).json({ error: error.message });
  }
});

users.post("/login", async (req, res, next) => {
  const loginPayload = {
    email: req.body.email,
    password: req.body.password,
  };

  try {
    const response = await axios.post(
      process.env.AUTHENTICATION_SERVICE + "/login/",
      loginPayload
    );

    res.status(200).json(response.data);
  } catch (error) {
    res.status(500).json({ error: error });
  }
});

users.post("/logout", async (req, res, next) => {
  const logoutPayload = {
    session_id: req.body.session_id,
  };

  try {
    const response = await axios.post(
      process.env.AUTHENTICATION_SERVICE + "/logout/",
      logoutPayload
    );

    res.status(200).json({ status: "success" });
  } catch (error) {
    res.status(500).json({ error: error });
  }
});

users.post("/reset", async (req, res, next) => {
  const lockPayload = {
    reason: "reset_password",
  };

  try {
    const response = await axios.put(
      process.env.AUTHENTICATION_SERVICE + "/users/lock?id=" + req.body.id,
      lockPayload
    );

    const resetPayload = {
      code: response.data.lock.code,
      password: req.body.password,
    };

    await axios.put(
      process.env.AUTHENTICATION_SERVICE + "/users/reset?id=" + req.body.id,
      resetPayload
    );
    res.status(200).json({ status: "success" });
  } catch (error) {
    res.status(500).json({ error: error });
  }
});

module.exports = users;
