import ConversationMenu from "./ConversationMenu";
import {useState} from "react";
import TemplateMenu from "./TemplateMenu";

function SideMenu(props) {
    const CONVERSATION = "Conversation";
    const TEMPLATE = "Template";
    const [sideMenu, setSideMenu] = useState(CONVERSATION);

    const menuComponents = {
        "Conversation" : <ConversationMenu {...props.conversationMenuProps}/>,
        "Template": <TemplateMenu {...props.templateProps}/>,
    }

    const handleContextChange = event => {
        props.setContext(event.target.checked)
        console.log("context changed to: ", event.target.checked)
    };

    return (
        <>
            <div className="flex flex-row justify-evenly items-center p-2">
                <button onClick={() => {
                    setSideMenu(CONVERSATION)
                }}
                        className={"font-semibold text-lg text-gray-700 hover:bg-gray-200 rounded-md p-2 capitalize font-poppins tracking-wide "+ (sideMenu === CONVERSATION ? "bg-gray-200" : "")}>Conversations
                </button>
                <button onClick={() => {
                    setSideMenu(TEMPLATE)
                }}
                        className={"font-semibold spaces-x-3 text-lg text-gray-700 hover:bg-gray-200 p-2 rounded-md capitalize font-poppins tracking-wide " + (sideMenu === TEMPLATE ? "bg-gray-200" : "")}>Templates
                </button>
            </div>

            {menuComponents[sideMenu]}
            <div className="items-center justify-center flex flex-row bg-gray-200 rounded-md h-10">
                <input id="checkbox" type="checkbox" value="" onChange={handleContextChange}
                       className="w-4 h-4 bg-gray-100 border-gray-300 accent-gray-700 rounded"/>
                <label htmlFor="checkbox" className="text-sm text-gray-700 font-medium ml-3">Send Context</label>
            </div>
        </>
    )
}
export default SideMenu;