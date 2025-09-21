import Logo from "@/components/icons/Logo";
import View from "@/components/View";

export default function PostDetails() {
  return (
    <div className="relative min-h-screen w-full flex flex-col items-center justify-center pb-[58px] sm:pb-0">
      <View />
      <div className="absolute inset-0 -z-10 flex px-4 items-center justify-center h-full opacity-20 w-full">
        <Logo className="w-150 h-150 lg:h-160 lg:w-160"/>
      </div>
    </div>
  );
}
