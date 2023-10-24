import React, { useState, useEffect } from "react";
import axios from "axios";

function App() {
  const [message, setMessage] = useState("");

  useEffect(() => {
    axios.get("http://localhost:8080/").then((response) => {
      setMessage(response.data);
    });
  }, []);

  return (
    <div className="App">
      <header className="App-header">{message}</header>
    </div>
  );
}

export default App;
