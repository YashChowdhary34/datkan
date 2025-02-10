require("dotenv").config();
const express = require("express");
const bodyparser = require("body-parser");
const cors = require("cors");
const mongoose = require("mongoose");

const app = express();

app.use(bodyparser.json());
app.use(cors());

mongoose
  .connect(process.env.MONGODB_API_KEY, {
    useNewUrlParser: true,
    useUnifiedTopology: true,
  })
  .then(console.log("Successfully connected to MongoDB!"))
  .catch((err) => console.error("MongoDB connection error: ", err));

app.get("/", (request, response) => {
  request.send("Backend is running!");
});

const PORT = process.env.PORT;
app.listen(PORT, () => console.log(`Server running on port -> ${PORT}`));
