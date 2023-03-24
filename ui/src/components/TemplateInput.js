import {createRef, useEffect, useRef, useState} from "react";
import SendLogo from "../assets.images/send.png"

function TemplateInputBox(props) {

    var inputValues = useRef(props.template.params.map(() => createRef()));
    var [query, setQuery] = useState("");

    function onSendClick() {
        console.log(inputValues.current);
        const queryString = getQueryString();
        console.log(queryString)
        inputValues.current.map((ref)=>{
            ref.current.value = ""
        })
        setQuery(getQueryString());
        props.onSubmit(query);
    }


    useEffect(
        () =>{
            console.log("chat window selected template", props.template)
        },
        []
    )

    useEffect(
        () =>{
            setQuery(getQueryString());
        },
        [props.template]
    )


    function getQueryString() {
        var tempString = ""
        for (let i = 0; i < props.template.params.length; i++) {
            tempString += props.template.parts[i]
            tempString += (inputValues.current[i].current.value !== "") ? inputValues.current[i].current.value : props.template.params[i]
        }
        tempString += props.template.parts[props.template.params.length]
        return tempString
    }

    return (
        <div className="w-full flex flex-col my-5  items-center px-5">
            <div className="w-full max-w-7xl flex items-center  shadow-lg rounded text-sm bg-gray-50 p-3">
                <div className="flex flex-col flex-grow">
                    {
                        props.template.params.map((param, index) => (
                            <textarea key={"query_" + param}
                                      ref={inputValues.current[index]}
                                      onChange={()=>{setQuery(getQueryString())}}
                                      className="m-1 block w-full p-4 rounded-md border border-gray-200 bg-gray-100 placeholder-gray-500 text-black text-base"
                                      placeholder={"Insert "+param+" here"}/>
                        ))
                    }

                </div>
                <img src={SendLogo} width="40" className="block h-10 stroke-1 ml-6" onClick={onSendClick} alt="send logo"/>
            </div>
            <div className="flex flex-row items-center w-full max-w-7xl bg-gray-50 p-6 mb-5 shadow-lg rounded-lg">
                <p className="flex-grow font-medium">{query}</p>
                <button onClick={() => {props.onCancel()}} className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
                    Close Template
                </button>
            </div>
        </div>


    );
}

export default TemplateInputBox;