import { Card } from "@nextui-org/react";
import { CardBody } from "@nextui-org/react";
import React from "react";

const TextOutput = ({ text }) => {
  return (
    <div className="w-full h-screen flex items-center justify-center bg-neutral-100 dark:bg-black">
      <Card className="w-1/3">
        <CardBody className="text-center">{text}</CardBody>
      </Card>
    </div>
  );
};

export default TextOutput;
