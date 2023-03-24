import {useRef} from "react";
import MarkdownView from "react-showdown";

function ChatBox(props) {
    let chatListRef = useRef(null);

    return (
        <>
            <div className="flex-grow flex flex-col-reverse items-center chat-container overflow-y-auto scroll-hide">
                {props.queries.map((query) => (
                    <>
                        <div className="w-full  prose prose-lg max-w-7xl flex flex-col">
                            <MarkdownView className="flex-grow mx-10 mb-5 rounded-lg shadow-lg px-6 text-base bg-gray-100 text-gray-900 font-medium" markdown={query.response}>
                            </MarkdownView>
                        </div>
                        <div className="w-full flex flex-col max-w-7xl">
                            <article className="flex-grow mx-10 mb-5 rounded-lg shadow-lg px-6 text-base bg-white py-6 text-gray-900 font-medium">
                                {query.query}
                            </article>
                        </div>


                    </>
                ))}
                <div ref={chatListRef}></div>
            </div>

        </>
    );
}

export default ChatBox;