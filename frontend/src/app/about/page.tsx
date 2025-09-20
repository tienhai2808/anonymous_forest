"use client";

import Mandala from "@/components/icons/Mandala";
import ButtonAbout from "@/components/ui/ButtonAbout";
import QRCodeGenerate from "@/components/ui/QRCode";
import Image from "next/image";
import { useState } from "react";
import { FaRegFaceRollingEyes, FaRegFaceSmileBeam } from "react-icons/fa6";
import { IoClose } from "react-icons/io5";
import { LuGithub, LuInstagram, LuLink2, LuShare } from "react-icons/lu";
import { PiHandsPraying } from "react-icons/pi";

export default function About() {
  const [showPopup, setShowPopup] = useState(false);

  const openPopup = () => {
    setShowPopup(true);
    document.body.style.overflow = "hidden";
  };

  const closePopup = () => {
    setShowPopup(false);
    document.body.style.overflow = "unset";
  };

  return (
    <div className="flex flex-col items-center justify-center md:justify-start py-4 bg-white h-[calc(100vh-58px)] dark:bg-black sm:h-screen sm:px-4">
      <div className="mb-4">
        <Mandala />
      </div>
      <div className="sm:w-full w-[90%] flex lg:flex-row flex-col justify-center gap-2 lg:gap-4 items-center">
        <div className="sm:rounded-3xl rounded-xl p-4 dark:bg-gray-700 h-full bg-gray-300">
          <div className="mb-2 font-medium sm:text-base text-sm">
            <FaRegFaceRollingEyes className="sm:text-2xl text-xl mb-1" />
            <span>Những điều cần lưu ý</span>
          </div>
          <div className="flex sm:flex-row flex-col items-center gap-2">
            <ol className="list-decimal list-inside sm:text-base text-sm">
              <li>Lời tâm sự sẽ được nhìn thấy nhưng không ai biết bạn</li>
              <li>Mỗi ngày bạn chỉ được tâm sự 5 lần</li>
              <li>Mỗi ngày bạn chỉ được hữu duyên đọc 10 lời tâm sự</li>
              <li>Thời gian đếm ngược 24 giờ kể từ lần đầu truy cập</li>
              <li>
                Tâm sự của bạn bị những người khác phản đối 3 lần thì sẽ bị xóa
              </li>
            </ol>
          </div>
        </div>
        <div className="sm:rounded-3xl rounded-xl p-4 dark:bg-gray-700 h-full bg-gray-300">
          <div className="mb-2 font-medium sm:text-base text-sm">
            <FaRegFaceSmileBeam className="sm:text-2xl text-xl mb-1" />
            <span>Chia sẻ Chốn An Yên tới mọi người</span>
          </div>
          <div className="flex items-center gap-2">
            <QRCodeGenerate url={window.location.href} size={125} />
            <div className="grid grid-cols-2 gap-2 p-2">
              <ButtonAbout
                label="copy"
                icon={LuLink2}
                link={window.location.href}
                isCopy={true}
              />
              <ButtonAbout
                label="share"
                icon={LuShare}
                link={window.location.href}
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
      </div>
      <div
        className="flex justify-center items-center mt-4 cursor-pointer"
        onClick={openPopup}
      >
        <span>Hữu duyên công đức... </span>
        <PiHandsPraying />
      </div>

      {showPopup && (
        <div className="fixed inset-0 bg-black bg-opacity-50 backdrop-blur-sm flex items-center justify-center z-50 p-4">
          <div className="relative bg-gray-100 dark:bg-neutral-900 rounded-2xl shadow-2xl max-w-md w-full max-h-[90vh] overflow-auto">
            <button
              onClick={closePopup}
              className="cursor-pointer absolute border-2 dark:border-white border-black sm:p-1 top-4 right-4 z-10 rounded-full p-0.5 transition-colors"
            >
              <IoClose className="text-xl text-gray-600 dark:text-gray-300" />
            </button>
            <div className="p-6">
              <div className="text-center mb-4">
                <h2 className="sm:text-xl font-semibold text-lg">
                  Hữu Duyên Công Đức
                </h2>
                <p className="text-sm text-gray-500">
                  Cảm ơn bạn đã ghé thăm Chốn An Yên
                </p>
              </div>
              <div className="flex justify-center mb-4">
                <div className="relative w-64 h-64 rounded-xl overflow-hidden">
                  <Image
                    src="/images/bank_qr.png"
                    alt="Hữu duyên công đức"
                    fill
                    className="object-cover rounded-3xl"
                    sizes="(max-width: 768px) 256px, 256px"
                  />
                </div>
              </div>
              <div className="text-center space-y-3">
                <p className="text-gray-500 text-sm leading-relaxed">
                  Công đức vô lượng, phước báu vô biên. Mong bạn luôn gặp may
                  mắn và bình an trong cuộc sống.
                </p>
              </div>
              <div className="mt-6 text-center flex justify-center">
                <button
                  onClick={closePopup}
                  className="border-2 cursor-pointer dark:border-white border-black flex items-center px-6 py-2 gap-1 rounded-full text-sm font-medium transition-colors"
                >
                  <span>Cảm ơn</span>
                  <PiHandsPraying />
                </button>
              </div>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
