"use client";

import { FaRegFaceRollingEyes } from "react-icons/fa6";

export default function Note() {
  return (
    <div className="sm:rounded-3xl rounded-xl p-4 dark:bg-gray-700/70 h-full bg-gray-300/70">
      <div className="mb-2 font-medium sm:text-base text-sm">
        <FaRegFaceRollingEyes className="sm:text-2xl text-xl mb-1" />
        <span>Những điều cần lưu ý</span>
      </div>
      <div className="flex sm:flex-row flex-col items-center gap-2">
        <ol className="list-decimal list-inside sm:text-base text-sm">
          <li>Mỗi ngày bạn chỉ được tâm sự 5 lần</li>
          <li>Mỗi ngày bạn chỉ được hữu duyên đọc 10 lời tâm sự</li>
          <li>Thời gian đếm ngược 24 giờ kể từ lần đầu truy cập</li>
          <li>Đường tới tâm sự của bạn chỉ tồn tại trong 7 ngày</li>
          <li>
            Tâm sự của bạn bị những người khác phản đối 3 lần sẽ bị xóa
          </li>
        </ol>
      </div>
    </div>
  );
}
