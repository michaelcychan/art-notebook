import axios from "axios";

export default axios.create({
  baseURL: "http://localhost:3001/v1/",
  headers: {
    "Content-type": "application/json"
  }
});