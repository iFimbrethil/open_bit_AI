import ChatInput from "./components/ChatInput";
import TextOutput from "./components/TextOutput";
import HeaderNavbar from "./components/Header";
import { NextUIProvider } from "@nextui-org/react";
import { useState } from "react";

function App() {
  const [inputText, setInputText] = useState(""); // State to store the input text

  const handleMessage = (message) => {
    setInputText(message); // Update the input text when a message is sent
  };

  return (
    <NextUIProvider>
      <div className="h-screen w-screen flex flex-col justify-end items-center relative bg-neutral-100">
        <HeaderNavbar className="w-full" />
        <TextOutput text={inputText} />
        <div className="mt-auto mb-5 w-1/3">
          <ChatInput
            message={inputText}
            setMessage={setInputText}
            onSendMessage={handleMessage}
          />
        </div>
      </div>
    </NextUIProvider>
  );
}

export default App;
