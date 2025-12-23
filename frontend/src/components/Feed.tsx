/* eslint-disable @typescript-eslint/no-explicit-any */
"use client";

import { useEffect, useRef, useState } from "react";
import Meditation from "./icons/Meditation";
import {
  getRandomPost,
  createPost,
  addEmpathy,
  addProtest,
  addComment,
} from "@/lib/api";
import { Post } from "@/shared/types";
import {
  PiChatCircleDots,
  PiHandHeart,
  PiHandPalm,
  PiCopy,
  PiCheck,
  PiX,
  PiSmileySad,
  PiSmiley,
  PiSmileyWink,
  PiArrowClockwise,
} from "react-icons/pi";
import { response } from "@/lib/constants";

export default function Feed() {
  const textareaRef = useRef<HTMLTextAreaElement | null>(null);
  const [post, setPost] = useState<Post | null>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);
  const [showAllComments, setShowAllComments] = useState<boolean>(false);
  const [responseMessage, setResponseMessage] = useState<string | null>(null);
  const [countdown, setCountdown] = useState<number | null>(null);
  const [submitting, setSubmitting] = useState<boolean>(false);
  const [showConfirmDialog, setShowConfirmDialog] = useState<boolean>(false);
  const [pendingContent, setPendingContent] = useState<string>("");
  const [postLink, setPostLink] = useState<string>("");
  const [linkCopied, setLinkCopied] = useState<boolean>(false);
  const [isCommentMode, setIsCommentMode] = useState<boolean>(false);

  const origin = process.env.NEXT_PUBLIC_BASE_URL!;

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

  const adjustHeight = () => {
    const textarea = textareaRef.current;
    if (!textarea) return;

    textarea.style.height = "auto";
    const newHeight = Math.min(textarea.scrollHeight, 120);
    textarea.style.height = `${newHeight}px`;
    textarea.style.overflowY = newHeight >= 120 ? "auto" : "hidden";
  };

  useEffect(() => {
    const textarea = textareaRef.current;
    if (!textarea) return;

    textarea.addEventListener("input", adjustHeight);

    return () => {
      textarea.removeEventListener("input", adjustHeight);
    };
  }, []);

  const getRandomResponseMessage = () => {
    const randomIndex = Math.floor(Math.random() * response.length);
    return response[randomIndex];
  };

  const startCountdown = (hasLink: boolean = false) => {
    const countdownTime = hasLink ? 10 : 5;
    setCountdown(countdownTime);
    const timer = setInterval(() => {
      setCountdown((prev) => {
        if (prev === 1) {
          clearInterval(timer);
          fetchNewPost();
          return null;
        }
        return prev ? prev - 1 : null;
      });
    }, 1000);
  };

  const fetchNewPost = async () => {
    try {
      setLoading(true);
      setError(null);
      setResponseMessage(null);
      setPostLink("");
      setIsCommentMode(false);
      const res = await getRandomPost();
      setPost(res.data.post);
      if (textareaRef.current) {
        setTimeout(adjustHeight, 0);
      }
    } catch (err: any) {
      if (err.response?.status === 404 || err.response?.status === 429) {
        setError(null);
        setPost(null);
        return;
      }
      setError(err.response?.data?.message || err.message || "Có lỗi xảy ra");
    } finally {
      setLoading(false);
    }
  };

  const copyToClipboard = () => {
    if (postLink) {
      navigator.clipboard.writeText(`${origin}/views/${postLink}`);
      setLinkCopied(true);
      setTimeout(() => setLinkCopied(false), 2000);
    }
  };

  const submitPost = async (content: string, getLinkValue: boolean = false) => {
    if (!content.trim()) return;

    try {
      setSubmitting(true);
      setError(null);
      const res = await createPost({ content, get_link: getLinkValue });
      setResponseMessage(getRandomResponseMessage());
      if (res.data && res.data.post_link) {
        setPostLink(res.data.post_link);
        startCountdown(true);
      } else {
        startCountdown(false);
      }
      if (textareaRef.current) {
        textareaRef.current.value = "";
        textareaRef.current.style.height = "40px";
        setTimeout(adjustHeight, 0);
      }
    } catch (err: any) {
      if (err.response?.status === 404 || err.response?.status === 429) {
        setError(null);
        return;
      }
      setError(
        err.response?.data?.message ||
          err.message ||
          "Có lỗi xảy ra khi gửi bài viết"
      );
    } finally {
      setSubmitting(false);
      setShowConfirmDialog(false);
    }
  };

  const submitComment = async (content: string) => {
    if (!content.trim() || !post) return;

    try {
      setSubmitting(true);
      setError(null);
      await addComment(post._id, { content });
      setResponseMessage(getRandomResponseMessage());
      startCountdown(false);
      if (textareaRef.current) {
        textareaRef.current.value = "";
        textareaRef.current.style.height = "40px";
        setTimeout(adjustHeight, 0);
      }
    } catch (err: any) {
      setError(
        err.response?.data?.message ||
          err.message ||
          "Có lỗi xảy ra khi gửi nhận xét"
      );
    } finally {
      setSubmitting(false);
    }
  };

  const handleAddEmotion = async (postId: string, isEmpathy: boolean) => {
    try {
      setSubmitting(true);
      setError(null);
      await (isEmpathy ? addEmpathy(postId) : addProtest(postId));
      setResponseMessage(getRandomResponseMessage());
      startCountdown(false);
    } catch (err: any) {
      setError(
        err.response?.data?.message ||
          err.message ||
          "Có lỗi xảy ra khi bày tỏ cảm xúc với bài viết"
      );
    } finally {
      setSubmitting(false);
    }
  };

  const handleConfirmSubmit = (getLinkValue: boolean) => {
    submitPost(pendingContent, getLinkValue);
  };

  const handleSkipPost = async () => {
    try {
      setLoading(true);
      setError(null);
      setResponseMessage(null);
      setPostLink("");
      setIsCommentMode(false);
      setShowAllComments(false);
      const res = await getRandomPost();
      setPost(res.data.post);
      if (textareaRef.current) {
        setTimeout(adjustHeight, 0);
      }
    } catch (err: any) {
      if (err.response?.status === 404 || err.response?.status === 429) {
        setError(null);
        setPost(null);
        return;
      }
      setError(err.response?.data?.message || err.message || "Có lỗi xảy ra");
    } finally {
      setLoading(false);
    }
  };

  const handleCommentClick = () => {
    setIsCommentMode(true);
    setTimeout(() => {
      textareaRef.current?.focus();
    }, 100);
  };

  const handleCancelComment = () => {
    setIsCommentMode(false);
    if (textareaRef.current) {
      textareaRef.current.value = "";
      textareaRef.current.style.height = "40px";
      setTimeout(adjustHeight, 0);
    }
  };

  const handleKeyDown = (e: React.KeyboardEvent<HTMLTextAreaElement>) => {
    if (e.key === "Enter" && !e.shiftKey) {
      e.preventDefault();
      const content = textareaRef.current?.value || "";
      if (content.trim()) {
        if (isCommentMode) {
          submitComment(content);
        } else {
          setPendingContent(content);
          setShowConfirmDialog(true);
        }
      }
    }
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    const content = textareaRef.current?.value || "";
    if (content.trim()) {
      if (isCommentMode) {
        submitComment(content);
      } else {
        setPendingContent(content);
        setShowConfirmDialog(true);
      }
    }
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
    <div className="sm:rounded-3xl rounded-xl h-full sm:text-base text-sm smp-4 p-2 sm:w-[500px] md:w-[600px] lg:w-[750px] w-[80%] dark:bg-black sm:border-4 border-3 dark:border-white border-black flex flex-col justify-between items-center">
      <div className="flex-1 mb-2 dark:border-white border-black sm:rounded-2xl rounded-xl sm:border-3 border-2 w-full overflow-y-auto">
        <div className="p-4">
          {showConfirmDialog && (
            <div className="text-center py-8 space-y-4">
              <h3 className="text-sm sm:text-base font-medium mb-4">
                Con có muốn nhìn lại tâm sự của mình không?
              </h3>
              <div className="flex justify-center gap-3">
                <button
                  onClick={() => handleConfirmSubmit(false)}
                  className="px-2 flex items-center justify-center gap-0.5 rounded-md cursor-pointer border-1 sm:border-2 border-none"
                >
                  <PiX />
                  Không
                </button>
                <button
                  onClick={() => handleConfirmSubmit(true)}
                  className="px-2 flex items-center justify-center gap-0.5 rounded-md cursor-pointer border-1 sm:border-2 dark:border-white border-black"
                >
                  <PiCheck />
                  Có
                </button>
              </div>
            </div>
          )}

          {loading && !showConfirmDialog && (
            <div className="flex items-center justify-center py-8">
              <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-current"></div>
              <span className="ml-3">Đang hữu duyên...</span>
            </div>
          )}

          {error && !showConfirmDialog && (
            <div className="text-center flex flex-col justify-center items-center py-8 text-red-500">
              <PiSmileySad className="mb-2 text-5xl sm:text-6xl" />
              <p className="sm:text-base text-sm">
                Ta đang không được khỏe, con hãy quay lại sau
              </p>
            </div>
          )}

          {responseMessage && !loading && !error && !showConfirmDialog && (
            <div className="text-center flex flex-col items-center py-4 space-y-4">
              {countdown !== null && (
                <div className="flex flex-col items-center text-gray-500">
                  <PiSmiley size={64} className="mb-2" />
                  <p className="sm:text-base text-sm">
                    {responseMessage}
                    {Array(countdown).fill(".").join("")}
                  </p>
                </div>
              )}
              {postLink && (
                <div className="mt-4 p-3 bg-gray-300 dark:bg-gray-700 rounded-lg w-full">
                  <p className="sm:text-sm text-xs mb-2">
                    Con đường tới tâm sự của con:
                  </p>
                  <div className="flex items-center gap-2 bg-gray-100 dark:bg-neutral-900 p-2 rounded border border-black dark:border-white">
                    <input
                      type="text"
                      value={`${origin}/views/${postLink}`}
                      readOnly
                      className="flex-1 bg-transparent outline-none text-sm"
                    />
                    <button
                      onClick={copyToClipboard}
                      className="p-1 hover:bg-gray-300 dark:hover:bg-gray-700 rounded cursor-pointer"
                      title="Sao chép"
                    >
                      <PiCopy size={18} />
                    </button>
                  </div>
                  {linkCopied && <p className="text-xs mt-1">Đã sao chép!</p>}
                </div>
              )}
            </div>
          )}

          {post &&
            !loading &&
            !error &&
            !responseMessage &&
            !showConfirmDialog && (
              <div className="space-y-3">
                <div className="prose dark:prose-invert max-w-none">
                  <p className="sm:text-base text-sm font-medium leading-relaxed whitespace-pre-wrap">
                    {post.content}
                  </p>
                </div>

                <div className="flex items-center justify-between sm:gap-0 gap-2 space-y-2">
                  <div className="flex items-center tex-base mb-0">
                    <button
                      title="Chuyển bài"
                      onClick={handleSkipPost}
                      disabled={loading || submitting || countdown !== null}
                      className="rounded-sm py-1 justify-between cursor-pointer flex items-center gap-0.5 hover:bg-gray-300 dark:hover:bg-gray-700 px-1"
                    >
                      <PiArrowClockwise />
                    </button>
                  </div>
                  <div className="flex items-center justify-end sm:gap-2 gap-1 space-y-2">
                    <div className="text-xs text-gray-500 dark:text-gray-400 my-0 text-right">
                      {formatDate(post.created_at)}
                    </div>
                    <div className="flex items-center sm:gap-2 tex-base">
                      <button
                        onClick={() => handleAddEmotion(post._id, true)}
                        title="Đồng cảm"
                        className="rounded-sm justify-between cursor-pointer flex items-center gap-0.5 hover:bg-gray-300 dark:hover:bg-gray-700 px-1"
                      >
                        <PiHandHeart />
                        <span>{post.empathy_count || 0}</span>
                      </button>
                      <button
                        onClick={() => handleAddEmotion(post._id, false)}
                        title="Phản đối"
                        className="rounded-sm justify-between cursor-pointer flex items-center gap-0.5 hover:bg-gray-300 dark:hover:bg-gray-700 px-1"
                      >
                        <PiHandPalm />
                        <span>{post.protest_count || 0}</span>
                      </button>
                      {post.comments && (
                        <button
                          onClick={handleCommentClick}
                          title="Nhận xét"
                          className="rounded-sm justify-between cursor-pointer flex items-center gap-0.5 hover:bg-gray-300 dark:hover:bg-gray-700 px-1"
                        >
                          <PiChatCircleDots />
                          <span>{post.comments.length || 0}</span>
                        </button>
                      )}
                    </div>
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
                          className="text-sm bg-gray-100 dark:bg-neutral-900 p-2 rounded-lg"
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
                            : `... và ${
                                post.comments.length - 3
                              } nhận xét khác`}
                        </p>
                      )}
                    </div>
                  </div>
                )}
              </div>
            )}

          {!post &&
            !loading &&
            !error &&
            !responseMessage &&
            !showConfirmDialog && (
              <div className="flex flex-col items-center py-8 text-gray-500">
                <PiSmileyWink className="mb-2 text-5xl sm:text-6xl" />
                <p className="sm:text-base text-sm">
                  Hôm nay hữu duyên tới đây thôi, hãy dành thời gian cho bản
                  thân
                </p>
              </div>
            )}
        </div>
      </div>

      {isCommentMode && (
        <div className="w-full flex items-center justify-between mb-2 px-2">
          <span className="text-xs sm:text-sm">Nhận xét lời tâm sự</span>
          <button
            onClick={handleCancelComment}
            className="text-xs sm:text-sm hover:bg-gray-300 dark:hover:bg-gray-700 p-1 rounded cursor-pointer"
            title="Hủy nhận xét"
          >
            <PiX />
          </button>
        </div>
      )}

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
          disabled={submitting || countdown !== null || showConfirmDialog}
        />
        <button
          type="submit"
          className="px-1 sm:px-2 text-sm sm:text-base cursor-pointer sm:rounded-2xl rounded-xl dark:border-white border-black border-2 sm:border-3 flex items-center justify-center"
          style={{ minHeight: "40px" }}
          disabled={submitting || countdown !== null || showConfirmDialog}
        >
          <Meditation />
        </button>
      </form>
    </div>
  );
}
