import React, { useState } from "react";
import "./App.css";

function Chat() {
  const [input, setInput] = useState("");
  const [messages, setMessages] = useState([]);

  // handle user input and send request to the API
  const handleInput = async () => {
    if (input.trim() !== "") {
      const response = await fetch("/chat", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ query: input }),
      });
      const data = await response.json();

      setMessages([
        ...messages,
        { text: input, sender: "user" },
        { text: data.response, sender: "bot" },
      ]);
      setInput("");
    }
  };

  const handleInputTest = () =>{
    setMessages([
        ...messages,
        { text: "hello", sender: "user" },
        { text: "chatgpt response", sender: "bot" }, 
    ])
  } 

  // handle input on Enter key press
  const handleKeyPress = (e) => {
    if (e.key === "Enter") {
      handleInput();
    }
  };

  return (
    <div className="App">
      <div className="chat-header">
        <h1>ChatGPT UI</h1>
      </div>
      <div className="chat-container">
        <div className="chat-history">
          {messages.map((message, index) => (
            <div
              key={index}
              className={
                message.sender === "user" ? "user-message" : "bot-message"
              }
            >
              <p>{message.text}</p>
            </div>
          ))}
        </div>
        <div className="chat-input">
          <input
            type="text"
            placeholder="Type your message..."
            value={input}
            onChange={(e) => setInput(e.target.value)}
            onKeyPress={handleKeyPress}
          />
          <button onClick={handleInputTest}>Send</button>
        </div>
      </div>
    </div>
  );
}

export default Chat;