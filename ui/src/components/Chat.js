import React, { useState } from "react";
import MarkdownView from 'react-showdown';
import "./Chat.css";
import ConversationMenu from "./ConversationMenu";

function Chat() {
  const [input, setInput] = useState("");
  const markdown = `
  Yes, I can generate the python code to generate fibonacci sequence.
Here is the code below:

\`\`\`
def fib(n):
    if n <= 0:
        return []
    elif n == 1:
        return [0]
    else:
        fib_list = [0, 1]
        while len(fib_list) < n:
            fib_list.append(fib_list[-1] + fib_list[-2])
        return fib_list
\`\`\`

This fibonacci function takes an integer 'n' as input and returns the first 'n' fibonacci numbers in a list starting from 0.
If 'n' is less than or equal to 0, it returns an empty list. If 'n' is 1, it returns a list containing only 0.
For larger values of 'n', the function uses a while loop to generate the fibonacci sequence and add them to the list until the length of the list becomes equal to 'n'.
  `;
  const [messages, setMessages] = useState([
    { text: "sample question", sender: "user" },
    { text: markdown, sender: "bot" },
    { text: "sample question", sender: "user" },
    { text: markdown, sender: "bot" },
  ]);

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

  const handleInputTest = (query) => {
    console.log(query);
    setMessages([
      ...messages,
      { text: "hello", sender: "user" },
      { text: "chatgpt response", sender: "bot" },
    ])
  };

  // handle input on Enter key press
  const handleKeyPress = (e) => {
    if (e.key === "Enter") {
      handleInput();
    }
  };

  return (
    <div className="flex flex-row h-screen">
      <div className="w-1/5 bg-red-200 h-full">
        <ConversationMenu/>
      </div>
      <div className="flex flex-col">
          <ConverstationList messages={messages}/>
          <InputBox messages={messages} bt_title="Submit" onSubmit={handleInputTest} />
      </div>
    </div>
  );
}

function ConverstationList(props) {
  console.log(props);
  
  return (
    <div className="flex-grow chat-container pb-10 mr-2 bg-gray-50 overflow-y-auto scroll-hide">
      <div className="chat-history">
        {props.messages.map((message, index) => (
          <div className="block w-full p-6 text-sm border border-gray-300 rounded-lg bg-gray-50 my-3 mx-2">
            {
            (message.sender === "user") ? (
              <article className="max-w-full text-base">
                {message.text}
            </article>
            ): (
              <MarkdownView className="prose max-w-full" markdown={message.text}>
            </MarkdownView>
            )
            } 
          </div>
        ))}
      </div>
    </div>
  );
}


function InputBox(props) {
  var [query, setQuery] = useState("");

  function onInputChange(event) {
    setQuery(event.target.value);
  }

  function onClick (){
    props.onSubmit(query);
    console.log(query);
    setQuery("");
  }

  return (
    
      <div className="w-full p-6 text-sm border border-gray-300 rounded-lg bg-gray-50 py-3 px-2">
        <input type="text" id="query" value={query} onChange={onInputChange}
          className="block w-full max-w-full p-4 pl-10 text-sm text-gray-900 border border-gray-300 rounded-lg bg-gray-50 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
          placeholder="Write a question" required /> 
        <button type="submit" onClick={onClick} className="text-white absolute right-2.5 bottom-2.5 bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
          {props.bt_title}
        </button>
      </div>
     

  );
}

export default Chat;