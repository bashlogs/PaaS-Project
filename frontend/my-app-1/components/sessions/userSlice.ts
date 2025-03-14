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

export const fetchUserData = createAsyncThunk(
  'user/fetchUserData',
  async (_, { rejectWithValue }) => {
    try {
      const response = await fetch('http://localhost:8000/dashboard', {
        method: 'GET',
        credentials: 'include', // Ensures cookies are sent
      });

      if (response.status === 401) {
        console.log("Unauthorized response detected, redirecting to login...");
        return rejectWithValue('Unauthorized');
      }

      if (!response.ok) {
        const errorText = await response.text();
        console.log("Error response body:", errorText);
        return rejectWithValue(`Failed to fetch user data: ${response.status} ${response.statusText}`);
      }

      const data = await response.json();
      console.log("User data fetched successfully:", data);
      return data;
    } catch (error) {
      console.error("Fetch error:", error);
      return rejectWithValue(error instanceof Error ? error.message : 'Unknown error');
    }
  }
);

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
        state.error = action.payload as string || action.error.message || 'Failed to fetch user data';
        console.log("Rejected action:", action);
      });
  },
});

export default userSlice.reducer;