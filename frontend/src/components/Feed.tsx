/* eslint-disable @typescript-eslint/no-explicit-any */
"use client";

import { useEffect, useRef, useState } from "react";
import Meditation from "./icons/Meditation";
import { getRandomPost } from "@/lib/api";
import { Post } from "@/shared/types";
import { PiChatCircleDots, PiHandHeart, PiHandPalm } from "react-icons/pi";

export default function Feed() {
  const textareaRef = useRef<HTMLTextAreaElement | null>(null);
  const [post, setPost] = useState<Post | null>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);
  const [showAllComments, setShowAllComments] = useState<boolean>(false);

  useEffect(() => {
    const fetchPost = async () => {
      try {
        setLoading(true);
        setError(null);
        const res = await getRandomPost();
        setPost(res.data.post);
      } catch (err: any) {
        if (err.response?.status === 404 || err.response?.status === 429) {
          setError(null);
          return;
        }
        setError(err.response?.data?.message || err.message || "Có lỗi xảy ra");
      } finally {
        setLoading(false);
      }
    };

    fetchPost();
  }, []);

  useEffect(() => {
    const textarea = textareaRef.current;
    if (!textarea) return;

    const adjustHeight = () => {
      textarea.style.height = "auto";
      const newHeight = Math.min(textarea.scrollHeight, 120);
      textarea.style.height = `${newHeight}px`;
      textarea.style.overflowY = newHeight >= 120 ? "auto" : "hidden";
    };

    textarea.addEventListener("input", adjustHeight);

    return () => {
      textarea.removeEventListener("input", adjustHeight);
    };
  }, []);

  const handleKeyDown = (e: React.KeyboardEvent<HTMLTextAreaElement>) => {
    if (e.key === "Enter" && !e.shiftKey) {
      e.preventDefault();
      console.log("Send message:", textareaRef.current?.value);

      if (textareaRef.current) {
        textareaRef.current.value = "";
        textareaRef.current.style.height = "40px";
      }
    }
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    console.log("Submit message:", textareaRef.current?.value);
  };

  const formatDate = (dateString: string) => {
    const date = new Date(dateString);
    const now = new Date();
    const diffTime = now.getTime() - date.getTime();
    const diffDays = Math.floor(diffTime / (1000 * 60 * 60 * 24));

    if (diffDays === 0) {
      return `Hôm nay lúc ${date.toLocaleTimeString("vi-VN", {
        hour: "2-digit",
        minute: "2-digit",
      })}`;
    } else if (diffDays === 1) {
      return `Hôm qua lúc ${date.toLocaleTimeString("vi-VN", {
        hour: "2-digit",
        minute: "2-digit",
      })}`;
    } else if (diffDays < 7) {
      return `${diffDays} ngày trước`;
    } else {
      return date.toLocaleString("vi-VN");
    }
  };

  const handleToggleComments = () => {
    setShowAllComments(!showAllComments);
  };

  return (
    <div className="sm:rounded-3xl rounded-xl h-full sm:text-base text-sm sm:px-4 sm:py-4 px-2 py-2 sm:w-[500px] md:w-[600px] lg:w-[750px] w-[80%] dark:bg-black sm:border-4 border-3 dark:border-white border-black flex flex-col justify-between items-center">
      <div className="flex-1 mb-2 dark:border-white border-black sm:rounded-2xl rounded-xl sm:border-3 border-2 w-full overflow-y-auto">
        <div className="sm:p-4 p-2">
          {loading && (
            <div className="flex items-center justify-center py-8">
              <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-current"></div>
              <span className="ml-3">Đang hữu duyên...</span>
            </div>
          )}

          {error && (
            <div className="text-center py-8">
              <p className="text-red-500 mb-4 sm:text-base text-sm">{error}</p>
            </div>
          )}

          {post && !loading && !error && (
            <div className="space-y-3">
              <div className="prose dark:prose-invert max-w-none">
                <p className="sm:text-base text-sm font-medium leading-relaxed whitespace-pre-wrap">
                  {post.content}
                </p>
              </div>

              <div className="flex items-center justify-end gap-2 space-y-2">
                <div className="text-xs text-gray-500 dark:text-gray-400 my-0">
                  {formatDate(post.created_at)}
                </div>

                <div className="flex items-center gap-2 tex-base">
                  <button
                    title="Đồng cảm"
                    className="rounded-sm justify-between cursor-pointer flex items-center gap-0.5 hover:bg-gray-300 dark:hover:bg-gray-700 px-1"
                  >
                    <PiHandHeart />
                    <span>{post.empathy_count || 0}</span>
                  </button>
                  <button
                    title="Phản đối"
                    className="rounded-sm justify-between cursor-pointer flex items-center gap-0.5 hover:bg-gray-300 dark:hover:bg-gray-700 px-1"
                  >
                    <PiHandPalm />
                    <span>{post.protest_count || 0}</span>
                  </button>
                  {post.comments && (
                    <button
                      title="Nhận xét"
                      className="rounded-sm justify-between cursor-pointer flex items-center gap-0.5 hover:bg-gray-300 dark:hover:bg-gray-700 px-1"
                    >
                      <PiChatCircleDots />
                      <span>{post.comments.length || 0}</span>
                    </button>
                  )}
                </div>
              </div>

              {post.comments && post.comments.length > 0 && (
                <div className="mt-4 pt-3 border-t border-gray-200 dark:border-gray-700">
                  <p className="sm:text-base text-sm text-gray-500 mb-2">
                    Xem người ta nói gì này:
                  </p>
                  <div className="space-y-2">
                    {(showAllComments
                      ? post.comments
                      : post.comments.slice(0, 3)
                    ).map((comment) => (
                      <div
                        key={comment._id}
                        className="text-sm bg-gray-100 dark:bg-gray-900 p-2 rounded-lg"
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
                          ? "Thu gọn nhận xét"
                          : `... và ${post.comments.length - 3} nhận xét khác`}
                      </p>
                    )}
                  </div>
                </div>
              )}
            </div>
          )}

          {!post && !loading && !error && (
            <div className="text-center py-8 text-gray-500">
              <p>
                Hôm nay hữu duyên tới đây thôi, hãy dành thời gian cho bản thân
              </p>
            </div>
          )}
        </div>
      </div>
      <form
        onSubmit={handleSubmit}
        className="h-auto flex gap-1 w-full items-end min-h-[40px]"
      >
        <textarea
          ref={textareaRef}
          className="sm:px-3 px-2 py-2 text-sm sm:text-base dark:border-white border-black border-2 sm:border-3 sm:rounded-2xl rounded-xl w-full resize-none outline-none dark:bg-black dark:text-white"
          placeholder="Nói cho ta nghe..."
          rows={1}
          style={{
            height: "40px",
            minHeight: "40px",
            maxHeight: "120px",
            overflowY: "hidden",
            lineHeight: "20px",
          }}
          onKeyDown={handleKeyDown}
        />
        <button
          type="submit"
          className="px-1 sm:px-2 text-sm sm:text-base cursor-pointer sm:rounded-2xl rounded-xl dark:border-white border-black border-2 sm:border-3 flex items-center justify-center"
          style={{ minHeight: "40px" }}
        >
          <Meditation />
        </button>
      </form>
    </div>
  );
}
