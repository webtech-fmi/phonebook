const customError = (defaultMsg, status) => {
  return class GoodError extends Error {
    constructor(msg, ...args) {
      super(msg || defaultMsg, ...args);
      this.status = status;
      Error.captureStackTrace(this, GoodError);
    }
  };
};

const BadRequest = customError("Bad Request", 400);
const Unauthorized = customError("Unauthorized", 401);
const NotFound = customError("Not found", 404);
const ExpectationFailed = customError("Unauthorized", 417);
const ServerError = customError("Internal server error", 500);

module.exports = {
  customError,
  BadRequest,
  Unauthorized,
  NotFound,
  ExpectationFailed,
  ServerError,
};
