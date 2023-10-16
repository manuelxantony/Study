import { createSlice } from "@reduxjs/toolkit";

const messagesSlice = createSlice({
  name: "messages",
  initialState: {
    messagesData: {},
  },
  reducers: {
    setMessagesData: (state, action) => {
      const exsistingMessagesData = state.messagesData;
      const { chatId, messagesData } = action.payload;
      exsistingMessagesData[chatId] = messagesData;
      state.messagesData = { ...exsistingMessagesData };
      console.log(
        "from state ^^^^^^^^^^^^^^^^***^^^^^^^^^^^^^",
        state.messagesData[chatId]
      );
    },
  },
});

export const setMessagesData = messagesSlice.actions.setMessagesData;
export default messagesSlice.reducer;
