"use client";

import { useState } from "react";
import { IconType } from "react-icons";
import { LuCheck } from "react-icons/lu";

interface ButtonAboutProps {
  label: string;
  icon: IconType;
  link?: string;
  isCopy?: boolean;
  isShare?: boolean;
}

export default function ButtonAbout({
  label,
  icon,
  link,
  isCopy,
  isShare,
}: ButtonAboutProps) {
  const [copied, setCopied] = useState<boolean>(false);

  const handleClick = async () => {
    try {
      if (isCopy && link) {
        await navigator.clipboard.writeText(link);
        setCopied(true);
        setTimeout(() => setCopied(false), 1500);
      } else if (isShare && link) {
        if (navigator.share) {
          await navigator.share({
            title: link,
            url: link,
          });
        } else {
          await navigator.clipboard.writeText(link);
          alert("Trình duyệt không hỗ trợ chia sẻ. Liên kết đã được sao chép!");
        }
      } else if (link) {
        window.open(link, "_blank", "noopener,noreferrer");
      }
    } catch (error) {
      console.error("Lỗi xử lý:", error);
    }
  };

  const RenderIcon = () => {
    const IconComp = icon;
    return <IconComp className="text-xl font-medium" />;
  };
  return (
    <button
      onClick={handleClick}
      className="rounded-md px-5 justify-between cursor-pointer flex flex-col py-2 items-center gap-0.5 bg-white dark:hover:bg-black hover:bg-gray-100 dark:bg-neutral-900 transition-colors"
    >
      {isCopy && copied ? <LuCheck className="text-xl font-medium"/>: RenderIcon()}
      <span className="text-sm font-medium">
        {isCopy && copied ? "Copied" : label}
      </span>
    </button>
  );
}
