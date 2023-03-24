import React, {useEffect, useRef, useState} from "react";
import MarkdownView from 'react-showdown';
import "./Chat.css";
import ConversationMenu from "./ConversationMenu";
import axios from 'axios';
import ChatInputBox from "./ChatInput";
import ChatBox from "./ChatBox";
import AddNewButton from "./AddNewButton";
import AddConversationPopup from "./AddConversationPopup";
import SideMenu from "./SideMenu";
import AddTemplatePopup from "./AddTemplatePopup";
import ChatWindow from "./ChatWindow";

function Conversation() {
    const token = localStorage.getItem('user_token')
    const headers = {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
        'Authorization': 'Bearer ' + token
    }

    const sample_temps = [{
        "id": 1,
        "name": "Sample Template",
        "template": {
            parts: ["Write a short mail regarding ", ". Make is short a sweet", ""],
            params: ["mail_topic", "email_topic"]
        }
    }]
    const [context, setContext] = useState(false)

    const CONVERSATION_URL = process.env.REACT_APP_API_ENDPOINT + "/api/conversation"
    const TEMPLATE_URL = process.env.REACT_APP_API_ENDPOINT + "/api/template"

    const [conversations, setConversations] = useState([]);
    const [templates, setTemplates] = useState([]);

    const [selectedTemplate, setSelectedTemplate] = useState(null);

    const [conversationId, setConversationId] = useState(-1)
    const [showAddConversationPopup, setShowAddConversationPopup] = useState(false);
    const [showAddTemplatePopup, setShowAddTemplatePopup] = useState(false);

    function getConversations() {
        axios
            .get(CONVERSATION_URL, {
                headers: headers
            })
            .then((res) => {
                console.log("Conversations: ", res.data)
                setConversations(res.data)
                console.log("Selected Conv: ", res.data[0].id)
                openConversation(res.data[0])
            })
            .catch((err) => {
                console.log(err);
            });
    }

    function getTemplates() {
        axios
            .get(TEMPLATE_URL, {
                headers: headers
            })
            .then((res) => {
                console.log("Templates: ", res.data)
                setTemplates(res.data)
            })
            .catch((err) => {
                console.log(err);
            });
    }

    useEffect(
        () => {
            getConversations()
            getTemplates()
        },
        []
    );

    function openConversation(conversationData) {
        setConversationId(conversationData.id)
    }

    function createConversation(conversationName) {
        if (conversationName === "") return
        setShowAddConversationPopup(false)
        axios.post(CONVERSATION_URL,
            {name: conversationName},
            {headers: headers}
        )
            .then(resp => {
                setConversations(resp.data)
            })
            .catch(error => console.log(error))
    }


    function templateFormat(template){
        const regex1 = /{[\w\d_]+}/gm;
        console.log("Template for format:", template);
        const iter = template.matchAll(regex1);
        const parts= template.split(regex1);
        const params = []
        for (let n = iter.next(); !n.done; n = iter.next()) {
            params.push(n.value[0])
        }

        return {
            "params" : params,
            "parts": parts
        }

    }

    function createTemplate(templateName, template) {
        if (templateName === "" || template === "") return
        setShowAddTemplatePopup(false)
        const formattedTemplate = templateFormat(template)
        console.log("Formatted template: " ,formattedTemplate);
        axios.post(TEMPLATE_URL,
            {name: templateName, parts: formattedTemplate.parts, params: formattedTemplate.params },
            {headers: headers}
        )
            .then(resp => {
                setTemplates(resp.data)
            })
            .catch(error => {
                console.log(error)
            })
    }



    let conversationMenuProps = {
        conversations: conversations,
        conversationClick: openConversation,
        addConversationClick: () => {
            setShowAddConversationPopup(true)
        },

    }


    let templateMenuProps = {
        templates: templates,
        addTemplateClick: () => {
            setShowAddTemplatePopup(true)
        },
        onTemplateSelect: (temp) => {
            console.log("Setting selected template: ", temp)
            setSelectedTemplate(temp)
        },
    }

    return (
        <div className="flex flex-row h-screen bg-white">

            <div className="flex w-1/6 h-full flex-col shadow bg-gray-50">
                <SideMenu setContext={setContext} templateProps={templateMenuProps} conversationMenuProps={conversationMenuProps}/>
            </div>

            <ChatWindow showAddConversationPopup={() => {
                setShowAddConversationPopup(true)
            }}
                        conversationId={conversationId}
                        template={selectedTemplate}
                        context={context}
                        onTemplateRemove={() => {
                            setSelectedTemplate(null)
                        }}
            />

            {
                showAddConversationPopup ? (<AddConversationPopup onClick={createConversation}
                                                                  onCancel={() => {
                                                                      setShowAddConversationPopup(false)
                                                                  }}
                />) : null
            }

            {
                showAddTemplatePopup ? (<AddTemplatePopup onClick={createTemplate} onCancel={() => {
                    setShowAddTemplatePopup(false)
                }}
                />) : null
            }

        </div>
    );
}


export default Conversation;
