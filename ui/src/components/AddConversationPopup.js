import {useRef} from "react";

function AddConversationPopup(props) {
    const conversationNameRef = useRef("");

    return (
        <div className="fixed inset-0 bg-gray-500 bg-opacity-75 flex items-center justify-center">
            <div className="bg-white p-6 w-1/6 rounded-lg flex flex-col items-center">
                <h2 className="text-xl font-bold mb-4">New conversation</h2>
                <p className="mb-4">Enter conversation name:</p>
                <input ref={conversationNameRef}
                       className="w-full border border-gray-300 rounded-lg mb-4 px-3 py-2" placeholder="Name"/>
                <div className="flex">
                    <button
                        onClick={() => {props.onCancel()}}
                        className="bg-gray-300 hover:bg-gray-400 text-gray-700 font-bold py-2 px-4 rounded mr-2">Cancel
                    </button>
                    <button onClick={() => {
                        props.onClick(conversationNameRef.current.value)
                    }} className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"> Create
                    </button>
                </div>

            </div>
        </div>
    )
}

export default AddConversationPopup;