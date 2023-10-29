import { Button } from "@nextui-org/react";
import { Input } from "@nextui-org/react";
import { useState } from "react";
import React from "react";

const ChatInput = ({ message, setMessage, onSendMessage }) => {
  const handleSubmit = (e) => {
    e.preventDefault();
    if (message !== "") {
      console.log("User message: ", message); // Handle the message here
      onSendMessage(message); // Call the callback function with the message
      setMessage("");
    }
  };

  const handleChange = (e) => {
    setMessage(e.target.value);
  };

  return (
    <div className="dark:bg-black flex gap-4 items-center">
      <form
        className="flex w-full space-x-2 bg-neutral-50 p-2 rounded-lg"
        onSubmit={handleSubmit}
      >
        <Input
          className="flex-grow dark:bg-gray-900 drop-shadow-sm"
          color="default"
          value={message}
          placeholder="Write your message..."
          onChange={handleChange}
        />
        <Button
          className="dark:bg-gray-800"
          color="default"
          size="md"
          onClick={handleSubmit}
        >
          Send
        </Button>
      </form>
    </div>
  );
};

export default ChatInput;
