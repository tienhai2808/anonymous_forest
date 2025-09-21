import Mandala from "@/components/icons/Mandala";
import Donate from "@/components/ui/Donate";
import Note from "@/components/ui/Note";
import Share from "@/components/ui/Share";

export default function About() {
  return (
    <div className="py-4 relative min-h-screen sm:px-4 pb-[58px] sm:pb-4">
      <div className="flex h-full flex-col items-center justify-center">
        <div className="sm:w-full w-[90%] flex lg:flex-row flex-col justify-center gap-2 lg:gap-4 items-center">
          <Note />
          <Share />
        </div>
        <Donate />
      </div>
      <div className="absolute inset-0 -z-10 flex px-2 items-center justify-center h-full opacity-20 w-full">
        <Mandala />
      </div>
    </div>
  );
}
