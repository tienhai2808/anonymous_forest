"use client";

import { FaRegFaceSmileBeam } from "react-icons/fa6";
import ButtonAbout from "./ButtonAbout";
import QRCodeGenerate from "./QRCode";
import { LuGithub, LuInstagram, LuLink2, LuShare } from "react-icons/lu";

export default function Share() {
  return (
    <div className="sm:rounded-3xl rounded-xl p-4 dark:bg-gray-700/70 h-full bg-gray-300/70">
      <div className="mb-2 font-medium sm:text-base text-sm">
        <FaRegFaceSmileBeam className="sm:text-2xl text-xl mb-1" />
        <span>Chia sẻ Chốn An Yên tới mọi người</span>
      </div>
      <div className="flex items-center gap-2">
        <QRCodeGenerate url={window.location.origin} size={125} />
        <div className="grid grid-cols-2 gap-2 p-2">
          <ButtonAbout
            label="copy"
            icon={LuLink2}
            link={window.location.origin}
            isCopy={true}
          />
          <ButtonAbout
            label="share"
            icon={LuShare}
            link={window.location.origin}
            isShare={true}
          />
          <ButtonAbout
            label="start"
            icon={LuGithub}
            link="https://github.com/tienhai2808/anonymous_forest"
          />
          <ButtonAbout
            label="follow"
            icon={LuInstagram}
            link="https://www.instagram.com/haict_08/"
          />
        </div>
      </div>
    </div>
  );
}
