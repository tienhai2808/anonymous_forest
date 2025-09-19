"use client";

interface Message {
  content: string;
}

export default function Alert({ content }: Message) {
  return (
    <div className="sm:rounded-3xl rounded-xl sm:text-[16px] text-sm sm:px-4 sm:py-2 px-3 py-1 sm:w-[350px] h-[110px] w-[250px] sm:font-medium font-normal dark:bg-black sm:border-4 border-3 dark:border-white border-black flex justify-center items-center">
      {content}
    </div>
  );
}
