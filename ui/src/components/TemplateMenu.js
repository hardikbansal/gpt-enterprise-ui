import {useRef, useState} from "react";
import AddNewButton from "./AddNewButton";

function TemplateMenu(props) {
    const [selectedTemplate, setSelectedTemplate] = useState(0)

    return (
        <>
            <div className="flex-wrap overflow-y-auto scroll-hide flex-grow">
                {
                    props.templates.map((template, index) => (
                        <div className="w-full flex">
                            <button onClick={() => {
                                setSelectedTemplate(index);
                                props.onTemplateSelect(template)
                            }}
                                    className={"flex flex-grow my-1 items-center space-x-3 text-gray-900 rounded-lg mx-3 px-4 py-2 font-medium capitalize hover:bg-gray-200 focus:shadow-outline " + (selectedTemplate === index ? "bg-gray-200" : "")}>


                                <span>{(template.name === "") ? "Undefined" : template.name}</span>
                            </button>
                        </div>
                    ))
                }
            </div>
            <div className="mt-5 w-full flex items-center justify-center mb-5">
                <AddNewButton onClick={() => {
                    props.addTemplateClick()
                }} name="Add Template"/>
            </div>
        </>
    );
}

export default TemplateMenu;
