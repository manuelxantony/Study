import { child, getDatabase, push, ref, update } from "firebase/database";
import { firebaseApp } from "../helper/firebaseHelper";

export const createChat = async (signedInUser, chatData) => {
  const newChatData = {
    ...chatData,
    createBy: signedInUser,
    updatedBy: signedInUser,
    createdTime: new Date().toISOString(),
    updatedTime: new Date().toISOString(),
  };

  const app = firebaseApp;
  const dbRef = ref(getDatabase(app));
  const chatNode = child(dbRef, "chats");
  const newChat = await push(chatNode, newChatData);

  const users = newChatData.users;

  for (let pos in users) {
    let userId = users[pos];
    const userNode = child(dbRef, `userChats/${userId}`);
    await push(userNode, newChat.key);
  }
  return newChat.key;
};

export const sendTextMessage = async (chatId, senderId, messageText) => {
  const messageData = {
    sentBy: senderId,
    sendDate: new Date().toISOString(),
    messageText,
  };

  const app = firebaseApp;
  const dbRf = ref(getDatabase(app));
  const messageRef = child(dbRf, `messages/${chatId}`);
  await push(messageRef, messageData);

  const chatRef = child(dbRf, `chats/${chatId}`);
  await update(chatRef, {
    updatedBy: senderId,
    updatedTime: new Date().toISOString(),
    lastMessageText: messageText,
  });
};
