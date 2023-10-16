import { createSlice } from "@reduxjs/toolkit";

const storedUsersSlice = createSlice({
  name: "storedUsersSlice",
  initialState: {
    storedUsersData: {},
  },
  reducers: {
    storedUsersData: (state, action) => {
      const newUsers = action.payload.newUsers;
      const exsistingUsersData = state.storedUsersData;

      for (let userId in newUsers) {
        exsistingUsersData[userId] = newUsers[userId];
      }

      state.storedUsersData = exsistingUsersData;
    },
  },
});

export const storedUsersData = storedUsersSlice.actions.storedUsersData;
export default storedUsersSlice.reducer;
