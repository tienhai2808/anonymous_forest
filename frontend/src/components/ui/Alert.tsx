"use client";

import React from "react";

interface Message {
  content: string;
}

export default function Alert({ content }: Message) {
  return (
    <div className="rounded-3xl px-4 py-2 w-[350px] h-[110px] font-medium dark:bg-black border-6 dark:border-white border-black flex justify-center items-center">
      {content}
    </div>
  );
}
