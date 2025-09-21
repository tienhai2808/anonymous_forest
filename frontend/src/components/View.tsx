/* eslint-disable @typescript-eslint/no-explicit-any */
"use client";

import { getPostByLink } from "@/lib/api";
import { Post } from "@/shared/types";
import { useParams } from "next/navigation";
import { useEffect, useState } from "react";
import Buddha from "./icons/Buddha";

export default function View() {
  const params = useParams();
  const postLink = params?.id as string;

  const [post, setPost] = useState<Post | null>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);
  const [showAllComments, setShowAllComments] = useState<boolean>(false);

  useEffect(() => {
    if (!postLink) return;

    const fetchPost = async () => {
      try {
        setLoading(true);
        setError(null);
        const res = await getPostByLink(postLink);
        setPost(res.data.post);
      } catch (err: any) {
        if (err.response?.status === 404) {
          setError(null);
          return;
        }
        setError(err.response?.data?.message || err.message || "Có lỗi xảy ra");
      } finally {
        setLoading(false);
      }
    };

    fetchPost();
  }, [postLink]);

  const formatDate = (dateStr: string): string => {
    const date = new Date(dateStr);
    return `${date.getDate().toString().padStart(2, "0")}/${(
      date.getMonth() + 1
    )
      .toString()
      .padStart(2, "0")}/${date.getFullYear()} ${date
      .getHours()
      .toString()
      .padStart(2, "0")}:${date.getMinutes().toString().padStart(2, "0")}`;
  };

  const handleToggleComments = () => {
    setShowAllComments(!showAllComments);
  };

  return (
    <div className="flex h-full items-center w-full justify-center">
      {loading && (
        <div className="flex items-center justify-center py-8">
          <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-current"></div>
          <span className="ml-3">Đang hữu duyên...</span>
        </div>
      )}

      {post && !loading && !error && (
        <div className="flex flex-col lg:flex-row gap-4 lg:gap-0 items-center h-full w-full p-4 pb-0 lg:pb-4">
          <div className="lg:flex-1 w-full h-full lg:p-2">
            <div className="dark:bg-gray-700/70 bg-gray-300/70 sm:rounded-2xl rounded-xl h-full sm:p-10 p-4 flex flex-col items-center">
              <div className="w-full">
                <p className="sm:text-lg text-base font-medium pb-2 underline underline-offset-2">
                  Lời tâm sự:
                </p>
                <div
                  className="text-sm bg-gray-100/70 dark:bg-neutral-900/70 p-2 sm:p-4 rounded-lg mb-2"
                >
                  <p className="sm:text-base text-sm text-gray-700 dark:text-gray-300">
                    {post.content}
                  </p>
                  <div className="sm:text-base text-sm text-end text-gray-500">
                    {formatDate(post.created_at)}
                  </div>
                </div>
                <div className="flex gap-y-0 flex-wrap gap-x-3">
                  <div>
                    <span className="mr-1 sm:text-base text-sm font-medium">
                      - Lượt đồng cảm:
                    </span>
                    <span className="sm:text-base text-sm mb-2">
                      {post.empathy_count}
                    </span>
                  </div>
                  <div>
                    <span className="mr-1 sm:text-base text-sm font-medium">
                      - Lượt phản đối:
                    </span>
                    <span className="sm:text-base text-sm mb-2">
                      {post.protest_count}
                    </span>
                  </div>
                  <div>
                    <span className="mr-1 sm:text-base text-sm font-medium">
                      - Lượt phán xét:
                    </span>
                    <span className="sm:text-base text-sm mb-2">
                      {post.comments.length}
                    </span>
                  </div>
                </div>
              </div>
            </div>
          </div>
          {post.comments && post.comments.length > 0 && (
            <div className="lg:w-1/2 w-full h-full lg:p-2 pb-4 lg:pb-2">
              <div className="dark:bg-gray-700/70 bg-gray-300/70 sm:rounded-2xl rounded-xl h-full sm:p-10 p-4 flex flex-col items-center">
                <div className="h-full w-full pb-2 overflow-hidden">
                  <p className="sm:text-lg text-base font-medium pb-2 underline underline-offset-2">
                    Lời phán xét:
                  </p>
                  <div className="h-full lg:overflow-y-auto lg:pb-8">
                    {(showAllComments
                      ? post.comments
                      : post.comments.slice(0, 3)
                    ).map((comment) => (
                      <div
                        key={comment._id}
                        className="text-sm bg-gray-100/70 dark:bg-neutral-900/70 py-2 sm:px-4 px-2 rounded-lg mb-2"
                      >
                        <p className="text-gray-700 dark:text-gray-300">
                          {comment.content}
                        </p>
                        <div className="text-xs text-end text-gray-500">
                          {formatDate(comment.created_at)}
                        </div>
                      </div>
                    ))}
                    {post.comments.length > 3 && (
                      <p
                        onClick={handleToggleComments}
                        className="text-xs sm:text-sm cursor-pointer text-gray-500 text-center hover:text-gray-700 dark:hover:text-gray-300 transition-colors"
                      >
                        {showAllComments
                          ? "Thu gọn phán xét"
                          : `... và ${post.comments.length - 3} phán xét khác`}
                      </p>
                    )}
                  </div>
                </div>
              </div>
            </div>
          )}
        </div>
      )}

      {error && !loading && (
        <div className="mx-10 px-4 py-2 flex items-center flex-col sm:flex-row gap-3 sm:gap-0 dark:bg-gray-700/90 bg-gray-300/90 sm:rounded-2xl rounded-xl">
          <Buddha />
          <p className="text-center sm:text-lg text-red-500 underline underline-offset-4">
            Ta đang không được khỏe, con hãy quay lại sau
          </p>
        </div>
      )}

      {!post && !error && !loading && (
        <div className="mx-10 px-4 py-2 flex items-center flex-col sm:flex-row gap-3 sm:gap-0 dark:bg-gray-700/90 bg-gray-300/90 sm:rounded-2xl rounded-xl">
          <Buddha />
          <p className="text-center sm:text-lg underline underline-offset-4">
            Có vẻ như con đang đi sai đường tới tâm sự
          </p>
        </div>
      )}
    </div>
  );
}
