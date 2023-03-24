import {useRef} from "react";

function AddTemplatePopup(props) {
    const templateNameRef = useRef("");
    const templateRef = useRef("");




    return (
        <div className="fixed inset-0 bg-gray-500 bg-opacity-75 flex items-center justify-center">
            <div className="bg-white p-6 w-3/6 rounded-lg flex flex-col items-center">
                <h2 className="text-xl font-bold mb-4">New Template</h2>
                <p className="mb-4">Enter template name:</p>
                <input ref={templateNameRef}
                       className="w-full border border-gray-300 rounded-lg mb-4 px-3 py-2" placeholder="Name"/>
                <label className="mb-4" htmlFor="template">Enter template here</label>
                <textarea id="template"  ref={templateRef}
                       className="w-full border border-gray-300 rounded-lg mb-4 px-3 py-2" placeholder="Info {param}"/>
                <div className="flex">
                    <button
                        onClick={() => {props.onCancel()}}
                        className="bg-gray-300 hover:bg-gray-400 text-gray-700 font-bold py-2 px-4 rounded mr-2">Cancel
                    </button>
                    <button onClick={() => {
                        props.onClick(templateNameRef.current.value, templateRef.current.value)
                    }} className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"> Create
                    </button>
                </div>

            </div>
        </div>
    )
}

export default AddTemplatePopup;