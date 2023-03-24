import {useRef, useState} from "react";
import AddNewButton from "./AddNewButton";

function ConversationMenu(props) {
    const [selectedChat, setSelectedChat] = useState(0)


    return (
        <>
            <div className="flex-wrap overflow-y-auto scroll-hide flex-grow">
                {
                    props.conversations.map((conversation, index) => (
                        <div className="w-full flex">
                            <button onClick={() => {
                                setSelectedChat(index);
                                props.conversationClick({id: conversation.id})
                            }}
                                    className={"flex flex-grow my-1 items-center space-x-3 text-gray-900 rounded-lg mx-3 px-4 py-2 font-medium capitalize hover:bg-gray-200 focus:shadow-outline " + (selectedChat === index ? "bg-gray-200" : "")}>


                                <span>{(conversation.name === "") ? "Undefined" : conversation.name}</span>
                            </button>
                        </div>
                    ))
                }
            </div>

            <div className="mb-5 w-full flex items-center justify-center">
                {/*<input ref={queryRef} maxLength={20} placeholder="Write a new chat" className="flex w-full p-2 text-sm items-centerspace-x-3 text-black placeholder-gray-400 rounded-md bg-gray-100" ></input>*/}
                <AddNewButton onClick={() => {
                    props.addConversationClick()
                }} name="Add Conversation"/>
            </div>

        </>
    );
}

export default ConversationMenu;
