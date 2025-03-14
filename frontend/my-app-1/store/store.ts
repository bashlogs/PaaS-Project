import { configureStore } from "@reduxjs/toolkit";
import userReducer from './userSlice';

export const store = configureStore({
  reducer: {
    user: userReducer, // ✅ Correct key
  }
});
// Infer types for useDispatch and useSelector
export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;