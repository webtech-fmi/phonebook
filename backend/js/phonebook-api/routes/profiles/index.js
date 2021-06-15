const express = require("express");
const profiles = express.Router();
const axios = require("axios");

profiles.use(express.json());

profiles.post("/get", async (req, res, next) => {
  try {
    const response = await axios.get(
      process.env.PROFILE_SERVICE +
        "/profiles/by-owner?id=" +
        req.body.session_id
    );

    res.status(200).json(response.data);
  } catch (error) {
    res.status(500).json({ error: error });
  }
});

profiles.post("/edit", async (req, res, next) => {
  try {
    const editPayload = {
      email: req.body.email,
      phone: req.body.phone,
      personal: req.body.personal,
      metadata: req.body.metadata,
    };

    const response = await axios.post(
      process.env.PROFILE_SERVICE + "/profiles/edit?id=" + req.body.session_id,
      editPayload
    );

    res.status(200).json(response.data);
  } catch (error) {
    res.status(500).json({ error: error });
  }
});

module.exports = profiles;
