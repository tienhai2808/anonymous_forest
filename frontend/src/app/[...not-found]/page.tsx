import Buddha from "@/components/icons/Buddha";

export default function NotFound() {
  return (
    <div className="min-h-screen w-full flex flex-col items-center justify-center pb-[58px] sm:pb-0">
      <div>
        <div className="sm:mx-10 px-4 flex items-center justify-center flex-row">
          <div className="flex-1">
            <p className="text-center sm:text-8xl text-6xl">404</p>
            <div className="text-center sm:text-lg underline underline-offset-4">
              Có vẻ như con đang lạc lối
            </div>
          </div>
          <Buddha />
        </div>
      </div>
    </div>
  );
}
