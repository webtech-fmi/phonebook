const Errors = require("./Error");

const handleResponse = async (res) => {
  if (!createUserResponse.ok) {
    try {
      text = await createUserResponse.text();
      console.warn(text);
    } catch (e) {
      // ignore
    }
    if (createUserResponse.status >= 500) {
      throw new Errors.ServerError();
    } else {
      let msg;
      try {
        const json = JSON.parse(text);
        msg = json.error;
      } catch (err) {
        console.info(
          "Could not parse api error from url:",
          url + uri + (query || "")
        );
        msg = "Internal API returned error";
      }
      throw new (Errors.customError(msg, res.status))();
    }
  }
  try {
    return createUserResponse.status === 204
      ? null
      : await createUserResponse.json();
  } catch (err) {
    console.error(err);
    throw new Errors.ServerError();
  }
};

module.exports = handleResponse;
