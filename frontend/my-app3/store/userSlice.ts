import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';

interface UserData {
  message: string;
  name: string;
  username: string;
  email: string;
}

interface UserState {
  userData: UserData | null;
  isLoading: boolean;
  error: string | null;
}

const initialState: UserState = {
  userData: null,
  isLoading: false,
  error: null,
};

export const fetchUserData = createAsyncThunk('user/fetchUserData', async () => {
  const response = await fetch('http://localhost:8000/dashboard', {
    method: 'GET',
    credentials: 'include', // Ensures cookies are sent
  });

  if (!response.ok) {
    throw new Error('Unauthorized');
  }

  const data = await response.json();
  return data;
});

const userSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchUserData.pending, (state) => {
        state.isLoading = true;
        state.error = null;
      })
      .addCase(fetchUserData.fulfilled, (state, action) => {
        state.userData = action.payload;
        state.isLoading = false;
      })
      .addCase(fetchUserData.rejected, (state, action) => {
        state.isLoading = false;
        state.error = action.error.message || 'Failed to fetch user data';
      });
  },
});

export default userSlice.reducer;