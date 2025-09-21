"use client";

import { useState } from "react";
import { IoClose } from "react-icons/io5";
import { PiHandsPraying } from "react-icons/pi";
import Image from "next/image";

export default function Donate() {
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
    <>
      <div
        className="flex justify-center dark:bg-black bg-white border-black sm:text-base text-sm items-center my-4 cursor-pointer rounded-2xl px-3 py-1 border-2 dark:border-white"
        onClick={openPopup}
      >
        <span>Hữu duyên công đức... </span>
        <PiHandsPraying />
      </div>

      {showPopup && (
        <div className="fixed inset-0 bg-black bg-opacity-50 backdrop-blur-sm flex items-center justify-center z-500 p-4">
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
    </>
  );
}
