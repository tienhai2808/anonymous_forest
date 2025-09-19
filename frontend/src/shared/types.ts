export interface ApiResponse<T = unknown> {
  message: string;
  data: T;
  status: number;
}

export interface PostResponse {
  post: Post;
}

export interface Post {
  _id: string;
  content: string;
  created_at: string;
  empathy_count: number;
  protest_count: number;
  comments: Comment[];
}

export interface Comment {
  _id: string;
  post_id: string;
  content: string;
  created_at: string;
}