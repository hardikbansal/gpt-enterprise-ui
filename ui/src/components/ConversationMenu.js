const ConversationMenu = (props) => {
    const chats = ["Dashboard"];
    return (
      <div class="w-full">
          <div class="flex items-center space-x-4 p-2">
              <button class="font-semibold text-lg text-gray-700 hover:bg-gray-200 rounded-md p-2 capitalize font-poppins tracking-wide">Chats</button>
              <button class="font-semibold spaces-x-3 text-lg text-gray-700 hover:bg-gray-200 p-2 rounded-md capitalize font-poppins tracking-wide">Prompts</button>
          </div>
          <div class="mt-5 text-sm">
            <a href="#" class="flex items-center space-x-3 text-gray-700 p-2 rounded-md font-medium hover:bg-gray-200 bg-gray-200 focus:shadow-outline">
              <span>➕ Add new</span>
            </a>
          </div>
          {chats.map((chat) => (
            <div class="mt-3">
                <a href="#" class="flex items-center space-x-3 text-gray-700 p-2 rounded-md font-medium hover:bg-gray-200 bg-gray-200 focus:shadow-outline">
                  <span>{chat}</span>
                </a>
            </div>
          ))}
          <div class="flex items-center space-x-3 mb-4 p-2 bg-gray-200 rounded-md absolute bottom-0">
            <input id="checkbox" type="checkbox" value="" class="w-4 h-4 bg-gray-100 border-gray-300 accent-gray-700 rounded focus:"/>
            <label for="checkbox" class="text-sm text-gray-700 font-medium ml-3">Send Context</label>
          </div>
      </div>
    );
  };
  
  export default ConversationMenu;
  