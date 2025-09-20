import {
  ApiResponse,
  CreatePostCommentData,
  CreatePostData,
  PostCreatedResponse,
  PostResponse,
} from "@/shared/types";
import axios, {
  AxiosInstance,
  AxiosResponse,
  InternalAxiosRequestConfig,
} from "axios";

const axiosInstance: AxiosInstance = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL || "localhost:5000/api/v1",
  withCredentials: true,
  headers: {
    "Content-Type": "application/json",
    Accept: "application/json",
  },
  timeout: 10000,
});

axiosInstance.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

axiosInstance.interceptors.response.use(
  (response: AxiosResponse) => {
    return response;
  },
  (error) => {
    return Promise.reject(error);
  }
);

export async function getRandomPost(): Promise<ApiResponse<PostResponse>> {
  const response = await axiosInstance
    .get<ApiResponse<PostResponse>>("/posts")
    .catch((error) => {
      throw error;
    });
  return response.data;
}

export async function createPost(
  data: CreatePostData
): Promise<ApiResponse<PostCreatedResponse>> {
  const response = await axiosInstance.post("/posts", data).catch((error) => {
    throw error;
  });
  return response.data;
}

export async function addEmpathy(postId: string): Promise<void> {
  await axiosInstance.patch(`/posts/${postId}/empathy`).catch((error) => {
    throw error;
  });
}

export async function addProtest(postId: string): Promise<void> {
  await axiosInstance.patch(`posts/${postId}/protest`).catch((error) => {
    throw error;
  });
}

export async function addComment(
  postId: string,
  data: CreatePostCommentData
): Promise<void> {
  await axiosInstance.post(`/posts/${postId}/comments`, data).catch((error) => {
    throw error;
  });
}
