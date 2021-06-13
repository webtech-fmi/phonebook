const development = require("../../configuration/development");

const configDB = development.db;
const Pool = require("pg").Pool;
const pool = new Pool(configDB);

const getUsers = async (request, response) => {
  try {
    const result = await pool.query("SELECT * FROM users ORDER BY id ASC");
    response.status(201).json(result.rows);
  } catch (error) {
    response.status(500).json(error);
  }
};

const getUserById = (request, response) => {
  const id = parseInt(request.params.id);

  pool.query("SELECT * FROM users WHERE id = $1", [id], (error, results) => {
    if (error) {
      throw error;
    }
    response.status(200).json(results.rows);
  });
};

const createUser = async (request, response) => {
  const { full_name, email, phone_number, password } = request.body;
  try {
    const result = await pool.query(
      "INSERT INTO users (full_name, email, password, phone_number) VALUES ($1, $2, $3, $4)",
      [full_name, email, password, phone_number]
    );
    response.status(201).send("User created!");
  } catch (error) {
    response.status(500).json({ error: error.detail });
  }
};

const updateUser = (request, response) => {
  const id = parseInt(request.params.id);
  const { name, email } = request.body;

  pool.query(
    "UPDATE users SET name = $1, email = $2 WHERE id = $3",
    [name, email, id],
    (error, results) => {
      if (error) {
        throw error;
      }
      response.status(200).send(`User modified with ID: ${id}`);
    }
  );
};

const deleteUser = (request, response) => {
  const id = parseInt(request.params.id);

  pool.query("DELETE FROM users WHERE id = $1", [id], (error, results) => {
    if (error) {
      throw error;
    }
    response.status(200).send(`User deleted with ID: ${id}`);
  });
};

module.exports = {
  getUsers,
  getUserById,
  createUser,
  updateUser,
  deleteUser,
};
