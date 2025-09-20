"use client";

import React from "react";
import Logo from "./icons/Logo";
import { useRouter } from "next/navigation";
import ButtonSideBar from "./ui/ButtonSideBar";
import { IoInformationCircleOutline } from "react-icons/io5";

export default function SideBar() {
  const router = useRouter();

  return (
    <div className="sm:static fixed bottom-0 left-0 right-0 sm:w-[88px] w-full h-auto sm:h-screen flex sm:flex-col items-center flex-row bg-gray-100 dark:bg-neutral-900">
      <div
        className="flex justify-center sm:px-5 sm:pt-5 sm:pb-4 p-2 pl-5 sm:pl-0 cursor-pointer"
        onClick={() => {
          router.push("/");
        }}
      >
        <Logo />
      </div>
      <div className="flex sm:flex-col flex-row sm:m-1 flex-1 justify-end pr-5 sm:pr-0 h-full w-full sm:h-auto">
        <div className="sm:gap-1 gap-4 flex sm:flex-col flex-row sm:m-1 m-0">
          <ButtonSideBar label="Thêm" icon={IoInformationCircleOutline} link="/about" />
          <ButtonSideBar icon="appearance" />
        </div>
      </div>
    </div>
  );
}
