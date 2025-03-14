import React, { useState } from "react";
import axios from "axios";
import "./App.css";

function App() {
  const [form, setForm] = useState({ name: "", email: "", feedback: "" });
  const [message, setMessage] = useState("");

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const backendUrl = process.env.REACT_APP_BACKEND_URL || window._env_.REACT_APP_BACKEND_URL;
      console.log("Version: 7.0 and Backend url: ", backendUrl);
      const response = await axios.post(`${backendUrl}/submit`, form);
      setMessage(response.data.message);
    } catch (error) {
      if (error.response) {
        // The request was made and the server responded with a status code
        // that falls out of the range of 2xx
        console.error("Error response:", error.response);
        setMessage(`Failed to submit form: ${error.response.data.message || error.response.statusText}`);
      } else if (error.request) {
        // The request was made but no response was received
        console.error("Error request:", error.request);
        setMessage("Failed to submit form: No response from server");
      } else {
        // Something happened in setting up the request that triggered an Error
        console.error("Error message:", error.message);
        setMessage(`Failed to submit form: ${error.message}`);
      }
    }
  };

  return (
    <div className="container">
      <h2>Feedback Form</h2>
      <form onSubmit={handleSubmit}>
        <input type="text" name="name" placeholder="Name" onChange={handleChange} required />
        <input type="email" name="email" placeholder="Email" onChange={handleChange} required />
        <textarea name="feedback" placeholder="Your feedback" onChange={handleChange} required></textarea>
        <button type="submit">Submit</button>
      </form>
      {message && <p className="message">{message}</p>}
    </div>
  );
}

export default App;