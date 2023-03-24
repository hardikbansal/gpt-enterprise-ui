import {useState} from "react";
import SendLogo from "../assets.images/send.png"

function ChatInputBox(props) {
    var [query, setQuery] = useState("");

    function onInputChange(event) {
        setQuery(event.target.value);
    }

    function onSendClick() {
        console.log("Querying with query: ", query);
        props.onSubmit(query);
        setQuery("");
    }

    return (
        <div className="w-full flex flex-col my-5  items-center px-5">
            <div className="w-full max-w-7xl flex items-center  shadow-lg rounded text-sm bg-gray-50 p-3">
                <textarea id="query" value={query} onChange={onInputChange}
                      className="block w-full p-4 rounded-md border border-gray-200 bg-gray-100 placeholder-gray-500 text-black text-base mr-6"
                      placeholder="Write a question" required/>
                <img src={SendLogo} width="40" className="block h-10 stroke-1" onClick={onSendClick} alt="send logo"/>
            </div>
        </div>


    );
}

export default ChatInputBox;