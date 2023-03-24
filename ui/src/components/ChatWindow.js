import AddNewButton from "./AddNewButton";
import ChatBox from "./ChatBox";
import ChatInputBox from "./ChatInput";
import React, {useEffect, useRef, useState} from "react";
import axios from "axios";
import TemplateInput from "./TemplateInput";

function ChatWindow(props) {
    const [queries, setQueries] = useState([
        // {query: "", response: "", id: -1, createdAt: ""}
    ]);
    const  [template, setTemplate] = useState()
    const token = localStorage.getItem('user_token')
    const headers = {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
        'Authorization': 'Bearer ' + token
    }
    const QUERY = process.env.REACT_APP_API_ENDPOINT + "/api/query"
    const CONVERSATION_QUERIES = process.env.REACT_APP_API_ENDPOINT + "/api/query?conversationId="

    function openConversation() {
        axios.get(CONVERSATION_QUERIES + props.conversationId, {headers: headers})
            .then(resp => {
                console.log("Get Queries for: ", resp.data)
                setQueries(resp.data)
                console.log("Selected conv id: ", props.conversationId);
            }).catch(error => console.log(error))
    }

    function doGptQuery(query) {
        console.log("Current conversation id to do query: ", props.conversationId)
        if (props.conversationId === -1) return
        // Api calling for getting conversations list
        axios.post(QUERY, {query: query, conversation_id: props.conversationId, context: props.context}, {headers: headers},)
            .then(resp => {
                console.log("Queries: ", resp.data)
                setQueries(resp.data);
            }).catch(error => console.log(error))
    }

    useEffect(
        () => {
            openConversation()
        },
        [props.conversationId]
    )

    useEffect(
        () => {
            setTemplate(props.template)
        },
        [props.template]
    )

    return (props.conversationId === -1) ?
        (<div className="flex items-center justify-center h-full w-5/6">
            <AddNewButton onClick={() => {
                props.showAddConversationPopup()
            }} name="Add Conversation"/>
        </div>)
        :
        (<div className="flex w-5/6 flex-col">
            <ChatBox queries={queries}/>
            {
                (template != null) ?
                    (<TemplateInput onSubmit={doGptQuery} template={template} onCancel={props.onTemplateRemove}/>)
                    : (<ChatInputBox onSubmit={doGptQuery} />)
            }
        </div>)
}

export default ChatWindow;