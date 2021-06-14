const express = require("express");
const contacts = express.Router();
const axios = require("axios");

contacts.use(express.json());

contacts.get("/get", async(req, res, next) => {
    try{
        const response = await axios.get(
            process.env.CONTACT_SERVICE + "contacts/by-owner?id" + req.body.session_id
        );
        res.status(200).json(response);
    } catch (error) {
        res.status(500).json({ error: error })
    }
});

contacts.get("/get", async (req, res, next) => {

    const contactId = {
        id: req.body.id
    }
    try {
      const response = await axios.get(
        process.env.CONTACT_SERVICE + "/contacts/by-id?id=" + req.body.session_id,
        contactId
      );
  
      res.status(200).json(response);
    } catch (error) {
      res.status(500).json({ error: error });
    }
});

contacts.post("/post", async(req, res, next) => {
    const contact = {
        email: req.body.email,
        personal: req.body.personal,
        phone: req.body.phone,
        metadata: req.body.metadata
    }

    try {
        const response = await axios.get(
          process.env.CONTACT_SERVICE + "/contacts/create" + req.body.session_id,
          contact
        );
    
        res.status(200).json(response);
      } catch (error) {
        res.status(500).json({ error: error });
      }
})

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
        const id = {
          id: req.body.id
        };
        
        const ids = {
          contacts: req.body.contacts.id
        }
    
        const response = await axios.post(
          process.env.CONTACT_SERVICE + "/contacts/merge?id=" + req.body.session_id,
          id,
          ids
        );
    
        res.status(200).json(response);
      } catch (error) {
        res.status(500).json({ error: error });
      }
});

module.exports = contacts;