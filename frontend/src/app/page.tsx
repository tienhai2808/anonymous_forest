import Feed from "@/components/Feed";
import Buddha from "@/components/icons/Buddha";
import Notification from "@/components/ui/Notification";
import { welcome } from "@/lib/quotes";

export default function Home() {
  const randomQuote = welcome[Math.floor(Math.random() * welcome.length)];

  return (
    <div className="flex flex-col sm:items-center items-start sm:justify-center justify-start py-4 bg-white h-[calc(100vh-58px)] dark:bg-black sm:h-screen  sm:px-4">
      <div className="h-2/5 w-full flex flex-col items-center justify-end">
        <div className="mb-0.5">
          <Notification content={randomQuote} />
        </div>
        <div className="">
          <Buddha />
        </div>
      </div>
      <div className="h-3/5 w-full">
        <div className="h-full w-full flex justify-center overflow-hidden">
          <Feed />
        </div>
      </div>
    </div>
  );
}
