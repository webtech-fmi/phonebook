require('dotenv').config();
const cors = require('cors');
const app = require("express")();
const routes = require("./routes");
const PORT = process.env.PORT || 3000;

//  Connect all our routes to our application
app.use(cors());
app.use("/", routes);

// Turn on that server!
app.listen(PORT, () => {
  console.log(`App listening on port ${PORT}`);
});
