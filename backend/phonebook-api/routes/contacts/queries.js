const development = require("../../configuration/development");

const configDB = development.db;
const Pool = require("pg").Pool;
const pool = new Pool(configDB);

const getContacts = async (request, response) => {
  try {
    const result = await pool.query(`SELECT * FROM contacts WHERE owner_id = ${request.params.owner_id} ORDER BY id ASC`);
    response.status(201).json(result.rows);
  } catch (error) {
    response.status(500).json(error);
  }
};

const getContactById = (request, response) => {
  const id = parseInt(request.params.id);

  pool.query("SELECT * FROM contacts WHERE id = $1", [id], (error, results) => {
    if (error) {
      throw error;
    }
    response.status(200).json(results.rows);
  });
};

const createContact = async (request, response) => {
  const { owner_id, email } = request.body; // personal, phone
  try {
    const result = await pool.query(
      `INSERT INTO contacts (owner_id, email, personal, phone, favorite) VALUES ($1, $2, $3, $4, $5)`,
      [owner_id, email, {}, {}, false]
    );
    response.status(201).send("Contact created!");
  } catch (error) {
    response.status(500).json({ error: error.detail });
  }
};

const deleteContact = (request, response) => {
  const id = parseInt(request.params.id);

  pool.query("DELETE FROM contacts WHERE id = $1", [id], (error, results) => {
    if (error) {
      throw error;
    }
    response.status(200).send(`Contact deleted with ID: ${id}`);
  });
};

module.exports = {
  getContacts,
  getContactById,
  createContact,
  deleteContact,
};
