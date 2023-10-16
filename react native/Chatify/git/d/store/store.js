import { configureStore } from "@reduxjs/toolkit";
import authSlice from "./authSlice"; // name change
import chatSlice from "./chatSlice";
import messagesSlice from "./messagesSlice";
import storedUsersSlice from "./storedUsersSlice";

export const store = configureStore({
  reducer: {
    auth: authSlice,
    storedUsers: storedUsersSlice,
    chats: chatSlice,
    messages: messagesSlice,
  },
});
