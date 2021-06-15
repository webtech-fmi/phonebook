const express = require("express");
const contacts = express.Router();
const axios = require("axios");

contacts.use(express.json());

contacts.get("/by-owner", async (req, res, next) => {
  try {
    const response = await axios.get(
      process.env.CONTACT_SERVICE + "/contacts/by-owner?id=" + req.query.id
    );
    res.status(200).json(response.data);
  } catch (error) {
    res.status(500).json({ error: error });
  }
});

contacts.get("/by-id", async (req, res, next) => {
  try {
    const response = await axios.get(
      process.env.CONTACT_SERVICE + "/contacts/by-id?id=" + req.query.id
    );

    res.status(200).json(response.data);
  } catch (error) {
    res.status(500).json({ error: error });
  }
});

// unverified
contacts.post("/create", async (req, res, next) => {
  const contact = {
    email: req.body.email,
    personal: req.body.personal,
    phone: req.body.phone,
    metadata: req.body.metadata,
  };
  try {
    const response = await axios.post(
      process.env.CONTACT_SERVICE + "/contacts/create?id=" + req.query.id,
      contact
    );

    res.status(200).json({ status: "success" });
  } catch (error) {
    res.status(500).json({ error: error });
  }
});

contacts.post("/edit", async (req, res, next) => {
  try {
    const editPayload = {
      email: req.body.email,
      phone: req.body.phone,
      personal: req.body.personal,
      metadata: req.body.metadata,
    };

    const response = await axios.post(
      process.env.CONTACT_SERVICE + "/contacts/edit?id=" + req.body.session_id,
      editPayload
    );

    res.status(200).json(response);
  } catch (error) {
    res.status(500).json({ error: error });
  }
});

contacts.post("/merge", async (req, res, next) => {
  try {
    const mergePayload = {
      main: req.body.main,
      contacts: req.body.contacts,
    };

    const response = await axios.post(
      process.env.CONTACT_SERVICE + "/contacts/merge?id=" + req.body.session_id,
      mergePayload
    );

    res.status(200).json({ status: "success" });
  } catch (error) {
    res.status(500).json({ error: error });
  }
});

contacts.post("/favourite", async (req, res, next) => {
  const favouritePayload = {
    session_id: req.body.session_id,
    id: req.body.id,
    favourite: req.body.favourite,
  };

  try {
    const response = await axios.post(
      process.env.CONTACT_SERVICE + "/contacts/favourite",
      favouritePayload
    );

    res.status(200).json(response.data);
  } catch (error) {
    res.status(500).json({ error: error });
  }
});

contacts.post("/delete", async (req, res, next) => {
  const deletePayload = {
    session_id: req.body.session_id,
    id: req.body.id,
  };

  try {
    const response = await axios.post(
      process.env.CONTACT_SERVICE + "/contacts/delete",
      deletePayload
    );

    res.status(200).json(response.data);
  } catch (error) {
    res.status(500).json({ error: error });
  }
});

module.exports = contacts;
