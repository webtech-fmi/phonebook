const development = require("../../configuration/development");

const configDB = development.db;
const Pool = require("pg").Pool;
const pool = new Pool(configDB);


//test
const getProfileByProfileId = (request, response) => {
    const id = parseInt(request.params.id);

    pool.query("SELECT * FROM profiles WHERE id = $1", [id], (error, results) => {
        if (error) {
          throw error;
        }
        response.status(200).json(results.rows);
      });
};

const getProfileByUserId = (request, response) => {
    const user_id = parseInt(request.params.id);

    pool.query("SELECT * FROM profiles WHERE user_id = $1", [user_id], (error, results) => {
        if (error) {
          throw error;
        }
        response.status(200).json(results.rows);
      });
};

const createProfile = async (request, response) => {
    const { email, personal, phone, metadata } = request.body;
    try {
        const result = await pool.query(
        "INSERT INTO profiles (email, personal, phone, metadata) VALUES ($1, $2, $3, $4)",
        [ email, personal, phone, metadata]
    );
        response.status(201).send("User created!");
    } catch (error) {
        response.status(500).json({ error: error.detail });
    }
};

const editProfile = async (request, response) => {
    //not done yet
};

const deleteProfile = (request, response) => {
    const id = parseInt(request.params.id);

    pool.query("DELETE FROM profiles WHERE id = $1", [id], (error, results) => {
        if (error) {
            throw error;
        }
        response.status(200).send(`Profile deleted with ID: ${id}`);
    });
};


module.exports = {
    getProfileByProfileId,
    getProfileByUserId,
    createProfile,
    editProfile,
    deleteProfile,
};