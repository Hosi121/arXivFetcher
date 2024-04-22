export interface ApiResponse<T> {
  status: 'success' | 'error';
  data?: T;
  message?: string;
}

// イメージを書いておく
// export type ApiResponse<T> = {
//   status: 'success' | 'error';
//   data?: T;
//   message?: string;
// };