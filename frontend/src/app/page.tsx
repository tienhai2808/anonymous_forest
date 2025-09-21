import Feed from "@/components/Feed";
import Buddha from "@/components/icons/Buddha";
import Notification from "@/components/ui/Notification";
import { welcome } from "@/lib/quotes";

export default function Home() {
  const randomQuote = welcome[Math.floor(Math.random() * welcome.length)];

  return (
    <div className="flex flex-col items-center justify-center sm:py-4 py-2 bg-white dark:bg-black min-h-screen sm:px-4 pb-[58px] sm:pb-4">
      <div className=" sm:min-h-auto min-h-[20vh] h-auto w-full flex flex-col items-center justify-end">
        <div className="mb-0.5">
          <Notification content={randomQuote} />
        </div>
        <div className="">
          <Buddha />
        </div>
      </div>
      <div className="sm:max-h-auto max-h-[60vh] w-full">
        <div className="h-full w-full flex justify-center overflow-hidden">
          <Feed />
        </div>
      </div>
    </div>
  );
}
