"use client";

import { useTheme } from "next-themes";
import { useRouter } from "next/navigation";
import { IconType } from "react-icons";
import { IoMoonOutline, IoSunnyOutline } from "react-icons/io5";

interface ButtonProps {
  label?: string;
  icon?: IconType | "appearance";
  link?: string;
  className?: string;
}

export default function ButtonSideBar({
  label,
  icon,
  link,
  className,
}: ButtonProps) {
  const { theme, setTheme } = useTheme();
  const router = useRouter();

  const handleClick = () => {
    if (icon === "appearance") {
      setTheme(theme === "dark" ? "light" : "dark");
    } else if (link) {
      router.push(link);
    }
  };

  const RenderIcon = () => {
    if (icon === "appearance") {
      return theme === "dark" ? <IoSunnyOutline /> : <IoMoonOutline />;
    }
    if (typeof icon === "function") {
      const IconComp = icon;
      return <IconComp className="text-xl" />;
    }
    return null;
  };

  const RenderLabel = () => {
    if (!label) {
      return theme === "dark" ? "Sáng" : "Tối";
    } else {
      return label;
    }
  };

  return (
    <button
      className={`rounded-sm w-[70px] sm:w-auto justify-between cursor-pointer flex flex-col py-2 items-center gap-0.5 bg-gray-100 dark:hover:bg-gray-700 hover:bg-gray-300 dark:bg-neutral-900 transition-colors  ${className}`}
      onClick={handleClick}
    >
      {RenderIcon()}
      <span className="text-sm">{RenderLabel()}</span>
    </button>
  );
}
