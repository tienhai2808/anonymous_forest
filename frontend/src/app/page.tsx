import Feed from "@/components/Feed";
import Buddha from "@/components/icons/Buddha";
import Alert from "@/components/ui/Alert";
import { quotes } from "@/data/quotes";

export default function Home() {
  const randomQuote = quotes[Math.floor(Math.random() * quotes.length)];

  return (
    <div className="flex flex-col items-center justify-center bg-white dark:bg-black h-screen pt-4 pb-4">
      <div className="h-2/5 w-full flex flex-col justify-center items-center">
        <div className="mb-0.5">
          <Alert content={randomQuote} />
        </div>
        <div className="">
          <Buddha height="180" width="180" />
        </div>
      </div>
      <div className="h-3/5">
        <div className="my-2 h-full">
          <Feed />
        </div>
      </div>
    </div>
  );
}
