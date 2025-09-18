"use client";

import React from "react";
import Logo from "./icons/Logo";
import Button from "./ui/Button";
import { CiHeart } from "react-icons/ci";
import { useRouter } from "next/navigation";

export default function SideBar() {
  const router = useRouter();

  return (
    <div className="w-[88px] h-screen flex flex-col bg-gray-100 dark:bg-neutral-900">
      <div className="flex justify-center px-5 pt-5 pb-4" onClick={() => {router.push("/")}}>
        <Logo height="60" width="60" />
      </div>
      <div className="flex flex-col m-1 flex-1 justify-end">
        <div className="gap-1 flex flex-col">
          <Button label="Donate" icon={CiHeart} link="/donate" />
          <Button icon="appearance" />
        </div>
      </div>
    </div>
  );
}
